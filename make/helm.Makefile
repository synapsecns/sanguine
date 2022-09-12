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

kind-install:
	@#Brew - MacOS
	@if [ "$(shell which kind)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install kind; fi;

# used for parsing yaml
yq-install:
	@#Github Actions
	@if [ "$(shell which yq)" = "" ] && [ "$(GITHUB_WORKFLOW)" != "" ]; then curl -sfL wget https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64 -O /usr/bin/yq && chmod +x /usr/bin/yq; fi;
	@#Brew - MacOS
	@if [ "$(shell which yq)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install yq; fi;

ct-install:
	@#Brew - MacOS
	@if [ "$(shell which ct)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install chart-testing; fi;

dependencies: yq-install helm-install ## install dependencies for all helm charts
	cd $(GIT_ROOT)
	$(GIT_ROOT)/make/scripts/helm_dependency.sh

lint: ct-install dependencies ## lints helm charts
	cd $(GIT_ROOT);	ct lint --all --validate-maintainers=false

test-install: ct-install kind-install## test chart installs on a local kubernetes cluster



# list helm charts, used for: https://github.com/helm/chart-testing/issues/226
list-chart-dirs: ## list all chart directories
	@eval $$(cd $(GIT_ROOT)/charts);	$(GIT_ROOT)/make/scripts/chart-dirs.sh
