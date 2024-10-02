module github.com/vxoid/yunroxy

go 1.19

require github.com/vxoid/yunroxy/updater v0.0.0

require github.com/vxoid/yunroxy/api v0.0.0

require github.com/vxoid/yunroxy/config v0.0.0

require github.com/vxoid/yunroxy/db v0.0.0

require github.com/vxoid/yunroxy/proxy v0.0.0

require github.com/vxoid/yunroxy/recaptcha v0.0.0

require github.com/vxoid/yunroxy/user v0.0.0

require (
	github.com/fatih/color v1.17.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	gorm.io/driver/sqlite v1.5.6 // indirect
	gorm.io/gorm v1.25.12 // indirect
)

replace github.com/vxoid/yunroxy/updater => ./updater

replace github.com/vxoid/yunroxy/api => ./api

replace github.com/vxoid/yunroxy/config => ./config

replace github.com/vxoid/yunroxy/db => ./db

replace github.com/vxoid/yunroxy/proxy => ./proxy

replace github.com/vxoid/yunroxy/recaptcha => ./recaptcha

replace github.com/vxoid/yunroxy/user => ./user
