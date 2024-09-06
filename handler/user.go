package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		
		errors := helper.FormatValidationErrors(err)
		errorMessage := gin.H{"errors": errors}
		
		response := helper.APIResponse("Invalid input", http.StatusUnprocessableEntity, "success", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Failed to registered", http.StatusBadRequest, "success", err.Error())
		c.JSON(http.StatusBadRequest, response)		
		return
	}

	formatter := user.FormatterUser(newUser, "token123")
	
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}