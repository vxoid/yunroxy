package engagemint

import (
	"net/url"

	"github.com/vxoid/yunroxy/updater/service"
)

type ProxyScrapeService struct{}

func (w *ProxyScrapeService) FetchProxies(proxy *url.URL) ([]*url.URL, error) {
	return GetProxies(proxy)
}

func (w *ProxyScrapeService) GetId() string {
	return "engagemint"
}

func GetService() service.Service {
	return &ProxyScrapeService{}
}
