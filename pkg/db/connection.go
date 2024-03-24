package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:root@tcp(localhost:3306)/database_Go"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	fmt.Println("database configured")
}