package mysql

import "database/sql"

func createVideosTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS videos (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name CHAR(32) NOT NULL,
		type INT UNSIGNED REFERENCES video_types(id),
		created_at DATE NOT NULL,
		INDEX(name)
	) DEFAULT CHARACTER SET utf8;`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func createVideoTypesTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS video_types (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name CHAR(32) NOT NULL,
		created_at DATE NOT NULL,
		INDEX(name)
	) DEFAULT CHARACTER SET utf8;`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertVideo(db *sql.DB, name string, typeId int64) (int64, error) {
	var sql = "INSERT INTO videos (name, type, created_at) value (?, ?, NOW())"
	r, err := db.Exec(sql, name, typeId)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectVideoId(db *sql.DB, name string) (int64, error) {
	var sql = "SELECT id FROM videos WHERE name = ? LIMIT 1"
	var id int64
	err := db.QueryRow(sql, name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func InsertVideoType(db *sql.DB, name string) (int64, error) {
	var sql = "INSERT INTO video_types (name, created_at) value (?, NOW())"
	r, err := db.Exec(sql, name)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectVideoType(db *sql.DB, name string) (int64, error) {
	var sql = "SELECT id FROM video_types WHERE name = ?"
	var id int64
	err := db.QueryRow(sql, name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
