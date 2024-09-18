package main

import (
	"log"
	"sync"

	"github.com/vxoid/yunroxy/api"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/updater"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go api.Api(config)
	go updater.Updater(config)

	wg.Wait()
}
