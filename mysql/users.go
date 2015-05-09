package mysql

import (
	"database/sql"
	"strconv"
)

func createUsersTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		password CHAR(32) NOT NULL,
		device INT UNSIGNED REFERENCES user_devices(id),
		INDEX(password)
	) DEFAULT CHARACTER SET utf8;`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func createUserDevicesTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS user_devices (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name CHAR(32) NOT NULL,
		os CHAR(32) NOT NULL
	) DEFAULT CHARACTER SET utf8;`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(password string, device int64) (int64, error) {
	return insertUser(sharedDB, password, device)
}

func insertUser(db *sql.DB, password string, device int64) (int64, error) {
	var sql = "INSERT INTO users (password, device) value (?, ?)"
	r, err := db.Exec(sql, password, device)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectUser(id int64) (string, error) {
	return selectUser(sharedDB, id)
}

func selectUser(db *sql.DB, id int64) (string, error) {
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
	return validateUser(sharedDB, int64(cid), password)
}

func ValidateUser(id int64, password string) error {
	return validateUser(sharedDB, id, password)
}

func validateUser(db *sql.DB, id int64, password string) error {
	var sql = "SELECT id FROM users WHERE id = ? AND password = ? LIMIT 1"
	var tid int
	err := db.QueryRow(sql, id, password).Scan(&tid)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserDevice(name string, os string) (int64, error) {
	return insertUserDevice(sharedDB, name, os)
}

func insertUserDevice(db *sql.DB, name string, os string) (int64, error) {
	var sql = "INSERT INTO user_devices (name, os) value (?, ?)"
	r, err := db.Exec(sql, name, os)
	if err != nil {
		return -1, nil
	}
	return r.LastInsertId()
}

func SelectUserDevice(name string, os string) (int64, error) {
	return selectUserDevice(sharedDB, name, os)
}

func selectUserDevice(db *sql.DB, name string, os string) (int64, error) {
	var sql = "SELECT id FROM user_devices WHERE name = ? AND os = ?"
	var id int64
	err := db.QueryRow(sql, name, os).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
