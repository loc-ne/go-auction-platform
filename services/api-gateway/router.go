package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/loc-ne/go-auction/services/api-gateway/config"
)

type Router struct {
	authProxy *httputil.ReverseProxy
}

func NewRouter(cfg *config.Config) *Router {
	authTarget, err := url.Parse(cfg.AuthServiceURL)
	if err != nil {
		log.Fatalf("URL parse error target AuthServiceURL: %v", err)
	}

	return &Router{
		authProxy: httputil.NewSingleHostReverseProxy(authTarget),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if strings.HasPrefix(path, "/api/auth") {
		req.URL.Path = strings.TrimPrefix(path, "/api/auth")
		r.authProxy.ServeHTTP(w, req)
		return
	}

	http.Error(w, "Service not found", http.StatusNotFound)
}
