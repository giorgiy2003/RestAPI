package main

import (
	"log"
	Handler "myapp/internal/handlers"
	Repository "myapp/internal/repository"

	"github.com/labstack/echo"
)

func main() {
	if err := Repository.OpenTable(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	router := echo.New()
	router.GET("/person", Handler.GetPersons)
	router.GET("/person/:id", Handler.GetById)
	router.POST("/person", Handler.PostPerson)
	router.DELETE("/person/:id", Handler.DeleteById)
	router.PUT("/person/:id", Handler.UpdatePersonById)
	router.Logger.Fatal(router.Start(":8080"))
}
