package db

import (
	"database/sql"
	"ent-three-layer/app/pkg/code"
	"ent-three-layer/app/pkg/options"
	v1 "ent-three-layer/app/user/srv/internal/data/v1"
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
	errors2 "ent-three-layer/pkg/errors"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	dbFactory v1.DataFactory
	once      sync.Once
)

type mysqlFactory struct {
	db *ent.Client
}

func (mf *mysqlFactory) User() v1.UserStore {
	return newUser(mf)
}

var _ v1.DataFactory = (*mysqlFactory)(nil)

func GetDBFactoryOr(mysqlOpts *options.MySQLOptions) (v1.DataFactory, error) {
	if mysqlOpts == nil && dbFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlOpts.Username,
			mysqlOpts.Password,
			mysqlOpts.Host,
			mysqlOpts.Port,
			mysqlOpts.Database)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return
		}
		db.SetMaxIdleConns(mysqlOpts.MaxOpenConnections)
		db.SetMaxOpenConns(mysqlOpts.MaxIdleConnections)
		db.SetConnMaxLifetime(mysqlOpts.MaxConnectionLifetime)

		drv := entsql.OpenDB("mysql", db)
		client := ent.NewClient(ent.Driver(drv))

		dbFactory = &mysqlFactory{
			db: client,
		}
	})

	if dbFactory == nil || err != nil {
		return nil, errors2.WithCode(code.ErrConnectDB, "failed to get mysql store factory")
	}
	return dbFactory, nil
}
