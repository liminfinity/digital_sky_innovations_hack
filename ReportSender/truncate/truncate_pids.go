package truncate

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func ConnectToDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TruncatePIDs(db *sql.DB) (sql.Result, error) {
	res, err := db.Exec("DELETE FROM pids")
	if err != nil {
		return nil, err
	}

	return res, nil
}
