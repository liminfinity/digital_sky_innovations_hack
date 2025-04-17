package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kurochkinivan/ReportSender/truncate"
)

func TestTruncatePIDs(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	defer os.Remove("test.db")

	countBefore := getCount(t, db)
	if countBefore != 1000 {
		t.Fatalf("expected 1000 records before truncate, got %d", countBefore)
	}

	_, err := truncate.TruncatePIDs(db)
	if err != nil {
		t.Fatalf("truncate failed: %v", err)
	}

	countAfter := getCount(t, db)
	if countAfter != 0 {
		t.Fatalf("expected 0 records after truncate, got %d", countAfter)
	}
}

func setupTestDB(t *testing.T) *sql.DB {
	dbPath := "test.db"
	os.Remove(dbPath)

	db, err := truncate.ConnectToDB(dbPath)
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE pids (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT,
		created_at DATETIME
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO pids(filename, created_at) VALUES (?, ?)")
	if err != nil {
		t.Fatalf("prepare insert failed: %v", err)
	}
	defer stmt.Close()

	for i := 0; i < 1000; i++ {
		_, err := stmt.Exec(fmt.Sprintf("file_%d.txt", i), time.Now())
		if err != nil {
			t.Fatalf("insert failed at row %d: %v", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("commit failed: %v", err)
	}

	return db
}

func getCount(t *testing.T, db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pids").Scan(&count)
	if err != nil {
		t.Fatalf("count query failed: %v", err)
	}
	return count
}
