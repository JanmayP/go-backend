package db

import "database/sql"

func Init() (*sql.DB, error) {
	return &sql.DB{}, nil
}
