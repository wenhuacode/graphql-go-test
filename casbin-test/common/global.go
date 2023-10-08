package common

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	adapter "github.com/casbin/ent-adapter"
	"github.com/casbin/ent-adapter/ent"
	"strings"
)

var (
	CasBinDB *casbin.Enforcer
)

func InitCasBinDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "Mybabysjk888888.,", "127.0.0.1", "3306", "casbin")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))
	a, err := adapter.NewAdapterWithClient(client)
	if err != nil {
		panic(err)
	}
	CasBin, err := casbin.NewEnforcer("model.conf", a)
	if err != nil {
		panic(err)
	}
	// url路径匹配
	CasBin.AddFunction("ParamsMatch", ParamsMatchFunc)
	// 加载策略
	err = CasBin.LoadPolicy()
	if err != nil {
		panic(err)
	}
	CasBinDB = CasBin
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

// 注册func到casbin
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}
