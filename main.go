package main

import (
	"fmt"
	"log"

	"login-backend/rest"
)

var banner []byte

//go:embed banner.txt
func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(banner))
	rest.Init(cfg.App.ServiceName, &cfg.App.GinRouter)
}
