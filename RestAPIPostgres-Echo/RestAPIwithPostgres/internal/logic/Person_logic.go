package Logic

import (
	"fmt"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"strconv"
)

//(`INSERT INTO "person" ("person_email", "person_phone", "person_firstName", "person_lastName") VALUES ($1, $2,$3,$4)`, p.Email, p.Phone, p.FirstName, p.LastName)
func Create(p Model.Person) error {
	sess := Repository.Connection.NewSession(nil)
	_, err := sess.InsertInto("person").Columns("person_email", "person_phone", "person_firstName", "person_lastName").Values(p.Email, p.Phone, p.FirstName, p.LastName).Exec()
	if err != nil {
		return err
	}
	return nil
}

//row, err := Repository.Connection.Query(`SELECT * FROM "person" WHERE "person_id" = $1`, person_id)
func ReadOne(id string) ([]Model.Person, error) {
	sess := Repository.Connection.NewSession(nil)
	person_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("Error: неверно введён параметр id: %v", err)
	}
	personInfo := []Model.Person{}
	err = sess.Select("*").From("person").Where("person_id = ?", person_id).LoadOne(&personInfo)
	if err != nil {
		return nil, err
	}
	return personInfo, nil
}

func Read() ([]Model.Person, error) {
	sess := Repository.Connection.NewSession(nil)
	personInfo := []Model.Person{}
	_, err := sess.Select("*").From("person").OrderBy("person_id").Load(&personInfo)
	if err != nil {
		return nil, err
	}
/*	row, err := Repository.Connection.Query(`SELECT * FROM "person" ORDER BY "person_id"`)
	if err != nil {
		return nil, err
	}*/
	//var personInfo = []Model.Person{}
	/*for personInfo.Next() {
		var p Model.Person
		err := row.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}*/
	return personInfo, nil
}

func Update(p Model.Person, id string) error {
	sess := Repository.Connection.NewSession(nil)
	if err := dataExist(id); err != nil {
		return err
	}
	_, err := sess.Update("person").SetMap(map[string]interface{}{"person_email": p.Email, "person_phone": p.Phone, "person_firstName": p.FirstName, "person_lastName": p.LastName}).Where("person_id = ?", id).Exec()
	if err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	sess := Repository.Connection.NewSession(nil)
	if err := dataExist(id); err != nil {
		return err
	}
	_, err := sess.DeleteFrom("person").Where("person_id = ?", id).Exec()
	if err != nil {
		return err
	}
	return nil
}

func dataExist(id string) error {
	persons, err := ReadOne(id)
	if err != nil {
		return err
	}
	if len(persons) == 0 {
		return fmt.Errorf("Error: записи с id = %s не существует", id)
	}
	return nil
}
