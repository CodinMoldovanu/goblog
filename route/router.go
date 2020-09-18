package route

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/codinmoldovanu/goblog/handlers"
	"github.com/dgrijalva/jwt-go"

	"github.com/codinmoldovanu/goblog/db"

	"github.com/foolin/goview/supports/echoview"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var jwtKey = []byte("secret")

//Init routing
func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"127.0.0.1", "http://localhost:8080", "https://blog.codin.ro"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Restricted group
	r := e.Group("/backend")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	r.POST("/posts/update", func(c echo.Context) error {
		handlers.UpdatePostHandler(c)
		return c.JSON(http.StatusOK, "")
	})

	r.POST("/posts/media", func(c echo.Context) error {
		fmt.Print(c)
		return c.JSON(http.StatusOK, "")
	})

	e.Renderer = echoview.Default()
	e.Static("/static", "static")

	posts := e.Group("/post")
	posts.GET("/:post_URL", func(c echo.Context) error {
		return c.JSON(http.StatusOK, db.GrabPost(c.Param("post_URL")))
	})
	e.GET("/contact", func(c echo.Context) error {
		return c.String(http.StatusOK, "renderings/contact.html")
	})

	auth := e.Group("/auth")
	auth.POST("/create", func(c echo.Context) error {
		err := handlers.AddUser(c)
		return err
	})

	auth.POST("/login", func(c echo.Context) error {
		err := db.Login(c)
		return err
	})

	backend := e.Group("/barosan")
	backend.Use(middleware.BasicAuth(db.CheckLoginValidator))

	backend.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "All good my dude")
	})

	e.POST("/create_account", func(c echo.Context) error {
		handlers.AddUser(c)
		return c.String(http.StatusOK, "Ok")
	})
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "backend/login.html", echo.Map{
			"title": "Codin Moldovanu",
		})
	})
	//
	///..
	// e.GET("/create_account", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "backend/register.html", echo.Map{
	// 		"title": "Codin Moldovanu",
	// 	})
	// })

	// e.GET("/", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "renderings/landing.html", echo.Map{
	// 		"title": "Codin Moldovanu",
	// 		"posts": db.GrabPosts(),
	// 	})
	// })

	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, db.GrabPosts())
	// })

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, db.GrabPosts())
	})

	return e
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	// user, err := c.Cookie("token").(*jwt.Token)
	// if err != nil {
	// 	return err
	// }
	// cookie, err := c.Cookie("token")
	// print(cookie.Value)
	// print("-----------------------------------------")
	// if err != nil {
	// 	return err
	// }
	// tokenString := cookie.Value
	// // claims := &jwt.Claims{}

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })

	// if err != nil {
	// 	return err
	// }

	// if !token.Valid {
	// 	return c.String(http.StatusUnauthorized, "no.")
	// }

	// print(token.Claims)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome"+" "+name+".")
	// return c.String(http.StatusOK, token.Raw)
}

func verifyToken(c echo.Context) bool {
	user := c.Get("user").(*jwt.Token)
	token, err := jwt.Parse(user.Raw, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Fatal("String wrong")
	}
	fmt.Printf(strconv.FormatBool(token.Valid))
	return token.Valid
}
