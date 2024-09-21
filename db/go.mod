module github.com/vxoid/yunroxy/db

go 1.19

require (
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.11
)

require golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
	github.com/vxoid/yunroxy/proxy v0.0.0
	golang.org/x/text v0.18.0 // indirect
)

replace github.com/vxoid/yunroxy/proxy => ../proxy
