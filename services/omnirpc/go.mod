module github.com/synapsecns/sanguine/services/omnirpc

go 1.19

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ../../core
	github.com/synapsecns/sanguine/ethergo => ../../ethergo
)

require (
	github.com/ImVexed/fasturl v0.0.0-20200730185836-b0c0fbead04e
	github.com/Soft/iter v0.1.0
	github.com/benbjohnson/immutable v0.4.0
	github.com/brianvoe/gofakeit/v6 v6.19.0
	github.com/buger/jsonparser v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.10.26
	github.com/gin-gonic/gin v1.8.1
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/go-resty/resty/v2 v2.7.0
	github.com/goccy/go-json v0.9.7
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hedzr/cmdr v1.10.49
	github.com/invopop/jsonschema v0.6.0
	github.com/ipfs/go-log v1.0.5
	github.com/jarcoal/httpmock v1.2.0
	github.com/jftuga/ellipsis v1.0.0
	github.com/josephburnett/jd v1.6.1
	github.com/jpillora/backoff v1.0.0
	github.com/lmittmann/w3 v0.10.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/puzpuzpuz/xsync v1.4.3
	github.com/rbretecher/go-postman-collection v0.9.0
	github.com/richardwilkes/toolbox v1.74.0
	github.com/stretchr/testify v1.8.1
	github.com/synapsecns/fasthttp-http2 v1.0.0
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000
	github.com/synapsecns/sanguine/ethergo v0.0.2
	github.com/urfave/cli/v2 v2.16.3
	github.com/valyala/fasthttp v1.41.0
	gitlab.com/1f320/x v0.3.0
	go.uber.org/automaxprocs v1.5.1
	golang.org/x/exp v0.0.0-20221012211006-4de253d81b95
	golang.org/x/sync v0.1.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.25.1
)

require (
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.0.0-20211129110424-6491aa3bf583 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.41.0-rc.3.0.20221208094921-9e0b0bfca62a // indirect
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/datadog-go/v5 v5.0.2 // indirect
	github.com/DataDog/go-tuf v0.3.0--fix-localmeta-fork // indirect
	github.com/DataDog/gostackparse v0.5.0 // indirect
	github.com/DataDog/sketches-go v1.2.1 // indirect
	github.com/Flaque/filet v0.0.0-20201012163910-45f684403088 // indirect
	github.com/LK4d4/trylock v0.0.0-20191027065348-ff7e133a5c54 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/VictoriaMetrics/fastcache v1.6.0 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/badoux/checkmail v0.0.0-20181210160741-9661bd69e9ad // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/briandowns/spinner v1.6.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.0 // indirect
	github.com/c-bata/go-prompt v0.2.6 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgrr/http2 v0.3.5 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/flynn/json5 v0.0.0-20160717195620-7620272ed633 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/gin-contrib/cors v1.4.0 // indirect
	github.com/gin-contrib/requestid v0.0.6 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-contrib/zap v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-github/v37 v37.0.0 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gosimple/slug v1.1.1 // indirect
	github.com/grafana-tools/sdk v0.0.0-20220919052116-6562121319fc // indirect
	github.com/graph-gophers/graphql-go v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-bexpr v0.1.10 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hedzr/log v1.5.53 // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/huin/goupnp v1.0.3 // indirect
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/influxdata/influxdb v1.8.3 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.5.1 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097 // indirect
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc // indirect
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a // indirect
	github.com/keep-network/keep-common v1.7.1-0.20211012131917-7102d7b9c6a0 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/manifoldco/promptui v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/pointerstructure v1.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/newrelic/go-agent/v3 v3.20.2 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrgin v1.1.3 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrzap v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/palantir/pkg v1.0.1 // indirect
	github.com/palantir/pkg/datetime v1.0.1 // indirect
	github.com/palantir/pkg/safejson v1.0.1 // indirect
	github.com/palantir/pkg/safeyaml v1.0.1 // indirect
	github.com/palantir/pkg/transform v1.0.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/peterh/liner v1.2.1 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/richardartoul/molecule v1.0.1-0.20221107223329-32cfee06a052 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	github.com/rung/go-safecast v1.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.4.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/subosito/gotenv v1.4.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/teivah/onecontext v1.3.0 // indirect
	github.com/tenderly/tenderly-cli v1.4.6 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/otel v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	go4.org/intern v0.0.0-20211027215823-ae77deb06f29 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20220617031537-928513b29760 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/term v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.44.1 // indirect
	gopkg.in/hedzr/errors.v3 v3.0.21 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.24.0 // indirect
	inet.af/netaddr v0.0.0-20220617031823-097006376321 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
)
