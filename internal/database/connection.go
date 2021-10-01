package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func MySQLConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot loading environment's variables")
	}

	user := os.Getenv("USUARIO")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORTDATABASE")
	dbname := os.Getenv("DATABASE")

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

func main() {
	db := MySQLConnection()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
