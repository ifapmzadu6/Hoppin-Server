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

func createActionTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS actions (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		video_id CHAR(32) UNIQUE NOT NULL,
		type INT UNSIGNED NOT NULL REFERENCES action_type(id),
		time INT UNSIGNED NOT NULL,
		start INT UNSIGNED NOT NULL,
		end INT UNSIGNED NOT NULL,
		INDEX(video_id)
	) DEFAULT CHARACTER SET utf8;`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func createActionTypeTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS action_types (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		string CHAR(32) UNIQUE NOT NULL
	) DEFAULT CHARACTER SET utf8;`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertAction(a Action) error {
	return insertAction(db, a)
}

func insertAction(db *sql.DB, a Action) error {
	var sql = "INSERT INTO actions (video_id, type, time, start, end) value (?, ?, ?, ?, ?)"

	_, err := db.Exec(sql, a.VideoId, a.Type.Id, a.Time, a.Start, a.End)
	if err != nil {
		return err
	}
	return nil
}
