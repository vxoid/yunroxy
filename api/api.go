package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vxoid/yunroxy/api/handlers/proxy"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
	yp "github.com/vxoid/yunroxy/proxy"
)

func Api(config *config.Config, validator *yp.ProxyValidator, db *db.YunroxyDb) {
	http.Handle("/proxy/random", &proxy.ProxyRandomHandler{Db: db, Validator: validator})

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
