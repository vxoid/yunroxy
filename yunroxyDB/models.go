package yunroxyDB

import "gorm.io/gorm"


type Proxy struct {
	gorm.Model
	ID uint
	Service string

}

type User struct {
	gorm.Model
	ID  uint
	ApiKey string

}
