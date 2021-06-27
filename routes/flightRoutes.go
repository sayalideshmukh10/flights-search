package routes

import (
	"flights-server/handlers"

	"github.com/gin-gonic/gin"
)

//Routes
func Init(o, r *gin.RouterGroup) {
	o.GET("/flights", handlers.FetchAllFlightsHandler())
	r.POST("/searchflights", handlers.FetchSourceAndDestinationFlightsHandler())
}
