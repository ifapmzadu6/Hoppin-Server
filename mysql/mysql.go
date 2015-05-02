package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:@/hoppin")
	if err != nil {
		panic(err)
	}
	return db
}

func Close(db *sql.DB) {
	db.Close()
}
