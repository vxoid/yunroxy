package engagemint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	yp "github.com/vxoid/yunroxy/proxy"
)

type ProxyResponse struct {
	Asn      string `json:"asn"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	Region   string `json:"region"`
}

func GetProxies(proxy *url.URL) ([]*url.URL, error) {
	client := http.Client{}
	if proxy != nil {
		if !yp.IsSsl(proxy) {
			return []*url.URL{}, yp.ErrProxyMustBeSSL
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	req, err := http.NewRequest(http.MethodGet, "https://broker.engagemintcreative.com/proxies", nil)
	if err != nil {
		return []*url.URL{}, err
	}
	req.Header.Set("Referer", "https://proxiware.com/")

	resp, err := client.Do(req)
	if err != nil {
		return []*url.URL{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []*url.URL{}, err
	}
	defer resp.Body.Close()

	var parsedResp []ProxyResponse
	err = json.Unmarshal(respBody, &parsedResp)
	if err != nil {
		return []*url.URL{}, err
	}

	var proxies []*url.URL
	for _, respProxy := range parsedResp {
		proxyUrl, err := yp.Parse(fmt.Sprintf("%s://%s:%d", respProxy.Protocol, respProxy.Ip, respProxy.Port))
		if err != nil {
			continue
		}
		proxies = append(proxies, proxyUrl)
	}

	return proxies, nil
}
