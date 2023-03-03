package routes

import (
	"fmt"
	"io"
	"log"
	config "login-backend/configuration"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	GrantType   string `form:"grant_type" json:"grant-type"`
	AuthCode    string `form:"authcode" json:"authcode"`
	RedirectUrl string `form:"redirect_url" json:"redirect-url"`
}


func (q *QueryParam) Token(c *gin.Context) {

	config := new(config.Param)

	param := url.Values{}
	param.Add("grant_type", "authorization_code")
	param.Add("code", "b107d0ac-14c5-42de-a837-b82a1567bcac.c8170fef-271b-455c-89b9-6879be075b76.1cc4bda5-5da5-41b8-89e4-5fc86f05f95f")
	param.Add("redirect_uri", "http://127.0.0.1:8085/test")

	request, err := http.NewRequest("POST", config.TokenUrl, strings.NewReader(param.Encode()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Basic bXktY2xpZW50OlRiY1ZkQ0RudXUya3JxZ044eXYzdEdkckFDSWZhV1Qw")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		fmt.Printf("client: error making http request: %s\n", err)
		return
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
