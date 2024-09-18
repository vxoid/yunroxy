package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vxoid/yunroxy/api/handlers/proxy"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
)

func Api(config *config.Config) {
	db, err := db.NewApiDb(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/proxy/random", &proxy.ProxyRandomHandler{Db: db})

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
