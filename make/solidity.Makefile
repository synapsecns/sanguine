# Note: this file is made to be symlinked to various folders where we use go for builds
# please use libraries if you wish to add folder-specific make functionality

default: help

# set variables
GIT_ROOT := $(shell git rev-parse --show-toplevel)
CURRENT_PATH := $(shell pwd)
RELPATH := $(shell realpath --relative-to="$(GIT_ROOT)" "$(CURRENT_PATH)")

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

# TODO: deployer utils requires jq, install

foundry-install:
	@if [ "$(shell which forge)" = "" ]; then curl -L https://foundry.paradigm.xyz | bash && foundryup; fi
