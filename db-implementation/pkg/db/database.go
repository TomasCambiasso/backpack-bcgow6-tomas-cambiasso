package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	engine = gin.Default()

	return engine, db
}


func InitDb() (*sql.DB, error) {
	configDB := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "storage",
	}
	txdb.Register("txdb", "mysql", configDB.FormatDSN())
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
