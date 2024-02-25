module github.com/synapsecns/sanguine/contrib/release-copier-action

go 1.21

require (
	github.com/brianvoe/gofakeit/v6 v6.27.0
	github.com/google/go-github/v41 v41.0.0
	github.com/migueleliasweb/go-github-mock v0.0.16
	github.com/sethvargo/go-githubactions v1.1.0
	github.com/stretchr/testify v1.8.4
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.17.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github/v50 v50.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/sethvargo/go-envconfig v0.8.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.29.2 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ../../core
)
