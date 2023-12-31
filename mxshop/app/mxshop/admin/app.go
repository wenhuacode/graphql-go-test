package admin

import (
	"github.com/hashicorp/consul/api"
	"mxshop/app/pkg/options"
	"mxshop/app/user/srv/config"
	gapp "mxshop/gmicro/app"
	"mxshop/pkg/app"
	"mxshop/pkg/log"

	"mxshop/gmicro/registry"
	"mxshop/gmicro/registry/consul"
)

func NewApp(basename string) *app.App {
	cfg := config.New()
	appl := app.NewApp("user",
		"mxshop",
		app.WithOptions(cfg),
		app.WithRunFunc(run(cfg)),
	)
	return appl
}

func NewRegistrar(registry *options.RegistryOptions) registry.Registrar {
	c := api.DefaultConfig()
	c.Address = registry.Address
	c.Scheme = registry.Scheme
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}

func NewUserApp(cfg *config.Config) (*gapp.App, error) {
	//初始化log
	log.Init(cfg.Log)
	defer log.Flush()

	//服务注册
	register := NewRegistrar(cfg.Registry)

	//生成rpc服务
	rpcServer, err := NewUserHTTPServer(cfg)
	if err != nil {
		return nil, err
	}

	return gapp.New(
		gapp.WithName(cfg.Server.Name),
		gapp.WithRestServer(rpcServer),
		gapp.WithRegistrar(register),
	), nil
}

func run(cfg *config.Config) app.RunFunc {
	return func(baseName string) error {
		userApp, err := NewUserApp(cfg)
		if err != nil {
			return err
		}

		//启动
		if err := userApp.Run(); err != nil {
			log.Errorf("run user app error: %s", err)
			return err
		}
		return nil
	}
}
