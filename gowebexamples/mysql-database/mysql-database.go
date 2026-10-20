package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func StartUpDatabaseTables(db *sql.DB) sql.Result {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return res
}

func GetDatabaseCon(user, pass, addr, dbName string) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=True", user, pass, addr, dbName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func TruncateTable(db *sql.DB, table string) {
	query := fmt.Sprintf("TRUNCATE TABLE `%s`", table)
	db.Exec(query)
}

type User struct {
	id        int64
	username  string
	password  string
	createdAt time.Time
}

func (u *User) Insert(db *sql.DB) (int64, error) {
	createdTime := time.Now()
	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
	result, err := db.Exec(query, u.username, u.password, createdTime)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	u.id = id
	u.createdAt = createdTime
	return id, nil
}

func (u *User) Get(db *sql.DB, id int64) error {
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&u.id, &u.username, &u.password, &u.createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\nUSERNAME: %s\nPASSWORD: %s\nCREATED: %s\n", u.id, u.username, u.password, u.createdAt.String())
}

// IT WOULD BE BETTER TO DO A REPOSITORY AND ETC... BUT I WANTED TO EXPLORE SOME STUFF WITH STRUCTS
func main() {
	db := GetDatabaseCon("MY_USER", "MY_PASSWORD", "127.0.0.1:3306", "MY_DATABASE")
	defer db.Close()
	defer TruncateTable(db, "users")

	StartUpDatabaseTables(db)
	u := User{username: "name", password: "pass"}
	uID, _ := u.Insert(db)
	fmt.Println(uID)
	fmt.Println(u.id == uID)
	var uGets User
	_ = uGets.Get(db, uID)
	fmt.Println("----------USER----------")
	fmt.Println(uGets)
}
