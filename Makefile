.PHONY: test test-verbose test-cover test-race test-bench clean test-unit test-integration test-mock test-all help

# Запуск всех тестов
test:
	go test ./...

# Подробный вывод тестов
test-verbose:
	go test -v ./...

# Тесты с покрытием кода
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report saved to coverage.html"

# Тесты на race conditions
test-race:
	go test -race ./...

# Бенчмарки
test-bench:
	go test -bench=. -benchmem ./...

# Запуск юнит-тестов
test-unit:
	go test -v ./internal/place/services/...

# Запуск интеграционных тестов
test-integration:
	go test -v ./tests/integration/...

# Очистка
clean:
	rm -f coverage.out coverage.html

# Запуск с моками
test-mock:
	REPO=fake go test -v ./...

# Полный цикл тестирования
test-all: test-race test-cover test-bench

# Справка
help:
	@echo "Available targets:"
	@echo "  test          - Run all tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-cover    - Run tests with coverage report"
	@echo "  test-race     - Run tests with race detection"
	@echo "  test-bench    - Run benchmarks"
	@echo "  test-unit     - Run unit tests only"
	@echo "  test-integration - Run integration tests"
	@echo "  test-mock     - Run tests with fake repository"
	@echo "  test-all      - Run complete test suite"
	@echo "  clean         - Clean coverage files"