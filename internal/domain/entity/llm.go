package entity

import (
	"fmt"
)

type OriginalPrompt struct {
	prompt string
}

func NewOriginalPrompt(prompt string) (*OriginalPrompt, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt cannot be empty")
	}
	return &OriginalPrompt{prompt: prompt}, nil
}

type FormatPrompt struct {
	systemPrompt string
	userPrompt   string
}

func NewFormatPrompt(systemPrompt, userPrompt string) (*FormatPrompt, error) {
	if systemPrompt == "" {
		return nil, fmt.Errorf("system prompt cannot be empty")
	}
	if userPrompt == "" {
		return nil, fmt.Errorf("user prompt cannot be empty")
	}
	return &FormatPrompt{
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
	}, nil
}

type OptimizedPrompt struct {
	systemPrompt string
	userPrompt   string
}

func NewOptimizedPrompt(original OriginalPrompt, template FormatPrompt, dbName string, schema string) (*OptimizedPrompt, error) {
	if schema == "" {
		return nil, fmt.Errorf("schema cannot be empty")
	}
	systemPrompt := fmt.Sprintf(template.systemPrompt, schema)
	userPrompt := fmt.Sprintf(template.userPrompt, dbName, original.prompt, dbName)
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
