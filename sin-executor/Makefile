include ../make/go.Makefile

generate: ## Generate code
	cd $(GIT_ROOT)/packages/contracts-communication && yarn build:go
	go generate ./...


