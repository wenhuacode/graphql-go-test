// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package srv

import (
	"mxshop/app/pkg/options"
	"mxshop/app/user/srv/controller/user"
	"mxshop/app/user/srv/data/v1/db"
	"mxshop/app/user/srv/service/v1"
	"mxshop/gmicro/app"
	"mxshop/pkg/log"
)

// Injectors from wire.go:

func initApp(nacosOptions *options.NacosOptions, logOptions *log.Options, serverOptions *options.ServerOptions, registryOptions *options.RegistryOptions, telemetryOptions *options.TelemetryOptions, mySQLOptions *options.MySQLOptions) (*app.App, error) {
	registrar := NewRegistrar(registryOptions)
	gormDB, err := db.GetDBFactoryOr(mySQLOptions)
	if err != nil {
		return nil, err
	}
	userStore := db.NewUsers(gormDB)
	userSrv := v1.NewUserService(userStore)
	userServer := user.NewUserServer(userSrv)
	nacosDataSource, err := NewNacosDataSource(nacosOptions)
	if err != nil {
		return nil, err
	}
	server, err := NewUserRPCServer(telemetryOptions, serverOptions, userServer, nacosDataSource)
	if err != nil {
		return nil, err
	}
	appApp, err := NewUserApp(logOptions, registrar, serverOptions, server)
	if err != nil {
		return nil, err
	}
	return appApp, nil
}
