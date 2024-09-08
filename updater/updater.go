package main

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/vxoid/yunroxy/proxy"
	"github.com/vxoid/yunroxy/updater/engagemint"
	"github.com/vxoid/yunroxy/updater/proxyscrape"
	"github.com/vxoid/yunroxy/updater/service"
)

const TICK = time.Minute * 15

var services = []service.Service{engagemint.GetService(), proxyscrape.GetService()} // Const

func tick(validator *proxy.ProxyValidator) {
	log.Printf("------ New TICK ------\n")
	removeBrokenProxies()
	fetchNewProxies(validator)
}

func fetchNewProxies(validator *proxy.ProxyValidator) {
	var wg sync.WaitGroup
	for _, s := range services {
		wg.Add(1)
		go func(s service.Service) {
			newProxies, err := s.FetchProxies(nil)
			if len(newProxies) < 1 {
				if err != nil {
					color.Red("[%s] throws error: %s", s.GetId(), err)
					return
				}
				color.Yellow("[%s]: no new proxies", s.GetId())
				return
			}

			if err != nil {
				// r, ok := err.(*service.RateError)
				// if ok {
				// restrict proxy here
				// }
				color.Red("[%s]: unknown error %s", s.GetId(), err)
			}

			var validatedNewProxies []*url.URL
			for _, p := range newProxies {
				err := validator.Validate(p)
				if err != nil {
					color.Red("[%s]: %s -> %s", s.GetId(), p, err)
					continue
				}

				color.Green("[%s]: %s succeed", s.GetId(), p)
				validatedNewProxies = append(validatedNewProxies, p)
			}

			color.Green("[%s] gifted new proxies: %v", s.GetId(), validatedNewProxies)
			wg.Done()
		}(s)
	}
}

func removeBrokenProxies() {

}

func main() {
	validator, err := proxy.NewValidator()
	if err != nil {
		color.Red("%s", err)
	}

	ticker := time.NewTicker(TICK)
	defer ticker.Stop()

	go tick(validator)

	for range ticker.C {
		go tick(validator)
	}
}
