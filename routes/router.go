package routes

import (
	"login-backend/configuration"
	//"login-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(serviceName string, host *configuration.RouterConfig, tokenUrl string) {

	router := gin.Default()

	//cors configuration

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Location", "Access-Control-Allow-Origin", "authorization", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//router.Use(middleware.Middleware())

	rt := new(Rest)
	rt.Auth.TokenUrl = tokenUrl

	router.GET("/gocloak", rt.VerifyToken)
	router.GET("/redirect", TestRedirect)
	router.GET("/token", rt.Token)
	router.Run(host.Router)

}
