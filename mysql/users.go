package mysql

import (
	"database/sql"
	"strconv"
)

type User struct {
	Id       int
	Password string
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

func InsertUser(password string) (int64, error) {
	return insertUser(sharedDB, password)
}

func insertUser(db *sql.DB, password string) (int64, error) {
	var sql = "INSERT INTO users (password) value (?)"

	r, err := db.Exec(sql, password)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func SelectUser(id int) (string, error) {
	return selectUser(sharedDB, id)
}

func selectUser(db *sql.DB, id int) (string, error) {
	var sql = "SELECT password FROM users WHERE id = ? LIMIT 1"

	var password string
	err := db.QueryRow(sql, id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func ValidateUserByStr(id string, password string) error {
	var cid, err = strconv.Atoi(id)
	if err != nil {
		return err
	}
	return validateUser(sharedDB, cid, password)
}

func ValidateUser(id int, password string) error {
	return validateUser(sharedDB, id, password)
}

func validateUser(db *sql.DB, id int, password string) error {
	var sql = "SELECT id FROM users WHERE id = ? AND password = ?"

	var tid int
	err := db.QueryRow(sql, id, password).Scan(&tid)
	if err != nil {
		return err
	}
	return nil
}
