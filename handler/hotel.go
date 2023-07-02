package handler

import (
	"log"
	"mygram/helper"
	"mygram/hotel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type hotelHandler struct {
	hotelService hotel.Service
}

func NewHotelHandler(hotelService hotel.Service) *hotelHandler {
	return &hotelHandler{hotelService}
}

func (h *hotelHandler) GetHotel(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("id"))

	sosmed, err := h.hotelService.GetHotels(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "eror to Get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosmed)
	c.JSON(http.StatusOK, response)
}

func (h *hotelHandler) GetRate(c *gin.Context){
	var input hotel.InputHotel

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	log.Println(input)
	hotel, err := h.hotelService.GetByInput(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "eror to Get")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (hotel))
	c.JSON(http.StatusOK, response)


}
