package main

import (
	"flag"
	"github.com/nankeen/protecc/proxies"
	"log"
	"net/http"
)

var proxyTargetHost = flag.String("target", "localhost", "Where the requests should be proxied to.")
var proxyListen = flag.String("listen", "localhost:1337", "Where the proxy server listens on")

func checkWhitelistMethod() {
}

func main() {
	flag.Parse()

	tp := proxies.NewTransparentHttpProxy(*proxyTargetHost)

	server := &http.Server{
		Addr: *proxyListen,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)
			// Only allow get methods
			if r.Method == http.MethodGet {
				tp.HttpHandler(w, r)
			} else {
				log.Printf("Blocking %s request from %s", r.Method, r.RemoteAddr)
			}
		}),
	}

	log.Printf("Listening on %s", *proxyListen)
	log.Printf("Relaying for target: %s", *proxyTargetHost)
	log.Fatal(server.ListenAndServe())
}
