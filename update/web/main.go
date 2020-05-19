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
		req.URL.Path = "/update"
	}

	service := web.NewService(
		web.Name("go.micro.web.update"),
	)

	service.Init()
	service.Handle("/", rp)
	service.Run()
}
