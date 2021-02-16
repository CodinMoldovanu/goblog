package main

import (
	"github.com/codinmoldovanu/goblog/route"

	"github.com/codinmoldovanu/goblog/db"
)

func main() {
	db.Setup()
	// cert := os.Getenv("certFile_goblog")
	// key := os.Getenv("keyFile_goblog")

	// db.LoginUser("user", "notMyPwd
	// e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.Logger.Fatal(e.Start(":1323"))
	router := route.Init()
	router.Logger.Fatal(router.Start(":80"))
	// router.Logger.Fatal(router.StartTLS(":80", cert, key))
	// router.Logger.Fatal(router.StartAutoTLS(":1333"))
}
