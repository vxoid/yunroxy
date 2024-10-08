package main

import (
	"log"
	"sync"

	"github.com/vxoid/yunroxy/api"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
	"github.com/vxoid/yunroxy/proxy"
	"github.com/vxoid/yunroxy/updater"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	validator, err := proxy.NewValidator()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewApiDb(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go api.Api(config, validator, db)
	go updater.Updater(config, validator, db)

	wg.Wait()
}
