package service

import "net/url"

type Service interface {
	FetchProxies(proxy *url.URL) ([]*url.URL, error)
	GetId() string
}
