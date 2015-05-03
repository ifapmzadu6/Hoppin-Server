package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/")
	if err != nil {
		return nil, err
	}
	err = CreateDataBase(db)
	if err != nil {
		return nil, err
	}
	db.Exec("USE hoppin")

	err = CreateActionTable(db)
	if err != nil {
		return nil, err
	}

	err = CreateActionTypeTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}

func CreateDataBase(db *sql.DB) error {
	var sql = "CREATE DATABASE IF NOT EXISTS hoppin DEFAULT CHARACTER SET utf8;"

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateActionTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS action (
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

func CreateActionTypeTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS action_type (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		string CHAR(32) UNIQUE NOT NULL
	) DEFAULT CHARACTER SET utf8;`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertAction(db *sql.DB, a Action) error {
	var sql = "INSERT INTO action (video_id, type, time, start, end) value (?, ?, ?, ?, ?)"

	_, err := db.Exec(sql, a.VideoId, a.Type.Id, a.Time, a.Start, a.End)
	if err != nil {
		return err
	}
	return nil
}
