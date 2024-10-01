package db

type Proxy struct {
	ID       uint `gorm:"primarykey"`
	Service  string
	ProxyUrl string `gorm:"unique;not null"`
}

type User struct {
	ID     uint   `gorm:"primarykey"`
	ApiKey []byte `gorm:"unique;not null"`
}
