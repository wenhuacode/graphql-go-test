package main

import (
	"math/rand"
	"mxshop/app/mxshop/admin"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	admin.NewApp("admin-server").Run()
}
