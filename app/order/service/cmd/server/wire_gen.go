// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/app/order/service/internal/biz"
	"kratos-demo/app/order/service/internal/conf"
	"kratos-demo/app/order/service/internal/data"
	"kratos-demo/app/order/service/internal/server"
	"kratos-demo/app/order/service/internal/service"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUseCase := biz.NewOrderUseCase(orderRepo, logger)
	orderService := service.NewOrderService(orderUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, orderService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
