package handler

import (
	"log"
	"mygram/cafe"
	"mygram/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cafeHandler struct {
	cafeService cafe.Service
}

func NewCafeHandler(cafeService cafe.Service) *cafeHandler {
	return &cafeHandler{cafeService}
}

func (h *cafeHandler) GetCafe(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("id"))

	sosmed, err := h.cafeService.GetCafes(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosmed)
	c.JSON(http.StatusOK, response)
}

func (h *cafeHandler) GetRate(c *gin.Context){
	var input cafe.InputCafe

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	log.Println(input)
	hotel, err := h.cafeService.GetByInput(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, hotel)
	c.JSON(http.StatusOK, response)
}


