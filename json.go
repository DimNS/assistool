package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONTab struct.
type JSONTab struct{}

// NewJSONTab creates a new JsonTab.
func NewJSONTab() *JSONTab {
	return &JSONTab{}
}

// Beautify beautify json.
func (t *JSONTab) Beautify(jsonString string, ntconvert bool) (string, error) {
	res, err := t.jsonBeautify(jsonString, ntconvert)
	if err != nil {
		return "", fmt.Errorf("json beautify: %v", err)
	}

	return res, nil
}

func (t *JSONTab) jsonBeautify(str string, ntv bool) (string, error) { //nolint:revive // it's ok
	str = strings.TrimSpace(str)
	if str == "" {
		return "", nil
	}

	buff := bytes.NewBuffer([]byte(str))
	var data any
	if err := json.NewDecoder(buff).Decode(&data); err != nil {
		return "", fmt.Errorf("decode json: %v", err)
	}

	processedData := t.bypassAny(data)
	val, err := json.MarshalIndent(processedData, "", "    ")
	if err != nil {
		return "", fmt.Errorf("marshal json: %v", err)
	}

	res := string(val)

	if ntv {
		res = strings.ReplaceAll(res, "\\n", "\n")
		res = strings.ReplaceAll(res, "\\t", "\t")
	}

	return res, nil
}

func (t *JSONTab) bypassAny(data any) any { //nolint:gocognit // it's ok
	switch v := data.(type) {
	case map[string]any:
		return t.bypassMap(v)
	case []any:
		return t.bypassArray(v)
	case string:
		if strings.TrimSpace(v) == "" {
			return v
		}

		trimmed := strings.TrimSpace(v)
		if (strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}")) ||
			(strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]")) {
			var parsedData any
			buff := bytes.NewBuffer([]byte(v))
			if err := json.NewDecoder(buff).Decode(&parsedData); err == nil {
				return t.bypassAny(parsedData)
			}
		}
		return v
	default:
		return v
	}
}

func (t *JSONTab) bypassMap(data map[string]any) map[string]any {
	res := make(map[string]any)

	for k, v := range data {
		res[k] = t.bypassAny(v)
	}

	return res
}

func (t *JSONTab) bypassArray(data []any) []any {
	res := make([]any, len(data))

	for i, v := range data {
		res[i] = t.bypassAny(v)
	}

	return res
}
