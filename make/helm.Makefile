default: help

GIT_ROOT := $(shell git rev-parse --show-toplevel)
YQ_VERSION := "v4.27.5"


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

# used for parsing yaml
yq-install:
	@#Github Actions
	@if [ "$(shell which yq)" = "" ] && [ "$(GITHUB_WORKFLOW)" != "" ]; then curl -sfL wget https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64 -O /usr/bin/yq && chmod +x /usr/bin/yq; fi;
	@#Brew - MacOS
	@if [ "$(shell which yq)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install yq; fi;

dependencies: yq-install helm-install ## install dependencies for all helm charts
	cd $(GIT_ROOT)
	$(GIT_ROOT)/make/scripts/helm_dependency.sh
