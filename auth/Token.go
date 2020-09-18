package auth

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/codinmoldovanu/goblog/helpers"

	"github.com/codinmoldovanu/goblog/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(username string, password string, c echo.Context) error {
	validate, err := db.CheckLoginValidator(username, password, c)
	log.Print("DATA IS : " + username + " " + password)
	user := db.GetUser(username)
	if err != nil {
		return err
	}
	log.Print("Validation: " + strconv.FormatBool(validate))

	if validate {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
		// print(claims["username"])
		// print("-------------------------------------------------------")
		// claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		helpers.SetUserToken(t, user.ID)

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"access": "Denied",
	})
}
