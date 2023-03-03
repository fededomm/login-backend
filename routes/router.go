package routes

import (
	"login-backend/configuration"

	"github.com/gin-gonic/gin"
)

func Init(serviceName string, host *configuration.RouterConfig, tokenUrl string) {

	router := gin.Default()
	//router.Use(middle.Middleware())
	rt := new(Rest)
	rt.Auth.TokenUrl = tokenUrl
	
	router.GET("/test", TestAuthCode)
	router.GET("/token", rt.Token)
	router.Run(host.Router)
}
