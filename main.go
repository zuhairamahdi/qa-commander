package main

import (
	"fmt"
	"qacommander/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Hello)
	fmt.Println("Server is running at port 1323")
	e.Logger.Fatal(e.Start(":1323"))
}
