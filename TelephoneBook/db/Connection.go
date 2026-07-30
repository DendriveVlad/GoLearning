package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB
var TableName string = "contacts"

func InitDB() error {
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	myAppDir := filepath.Join(appDataDir, "TelephoneBook")
	err = os.MkdirAll(myAppDir, os.ModePerm)
	if err != nil {
		return err
	}
	dbPath := filepath.Join(myAppDir, "storage.db")
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}
	if err := initTable(); err != nil {
		return err
	}
	return nil
}

func initTable() error {
	var err error
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS " + TableName + " (phone TEXT PRIMARY KEY, name TEXT)")
	if err != nil {
		return err
	}
	return nil
}
