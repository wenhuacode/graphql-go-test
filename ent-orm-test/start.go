package start

import (
	"context"
	"ent-orm-test/ent/migrate"
	"log"
	"time"

	"database/sql"
	"ent-orm-test/ent"
	_ "ent-orm-test/ent/runtime"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open() (*ent.Client, error) {
	dsn := "root:Mybabysjk888888.,@tcp(127.0.0.1:3306)/ent-orm-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)

	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	//Run the auto migration tool.
	if err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil

}
