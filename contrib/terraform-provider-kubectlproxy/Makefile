include ../../make/go.Makefile

install-plugin-local: # will install the terraform provider as a local plugin for testing.
	./scripts/build-tf.sh

run-example: install-plugin-local cleanup-examples # runs an example
	echo "running terraform init, if this fails, you might have to specify amd64 as the arch before using terraform, please see: https://github.com/tfutils/tfenv/issues/337"
	echo "on osx arm64, you can run run-example-m1 as a workaround."
	cd examples && terraform init

cleanup-examples:
	rm -rf examples/.terraform rm -rf examples/.terraform.lock.hcl

run-example-m1: install-plugin-local cleanup-examples # runs an example on osx arm64
	./scripts/add-tfmac.sh
	source ~/.zshrc
	echo "please run: cd examples && tfmac init"


tfenv-install:
	@#Brew - MacOS
	@if [ "$(shell which tflint)" = "" ] && [ "$(shell which brew)" != "" ]; then brew install rflint; fi;
	# default
	@if [ "$(shell which tflint)" = "" ]; then curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash; fi;


lint-tf: tfenv-install ## Run golangci-lint and go fmt ./...
	cd examples && tflint --init
	cd examples && tflint
