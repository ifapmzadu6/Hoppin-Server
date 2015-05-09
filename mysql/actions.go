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
	Id   int64
	Name string
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
		name CHAR(32) NOT NULL
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

func InsertActionType(name string) (int64, error) {
	return insertActionType(sharedDB, name)
}

func insertActionType(db *sql.DB, name string) (int64, error) {
	var sql = "INSERT INTO action_types (name) value (?)"
	r, err := db.Exec(sql, name)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func SelectActionTypeById(id int) (string, error) {
	return selectActionTypeById(sharedDB, id)
}

func selectActionTypeById(db *sql.DB, id int) (string, error) {
	var sql = "SELECT name FROM action_types WHERE id = ? LIMIT 1"
	var name string
	err := db.QueryRow(sql, id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func SelectActionTypeByName(name string) (int, error) {
	return selectActionTypeByName(sharedDB, name)
}

func selectActionTypeByName(db *sql.DB, name string) (int, error) {
	var sql = "SELECT name FROM action_types WHERE id = ? LIMIT 1"
	var id int
	err := db.QueryRow(sql, id).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
