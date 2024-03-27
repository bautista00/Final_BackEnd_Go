package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {

	dataSource := fmt.Sprintf("root:1234@tcp(localhost:3306)/database_Go")
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	fmt.Println("database configured")
}
