package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando el archivo .env")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error al conectar con la base de datos: %v", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("No se pudo hacer ping a la base de datos: %v", err))
	}

	fmt.Println("Conexión a la base de datos establecida correctamente.")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}
