package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vxoid/yunroxy/db"
	"github.com/vxoid/yunroxy/proxy"
)

type ProxyRandomHandler struct {
	Db        *db.YunroxyDb
	Validator *proxy.ProxyValidator
}

type ResponseGetProxy struct {
	ProxyURL string `json:"proxy_url"`
}
type ErrorGetProxy struct {
	Error string `json:"error"`
}

func (h *ProxyRandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKeyHex := r.URL.Query().Get("api_key")
	apiKey, err := db.ParseApiKey(apiKeyHex)
	if err != nil {
		var resp ErrorGetProxy
		resp.Error = fmt.Sprintf("Please pass valid api_key to complete authentication.")
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	proxyURL, err := h.Db.GetRandomProxy(h.Validator, apiKey)
	if err != nil {
		var resp ErrorGetProxy
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var proxy ResponseGetProxy
	proxy.ProxyURL = proxyURL.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxy)
}
