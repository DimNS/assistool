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
	data := map[string]any{}

	if err := json.NewDecoder(buff).Decode(&data); err != nil {
		return "", fmt.Errorf("decode json: %v", err)
	}

	val, err := json.MarshalIndent(t.bypassMap(data), "", "    ")
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

func (t *JSONTab) bypassMap(data map[string]any) map[string]any { //nolint:gocognit // it's ok
	res := map[string]any{}

	for k, v := range data {
		if vv, ok := v.(map[string]any); ok {
			res[k] = t.bypassMap(vv)

			continue
		}

		if vv, ok := v.(string); ok {
			if (strings.HasPrefix(vv, "{") && strings.HasSuffix(vv, "}")) ||
				strings.HasPrefix(vv, "[") && strings.HasSuffix(vv, "]") {
				b := bytes.NewBuffer([]byte(vv))
				j := map[string]any{}

				if err := json.NewDecoder(b).Decode(&j); err != nil {
					fmt.Printf("Json decode: %v\n", err)

					return nil
				}

				res[k] = t.bypassMap(j)

				continue
			}

			res[k] = vv

			continue
		}

		res[k] = v
	}

	return res
}
