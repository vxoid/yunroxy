package yunroxyDB

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiDb struct {
	db  *gorm.DB
}

func NewApiDb(dbPath string) (*ApiDb, error){
	var db, Err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if  Err != nil { return nil, Err }
	db.AutoMigrate(&User{}, &Proxy{})
	
	return &ApiDb{db : db}, nil
}

func (slf ApiDb) IsApiKey(key string) bool {
	var user User
	res := slf.db.Take(&user, "api_key = ?", key)
	return res.Error == nil
}

func (slf ApiDb) AddProxy(serviceUrl string, proxyUrl string){
	slf.db.Create(&Proxy{Service: serviceUrl,  ProxyUrl: proxyUrl})
}

func (slf ApiDb) DelProxy(proxyUrl string){	
	var proxy Proxy
	res := slf.db.Take(&proxy, "proxy_url = ?", proxyUrl)
	if res.Error != nil{
		return
	} 
	slf.db.Unscoped().Delete(Proxy{ProxyUrl: proxyUrl})
}
func (slf ApiDb) GetProxy() *gorm.DB{
	var proxy Proxy
	return slf.db.Take(&proxy)
}