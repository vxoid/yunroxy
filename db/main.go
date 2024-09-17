package main

import (
	"encoding/json"
	"net/http"
	"time"

	yunroxyDB "github.com/vxoid/yunroxy/yunroxyDb"
)

var sql, Err = yunroxyDB.NewApiDb("yunroxyDB/db.db")

type ResponseGetProxy struct {
	ProxyURL string `json:"proxy_url"`
}

func getProxy(w http.ResponseWriter, r *http.Request) {
	apiCheck := r.URL.Query().Get("api_key")
	if !sql.IsApiKey(apiCheck) {
		return
	}
	proxyURL := sql.GetProxy()
	var proxy ResponseGetProxy
	proxy.ProxyURL = proxyURL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxy)
}

func main() {

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/getProxy", getProxy)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
