// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/zh1lu0/gocamp/internal/biz"
	"github.com/zh1lu0/gocamp/internal/conf"
	"github.com/zh1lu0/gocamp/internal/data"
	"github.com/zh1lu0/gocamp/internal/server"
	"github.com/zh1lu0/gocamp/internal/service"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
