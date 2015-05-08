package mysql

import "database/sql"

type Action struct {
	Id      int
	VideoId string
	Type    ActionType
	Time    int
	Start   int
	End     int
	UserId  int
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
		user_id INT UNSIGNED NOT NULL REFERENCES users(id),
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

func InsertAction(a Action, userId int) error {
	return insertAction(sharedDB, a, userId)
}

func insertAction(db *sql.DB, a Action, userId int) error {
	var sql = "INSERT INTO actions (video_id, type, time, start, end, user_id) value (?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(sql, a.VideoId, a.Type.Id, a.Time, a.Start, a.End, userId)
	if err != nil {
		return err
	}
	return nil
}
