package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Println("This is a Good Start", e)

	e.GET("/",func (c echo.Context) error{
		return c.String(http.StatusOK, "SUCCESS!")
	})

	e.Start(":8000") // Start Server
}
