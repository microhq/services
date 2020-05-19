package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/micro/go-micro/v2/web"
)

func main() {
	rp := new(httputil.ReverseProxy)
	rp.Director = func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = "go.micro.mu"
		req.URL.Path = "/usage"
	}

	service := web.NewService(
		web.Name("go.micro.web.usage"),
	)

	service.Init()
	service.Handle("/", rp)
	service.Run()
}
