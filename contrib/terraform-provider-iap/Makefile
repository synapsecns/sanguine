include ../../make/go.Makefile

install-plugin-local: # will install the terraform provider as a local plugin for testing.
	GODEBUG=asyncpreemptoff=1  go build .
	## TODO: since this is just for local testing, the binary here is produced independent of architecture. This can be fixed with a go command of some kind to get the os + arch
	## TODO: use a variable here
	rm -rf ~/.terraform.d/plugins/example-iap.com/provider/iap/1.0.0/darwin_arm64
	mkdir -p ~/.terraform.d/plugins/example-iap.com/provider/iap/1.0.0/darwin_arm64
	cp terraform-provider-iap ~/.terraform.d/plugins/example-iap.com/provider/iap/1.0.0/darwin_arm64

run-example: install-plugin-local # runs an example
	rm -rf examples/.terraform rm -rf examples/.terraform.lock.hcl
	cd examples && terraform init
