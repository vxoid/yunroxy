package updater

import (
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
	"github.com/vxoid/yunroxy/proxy"
	"github.com/vxoid/yunroxy/updater/engagemint"
	"github.com/vxoid/yunroxy/updater/proxyscrape"
	"github.com/vxoid/yunroxy/updater/service"
)

const TICK = time.Minute * 15
const UPDATER_API_KEY = "0xb37392919543d51f607e4db8114fc448054b7d93692fa91f6827e3913c9b39ab339dfe48066cb4e60d58fb2bb00e8adc61592852500f71a9dfc78b716bc73e54c06a38c6b7a7ddd79a1ab8f1ce199db33773f52cc4979f77828b2faa80716a43cacf1a11b6b0baa8781ba276110a0e80d79af421a1872df7c481bfc0cff5d17f17c3fdac2f2c5ba7a588f8ae34c8a9ae150a3b6763bfa87132a8dcb1556b2292f539c013d008372fdcfb6c0e589575166507a2adc934330f5ecdc06747cd5af4663b8402f7f5408a37f754430bb5c31bae63768528429a1c282ec78243d1c8293df677"

var services = []service.Service{engagemint.GetService(), proxyscrape.GetService()} // Const

func tick(database *db.ApiDb, validator *proxy.ProxyValidator) {
	log.Printf("------ New TICK ------\n")
	removeBrokenProxies()
	fetchNewProxies(database, validator)
}

func fetchNewProxies(database *db.ApiDb, validator *proxy.ProxyValidator) {
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
			}

			if err != nil {
				// r, ok := err.(*service.RateError)
				// if ok {
				// restrict proxy here
				// }
				color.Red("[%s]: unknown error %s", s.GetId(), err)
				return
			}

			for _, p := range newProxies {
				err := validator.Validate(p)
				if err != nil {
					color.Yellow("[%s]: %s -> %s", s.GetId(), p, err)
					continue
				}

				database.AddProxy(s.GetId(), p.String())
				color.Green("[%s]: %s succeed", s.GetId(), p)
			}

			wg.Done()
		}(s)
	}
}

func removeBrokenProxies() {

}

func Updater(config *config.Config) {
	validator, err := proxy.NewValidator()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewApiDb(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(TICK)
	defer ticker.Stop()

	go tick(db, validator)

	for range ticker.C {
		go tick(db, validator)
	}
}
