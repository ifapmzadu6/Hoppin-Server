package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	var query string
	if os.Getenv("DEBUG") == "1" {
		query = "root:@/" + os.Getenv("DB_NAME")
	} else {
		query = "root:" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_ADDR") + ":3306)/" + os.Getenv("DB_NAME")
	}
	db, err := sql.Open("mysql", query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if _, err := db.Exec("USE hoppin"); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}

func CreateDataBase(db *sql.DB) error {
	var sql = "CREATE DATABASE IF NOT EXISTS hoppin DEFAULT CHARACTER SET utf8;"
	_, err := db.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) error {
	if err := createActionTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createActionTypeTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createUsersTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createUserDevicesTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createUserOSTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createVideosTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := createVideoTypesTable(db); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
