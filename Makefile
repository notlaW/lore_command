
# TESTING TASKS
# Test golang app
test: ## Run all test files
	go test ./...

# DOCKER TASKS
# Build the container
coverage: ## Build the release and develoment container. The development
	go test -cover ./...