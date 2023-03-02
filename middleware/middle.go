package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		latency := time.Since(t)
		// Prima della Chiamata
		log.Print("Inizio la Chiamata")

		c.Next()

		// Dopo la Chiamata
		log.Println("Termino la Chiamata")
		log.Println(latency)
	}

}
