# Note: this file is made to be symlinked to various folders where we use go for builds
# please use libraries if you wish to add folder-specific make functionality

default: help

# set variables
GIT_ROOT := $(shell git rev-parse --show-toplevel)
CURRENT_PATH := $(shell pwd)



help: ## This help dialog.
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "%-30s %s\n" $$help_command $$help_info ; \
	done

# TODO tag a version
golangci-install:
	@#Travis (has sudo)
	@if [ "$(shell which golangci-lint)" = "" ] && [ $(TRAVIS) ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b && sudo cp ./bin/golangci-lint $(go env GOPATH)/bin/; fi;
	@#AWS CodePipeline
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(CODEBUILD_BUILD_ID)" != "" ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin; fi;
	@#Github Actions
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(GITHUB_WORKFLOW)" != "" ]; then curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin; fi;
	@#Brew - MacOS
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install golangci-lint; fi;


lint: golangci-install ## Run golangci-lint and go fmt ./...
	go mod tidy
	go fmt ./...
	cd $(GIT_ROOT)
	go work sync
	cd $(CURRENT_PATH)
	@golangci-lint run --fix --config=$(GIT_ROOT)/.golangci.yml
