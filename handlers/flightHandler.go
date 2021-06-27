package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"flights-server/models"
	"flights-server/services"
)

func FetchAllFlightsHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		allFlights := services.FetchAllFlightsService()

		if len(allFlights.Flights) != 0 {
			c.JSON(http.StatusOK, allFlights.Flights)
		} else {
			result := models.Result{
				Status:  1,
				Message: "no data found",
			}
			c.JSON(http.StatusNotFound, result)
		}
	}
}

func FetchSourceAndDestinationFlightsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.UserInput{}
		c.Bind(&requestBody)
		flights := services.FetchSourceAndDestinationFlightsService(requestBody)

		if len(flights.Flights) != 0 {
			c.JSON(http.StatusOK, flights.Flights)
		} else {
			result := models.Result{
				Status:  1,
				Message: "flights not found",
			}

			c.JSON(http.StatusNotFound, result)
		}

	}
}
