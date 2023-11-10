# Note: this file is made to be symlinked to various folders where we use go for builds
# please use libraries if you wish to add folder-specific make functionality

default: help

# set variables
GIT_ROOT := $(shell git rev-parse --show-toplevel)
CURRENT_PATH := $(shell pwd)
RELPATH := $(shell perl -e 'use Cwd "abs_path"; use File::Spec; print File::Spec->abs2rel("$(shell pwd)", "$(GIT_ROOT)")')


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


lint: ## Run golangci-lint and go fmt ./...
	go mod tidy
	go fmt ./...
	cd $(GIT_ROOT)
	go work sync
	cd $(CURRENT_PATH)
	# Note: when we upgrade, we can use either the brew version (needs to stay at latest). If we decide to stay with docker, we can use gomemlimit instead of a constant heap size.
	# TODO: investigate why this is so much slower than local install
	docker run -t --rm -v $(shell go env GOCACHE):/cache/go -v ${GOPATH}/pkg:/go/pkg -e GOGC=2000  -e GOCACHE=/cache/go -v ~/.cache/golangci-lint/:/root/.cache -v "$(GIT_ROOT)":/app -w "/app/$(RELPATH)" golangci/golangci-lint:v1.48.0 golangci-lint run -v --fix

docker-clean: ## stops and removes all containers at once
	docker ps -aq | xargs docker stop | xargs docker rm
	docker network prune

lint-legacy:
	go mod tidy
	go fmt ./...
	cd $(GIT_ROOT)
	go work sync
	cd $(CURRENT_PATH)
	@golangci-lint run --fix --config=$(GIT_ROOT)/.golangci.yml
