package entity

import "fmt"

type UserDBInfo struct {
	name   string // e.g. "MySQL", "PostgreSQL"
	schema string
}

func NewUserDBInfo(name, schema string) (*UserDBInfo, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if schema == "" {
		return nil, fmt.Errorf("schema cannot be empty")
	}
	return &UserDBInfo{
		name:   name,
		schema: schema,
	}, nil
}
