#!/bin/bash
include .env

# run the client
run: ## calling the cmd to run the client.
	@echo "\033[2m→ Running the command line executable...\033[0m"
	@go run cmd/main.go

lint:
	@echo "\033[2m→ Running linter...\033[0m"
	@golangci-lint run --config .golangci.yaml

test:
	@echo "Go tests of this project"
	@go test ./...