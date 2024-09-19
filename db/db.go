package db

import (
	"encoding/hex"
	"fmt"
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ApiDb struct {
	Db *gorm.DB
}

func NewApiDb(dbPath string) (*ApiDb, error) {
	var db, Err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if Err != nil {
		return nil, Err
	}
	db.AutoMigrate(&User{}, &Proxy{})
	return &ApiDb{Db: db}, nil
}

func (slf ApiDb) IsApiKey(key string) bool {
	var user User
	return slf.Db.First(&user, "api_key = ?", key).Error == nil
}

func (slf ApiDb) AddProxy(serviceUrl string, proxyUrl string) {
	slf.Db.Create(&Proxy{Service: serviceUrl, ProxyUrl: proxyUrl})
}

func (slf ApiDb) DelProxy(proxyUrl string) {
	var proxy Proxy
	res := slf.Db.Take(&proxy, "proxy_url = ?", proxyUrl)
	if res.Error != nil {
		return
	}
	slf.Db.Unscoped().Delete(Proxy{ProxyUrl: proxyUrl})
}

func (slf ApiDb) GetProxy(apikey string) (string, error) {
	if !slf.IsApiKey(apikey) {
		return "", fmt.Errorf("invalid api key")
	}
	var proxy, last Proxy
	max := slf.Db.Last(&last)
	if max.Error != nil {
		return "", max.Error
	}
	res := slf.Db.Find(&proxy, randRange(1, int(last.ID)))
	if res.Error != nil {
		return "", res.Error
	}
	return proxy.ProxyUrl, nil
}

func (slf ApiDb) GetAllProxies() []string {
	var proxies []string
	var proxy Proxy
	var len = slf.Db.Find(&proxy).RowsAffected
	for i := 0; i < int(len); i++ {
		proxies = append(proxies, proxy.ProxyUrl)
	}
	return proxies
}

func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func CreateApiKey() (string, error) {
	bytes := make([]byte, 256)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
