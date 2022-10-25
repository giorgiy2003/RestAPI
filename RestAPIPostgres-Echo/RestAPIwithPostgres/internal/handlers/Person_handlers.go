package Handler

import (
	"fmt"
	"log"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	"net/http"

	"github.com/labstack/echo"
)

func PostPerson(c echo.Context) error {
	var newPerson Model.Person
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err := Logic.Create(newPerson)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Println("Добавлена запись", newPerson)
	return c.JSON(http.StatusCreated, newPerson)
}

func GetPersons(c echo.Context) error {
	persons, err := Logic.Read()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	return c.JSON(http.StatusOK, persons)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	persons, err := Logic.ReadOne(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Println(persons)
	return c.JSON(http.StatusOK, persons)
}

func DeleteById(c echo.Context) error {
	id := c.Param("id")
	err := Logic.Delete(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Printf("Запись с id = %s  успешно удалена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно удалена", id))
}

func UpdatePersonById(c echo.Context) error {
	var newPerson Model.Person
	id := c.Param("id")
	newPerson.Email = c.FormValue("email")
	newPerson.Phone = c.FormValue("phone")
	newPerson.FirstName = c.FormValue("firstName")
	newPerson.LastName = c.FormValue("lastName")
	err := Logic.Update(newPerson, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	log.Printf("Запись с id = %s  успешно обновлена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно обновлена", id))
}
