package mysql

import "database/sql"

type Users struct {
	user_id  int
	password string
}

func createUsersTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		password CHAR(32) NOT NULL
	) DEFAULT CHARACTER SET utf8;`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(password string) error {
	return insertUser(sharedDB, password)
}

func insertUser(db *sql.DB, password string) (id, error) {
	var sql = "INSERT INTO users (password) value (?)"

	r, err := db.Exec(sql, password)
	if err != nil {
		return nil, err
	}
	return r.LastInsertId()
}
