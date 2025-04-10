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

func (t *GoStructTab) goStructBeautify(input string) (string, error) { //nolint:gocognit //nolint:gocyclo // it's ok
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

		switch {
		case char == '{':
			// Проверяем пустую структуру {}
			if i+1 < len(input) && input[i+1] == '}' {
				result.WriteString("{}")
				i++
				continue
			}

			stack = append(stack, indentLevel)
			indentLevel++
			result.WriteByte(char)
			result.WriteByte('\n')
			result.WriteString(strings.Repeat(" ", indentLevel*indentSize))

		case char == '}':
			if len(stack) > 0 {
				indentLevel = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			result.WriteByte('\n')
			result.WriteString(strings.Repeat(" ", indentLevel*indentSize))
			result.WriteByte(char)

		case char == '(':
			// Проверяем, не начинается ли это time.Date
			if i > 4 && strings.HasSuffix(input[:i], "Date") {
				inDateFunction = true
			}

			result.WriteByte(char)

		case char == ')':
			if inDateFunction {
				inDateFunction = false
			}

			result.WriteByte(char)

		case char == ',':
			result.WriteByte(char)

			if inDateFunction { //nolint:gocritic // it's ok
				// Для time.Date оставляем запятую с пробелом
				result.WriteByte(' ')
			} else if len(stack) > 0 && stack[len(stack)-1] == indentLevel-1 { //nolint:gocritic // it's ok
				// Для вложенных структур делаем перенос
				result.WriteByte('\n')
				result.WriteString(strings.Repeat(" ", indentLevel*indentSize))
			} else {
				// Для верхнего уровня делаем перенос
				result.WriteByte('\n')
				result.WriteString(strings.Repeat(" ", indentLevel*indentSize))
			}

			// Пропускаем пробел после запятой, если он есть
			if i+1 < len(input) && input[i+1] == ' ' {
				i++
			}

		default:
			result.WriteByte(char)
		}
	}

	return result.String(), nil
}
