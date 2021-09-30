package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@/golang") // conexion a base de datos local sin password

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping() // Verifica  errores en la conección
	if err != nil {
		panic(err.Error())
	}

	return db
}
