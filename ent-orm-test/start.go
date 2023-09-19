package main

import (
	"context"
	"database/sql"
	"ent-orm-test/ent"
	"ent-orm-test/ent/migrate"
	"ent-orm-test/ent/user"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName("wenhua22").
		SetAge(30).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("wenhua22")).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user was queried: ", u)
	return u, nil
}

func Open() (*ent.Client, error) {
	dsn := "root:Mybabysjk888888.,@tcp(127.0.0.1:3306)/ent-orm-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	client, err := Open()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//CreateUser(context.Background(), client)
	queryUser, err := QueryUser(context.Background(), client)
	fmt.Println(queryUser[0].Name)
}
