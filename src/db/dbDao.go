package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConn() *sql.DB{
	db, err := sql.Open("mysql", "test:123456@tcp(172.16.6.85:3306)/jobmanager?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return nil
	}
	return db
}
func CloseConn(conn *sql.DB){
	conn.Close()
}