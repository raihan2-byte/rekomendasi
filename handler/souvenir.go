package handler

import (
	"log"
	"mygram/helper"
	"mygram/souvenir"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type souvenirHandler struct {
	souvenirService souvenir.Service
}

func NewSouvenirHandler(souvenirService souvenir.Service) *souvenirHandler {
	return &souvenirHandler{souvenirService}
}

func (h *souvenirHandler) GetSouvenir(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("id"))

	sosmed, err := h.souvenirService.GetSouvenir(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to Get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosmed)
	c.JSON(http.StatusOK, response)
}

func (h *souvenirHandler) GetRate(c *gin.Context) {
	var input souvenir.InputSouvenir

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	log.Println(input)
	hotel, err := h.souvenirService.GetByInput(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to Get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, hotel)
	c.JSON(http.StatusOK, response)
}