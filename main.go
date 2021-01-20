package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Book struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Year  int    `json: "year"`
}

func main() {
	fmt.Println("welcome to the server")

	//configuration of the server
	serverConfig()
}

func serverConfig() {
	//port of the application
	httpAddr := ":5000"

	fmt.Println("server is running in the port", httpAddr)

	e := echo.New()
	//func of the routes
	e.GET("/", helloHandler)
	e.GET("/books/:id", listBookByIDHandler)

	e.POST("/newbook", newBookHandler)

	//start the application
	e.Start(httpAddr)

}

// HANDLERS

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func listBookByIDHandler(c echo.Context) error {
	//id := c.QueryParam("id")

	id := c.Param("id")

	return c.String(http.StatusOK, fmt.Sprintf("Found this id: %s", id))
}

func newBookHandler(c echo.Context) error {

	newBook := Book{}
	//defer c.Request().Body.Close()

	body, err := ioutil.ReadAll(c.Request().Body())
	if err != nil {
		log.Fatal("Failed to read the request body: %s", err)
		return c.String(http.StatusInternalServerError, "não deu")
	}

	err = json.Unmarshal(body, &newBook)
	if err != nil {
		log.Fatal("Failed on Unmarshal: %s", err)
		return c.String(http.StatusInternalServerError, "não rolou")
	}

	log.Printf("thi is your book: %v", newBook)

	return c.String(http.StatusCreated, "created")
}
