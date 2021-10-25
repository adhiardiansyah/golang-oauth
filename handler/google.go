package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-oauth/auth"
	"golang-oauth/helper"
	"golang-oauth/user"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type googleHandler struct {
	googleOauthConfig *oauth2.Config
	authService       auth.Service
	userService       user.Service
}

func NewGoogleHandler(googleOauthConfig *oauth2.Config, authService auth.Service, userService user.Service) *googleHandler {
	return &googleHandler{googleOauthConfig: googleOauthConfig, authService: authService, userService: userService}
}

func (h *googleHandler) Login(c *gin.Context) {
	url := h.googleOauthConfig.AuthCodeURL("login")
	c.Redirect(http.StatusMovedPermanently, url)
}

func (h *googleHandler) CallBack(c *gin.Context) {
	if c.Request.FormValue("state") != "login" {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	token, err := h.googleOauthConfig.Exchange(context.Background(), c.Request.FormValue("code"))
	if err != nil {
		response := helper.APIResponse("Could not get token", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		response := helper.APIResponse("Could not create request", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

	input := user.RegisterUserInput{}

	fmt.Println(string(content))

	json.Unmarshal([]byte(string(content)), &input)

	IsEmailExist, err := h.userService.IsEmailExist(input.Email)
	if err != nil {
		errorMessage := gin.H{"error": "Server error"}

		response := helper.APIResponse("Email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !IsEmailExist {
		_, err := h.userService.RegisterUser(input)
		if err != nil {
			response := helper.APIResponse("Register account failed", "error", http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	loggedInUser, err := h.userService.SignIn(user.SignInInput{
		Email: input.Email,
	})
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	jwtToken, err := h.authService.GenerateToken(loggedInUser.Uuid)
	if err != nil {
		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, jwtToken)

	response := helper.APIResponse("Successfully logged in", "success", http.StatusOK, formatter)

	c.JSON(http.StatusOK, response)
}
