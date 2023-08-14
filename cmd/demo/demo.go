package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/rshulabs/micro-frame/internal/demo"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	demo.NewApp("demo-server").Run()
}
