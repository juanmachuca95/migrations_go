package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnectionDatabase2() *sql.DB {
	user := os.Getenv("USUARIO")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORTDATABASE")
	dbname := os.Getenv("DATABASE2")

	log.Println(user, password, host, port, dbname)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)) // conexion a base de datos local sin password

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping() // Verifica  errores en la conección
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
