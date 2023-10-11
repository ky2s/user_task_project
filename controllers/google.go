package controllers

import (
	"fmt"
	"net/http"
	"user_task_project/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleController interface {
	GoogleLogin(c *gin.Context)
}

type googleController struct {
	userMod models.UserModels
}

func NewGoogleController(googleModels models.UserModels) GoogleController {
	return &googleController{
		userMod: googleModels,
	}
}

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "525486579481-3bm754i69vhkmmbnl6iiafq4nfcivck2.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}

func (ctr *googleController) GoogleLogin(c *gin.Context) {

	gConfig := SetupConfig()
	url := gConfig.AuthCodeURL("randomstate")

	fmt.Println(url)
	c.Redirect(http.StatusSeeOther, url)

	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  true,
	// 	"message": "Success created data",
	// })
	// return
}
