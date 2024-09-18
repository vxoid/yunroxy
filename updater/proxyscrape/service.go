package proxyscrape

import (
	"net/url"

	"github.com/vxoid/yunroxy/updater/service"
)

type ProxyScrapeService struct{}

func (w *ProxyScrapeService) FetchProxies(proxy *url.URL) ([]*url.URL, error) {
	proxies, _, err := GetProxies(nil, proxy)
	return proxies, err
}

func (w *ProxyScrapeService) GetId() string {
	return "https://proxyscrape.com/"
}

func GetService() service.Service {
	return &ProxyScrapeService{}
}
