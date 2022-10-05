include ../make/go.Makefile

install-tenderly: ## installs tenderly
	@if [ "$(shell which tenderly)" = "" ] && [ "$(shell which brew)" != "" ]; then brew tenderly/tenderly; brew install tenderly; fi;
	echo "make sure to login to tenderly!"
