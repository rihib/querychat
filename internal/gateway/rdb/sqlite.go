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

func NewSQLite3(filepath string) (*SQLite3, error) {
	if filepath == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &SQLite3{DB: db}, nil
}

func (s *SQLite3) Exec(output entity.LLMOutput) ([]map[string]interface{}, error) {
	rows, err := s.DB.Query(output.Query())
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}
	datas := []map[string]interface{}{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}
		datas = append(datas, rowMap)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get rows: %v", err)
	}

	return datas, nil
}
