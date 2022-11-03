package Handler

import (
	"context"
	"fmt"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"net/http"
	"time"

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
		Logic.Log.Error(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	Logic.Log.Info("Добавлена новая запись")
	return c.JSON(http.StatusCreated, newPerson)
}

func GetPersons(c echo.Context) error {
	persons, err := Logic.Read()
	if err != nil {
		Logic.Log.Error(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	Logic.Log.Info("Выведены все записи")
	return c.JSON(http.StatusOK, persons)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	persons, err := Logic.ReadOne(id)
	if err != nil {
		Logic.Log.Error(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	Logic.Log.Infof("Выведена Запись с id = %s", id)
	return c.JSON(http.StatusOK, persons)
}

func DeleteById(c echo.Context) error {
	id := c.Param("id")
	err := Logic.Delete(id)
	if err != nil {
		Logic.Log.Error(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	Logic.Log.Infof("Запись с id = %s  успешно удалена", id)
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
		Logic.Log.Error(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}
	Logic.Log.Infof("Запись с id = %s  успешно обновлена", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("Запись с id = %s  успешно обновлена", id))
}

//Middleware
func ConnectDB(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		errorCh := make(chan error)
		err := Repository.OpenTable()
		go check(ctx, errorCh)
		//time.Sleep(3 * time.Second) //Используем для имитации долгого подключения к БД
		errorCh <- err
		return next(c)
	}
}

func check(ctx context.Context, errorCh chan error) {
	for {
		select {
		case <-ctx.Done():
			Logic.Log.Fatalf("Timed out: %v", ctx.Err())
			return
		case err := <-errorCh:
			if err != nil {
				Logic.Log.Errorf("Возникла ошибка... %v", err)
				return
			}
		default:
			fmt.Println("Trying to connect to database...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
