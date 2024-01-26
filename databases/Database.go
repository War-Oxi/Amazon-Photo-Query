package database

import (
	"database/sql"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)
func ConnectDB() *sql.DB{
	err := godotenv.Load("./env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	dbName := os.Getenv("DB_NAME")

	// 데이터베이스 연결
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", dbUser, dbPassword, dbEndpoint, dbName))
	// fmt.Printf("%T\n", db) //sql.DB
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// 데이터베이스 연결 테스트
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	return db
}