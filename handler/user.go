package handler

import (
	"golang-oauth/auth"
	"golang-oauth/helper"
	"golang-oauth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.SignInInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.SignIn(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.Uuid)
	if err != nil {
		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIResponse("Successfully logged in", "success", http.StatusOK, formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	foundUser, err := h.userService.GetUserByUuid(currentUser.Uuid)
	if err != nil {
		response := helper.APIResponse("Get user failed", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(foundUser, "")

	response := helper.APIResponse("User data", "success", http.StatusOK, formatter)

	c.JSON(http.StatusOK, response)
}
