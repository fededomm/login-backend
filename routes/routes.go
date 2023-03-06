package routes

import (
	"fmt"
	"io"
	"log"
	"login-backend/configuration"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type QueryParam struct {
	AuthCode    string `form:"authcode" json:"authcode"`
	RedirectUrl string `form:"redirect_url" json:"redirect-url"`
}
type Rest struct {
	Auth configuration.Param `mapstructure:"param"`
}

func (r *Rest) Token(c *gin.Context) {

	param := url.Values{}
	param.Add("grant_type", "authorization_code")
	param.Add("code", "")
	param.Add("redirect_uri", "http://127.0.0.1:8085/test")

	request, err := http.NewRequest("POST", r.Auth.TokenUrl, strings.NewReader(param.Encode()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatal(err)
		return
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

func (r *Rest) TestRedirect(c *gin.Context) {
	c.Header("Location", "http://localhost:8443/realms/my-realm/protocol/openid-connect/auth?response_type=code&client_id=my-client&redirect_uri=http://127.0.0.1:8085/x")
	c.Header("Access-Control-Allow-Origin", "http://localhost:8443/")
	c.Status(302)
}
