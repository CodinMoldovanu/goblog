package handlers

import (
	"log"

	"github.com/codinmoldovanu/goblog/auth"
	"github.com/codinmoldovanu/goblog/db"
	"github.com/labstack/echo"
)

//Salt thingy temporarily exposed
var Salt = "secretsalt"

//AddUser into db
func AddUser(c echo.Context) error {
	username := c.FormValue("email")
	password := c.FormValue("password")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")

	if username == "test" {
		print("YES IT'S TEST DUDE -------------------------------------------------------------------------------------------------")
	}

	if username == "" {
		print("IT'S MOTHERFUCKING EMPTY --------------------------------__________________________________________________")
	}
	log.Print("username " + username)
	user := db.CreateUser(username, password, firstname, lastname)
	err := auth.CreateToken(user.Username, user.Password, c)
	return err
}
