package handlers

import (
	"flights-server/helpers"
	"flights-server/models"
	"flights-server/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := models.Login{}
		c.Bind(&requestBody)

		isValid, err := services.ValidateCredentials(requestBody)

		log.Print(isValid, err)

		if isValid {

			token, err := helpers.GenerateToken(requestBody.Username, requestBody.Password, 24*time.Hour)
			if err != nil {
				log.Print("error while generating token:", err)
				// return err
			}

			c.Header("Authorization", token)
			c.JSON(http.StatusOK, token)
		} else {

			result := models.Result{
				Status:  1,
				Message: "unauthorized user",
			}

			c.JSON(http.StatusNotFound, result)
		}

	}

}
