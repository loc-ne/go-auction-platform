package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/loc-ne/go-auction/services/api-gateway/internal/config"
)

type Router struct {
	authProxy *httputil.ReverseProxy
	productProxy *httputil.ReverseProxy
	biddingProxy *httputil.ReverseProxy
}

func NewRouter(cfg *config.Config) *Router {
	authTarget, err := url.Parse(cfg.AuthServiceURL)
	if err != nil {
		log.Fatalf("URL parse error target AuthServiceURL: %v", err)
	}

	productTarget, err := url.Parse(cfg.ProductServiceURL)
	if err != nil {
		log.Fatalf("URL parse error target ProductServiceURL: %v", err)
	}

	biddingTarget, err := url.Parse(cfg.BiddingServiceURL)
	if err != nil {
		log.Fatalf("URL parse error target BiddingServiceURL: %v", err)
	}

	return &Router{
		authProxy: httputil.NewSingleHostReverseProxy(authTarget),
		productProxy: httputil.NewSingleHostReverseProxy(productTarget),
		biddingProxy: httputil.NewSingleHostReverseProxy(biddingTarget),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := req.URL.Path

	if strings.HasPrefix(path, "/api/v1/auth") {
		r.authProxy.ServeHTTP(w, req)
		return
	}
	if strings.HasPrefix(path, "/api/v1/product") || strings.HasPrefix(path, "/api/v1/media") {
		r.productProxy.ServeHTTP(w, req)
		return
	}
	if strings.HasPrefix(path, "/api/v1/bids") {
		r.biddingProxy.ServeHTTP(w, req)
		return
	}

	http.Error(w, "Service not found", http.StatusNotFound)
}
