package webshare

import (
	"fmt"
	"net/url"

	v3 "github.com/vxoid/yunroxy/recaptcha/v3"
	"github.com/vxoid/yunroxy/updater/service"
	"github.com/vxoid/yunroxy/user"
)

type WebShareService struct{}

func (w *WebShareService) FetchProxies(proxy *url.URL) ([]*url.URL, error) {
	recaptcha, err := v3.New(ReCaptchaSite, ReCaptchaSiteKey, proxy)
	if err != nil {
		return []*url.URL{}, err
	}

	return GenAccountsWhilePossible(recaptcha, proxy)
	// recaptcha, err = v3.Bypass(ReCaptchaSite, ReCaptchaSiteKey, proxy)
	// if err != nil {
	// 	return []*url.URL{}, err
	// }

	// proxies2, err := GenAccountsWhilePossible(recaptcha, proxy)
	// return append(proxies1, proxies2...), err
}

func GenAccountsWhilePossible(recaptcha *v3.ReCaptchaV3, proxy *url.URL) ([]*url.URL, error) {
	var proxies []*url.URL
	for {
		account, err := CreateAccount(user.NewRandom(), recaptcha, proxy)
		if err != nil {
			return proxies, err
		}

		fmt.Printf("token: %s\n", account.GetToken())
	}
}

func (w *WebShareService) GetId() string {
	return "webshare"
}

func GetService() service.Service {
	return &WebShareService{}
}
