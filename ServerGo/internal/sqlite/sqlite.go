package sqlite

import "database/sql"

var DB *sql.DB

func InitDataBase(address string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", address)
	if err != nil {
		return db, err
	}
	return db, nil
}
