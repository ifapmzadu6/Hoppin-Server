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
		os INT UNSIGNED REFERENCES user_oss(id),
		created_at DATETIME NOT NULL,
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
		device CHAR(32) NOT NULL,
		INDEX(device)
	) DEFAULT CHARACTER SET utf8;`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func createUserOSTable(db *sql.DB) error {
	var sql = `
	CREATE TABLE IF NOT EXISTS user_oss (
		id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
		os CHAR(32) NOT NULL,
		INDEX(os)
	) DEFAULT CHARACTER SET utf8;
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(db *sql.DB, password string, udid int64, uoid int64) (int64, error) {
	var sql = "INSERT INTO users (password, device, os, created_at) value (?, ?, ?, NOW())"
	r, err := db.Exec(sql, password, udid, uoid)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectUser(db *sql.DB, id int64) (string, error) {
	var sql = "SELECT password FROM users WHERE id = ? LIMIT 1"
	var password string
	err := db.QueryRow(sql, id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func ValidateUserByStr(db *sql.DB, id string, password string) (int64, error) {
	var cid, err = strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	nid := int64(cid)
	return nid, ValidateUser(db, nid, password)
}

func ValidateUser(db *sql.DB, id int64, password string) error {
	var sql = "SELECT id FROM users WHERE id = ? AND password = ? LIMIT 1"
	var tid int
	err := db.QueryRow(sql, id, password).Scan(&tid)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserDevice(db *sql.DB, device string) (int64, error) {
	var sql = "INSERT INTO user_devices (device) value (?)"
	r, err := db.Exec(sql, device)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectUserDevice(db *sql.DB, device string) (int64, error) {
	var sql = "SELECT id FROM user_devices WHERE device = ?"
	var id int64
	err := db.QueryRow(sql, device).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func InsertUserOS(db *sql.DB, os string) (int64, error) {
	var sql = "INSERT INTO user_oss (os) value (?)"
	r, err := db.Exec(sql, os)
	if err != nil {
		return -1, err
	}
	return r.LastInsertId()
}

func SelectUserOS(db *sql.DB, os string) (int64, error) {
	var sql = "SELECT id FROM user_oss WHERE os = ?"
	var id int64
	err := db.QueryRow(sql, os).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
