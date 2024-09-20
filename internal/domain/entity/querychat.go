package entity

import (
	"encoding/json"
	"fmt"
)

type QueryChatConfig struct {
	prompt       string
	systemPrompt string
	userPrompt   string
	dbName       string
	schema       string
}

func NewQueryChatConfig(prompt, systemPrompt, userPrompt, dbName, schema string) (*QueryChatConfig, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt cannot be empty")
	}
	if systemPrompt == "" {
		return nil, fmt.Errorf("system prompt cannot be empty")
	}
	if userPrompt == "" {
		return nil, fmt.Errorf("user prompt cannot be empty")
	}
	if dbName == "" {
		return nil, fmt.Errorf("db name cannot be empty")
	}
	if schema == "" {
		return nil, fmt.Errorf("schema cannot be empty")
	}

	return &QueryChatConfig{
		prompt:       prompt,
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
		dbName:       dbName,
		schema:       schema,
	}, nil
}

func (qcc *QueryChatConfig) Prompt() string {
	return qcc.prompt
}

func (qcc *QueryChatConfig) SystemPrompt() string {
	return qcc.systemPrompt
}

func (qcc *QueryChatConfig) UserPrompt() string {
	return qcc.userPrompt
}

func (qcc *QueryChatConfig) DBName() string {
	return qcc.dbName
}

func (qcc *QueryChatConfig) Schema() string {
	return qcc.schema
}

type VisualizableData struct {
	datas []map[string]interface{}
	chart map[string]string
}

func NewVisualizableData(datas []map[string]interface{}, output LLMOutput) (*VisualizableData, error) {
	if datas == nil {
		return nil, fmt.Errorf("datas cannot be nil")
	}

	dataBytes := []byte(output.Data())
	if !json.Valid(dataBytes) {
		return nil, fmt.Errorf("provided data is not valid JSON")
	}

	var chart map[string]interface{}
	err := json.Unmarshal(dataBytes, &chart)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	expectedKeys := map[string]struct{}{"type": {}, "x": {}, "y": {}}
	for key := range expectedKeys {
		if _, ok := chart[key]; !ok {
			return nil, fmt.Errorf("missing expected key: %s", key)
		}
	}
	for key := range chart {
		if _, ok := expectedKeys[key]; !ok {
			return nil, fmt.Errorf("unexpected key in JSON data: %s", key)
		}
	}

	var cleanedChart = make(map[string]string)
	for key, value := range chart {
		cleanedChart[key] = fmt.Sprintf("%v", value)
	}

	return &VisualizableData{
		datas: datas,
		chart: cleanedChart,
	}, nil
}

func (vd *VisualizableData) Datas() []map[string]interface{} {
	return vd.datas
}

func (vd *VisualizableData) Chart() map[string]string {
	return vd.chart
}
