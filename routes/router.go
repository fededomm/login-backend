package routes

import (
	"login-backend/configuration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(serviceName string, host *configuration.RouterConfig, tokenUrl string) {

	router := gin.Default()

	//cors configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5500/html", "http://localhost:8443"}
	config.AllowHeaders = []string{"Origin", "Location"}
	config.AllowMethods = []string{"GET", "POST", "PUT"}
	router.Use(cors.New(config))

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:5500/html", "http://localhost:8443"},
	//	AllowMethods:     []string{"GET", "POST", "PUT"},
	//	AllowHeaders:     []string{"Origin", "Location"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "http://localhost:5500/html"
	//	},
	//}))
	rt := new(Rest)
	rt.Auth.TokenUrl = tokenUrl

	router.GET("/redirect", rt.TestRedirect)
	router.GET("/token", rt.Token)
	router.Run(host.Router)
}
