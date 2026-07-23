package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

var DB *sql.DB

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
	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}
	return nil
}

// TODO: Нужно сделать создание таблички
