package webshare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	yp "github.com/vxoid/yunroxy/proxy"
	v3 "github.com/vxoid/yunroxy/recaptcha/v3"
	"github.com/vxoid/yunroxy/updater/service"
	"github.com/vxoid/yunroxy/user"
)

type RegisterPayload struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Recaptcha   string `json:"recaptcha"`
	TosAccepted bool   `json:"tos_accepted"`
}

type RegisterResponse struct {
	Token                string `json:"tokenn"`
	LoggedInExistingUser bool   `json:"logged_in_existing_user"`
}

func CreateAccount(user *user.UserCredentials, recaptcha *v3.ReCaptchaV3, proxy *url.URL) (*WebShare, error) {
	if proxy != nil && !yp.IsSsl(proxy) {
		return nil, yp.ErrProxyMustBeSSL
	}

	client := http.Client{}
	if proxy != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	if recaptcha == nil {
		var err error
		recaptcha, err = v3.New("https://proxy2.webshare.io:443", ReCaptchaSiteKey, proxy)
		if err != nil {
			return nil, err
		}
	}

	token := recaptcha.GetToken()
	payload, err := json.Marshal(RegisterPayload{Email: user.GetEmail(), Password: user.GetPassword(), Recaptcha: token, TosAccepted: true})
	if err != nil {
		return nil, err
	}

	resp, err := client.Post("https://proxy.webshare.io/api/v2/register/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if strings.Contains(string(respBody), "captcha_invalid") {
			return nil, service.NewReCaptchaInvalidError(token, string(respBody))
		}

		return nil, fmt.Errorf("%s: %s", resp.Status, string(respBody))
	}

	var respData RegisterResponse
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}

	return &WebShare{token: respData.Token}, nil
}
