package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func begin(c echo.Context) error {
	return c.String(http.StatusOK, "SUCCESS!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data") // parameter

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is : %s\nand his type is : %s,\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name" : catName,
			"type" : catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error" : "You need to let us know if you want JSON or STRING data.",
	}) // Neither of those condition's match than return this
}

func main() {
	fmt.Println("It's a good beginning for me.")

	e := echo.New() // create an instance of Echo

	e.GET("/", begin) 		 // Endpoint1
	e.GET("/cats/:data", getCats) // Endpoint2

	// Routes
	e.Start(":8000") // Start Server
}

// Get information from the URL