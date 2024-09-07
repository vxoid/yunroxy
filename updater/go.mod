module github.com/vxoid/yunroxy/updater

go 1.19

require github.com/vxoid/yunroxy/recaptcha v0.0.0
require github.com/vxoid/yunroxy/user v0.0.0

require (
	github.com/fatih/color v1.17.0
	github.com/vxoid/yunroxy/proxy v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

replace github.com/vxoid/yunroxy/recaptcha => ../recaptcha

replace github.com/vxoid/yunroxy/proxy => ../proxy
replace github.com/vxoid/yunroxy/user => ../user
