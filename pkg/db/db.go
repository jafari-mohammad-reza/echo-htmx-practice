package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

func SeedInit() error {
	db, err = sql.Open("sqlite3", "file:identifier.sqlite")
	if err != nil {
		return err
	}
	err = createTables()
	if err != nil {
		return err
	}
	func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	return nil
}

func createTables() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Blog(
	   id int NOT NULL,
	   title varchar(255) NOT NUll,
	   primary key (id)
	)`)

	if err != nil {
		return err
	}
	return nil
}
