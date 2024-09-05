package proxy

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"golang.org/x/exp/slices"
)

const CheckBrokenRetries = 3

var ErrProxyMustBeSSL = errors.New("this function accepts SSL only proxies, proxy must be SSL")

func IsSsl(proxy *url.URL) bool {
	return proxy.Scheme == "https" || proxy.Scheme == "socks4" || proxy.Scheme == "socks5"
}

func GetSupportedProtocols() []string {
	return []string{"http", "https", "socks4", "socks5"}
}

func NewProxy(proto string, hostname string, port uint16, user *url.Userinfo) (*url.URL, error) {
	if !slices.Contains(GetSupportedProtocols(), proto) {
		return nil, fmt.Errorf("\"%s\" protocol is not supported", proto)
	}
	return &url.URL{Scheme: proto, User: user, Host: hostname + ":" + strconv.FormatUint(uint64(port), 10)}, nil
}
