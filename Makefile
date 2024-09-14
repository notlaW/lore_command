
# TESTING TASKS
# Test golang app
test: ## Run all test files
	go test ./...

# DOCKER TASKS
# Build the container
coverage: ## Build the release and develoment container. The development
	go test -cover ./...

build-container: ## Generate container `{version}` tag
	@echo 'build container'
	docker build -t lore_command .


tag-latest: ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag lore_command 381492125334.dkr.ecr.us-west-2.amazonaws.com/lore_command:latest

repo-login: ## Auto login to AWS-ECR unsing aws-cli
	aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 381492125334.dkr.ecr.us-west-2.amazonaws.com

push: # Push the container
	@echo 'lore command to repo'
	docker push 381492125334.dkr.ecr.us-west-2.amazonaws.com/lore_command:latest

