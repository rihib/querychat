package entity

import "fmt"

type UserDBInfo struct {
	name     string // e.g. "MySQL", "PostgreSQL"
	filepath string
	schema   string
}

func NewUserDBInfo(name, filepath, schema string) (*UserDBInfo, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if filepath == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}
	if schema == "" {
		return nil, fmt.Errorf("schema cannot be empty")
	}
	return &UserDBInfo{
		name:     name,
		filepath: filepath,
		schema:   schema,
	}, nil
}

func (info *UserDBInfo) Name() string {
	return info.name
}

func (info *UserDBInfo) Filepath() string {
	return info.filepath
}

func (info *UserDBInfo) Schema() string {
	return info.schema
}
