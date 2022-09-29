// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-demo/app/order/service/internal/biz"
	"kratos-demo/app/order/service/internal/conf"
	"kratos-demo/app/order/service/internal/data"
	"kratos-demo/app/order/service/internal/server"
	"kratos-demo/app/order/service/internal/service"
)

func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
