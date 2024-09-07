package proxyscrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	yp "github.com/vxoid/yunroxy/proxy"
)

const pageSize = 100

type GetProxiesResponse struct {
	ShownRecords uint                      `json:"shown_records"`
	TotalRecords uint                      `json:"total_records"`
	Limit        uint                      `json:"limit"`
	Skip         uint                      `json:"skip"`
	NextPage     bool                      `json:"nextpage"`
	Proxies      []GetProxiesResponseProxy `json:"proxies"`
}

type GetProxiesResponseProxy struct {
	Proxy    string `json:"proxy"`
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Ssl      bool   `json:"ssl"`
}

func GetProxies(page *int, proxy *url.URL) ([]*url.URL, uint, error) {
	client := http.Client{}
	if proxy != nil {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	skip := 0
	if page != nil {
		skip = pageSize * *page
	}
	link := fmt.Sprintf("https://api.proxyscrape.com/v3/free-proxy-list/get?request=getproxies&skip=%d&proxy_format=protocolipport&format=json", skip)
	if proxy != nil && !yp.IsSsl(proxy) {
		link = fmt.Sprintf("http://api.proxyscrape.com/v3/free-proxy-list/get?request=getproxies&skip=%d&proxy_format=protocolipport&format=json", skip)
	}

	if page != nil {
		link = fmt.Sprintf("%s&limit=%d", link, pageSize)
	}
	resp, err := client.Get(link)
	if err != nil {
		return []*url.URL{}, 0, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []*url.URL{}, 0, err
	}
	defer resp.Body.Close()

	var parsedResp GetProxiesResponse
	err = json.Unmarshal(respBody, &parsedResp)
	if err != nil {
		return []*url.URL{}, 0, err
	}

	var proxies []*url.URL
	for _, proxy := range parsedResp.Proxies {
		proxyUrl, err := yp.Parse(proxy.Proxy)
		if err != nil {
			continue
		}
		proxies = append(proxies, proxyUrl)
	}

	return proxies, parsedResp.TotalRecords, nil
}
