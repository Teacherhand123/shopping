package main

import (
	"fmt"
	"shopping/conf"
	"shopping/routers"
)

func main() {
	fmt.Println("Hello world")
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
