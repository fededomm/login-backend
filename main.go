package main

import (
	_ "embed"
	"fmt"
	"log"

	"login-backend/routes"
)

//go:embed banner.txt
var banner []byte

func main() {

	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(banner))
	routes.Init(cfg.App.ServiceName, &cfg.App.GinRouter, cfg.App.Auth.TokenUrl)
}
