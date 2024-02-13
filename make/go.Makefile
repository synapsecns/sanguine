r# Note: this file is made to be symlinked to various folders where we use go for builds
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

docker-clean: ## stops and removes all containers at once
	docker ps -aq | xargs docker stop | xargs docker rm
	docker network prune

lint: ## lint lints the code with golangci-lint
	go mod tidy
	go fmt ./...
	cd $(GIT_ROOT)
	go work sync
	cd $(CURRENT_PATH)
	@golangci-lint run --fix --config=$(GIT_ROOT)/.golangci.yml
