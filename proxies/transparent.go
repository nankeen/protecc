package proxies

import (
	"net/http"
	"net/http/httputil"
)

type TransparentHttpProxy struct{ ProxyTargetHost string }

func NewTransparentHttpProxy(proxyTargetHost string) *TransparentHttpProxy {
	return &TransparentHttpProxy{ProxyTargetHost: proxyTargetHost}
}

// This type implements http.Handler
func (tp *TransparentHttpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	director := func(target *http.Request) {
		target.URL.Scheme = "http"
		target.URL.Path = r.URL.Path

		target.URL.Host = tp.ProxyTargetHost
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
