package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	AuthCode := "code=" + ""
	body := "grant_type=authorization_code&" + AuthCode + "&redirect_uri=http://127.0.0.1:8085/x"
	request, err := http.NewRequest("POST", "http://localhost:8443/realms/my-realm/protocol/openid-connect/token?", strings.NewReader(body))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Basic bXktY2xpZW50OlJsUlZjNmcwOHFTZXBvNUIzZ2dxM1Q5bkF1QzRFYm1o")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error decode reader")
	}
	log.Printf("%v", resp)

	c.JSON(200, gin.H{"body": b})
}
func TestAuthCode(c *gin.Context) {
	fmt.Printf("%v", c.Params)

}
