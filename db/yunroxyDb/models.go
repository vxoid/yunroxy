package yunroxyDB

type Proxy struct {
	ID       uint `gorm:"primarykey"`
	Service  string
	ProxyUrl string
}

type User struct {
	ID     uint `gorm:"primarykey"`
	ApiKey string
}
