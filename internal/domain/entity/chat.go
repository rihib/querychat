package entity

import (
	"encoding/json"
	"fmt"
)

/*
ChatConfig is a configuration for the chat.
Mainly used for constructing the optimized prompt.
*/
type ChatConfig struct {
	prompt       string
	systemPrompt string
	userPrompt   string
	dbName       string
	schema       string
}

func NewChatConfig(prompt, dbName, schema string) (*ChatConfig, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt cannot be empty")
	}
	if dbName == "" {
		return nil, fmt.Errorf("db name cannot be empty")
	}
	if schema == "" {
		return nil, fmt.Errorf("schema cannot be empty")
	}
	return &ChatConfig{
		prompt:       prompt,
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
		dbName:       dbName,
		schema:       schema,
	}, nil
}

func (cc *ChatConfig) Prompt() string {
	return cc.prompt
}

func (cc *ChatConfig) DBName() string {
	return cc.dbName
}

/*
Optimized prompt is a prompt that is optimized based on original prompt.
Prompt is composed of two parts: system prompt and user prompt.
System prompt is the part that tells the LLM the prerequisites.
User prompt is the part that asks the actual question.
*/
type OptimizedPrompt struct {
	systemPrompt string
	userPrompt   string
}

func NewOptimizedPrompt(cc ChatConfig) (*OptimizedPrompt, error) {
	systemPrompt := fmt.Sprintf(cc.systemPrompt, cc.schema)
	userPrompt := fmt.Sprintf(cc.userPrompt, cc.dbName, cc.prompt, cc.dbName)
	return &OptimizedPrompt{
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
	}, nil
}

func (optimized *OptimizedPrompt) SystemPrompt() string {
	return optimized.systemPrompt
}

func (optimized *OptimizedPrompt) UserPrompt() string {
	return optimized.userPrompt
}

/*
LLMOutput is the output of the LLM.

Example:

	{
		query: "SELECT user_name, SUM(amount) AS total_amount FROM purchases GROUP BY user_id, user_name".
		chart: {"type": "bar", "x": "UserName", "y": "TotalAmount"},
	}
*/
type LLMOutput struct {
	query string
	chart string
}

func NewLLMOutput(query string, chart string) (*LLMOutput, error) {
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}
	if chart == "" {
		return nil, fmt.Errorf("chart cannot be empty")
	}
	return &LLMOutput{
		query: query,
		chart: chart,
	}, nil
}

func (output *LLMOutput) Query() string {
	return output.query
}

func (output *LLMOutput) Chart() string {
	return output.chart
}

/*
VisualizableData is the data that can be visualized.

Example:

	{
		"chart": {"type": "bar", "x": "UserName", "y": "TotalAmount"},
		"datas": [{"UserName": "Alice", "TotalAmount": 100},{"UserName": "Bob", "TotalAmount": 200}],
		"executedQuery": "SELECT user_name, SUM(amount) AS total_amount FROM purchases GROUP BY user_id, user_name"
	}
*/
type VisualizableData struct {
	chart         map[string]string
	datas         []map[string]interface{}
	executedQuery string
}

func NewVisualizableData(output LLMOutput, datas []map[string]interface{}) (*VisualizableData, error) {
	if datas == nil {
		return nil, fmt.Errorf("datas cannot be nil")
	}

	chartBytes := []byte(output.chart)
	if !json.Valid(chartBytes) {
		return nil, fmt.Errorf("provided chart is not valid JSON")
	}

	var chart map[string]interface{}
	err := json.Unmarshal(chartBytes, &chart)
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
			return nil, fmt.Errorf("unexpected key in JSON chart: %s", key)
		}
	}

	var cleanedChart = make(map[string]string)
	for key, value := range chart {
		cleanedChart[key] = fmt.Sprintf("%v", value)
	}

	return &VisualizableData{
		chart:         cleanedChart,
		datas:         datas,
		executedQuery: output.query,
	}, nil
}

func (vd *VisualizableData) Chart() map[string]string {
	return vd.chart
}

func (vd *VisualizableData) Datas() []map[string]interface{} {
	return vd.datas
}

func (vd *VisualizableData) ExecutedQuery() string {
	return vd.executedQuery
}
