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

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Tiger struct {
	Name string `json:"name"`
	Type string `json:"Type"`
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

func addCat(c echo.Context) error{
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

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil{
		log.Printf("Failed processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog!")
}

func addTiger(c echo.Context) error {
	tiger := Tiger{}

	err := c.Bind(&tiger)  // Bind 3rd Party Methods
	if err != nil {
		log.Printf("Failed processing addTiger request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your tiger: %#v", tiger)
	return c.String(http.StatusOK, "We got your Tiger! Here! Take it.")
}

func main() {
	fmt.Println("It's a good beginning for me.")

	e := echo.New() // create an instance of Echo

	e.GET("/", begin) 		 // Endpoint1
	e.GET("/cats/:id", getCats) // Endpoint2

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/tigers", addTiger)

	// Routes
	e.Start(":8000") // Start Server
}

// Get information from the URL