//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	orderclient "kratos-demo/internal/client/order"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/registrar"
	orderserver "kratos-demo/internal/server/order"
	orderservice "kratos-demo/internal/service/order"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registrar, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		orderservice.ProviderSet,
		orderserver.ProviderSet,
		orderclient.ProviderSet,
		registrar.ProviderSet,
		newApp))
}
