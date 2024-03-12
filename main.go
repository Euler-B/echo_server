package main

import (

	"github.com/Euler-B/echo_server/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Euler-B/echo_server/handlers"
)

func main() {
	//echo instance
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routes
	e.GET("/", handlers.Hello)

	// start database 
	database.ConectDB()

	//start server 
	e.Logger.Fatal(e.Start(":8080"))
}
