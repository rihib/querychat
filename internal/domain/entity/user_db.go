package entity

import (
	"fmt"
	"os"
)

type UserDBInfo struct {
	name     string // e.g. "MySQL", "PostgreSQL"
	filepath string
	schema   string
}

func NewUserDBInfo(name, filepath, schemapath string) (*UserDBInfo, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if filepath == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}
	if schemapath == "" {
		return nil, fmt.Errorf("schemapath cannot be empty")
	}
	schema, err := os.ReadFile(schemapath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return &UserDBInfo{
		name:     name,
		filepath: filepath,
		schema:   string(schema),
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
