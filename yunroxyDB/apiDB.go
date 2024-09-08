package yunroxyDB

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApiDb struct {
	db  *gorm.DB
}

func  NewApiDb(dbPath string) (*ApiDb, error){
	var db, Err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if  Err != nil { return nil, Err }
	db.AutoMigrate(&User{}, &Proxy{})
	
	return &ApiDb{db : db}, Err
}

func (slf ApiDb) IsApiKey(key string) bool {
	var user User
	res := slf.db.Take(&user, "api_key = ?", key)
	return res.Error == nil
}