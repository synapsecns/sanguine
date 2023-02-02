module github.com/synapsecns/sanguine/contrib/git-changes-action

go 1.19

require (
	github.com/Flaque/filet v0.0.0-20201012163910-45f684403088
	github.com/brianvoe/gofakeit/v6 v6.19.0
	github.com/ethereum/go-ethereum v1.10.26
	github.com/go-git/go-git/v5 v5.4.2
	github.com/kendru/darwin/go/depgraph v0.0.0-20221105232959-877d6a81060c
	github.com/otiai10/copy v1.9.0
	github.com/stretchr/testify v1.8.1
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000
	github.com/vishalkuo/bimap v0.0.0-20220726225509-e0b4f20de28b
	github.com/xlab/treeprint v1.1.0
	golang.org/x/mod v0.7.0
)

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.25.5 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ../../core
)
