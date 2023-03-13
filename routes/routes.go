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
	"github.com/go-resty/resty/v2"
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
	param.Add("code", "343e419a-fbd6-410f-81e2-b5f9f27cb2ac.126eca88-872c-4a31-bf35-180b8a761895.1cc4bda5-5da5-41b8-89e4-5fc86f05f95f")
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

	b, err := io.ReadAll(resp.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		fmt.Print("Error decode reader")
	}
	log.Printf("%v", resp)
	log.Println(b)
	c.JSON(200, gin.H{"body": b})
}

func TestRedirect(c *gin.Context) {

	c.Header("Location", "http://localhost:8443/realms/my-realm/protocol/openid-connect/auth?response_type=code&client_id=my-client&redirect_uri=http://127.0.0.1:8085/")
	c.Status(302)

	authToken, _ := c.GetQuery("code")
	if authToken == "" {
		c.JSON(404, gin.H{"message": "code not found"})
		c.Abort()
	}
	client := resty.New()
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(authToken).
		Get("/api/v1/token")

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	log.Println(authToken)
}
func Login(authCode string) {
	var c gin.Context
	authToken, _ := c.GetQuery("code")
	if authToken == "" {
		c.JSON(404, gin.H{"message": "code not found"})
		c.Abort()
	}
	client := resty.New()
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(authToken).
		Get("/api/v1/token")

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	log.Println(authToken)
}
