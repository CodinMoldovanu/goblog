package handlers

import (
	"fmt"
	"log"

	"github.com/codinmoldovanu/goblog/db"
	"github.com/codinmoldovanu/goblog/models"
	"github.com/labstack/echo"
)

//UpdatePostHandler is a handler to pass data to the DB Content Updater
func UpdatePostHandler(c echo.Context) {
	post := models.Post{}
	maps := echo.Map{}
	// json.Unmarshal()
	err := c.Bind(&post)

	if err != nil {
		log.Print("Something went wrong in the UpdatePostHandler")
		log.Print(err)
	}
	fmt.Println(maps)
	db.UpdatePost(post)
}

//UploadFormImages handles exactly what it says it does.
func UploadFormImages(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Print(err)
		return err
	}
	files := form.File["images"]
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			log.Print(err)
			return err
		}
		defer src.Close()

	}
	return err
}
