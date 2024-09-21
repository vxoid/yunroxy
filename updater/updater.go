package updater

import (
	"log"
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
const UPDATER_API_KEY = "0x52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c64981855ad8681d0d86d1e91e00167939cb6694d2c422acd208a0072939487f6999eb9d18a44784045d87f3c67cf22746e995af5a25367951baa2ff6cd471c483f15fb90badb37c5821b6d95526a41a9504680b4e7c8b763a1b1d49d4955c8486216325253fec738dd7a9e28bf921119c160f0702448615bbda08313f6a8eb668d20bf5059875921e668a5bdf2c7fc4844592d2572bcd0668d2d6c52f5054e2d0836bf84c7174cb7476364cc3dbd968b0f7172ed85794bb358b0c3b525da1786f9fff094279db1944ebd7a19d0f7bbacbe0255aa5b7d44bec40f84c892b9bffd436"

var services = []service.Service{engagemint.GetService(), proxyscrape.GetService()} // Const

func tick(database *db.YunroxyDb, validator *proxy.ProxyValidator) {
	log.Printf("------ New TICK ------\n")
	go removeBrokenProxies(database, validator)
	go fetchNewProxies(database, validator)
}

func fetchNewProxies(database *db.YunroxyDb, validator *proxy.ProxyValidator) {
	for _, s := range services {
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

				database.AddProxy(s.GetId(), p)
				color.Green("[%s]: %s succeed", s.GetId(), p)
			}
		}(s)
	}
}

func removeBrokenProxies(database *db.YunroxyDb, validator *proxy.ProxyValidator) {
	proxies, err := database.GetAllProxies()
	if err != nil {
		color.Red("couldn't fetch proxies from db.")
		return
	}
	for _, proxyUrl := range proxies {
		err := validator.Validate(proxyUrl)
		if err != nil {
			color.Yellow("removing [%s]: %s", proxyUrl.String(), err)
			database.DeleteProxy(proxyUrl)
		}
	}
}

func Updater(_ *config.Config, validator *proxy.ProxyValidator, db *db.YunroxyDb) {
	ticker := time.NewTicker(TICK)
	defer ticker.Stop()

	go tick(db, validator)

	for range ticker.C {
		go tick(db, validator)
	}
}
