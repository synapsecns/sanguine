default: help

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

check_reset:
	@echo -n "Are you sure? This action will delete any files not associated with the git repo, reset all submodules and reset any unchanged work. Files like node_modules, flattened files, etc will be deleted [y/N] " && read ans && [ $${ans:-N} = y ]

reset-submodules: ## hard resets all submodules
	git submodule foreach --recursive git clean -xfd
	git submodule foreach --recursive git reset --hard
	git submodule update --init --recursive

full-reset: check_reset reset-submodules ## Fully reset the repository to it's cloned state
	git clean -xfdf
	git reset --hard

go-  :
	find . \( -name vendor -o -name '[._].*' -o -name node_modules \) -prune -o -name go.mod -print | sed 's:/go.mod$::'

.PHONY: full-reset check_reset
