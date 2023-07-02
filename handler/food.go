package handler

import (
	"log"
	food "mygram/Food"
	"mygram/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	foodService food.Service
}

func NewFoodHandler(foodService food.Service) *foodHandler {
	return &foodHandler{foodService}
}

func (h *foodHandler) GetFoods(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("id"))

	sosmed, err := h.foodService.GetFoods(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosmed)
	c.JSON(http.StatusOK, response)
}

func (h *foodHandler) GetRate(c *gin.Context) {
	var input food.InputFood

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	log.Println(input)
	hotel, err := h.foodService.GetByInput(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, hotel)
	c.JSON(http.StatusOK, response)
}
