package handler

import (
	"mygram/helper"
	"mygram/touris"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tourisHandler struct {
	tourisService touris.Service
}

func NewTourisHandler(tourisService touris.Service) *tourisHandler {
	return &tourisHandler{tourisService}
}

func (h *tourisHandler) GetTouris(c *gin.Context){
	ID, _ := strconv.Atoi(c.Query("id"))

	sosmed, err := h.tourisService.GetTouris(ID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to Get Touris")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosmed)
	c.JSON(http.StatusOK, response)
}

func (h *tourisHandler) GetRate(c *gin.Context){
	var input touris.InputGetTouris

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	hotel, err := h.tourisService.GetRating(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to Get touris")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, hotel)
	c.JSON(http.StatusOK, response)


}
