package rest

import (
	"login-backend/configuration"
	"login-backend/routes"

	"github.com/gin-gonic/gin"
)

func Init(serviceName string, host *configuration.RouterConfig) {

	router := gin.Default()
	//router.Use(middle.Middleware())

	rt := new(routes.QueryParam)

	router.GET("/test", routes.TestAuthCode)
	router.GET("/token", rt.Token)
	router.Run(host.Router)
}
