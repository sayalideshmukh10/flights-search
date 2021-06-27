package main

import (
	"flights-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	Server()
}

func Server() {

	router := gin.Default()

	router.GET("/", CheckStatus())

	middleware.InitMiddleware(router)

	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}

	//ListenAndServe starts an HTTP server with a given address and handler.
	s.ListenAndServe()

}

func CheckStatus() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is Running")
	}
}
