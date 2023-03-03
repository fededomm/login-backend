package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	GrantType   string `form:"grant-type"`
	AuthCode    string `form:"authcode"`
	RedirectUrl string `form:"redirect-url"`
}

func (q *QueryParam) Token(c *gin.Context) {
	var QueryParam QueryParam
	if err := c.ShouldBindQuery(&QueryParam); err != nil {
		log.Fatal(err)
	} else {
		log.Println(QueryParam)
	}
	request, err := http.NewRequest("POST", "http://localhost:8443/realms/my-realm/protocol/openid-connect/token?", strings.NewReader("ciao"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Basic bXktY2xpZW50OlRiY1ZkQ0RudXUya3JxZ044eXYzdEdkckFDSWZhV1Qw")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		fmt.Printf("client: error making http request: %s\n", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		fmt.Print("Error decode reader")
	}
	log.Printf("%v", resp)

	c.JSON(200, gin.H{"body": b})
}
func TestAuthCode(c *gin.Context) {
	fmt.Printf("%v", c.Params)

}
