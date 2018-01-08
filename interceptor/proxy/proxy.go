package interceptor

import (
	"net/http/httputil"
	"net/url"
)

// Proxy asdfs
func Proxy() {
	proxyURL, _ := url.Parse("https://jsonplaceholder.typicode.com")
	revProxy := httputil.NewSingleHostReverseProxy(proxyURL)
	revProxy.Director()
}
