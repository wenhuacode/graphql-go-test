package srv

import (
	v12 "ent-three-layer/api/user/v1"
	"ent-three-layer/app/user/srv/config"
	"ent-three-layer/app/user/srv/internal/controller/user"
	"ent-three-layer/app/user/srv/internal/data/v1/db"
	v1 "ent-three-layer/app/user/srv/internal/service/v1"
	"ent-three-layer/gmicro/core/trace"
	"ent-three-layer/gmicro/server/rpcserver"
	"ent-three-layer/pkg/log"
	"fmt"
)

func NewUserRPCServer(cfg *config.Config) (*rpcserver.Server, error) {
	//初始化open-telemetry的exporter
	trace.InitAgent(trace.Options{
		cfg.Telemetry.Name,
		cfg.Telemetry.Endpoint,
		cfg.Telemetry.Sampler,
		cfg.Telemetry.Batcher,
	})

	//有点繁琐，wire， ioc-golang
	dataFactory, err := db.GetDBFactoryOr(cfg.MySQLOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	srvFactory := v1.NewService(dataFactory)
	userServer := user.NewUserService(srvFactory)
	rpcAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	grpcServer := rpcserver.NewServer(rpcserver.WithAddress(rpcAddr))

	v12.RegisterUserServer(grpcServer.Server, userServer)

	return grpcServer, nil
}
