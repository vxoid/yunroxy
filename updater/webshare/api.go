package webshare

const ReCaptchaSiteKey = "6LeHZ6UUAAAAAKat_YS--O2tj_by3gv3r_l03j9d"
const ReCaptchaSite = "https://proxy2.webshare.io:443"

type WebShare struct {
	token string
}

func (w *WebShare) GetToken() string {
	return w.token
}
