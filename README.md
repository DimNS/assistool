# Assistool

[![Go Report Card](https://goreportcard.com/badge/github.com/dimns/assistool?style=flat-square)](https://goreportcard.com/report/github.com/dimns/assistool)
[![Audit](https://github.com/dimns/assistool/actions/workflows/audit.yml/badge.svg?branch=master)](https://github.com/dimns/assistool/actions/workflows/audit.yml)
![GitHub release](https://img.shields.io/github/v/release/dimns/assistool)
![License](https://img.shields.io/github/license/dimns/assistool.svg)

## Comments

-   The use of `fmt.Errorf("error: %v", err)` instead of `fmt.Errorf("error: %w", err)` is due to the fact that this error is not unwrapped anywhere above.

- ULID
  - добавить multiline где генерировать сразу 20 строк
- подключить либу https://github.com/brianvoe/gofakeit
