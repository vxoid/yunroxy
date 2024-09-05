package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

func GetIp(proxy *url.URL) (net.IP, error) {
	link := "http://api.ipify.org"

	client := http.Client{}
	if proxy != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
		if IsSsl(proxy) {
			link = "https://api.ipify.org"
		}
	}

	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respIp := net.ParseIP(string(respBody))
	if respIp == nil {
		return nil, fmt.Errorf("api.ipify.org returned value not parsable to ip '%s'", string(respBody))
	}

	return respIp, nil
}
