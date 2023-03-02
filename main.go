package main

import (
	_ "embed"
	"fmt"
	"log"

	"login-backend/rest"
)

//go:embed banner.txt
var banner []byte

func main() {
	
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(banner))
	rest.Init(cfg.App.ServiceName, &cfg.App.GinRouter)
}
