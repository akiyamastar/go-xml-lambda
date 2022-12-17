package database

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
    err := godotenv.Load()
    if err != nil {
        fmt.Println(err.Error())
    }

    user := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    database_name := os.Getenv("DB_DATABASE")

	dbconf := user + ":" 
    dbconf += password + "@tcp(" 
    dbconf += host + ":" + port + ")/" 
    dbconf += database_name + "?charset=utf8mb4"

    db, err := sql.Open("mysql", dbconf)
    if err != nil {
        fmt.Println(err.Error())
    }
    return db
}