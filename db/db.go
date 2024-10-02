package db

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/proxy"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type YunroxyDb struct {
	Db *gorm.DB
}

func NewApiDb(dbPath string) (*YunroxyDb, error) {
	projectDir, err := config.GetProjectDir()
	if err != nil {
		return nil, err
	}

	var db, Err = gorm.Open(sqlite.Open(filepath.Join(projectDir, dbPath)), &gorm.Config{})
	if Err != nil {
		return nil, Err
	}

	db.AutoMigrate(&User{}, &Proxy{})
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA locking_mode = NORMAL;")
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	return &YunroxyDb{Db: db}, nil
}

func (slf *YunroxyDb) GetUserByApiKey(apiKey []byte) (*User, error) {
	var users []User
	result := slf.Db.Limit(1).Find(&users, "api_key = ?", apiKey)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) < 1 {
		return nil, fmt.Errorf("'%s' is not valid api key", apiKey)
	}

	return &users[0], nil
}

func (slf *YunroxyDb) AddProxy(serviceUrl string, proxyUrl *url.URL) error {
	tx := slf.Db.Begin()

	if err := tx.Create(&Proxy{Service: serviceUrl, ProxyUrl: proxyUrl.String()}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (slf *YunroxyDb) DeleteProxy(proxyUrl *url.URL) error {
	return slf.deleteProxy(proxyUrl.String())
}

func (slf *YunroxyDb) deleteProxy(proxyUrl string) error {
	tx := slf.Db.Begin()

	err := tx.Exec("DELETE FROM proxies WHERE proxy_url = ?", proxyUrl).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (slf *YunroxyDb) parseProxy(proxyAssoc Proxy) (*url.URL, error) {
	proxyUrl, err := proxy.Parse(proxyAssoc.ProxyUrl)
	if err != nil {
		slf.deleteProxy(proxyAssoc.ProxyUrl)
		return nil, err
	}
	return proxyUrl, nil
}

func (y *YunroxyDb) Validate(validator *proxy.ProxyValidator, proxy *url.URL) error {
	err := validator.Validate(proxy)
	if err != nil {
		delErr := y.DeleteProxy(proxy)
		if delErr != nil {
			log.Printf("Cannot delete %s from the db: %s", proxy, delErr)
		}
	}

	return err
}

func (slf *YunroxyDb) GetRandomProxy(validator *proxy.ProxyValidator, apiKey []byte) (*url.URL, error) {
	_, err := slf.GetUserByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	var count int64
	slf.Db.Model(&Proxy{}).Count(&count)

	if count < 1 {
		return nil, fmt.Errorf("no proxies available (empty db)")
	}

	rand.Seed(time.Now().UnixNano())
	randomOffset := rand.Intn(int(count))

	var proxyAssoc Proxy

	if err := slf.Db.Offset(randomOffset).Limit(1).Find(&proxyAssoc).Error; err != nil {
		return nil, err
	}

	proxyUrl, err := slf.parseProxy(proxyAssoc)
	if err != nil {
		return nil, err
	}

	err = slf.Validate(validator, proxyUrl)
	if err != nil {
		return slf.GetRandomProxy(validator, apiKey)
	}

	return proxyUrl, nil
}

func (slf *YunroxyDb) GetAllProxies() ([]*url.URL, error) {
	var proxies []Proxy
	res := slf.Db.Find(&proxies)
	if res.Error != nil {
		return nil, res.Error
	}

	var result []*url.URL
	for _, proxyAssoc := range proxies {
		proxyUrl, err := slf.parseProxy(proxyAssoc)
		if err != nil {
			return nil, err
		}

		result = append(result, proxyUrl)
	}
	return result, nil
}

func (slf *YunroxyDb) CreateApiKey() ([]byte, error) {
	bytes := make([]byte, 256)
	for i := 0; i < 16; i++ {
		UUID := uuid.New()
		copy(bytes[i*16:], UUID[:])
	}

	res := slf.Db.Create(&User{ApiKey: bytes})
	if res.Error != nil {
		return nil, res.Error
	}
	return bytes, nil
}

func (y *YunroxyDb) RemoveApiKey(apiKey []byte) error {
	var users []User
	result := y.Db.Where("api_key = ?", apiKey).Find(&users)
	if result.Error != nil {
		return result.Error
	}

	if len(users) < 1 {
		return fmt.Errorf("invalid api key '%s'", hex.EncodeToString(apiKey))
	}

	return y.Db.Where("api_key = ?", apiKey).Delete(&User{}).Error
}

func ParseApiKey(apiKeyHex string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(apiKeyHex, "0x"))
}
