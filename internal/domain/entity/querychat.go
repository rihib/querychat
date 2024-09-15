package entity

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type VisualizableData struct {
	datas []map[string]interface{}
	chart map[string]interface{}
}

func NewVisualizableData(rows *sql.Rows, output LLMOutput) (*VisualizableData, error) {
	if rows == nil {
		return nil, fmt.Errorf("rows cannot be nil")
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Printf("failed to get columns: %v", err)
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}
	datas := []map[string]interface{}{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := 0; i < len(columns); i++ {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Printf("failed to scan rows: %v", err)
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			var value interface{}
			val := values[i]

			b, ok := val.([]byte)
			if ok {
				value = string(b)
			} else {
				value = val
			}

			rowMap[colName] = value
		}
		datas = append(datas, rowMap)
	}
	if err := rows.Err(); err != nil {
		log.Printf("failed to get rows: %v", err)
		return nil, fmt.Errorf("failed to get rows: %v", err)
	}
	defer rows.Close()

	var chart map[string]interface{}
	err = json.Unmarshal([]byte(output.Data()), &chart)
	if err != nil {
		log.Printf("failed to unmarshal JSON: %v", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &VisualizableData{
		datas: datas,
		chart: chart,
	}, nil
}
