package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"kratos-demo/app/order/service/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(c *conf.Registry) registry.Registrar {
	cfg := consulAPI.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme
	cli, err := consulAPI.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return consul.New(cli, consul.WithHealthCheck(false))
}
