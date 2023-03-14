//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/lmxia/highway/internal/app/api"
	"github.com/lmxia/highway/internal/app/dao"
	"github.com/lmxia/highway/internal/app/module/adapter"
	"github.com/lmxia/highway/internal/app/router"
	"github.com/lmxia/highway/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
