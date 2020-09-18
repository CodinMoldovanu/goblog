package db

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/codinmoldovanu/goblog/models"

	"github.com/codinmoldovanu/goblog/helpers"

	"golang.org/x/crypto/bcrypt"
)

var salt = ""

//CreateUser in database and return a struct with it
func CreateUser(username string, password string, firstname string, lastname string) models.User {
	db := createConn()
	defer db.Close()
	insert, err := db.Prepare(`INSERT INTO users(username, password, firstname, lastname, created) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(insert)
	hpw, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	creationTime := helpers.GetTimestampTz()
	user := models.User{}
	insert.QueryRow(username, hpw, firstname, lastname, creationTime).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	return user
}

//CheckLogin checks user and pwd and returns boolean
func CheckLogin(username string, password string) bool {
	db := createConn()
	defer db.Close()
	var dbpwd string
	loggedUser := models.User{}
	var result = false
	db.QueryRow(`SELECT password FROM users WHERE username = $1`, username).Scan(&dbpwd)
	if checkPasswordHash(password, dbpwd) {
		db.QueryRow(`SELECT * FROM users WHERE username = $1`, username).Scan(&loggedUser.ID, &loggedUser.Username, &loggedUser.Password, &loggedUser.FirstName, &loggedUser.LastName)
		result = true
	}
	fmt.Println(result)
	return result
}

//CheckLoginValidator checks user and pwd and returns boolean
func CheckLoginValidator(username string, password string, c echo.Context) (bool, error) {
	db := createConn()
	defer db.Close()
	var dbpwd string
	loggedUser := models.User{}
	db.QueryRow(`SELECT password FROM users WHERE username = $1`, username).Scan(&dbpwd)
	if checkPasswordHash(password, dbpwd) {
		db.QueryRow(`SELECT * FROM users WHERE username = $1`, username).Scan(&loggedUser.ID, &loggedUser.Username, &loggedUser.Password, &loggedUser.FirstName, &loggedUser.LastName)
		return true, nil
	}
	// fmt.Println(result)
	return false, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUser as a struct
func GetUser(username string) models.User {
	db := createConn()
	defer db.Close()
	user := models.User{}
	err := db.QueryRow(`SELECT ID, username, password, firstname, lastname FROM users WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

//Login with token
func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	if CheckLogin(username, password) != true {
		return errors.New("Unauthorized")
	}
	user := GetUser(username)
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.FirstName
	claims["username"] = user.Username
	claims["id"] = user.ID
	expires := time.Now().Add(time.Hour * 72).Unix()
	claims["exp"] = expires

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   t,
		Expires: time.Unix(expires, 0),
	})
	print("GOT HEREEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
	print(t)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
