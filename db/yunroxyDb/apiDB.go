package yunroxyDB

import (
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiDb struct {
	Db *gorm.DB
}

func NewApiDb(dbPath string) (*ApiDb, error) {
	var db, Err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if Err != nil {
		return nil, Err
	}
	db.AutoMigrate(&User{}, &Proxy{})
	return &ApiDb{Db: db}, nil
}

func (slf ApiDb) IsApiKey(key string) bool {
	var user User
	res := slf.Db.Take(&user, "api_key = ?", key)
	return res.Error == nil
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

func (slf ApiDb) GetProxy() string {
	var proxy, last Proxy
	max := slf.Db.Last(&last)
	res := slf.Db.Find(&proxy, randRange(1, int(last.ID)))

	if res.Error != nil && max.Error != nil {
		return ""
	}
	return proxy.ProxyUrl
}

func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
