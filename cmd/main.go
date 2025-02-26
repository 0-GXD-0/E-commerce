package main

import (
	"E-commerce/conf"
	"E-commerce/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
