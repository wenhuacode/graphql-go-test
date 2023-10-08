package main

import (
	"casbin-test/common"
	"casbin-test/router"
)

func main() {
	// 初始化DB
	common.InitCasBinDB()

	// 初始化路由
	Router := router.InitRouter()
	err := Router.Run(":8099")
	if err != nil {
		panic(err)
	}
}
