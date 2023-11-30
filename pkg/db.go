package pkg

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

func InitDatabase() error {
	db, err = sql.Open("sqlite3", "file:identifier.sqlite")
	if err != nil {
		return err
	}
	err = createTables()
	if err != nil {
		return err
	}
	err = createSeedData()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Blog(
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL
)`)

	if err != nil {
		return err
	}
	println("Tables created successfully")
	return nil
}

func createSeedData() error {
	// Check if the seed data already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Blog WHERE title IN ('Test1', 'Test2')").Scan(&count)
	if err != nil {
		return err
	}

	// If count is less than the number of seed entries, insert the missing data
	if count < 2 {
		_, err := db.Exec(`
            INSERT INTO Blog (title)
            SELECT 'Test1' WHERE NOT EXISTS (SELECT 1 FROM Blog WHERE title = 'Test1');
            INSERT INTO Blog (title)
            SELECT 'Test2' WHERE NOT EXISTS (SELECT 1 FROM Blog WHERE title = 'Test2');
        `)
		if err != nil {
			return err
		}
		println("Seed created successfully")
	} else {
		println("Seed data already exists")
	}

	return nil
}
