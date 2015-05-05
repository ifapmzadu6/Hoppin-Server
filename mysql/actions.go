package mysql

import "database/sql"

type Action struct {
	Id      int
	VideoId string
	Type    ActionType
	Time    int
	Start   int
	End     int
}

type ActionType struct {
	Id     int
	String string
}

func createActionTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS actions (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		video_id CHAR(32) NOT NULL,
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
	return insertAction(sharedDB, a)
}

func insertAction(db *sql.DB, a Action) error {
	var sql = "INSERT INTO actions (video_id, type, time, start, end) value (?, ?, ?, ?, ?)"

	_, err := db.Exec(sql, a.VideoId, a.Type.Id, a.Time, a.Start, a.End)
	if err != nil {
		return err
	}
	return nil
}
