package helpers

import (
	"errors"
	"flights-server/models"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var GlobalJWTKey string

func init() {
	GlobalJWTKey = "FLIGHT"
}

type jwtCustomClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string, expirationTime time.Duration) (string, error) {

	claims := jwtCustomClaim{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	//Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Generate encode token and send it as response
	t, err := token.SignedString([]byte(GlobalJWTKey))

	if err != nil {
		log.Print("Error during token generation", err)
	}

	return t, nil
}

//GetLoginFromToken login object from JWT Token
func GetLoginFromToken(c *gin.Context) (models.Login, error) {

	login := models.Login{}

	decodedToken, err := DecodeToken(c.GetHeader("Authorization"), GlobalJWTKey)

	if err != nil {
		return login, errors.New("GetLoginFromToken - unable to decode token")
	}

	//login name is compulsory fields
	if decodedToken["username"] == nil || decodedToken["username"] == "" {
		return login, errors.New("GetLoginFromToken - login id not found")

	}
	login.Username = decodedToken["username"].(string)
	login.Password = decodedToken["password"].(string)
	return login, nil

}

func DecodeToken(tokenFromRequest, jwtKey string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil

	})

	if err != nil {
		log.Print("error while parsing JWT Token: ", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error while getting claims")
	}
	return claims, nil

}
