package Repository

import (
	"fmt"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

const (
	//host     = "database" //используем в случае подключения к базе данных из контейнера
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "4650"
	dbname   = "postgres"
	sslmode  = "disable"
)

var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

var Connection *dbr.Connection

func OpenTable() error {
	var err error
	Connection, err = dbr.Open("postgres", connectionString, nil)
	if err != nil {
		return err
	}
	sess := Connection.NewSession(nil)
	table, err := sess.Query(`CREATE TABLE IF NOT EXISTS person
	(
		"person_id" serial PRIMARY KEY,
		"person_email" character varying(32),
		"person_phone" character varying(32),
		"person_firstName" character varying(32),
		"person_lastName" character varying(32)
	)`)
	if err != nil {
		return err
	}
	defer table.Close()

	return nil
}