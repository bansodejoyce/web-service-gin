package user

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// albums slice to seed record album data.
var user = User{Username: "joycebansode", Password: "QWERT12345"}

func Authenticate(c *gin.Context) {
	var u User
	if err := c.BindJSON(&u); err != nil {
		return
	}
	fmt.Println("user -------", u)
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(username string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //add in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(token string) bool {

	// token, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return key, nil
	// })
	return true
}
