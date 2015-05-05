package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Open() error {
	tdb, err := openDataBase()
	if err != nil {
		return err
	}
	db = tdb
	return nil
}

func openDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@"+os.Getenv("MYSQL_PORT_3306_TCP_PROTO")+"("+os.Getenv("MYSQL_PORT_3306_TCP_ADDR")+":"+os.Getenv("MYSQL_PORT_3306_TCP_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		return nil, err
	}

	err = createDataBase(db)
	if err != nil {
		return nil, err
	}
	db.Exec("USE hoppin")

	err = createActionTable(db)
	if err != nil {
		return nil, err
	}

	err = createActionTypeTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close() error {
	return closeDataBase(db)
}

func closeDataBase(db *sql.DB) error {
	return db.Close()
}

func createDataBase(db *sql.DB) error {
	var sql = "CREATE DATABASE IF NOT EXISTS hoppin DEFAULT CHARACTER SET utf8;"

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
