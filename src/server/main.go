package main

import (
	"os"
	"runtime"
	"./router"
)

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}


func main() {
	//registry router
	r := router.Registry()
	//execute service
	router.Run(r)
}
