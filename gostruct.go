package main

import (
	"fmt"
	"strings"
)

// GoStructTab struct.
type GoStructTab struct{}

// NewGoStructTab creates a new GoStructTab.
func NewGoStructTab() *GoStructTab {
	return &GoStructTab{}
}

// Beautify beautify go struct.
func (t *GoStructTab) Beautify(gostruct string) (string, error) {
	res, err := t.goStructBeautify(gostruct)
	if err != nil {
		return "", fmt.Errorf("go struct beautify: %v", err)
	}

	return res, nil
}

func (*GoStructTab) goStructBeautify(input string) (string, error) { //nolint:gocognit,gocyclo,cyclop // it's ok
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil
	}

	var (
		result strings.Builder

		indentLevel = 0
		indentSize  = 2

		// Стек для отслеживания вложенности структур
		stack = []int{}

		inDateFunction = false
	)

	for i := 0; i < len(input); i++ {
		char := input[i]

		switch char {
		case '{':
			// Проверяем пустую структуру {}
			if i+1 < len(input) && input[i+1] == '}' {
				_, err := result.WriteString("{}")
				if err != nil {
					return "", fmt.Errorf("write string: %v", err)
				}

				i++
				continue
			}

			stack = append(stack, indentLevel)
			indentLevel++

			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}
			if err := result.WriteByte('\n'); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}
			if _, err := result.WriteString(strings.Repeat(" ", indentLevel*indentSize)); err != nil {
				return "", fmt.Errorf("write string: %v", err)
			}

		case '}':
			if len(stack) > 0 {
				indentLevel = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			if err := result.WriteByte('\n'); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}
			if _, err := result.WriteString(strings.Repeat(" ", indentLevel*indentSize)); err != nil {
				return "", fmt.Errorf("write string: %v", err)
			}
			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}

		case '(':
			// Проверяем, не начинается ли это time.Date
			if i > 4 && strings.HasSuffix(input[:i], "Date") {
				inDateFunction = true
			}

			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}

		case ')':
			if inDateFunction {
				inDateFunction = false
			}

			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}

		case ',':
			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}

			if inDateFunction { //nolint:gocritic,nestif // it's ok
				// Для time.Date оставляем запятую с пробелом
				if err := result.WriteByte(' '); err != nil {
					return "", fmt.Errorf("write byte: %v", err)
				}
			} else if len(stack) > 0 && stack[len(stack)-1] == indentLevel-1 { //nolint:gocritic,revive // it's ok
				// Для вложенных структур делаем перенос
				if err := result.WriteByte('\n'); err != nil {
					return "", fmt.Errorf("write byte: %v", err)
				}
				if _, err := result.WriteString(strings.Repeat(" ", indentLevel*indentSize)); err != nil {
					return "", fmt.Errorf("write string: %v", err)
				}
			} else {
				// Для верхнего уровня делаем перенос
				if err := result.WriteByte('\n'); err != nil {
					return "", fmt.Errorf("write byte: %v", err)
				}
				if _, err := result.WriteString(strings.Repeat(" ", indentLevel*indentSize)); err != nil {
					return "", fmt.Errorf("write string: %v", err)
				}
			}

			// Пропускаем пробел после запятой, если он есть
			if i+1 < len(input) && input[i+1] == ' ' {
				i++
			}

		default:
			if err := result.WriteByte(char); err != nil {
				return "", fmt.Errorf("write byte: %v", err)
			}
		}
	}

	return result.String(), nil
}
