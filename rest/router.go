package rest

import (
	"login-backend/configuration"
	middle "login-backend/middleware"
	"login-backend/routes"

	"github.com/gin-gonic/gin"
)

func Init(serviceName string, host *configuration.RouterConfig) {
	router := gin.New()
	router.Use(middle.Middleware())

	router.GET("/test", routes.TestAuthCode)
	router.GET("/token", routes.Token)
	router.Run(host.Router)
}
