package models

import "encoding/xml"

//User Model
type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
	Settings  xml.CharData
}
