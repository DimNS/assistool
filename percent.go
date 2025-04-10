package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	oneHundred = 100
)

// PercentTab struct.
type PercentTab struct{}

// NewPercentTab creates a new PercentTab.
func NewPercentTab() *PercentTab {
	return &PercentTab{}
}

// Calculate calculates the difference between two numbers.
func (t *PercentTab) Calculate(input1, input2 string) (string, error) {
	r, err := t.calculate(input1, input2)
	if err != nil {
		return "", fmt.Errorf("calculate: %v", err)
	}

	return fmt.Sprintf("Result: %.2f%%", r), nil
}

func (t *PercentTab) calculate(s1, s2 string) (float64, error) {
	s1 = strings.TrimSpace(s1)
	if s1 == "" {
		return 0, fmt.Errorf("first value is empty")
	}

	s2 = strings.TrimSpace(s2)
	if s2 == "" {
		return 0, fmt.Errorf("second value is empty")
	}

	s1 = strings.ReplaceAll(s1, ",", ".")
	s2 = strings.ReplaceAll(s2, ",", ".")

	n1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		return 0, fmt.Errorf("first value is not a number: %v", err)
	}

	n2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return 0, fmt.Errorf("second value is not a number: %v", err)
	}

	return math.Abs(n1-n2) / math.Max(n1, n2) * oneHundred, nil
}
