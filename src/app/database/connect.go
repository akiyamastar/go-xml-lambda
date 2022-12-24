package database

import (
    "database/sql"
    "fmt"
    "os"
    "log"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

type Job struct {
    OriginalUrl string
}

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

func GetUrls() []string {
    var dbUrlList []string
    db := Connect()
    dbErr := db.Ping()
    if dbErr != nil {
        fmt.Println("データベース接続失敗")
        return dbUrlList
    } else {
        fmt.Println("データベース接続成功")
    }
    selectSql := os.Getenv("DB_SELECT_SQL")
    rows, err := db.Query(selectSql)
    defer rows.Close()
    if err != nil {
        // log.Fatal(err)
        fmt.Println("データベースエラー")
    }
    // fmt.Println(rows)
    for rows.Next() {
        u := &Job{}
        // if err := rows.Scan(&u.ID, &u.Name, &u.Profile, &u.Created, &u.Updated); err != nil {
        if err := rows.Scan(&u.OriginalUrl); err != nil {
            fmt.Println("urlエラー")
        }
        dbUrlList = append(dbUrlList, u.OriginalUrl)
    }
    fmt.Printf("database url count: %d\n", len(dbUrlList))
    return dbUrlList
}

func ReinsertNewUrls(urlList []string) {
    db := Connect()
    dbErr := db.Ping()
    if dbErr != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }

    fmt.Println("Delete New List Urls!")
    deleteSql := os.Getenv("DB_DELETE_LIST_SQL")
    if _, err := db.Exec(deleteSql); err != nil {
        log.Fatal("Exec error: ", err)
        return
    }
    
    fmt.Println("Insert New List Urls!")
    insert := os.Getenv("DB_INSERT_LIST_SQL")

    values := make([]interface{}, 0, len(urlList))
    for index, value := range urlList {
        insert += fmt.Sprintf(`(%v, ?),`, index)
        // insert += fmt.Sprintf(`( ? ),`)
        values = append(values, value)
    }
    insert = insert[:len(insert)-1]
    stmt, err := db.Prepare(insert)
    if err != nil {
        log.Fatal("Prepare error: ", err)
    }
    if _, err := stmt.Exec(values...); err != nil {
        log.Fatal("Exec error: ", err)
    }
}