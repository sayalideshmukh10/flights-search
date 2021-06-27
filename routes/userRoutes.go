package routes

import (
	"flights-server/handlers"

	"github.com/gin-gonic/gin"
)

func InitUser(o, r *gin.RouterGroup) {

	o.POST("/login", handlers.Login())

}
