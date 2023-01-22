include ../../make/go.Makefile

install-plugin-local: # will install the terraform provider as a local plugin for testing.
	./scripts/build-tf.sh

run-example: install-plugin-local # runs an example
	rm -rf examples/.terraform rm -rf examples/.terraform.lock.hcl
	echo "running terraform init, if this fails, you might have to specify amd64 as the arch before using terraform, please see: https://github.com/tfutils/tfenv/issues/337"
	cd examples && terraform init
