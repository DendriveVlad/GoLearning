package main

import (
	"TelephoneBook/db"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			derr := DB.Ping()
			if derr != nil {
				panic(derr)
			}
		}
	}(db.DB)

	err = hub()
	if err != nil {
		fmt.Println(err)
	}
}
