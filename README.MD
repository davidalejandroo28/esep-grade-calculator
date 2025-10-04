# ESEP Grade Calculator (Go)

A small Go library that computes a student’s final letter grade based on weighted averages:
- Assignments: 50%
- Exams: 35%
- Essays: 15%

Letter scale:
- A ≥ 90, B ≥ 80, C ≥ 70, D ≥ 60, F < 60

## Prerequisites
- Go 1.x installed: `go version`
- (Recommended) VS Code + “Go” extension by the Go Team (for inline test running)

## Project Layout
- `grade-calculator/grade_calculator.go`: Library implementation
- `grade-calculator/grade_calculator_test.go`: Unit tests
- `grade-calculator/go.mod`: Go module definition

## Run Tests
```bash
cd grade-calculator
go test
go test -coverage
