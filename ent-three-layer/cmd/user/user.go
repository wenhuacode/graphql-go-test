package main

import (
	"ent-three-layer/app/user/srv"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	srv.NewApp("user-server").Run()
}
