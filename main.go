package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"github.com/labstack/echo/v4"
)

type Cat struct{
	Name string `json:"name"`
	Type string `json:"type"`
}

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

func addCats(c echo.Context) error{
	cat := Cat{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body) // b : body
	if err != nil{
		log.Printf("Failed reading the request body : %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat) // Memory address of cat
	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat!")
}

func main() {
	fmt.Println("It's a good beginning for me.")

	e := echo.New() // create an instance of Echo

	e.GET("/", begin) 		 // Endpoint1
	e.GET("/cats/:data", getCats) // Endpoint2

	e.POST("/cats", addCats)

	// Routes
	e.Start(":8000") // Start Server
}

// Get information from the URL