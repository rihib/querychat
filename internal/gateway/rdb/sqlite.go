package rdb

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type SQLite struct {
	DB *sql.DB
}

func NewSQLite() (*SQLite, error) {
	db, err := sql.Open("sqlite3", DB_FILE_PATH)
	if err != nil {
		db.Close()
		slog.Error("failed to open database", "msg", err.Error())
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &SQLite{DB: db}, nil
}

func (s *SQLite) Exec(query string) (*sql.Rows, error) {
	defer s.DB.Close()
	rows, err := s.DB.Query(query)
	if err != nil {
		slog.Error("failed to query database", "msg", err.Error())
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	return rows, nil
}

const (
	SCHEMA_FILE_PATH = "schema.sql"
	DB_FILE_PATH     = "chinook.db"
)
