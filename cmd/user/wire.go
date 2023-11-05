//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	userbiz "kratos-demo/internal/biz/user"
	"kratos-demo/internal/conf"
	userdata "kratos-demo/internal/data/user"
	"kratos-demo/internal/registrar"
	userserver "kratos-demo/internal/server/user"
	userservice "kratos-demo/internal/service/user"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registrar, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		userserver.ProviderSet,
		userdata.ProviderSet,
		userbiz.ProviderSet,
		userservice.ProviderSet,
		registrar.ProviderSet,
		newApp))
}
