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
	CreateDataBase(db)
	db.Exec("USE hoppin")

	CreateActionTable(db)

	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}

/*
type Action struct {
	Id      int
	VideoId string
	Type    ActionType
	Time    uint16
	Start   uint16
	End     uint16
}
type ActionType struct {
    Id     int
    String string
}
*/

func CreateDataBase(db *sql.DB) {
	var sql = "CREATE DATABASE IF NOT EXISTS hoppin DEFAULT CHARACTER SET utf8;"

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func CreateActionTable(db *sql.DB) {
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
		panic(err)
	}
}

func CreateActionTypeTable(db *sql.DB) {
	var sql = `
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func InsertAction(a *Action) {
	//var sql = "INSERT INTO video (id, videoId, type, time, start, end) value (?, ?, ?, ?)"

	//db.Exec(sql, Action.VideoId, Action.Type, Action.Time, Action.Start, Action.End)
}
