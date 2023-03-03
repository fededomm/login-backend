package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"login-backend/configuration"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	GrantType   string `form:"grant-type" json:"grant-type"`
	AuthCode    string `form:"authcode" json:"authcode"`
	RedirectUrl string `form:"redirect-url" json:"redirect-url"`
}

func (q *QueryParam) Token(c *gin.Context) {
	QueryParam := new(QueryParam)
	conf := new(configuration.Param)

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(c.ShouldBindQuery(&QueryParam))
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", conf.TokenUrl, (buf))

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
