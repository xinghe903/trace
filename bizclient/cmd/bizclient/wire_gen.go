// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bizclient/internal/biz"
	"bizclient/internal/conf"
	"bizclient/internal/data"
	"bizclient/internal/server"
	"bizclient/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewEtcdDiscovery(confData)
	dataData, cleanup, err := data.NewData(confData, logger, discovery)
	if err != nil {
		return nil, nil, err
	}
	bizServerRepo := data.NewBizServerRepo(dataData, logger)
	bizClientUsecase := biz.NewBizClientUsecase(bizServerRepo, logger)
	bizClientService := service.NewBizClientService(bizClientUsecase)
	grpcServer := server.NewGRPCServer(confServer, bizClientService, logger)
	httpServer := server.NewHTTPServer(confServer, bizClientService, logger)
	registrar := data.NewEtcdRegistrar(confData)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
