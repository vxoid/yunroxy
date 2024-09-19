package proxy

import (
	"encoding/json"
	"net/http"

	"github.com/vxoid/yunroxy/db"
)

type ProxyRandomHandler struct {
	Db *db.ApiDb
}

type ResponseGetProxy struct {
	ProxyURL string `json:"proxy_url"`
}
type ErrorGetProxy struct {
	Error string `json:"error"`
}

func (h *ProxyRandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proxyURL, err := h.Db.GetProxy("")
	if err != nil {
		var resp ErrorGetProxy
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var proxy ResponseGetProxy
	proxy.ProxyURL = proxyURL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxy)
}
