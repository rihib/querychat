package rdb

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rihib/querychat/internal/domain/entity"
)

type SQLite3 struct {
	DB *sql.DB
}

func NewSQLite3(info *entity.UserDBInfo) (*SQLite3, error) {
	db, err := sql.Open(info.Name(), info.Filepath())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &SQLite3{DB: db}, nil
}

func (s *SQLite3) Exec(output entity.LLMOutput) (*sql.Rows, error) {
	rows, err := s.DB.Query(output.Query())
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	return rows, nil
}
