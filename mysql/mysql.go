package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var sharedDB *sql.DB

func Open() error {
	db, err := openDataBase()
	if err != nil {
		return err
	}
	sharedDB = db
	return nil
}

func openDataBase() (*sql.DB, error) {
	var query string
	if os.Getenv("DEBUG") == "1" {
		query = "root:@/" + os.Getenv("DB_NAME")
	} else {
		query = "root:password@" + os.Getenv("MYSQL_PORT_3306_TCP_PROTO") + "(" + os.Getenv("MYSQL_PORT_3306_TCP_ADDR") + ":" + os.Getenv("MYSQL_PORT_3306_TCP_PORT") + ")/" + os.Getenv("DB_NAME")
	}
	db, err := sql.Open("mysql", query)
	if err != nil {
		return nil, err
	}

	if err := createDataBase(db); err != nil {
		return nil, err
	}
	if _, err := db.Exec("USE hoppin"); err != nil {
		return nil, err
	}

	if err := createActionTable(db); err != nil {
		return nil, err
	}
	if err := createActionTypeTable(db); err != nil {
		return nil, err
	}
	if err := createUsersTable(db); err != nil {
		return nil, err
	}
	if err := createUserDevicesTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func Close() error {
	return closeDataBase(sharedDB)
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
