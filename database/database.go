package database

import (
	"database/sql"
)

func InitDb() *sql.DB {
	connectionString := "root@tcp(localhost:3306)/nortwind"

	databaseConnection, error := sql.Open("mysql", connectionString)

	defer databaseConnection.Close() //se va a cerrar cuando se dejede usar

	if error != nil {
		panic(error.Error()) //manejo de errores
	}

	return databaseConnection
}
