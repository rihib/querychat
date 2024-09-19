package entity

import (
	"encoding/json"
	"fmt"
)

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
