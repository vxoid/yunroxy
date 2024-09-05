package proxy

import (
	"fmt"
	"net"
	"net/url"
)

type ProxyValidator struct {
	selfIp net.IP
}

func NewValidator() (*ProxyValidator, error) {
	ip, err := GetIp(nil)
	if err != nil {
		return nil, err
	}

	return &ProxyValidator{selfIp: ip}, nil
}

func (pv *ProxyValidator) Validate(proxy *url.URL) (bool, []error) {
	var errs []error
	for i := 0; i < CheckBrokenRetries; i++ {
		err := pv.TryValidate(proxy)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return len(errs) < CheckBrokenRetries, errs
}

func (pv *ProxyValidator) TryValidate(proxy *url.URL) error {
	if proxy == nil {
		return fmt.Errorf("proxy must not be nil")
	}

	ip, err := GetIp(proxy)
	if err != nil {
		return err
	}

	if ip.Equal(pv.selfIp) {
		return fmt.Errorf("api.ipify.org returns %s (%s) which is local machine ip", ip.String(), proxy.String())
	}

	return nil
}
