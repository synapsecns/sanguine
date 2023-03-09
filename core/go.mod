module github.com/synapsecns/sanguine/core

go 1.19

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/ImVexed/fasturl v0.0.0-20230304231329-4e41488060f3
	github.com/LK4d4/trylock v0.0.0-20191027065348-ff7e133a5c54
	github.com/Soft/iter v0.1.0
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/brianvoe/gofakeit/v6 v6.20.1
	github.com/c-bata/go-prompt v0.2.6
	github.com/cheekybits/genny v1.0.0
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/ethereum/go-ethereum v1.11.3
	github.com/fatih/structtag v1.2.0
	github.com/filipemendespi/newrelic-context v0.0.0-20200414012421-889cae3d7978
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-contrib/requestid v0.0.6
	github.com/gin-contrib/zap v0.1.0
	github.com/gin-gonic/gin v1.9.0
	github.com/go-git/go-git/v5 v5.6.0
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/go-kit/kit v0.12.0
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0
	github.com/grafana-tools/sdk v0.0.0-20220919052116-6562121319fc
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/ipfs/go-log v1.0.5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/newrelic/go-agent/v3 v3.20.3
	github.com/newrelic/go-agent/v3/integrations/nrgin v1.1.3
	github.com/newrelic/go-agent/v3/integrations/nrzap v1.0.1
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/common v0.42.0
	github.com/prometheus/exporter-toolkit v0.7.1
	github.com/prometheus/pushgateway v1.4.3
	github.com/rung/go-safecast v1.0.1
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0
	github.com/stretchr/testify v1.8.2
	github.com/temoto/robotstxt v1.1.2
	github.com/urfave/cli/v2 v2.25.0
	go.uber.org/zap v1.24.0
	golang.org/x/exp v0.0.0-20230307190834-24139beb5833
	golang.org/x/sync v0.1.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.48.0
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.6
	k8s.io/apimachinery v0.26.2
)

require (
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.43.1 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.43.1 // indirect
	github.com/DataDog/datadog-go/v5 v5.3.0 // indirect
	github.com/DataDog/go-tuf v0.3.0--fix-localmeta-fork // indirect
	github.com/DataDog/gostackparse v0.6.0 // indirect
	github.com/DataDog/sketches-go v1.4.1 // indirect
	github.com/LK4D4/trylock v0.0.0-20191027065348-ff7e133a5c54 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/VictoriaMetrics/fastcache v1.12.1 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/bytedance/sonic v1.8.3 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cloudflare/circl v1.3.2 // indirect
	github.com/cockroachdb/pebble v0.0.0-20230307175142-1c38c4cb89d4 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/getsentry/sentry-go v0.19.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.4.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20230228050547-1710fef4ab10 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gosimple/slug v1.1.1 // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/huin/goupnp v1.1.0 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/matryer/is v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mattn/go-tty v0.0.4 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/outcaste-io/ristretto v0.2.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/richardartoul/molecule v1.0.1-0.20221107223329-32cfee06a052 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.5.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/skeema/knownhosts v1.1.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.2.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go4.org/intern v0.0.0-20230205224052-192e9f60865c // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20230221090011-e4bae7ad2296 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.29.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	inet.af/netaddr v0.0.0-20220811202034-502d2d690317 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448 // indirect
)
