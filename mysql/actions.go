package mysql

import "database/sql"

func createActionTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS actions (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		video_id INT UNSIGNED NOT NULL REFERENCES videos(id),
		type_id INT UNSIGNED NOT NULL REFERENCES action_type(id),
		user_id INT UNSIGNED NOT NULL REFERENCES users(id),
		time INT UNSIGNED NOT NULL,
		start INT UNSIGNED NOT NULL,
		end INT UNSIGNED NOT NULL,
		created_at DATETIME NOT NULL
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
		name CHAR(32) UNIQUE NOT NULL,
		created_at DATETIME NOT NULL,
		INDEX(name)
	) DEFAULT CHARACTER SET utf8;`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertAction(db *sql.DB, videoId int64, typeId int64, userId int64, time int, start int, end int) error {
	var sql = "INSERT INTO actions (video_id, type_id, user_id, time, start, end, created_at) value (?, ?, ?, ?, ?, ?, NOW())"
	_, err := db.Exec(sql, videoId, typeId, userId, time, start, end)
	if err != nil {
		return err
	}
	return nil
}

func InsertActionType(db *sql.DB, name string) (int64, error) {
	var sql = "INSERT INTO action_types (name, created_at) value (?, NOW())"
	r, err := db.Exec(sql, name)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func SelectActionTypeById(db *sql.DB, id int) (string, error) {
	var sql = "SELECT name FROM action_types WHERE id = ? LIMIT 1"
	var name string
	err := db.QueryRow(sql, id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func SelectActionTypeByName(db *sql.DB, name string) (int, error) {
	var sql = "SELECT id FROM action_types WHERE name = ? LIMIT 1"
	var id int
	err := db.QueryRow(sql, name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
