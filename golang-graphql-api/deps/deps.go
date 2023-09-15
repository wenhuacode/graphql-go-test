package deps

import (
	"blacheapi/config"
	"blacheapi/dal"
	"blacheapi/logger"
	"blacheapi/migration"
	"blacheapi/services/redis"

	"github.com/pkg/errors"
)

type Dependencies struct {
	// DAL - Database Access Layer
	DAL *dal.DAL

	// RAL - Redis Access Layer
	Redis *redis.RAL
}

// New initiliazes the project dependencies based on the config
func New(cfg *config.Config) (*Dependencies, error) {

	// 连接数据库
	ndal, err := dal.New(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set up data access layer")
	}
	logger.GetLogger().Info("[DB]: OK")

	// 数据库表更新
	createErr := migration.CreateTables()
	if createErr != nil {
		logger.GetLogger().Sugar().Fatalf("[DB]: unable to create new tables: %v", createErr.Error())
		return nil, createErr
	}
	logger.GetLogger().Info("[DB]: tables updated")

	// 数据库迁移
	migrateErr := migration.Migrate(cfg)
	if migrateErr != nil {
		logger.GetLogger().Sugar().Fatalf("[DB]: unable to migrate schema: %v", migrateErr.Error())
		return nil, migrateErr
	}
	logger.GetLogger().Info("[DB]: migration competed")

	// 连接redis
	redisS, err := redis.New(cfg)
	if err != nil {
		logger.GetLogger().Sugar().Fatalf("[REDIS]: unable to connect to redis: %v", err.Error())
		return nil, err
	}
	logger.GetLogger().Info("[REDIS]: OK")

	deps := &Dependencies{
		DAL:   ndal,
		Redis: redisS,
	}
	return deps, nil
}
