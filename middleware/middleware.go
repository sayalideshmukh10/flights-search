package middleware

import (
	"flights-server/helpers"
	"flights-server/models"
	"flights-server/routes"
	"flights-server/services"
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(g *gin.Engine) {

	g.Use(cors.Default()) //CORS Request

	//open
	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())

	//restricted
	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())

	routes.Init(o, r)
	routes.InitUser(o, r)

}

func OpenRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func RestrictedRequestMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		login, err := helpers.GetLoginFromToken(c)

		if err != nil {
			log.Print("Token not available:", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}

		if strings.Trim(token, "") == "" {
			log.Print("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}

		user := models.Login{}
		user.Username = login.Username
		user.Password = login.Password

		isValid, usererr := services.ValidateCredentials(user)
		if usererr != nil || !isValid {
			log.Print("Failed to validate user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		}

		c.Next()

	}
}
