// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bizserver/internal/biz"
	"bizserver/internal/conf"
	"bizserver/internal/data"
	"bizserver/internal/server"
	"bizserver/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	helloRepo := data.NewHelloRepo(dataData, logger)
	bizServerUsercase := biz.NewBizServerUsercase(helloRepo, logger)
	bizServerService := service.NewBizServerService(bizServerUsercase)
	grpcServer := server.NewGRPCServer(confServer, bizServerService, logger)
	httpServer := server.NewHTTPServer(confServer, bizServerService, logger)
	registrar := data.NewEtcdRegistrar(confData)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
