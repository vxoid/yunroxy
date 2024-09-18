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

func (h *ProxyRandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiCheck := r.URL.Query().Get("api_key")
	if !h.Db.IsApiKey(apiCheck) {
		return
	}
	proxyURL := h.Db.GetProxy()
	var proxy ResponseGetProxy
	proxy.ProxyURL = proxyURL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxy)
}
