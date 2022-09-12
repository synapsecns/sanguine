default: help

GIT_ROOT := $(shell git rev-parse --show-toplevel)


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

helm-install:
	@#Brew - MacOS
	@if [ "$(shell which helm)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install helm; fi;

dependencies: helm-install ## install dependencies for all helm charts
	cd $(GIT_ROOT)
	$(GIT_ROOT)/make/scripts/helm_dependency.sh
