package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	"github.com/micro/services/m3o/api/handler"
	pb "github.com/micro/services/m3o/api/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.m3o"),
		micro.Version("latest"),
	)
	service.Init()

	pb.RegisterAccountsHandler(service.Server(), handler.NewAccounts(service))
	pb.RegisterProjectsHandler(service.Server(), handler.NewProjects(service))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
