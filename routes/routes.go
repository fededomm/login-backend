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
	param.Add("code", "1c665384-3dda-4a78-97b8-ebfbf1cddad7.e1272ca4-5271-4493-ae0d-baf4730e651f.1cc4bda5-5da5-41b8-89e4-5fc86f05f95f")
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
func (r *Rest) TestAuthCode(c *gin.Context) {
	fmt.Printf("%v", c.Params)
}
