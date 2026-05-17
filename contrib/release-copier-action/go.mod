module github.com/synapsecns/sanguine/contrib/release-copier-action

go 1.26.3

require (
	github.com/brianvoe/gofakeit/v6 v6.27.0
	github.com/google/go-github/v41 v41.0.0
	github.com/migueleliasweb/go-github-mock v1.5.0
	github.com/sethvargo/go-githubactions v1.1.0
	github.com/stretchr/testify v1.11.1
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.36.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-github/v73 v73.0.0 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/sethvargo/go-envconfig v0.8.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	golang.org/x/crypto v0.50.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.35.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20251002143259-bc988d571ff4 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ../../core
)
