package db

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/vxoid/yunroxy/proxy"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type YunroxyDb struct {
	Db *gorm.DB
}

func NewApiDb(dbPath string) (*YunroxyDb, error) {
	var db, Err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if Err != nil {
		return nil, Err
	}
	db.AutoMigrate(&User{}, &Proxy{})
	return &YunroxyDb{Db: db}, nil
}

func (slf YunroxyDb) GetUserByApiKey(apiKeyHex string) (*User, error) {
	apiKey, err := hex.DecodeString(apiKeyHex)
	if err != nil {
		return nil, err
	}

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

func (slf YunroxyDb) AddProxy(serviceUrl string, proxyUrl *url.URL) {
	slf.Db.Create(&Proxy{Service: serviceUrl, ProxyUrl: proxyUrl.String()})
}

func (slf YunroxyDb) DeleteProxy(proxyUrl *url.URL) error {
	return slf.deleteProxy(proxyUrl.String())
}

func (slf YunroxyDb) deleteProxy(proxyUrl string) error {
	return slf.Db.Where("proxy_url = ?", proxyUrl).Delete(&Proxy{}).Error
}

func (slf YunroxyDb) GetRandomProxy(apiKeyHex string) (*url.URL, error) {
	_, err := slf.GetUserByApiKey(apiKeyHex)
	if err != nil {
		return nil, err
	}

	var proxyAssoc, last Proxy
	max := slf.Db.Last(&last)
	if max.Error != nil {
		return nil, max.Error
	}

	res := slf.Db.Find(&proxyAssoc, randRange(1, int(last.ID)))
	if res.Error != nil {
		return nil, res.Error
	}
	return proxy.Parse(proxyAssoc.ProxyUrl)
}

func (slf YunroxyDb) GetAllProxies() ([]*url.URL, error) {
	var proxies []Proxy
	res := slf.Db.Find(&proxies)
	if res.Error != nil {
		return nil, res.Error
	}

	var result []*url.URL
	for _, proxyAssoc := range proxies {
		proxyUrl, err := proxy.Parse(proxyAssoc.ProxyUrl)
		if err != nil {
			slf.deleteProxy(proxyAssoc.ProxyUrl)
			continue
		}
		result = append(result, proxyUrl)
	}
	return result, nil
}

func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (slf YunroxyDb) CreateApiKey() ([]byte, error) {
	bytes := make([]byte, 256)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}

	res := slf.Db.Create(&User{ApiKey: bytes})
	if res.Error != nil {
		return nil, res.Error
	}
	return bytes, nil
}
