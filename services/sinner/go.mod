module github.com/synapsecns/sanguine/services/sinner

go 1.20

replace (
	github.com/Yamashou/gqlgenc => github.com/synapsecns/gqlgenc v0.10.0-hotfix
	// later versions give errors on uint64 being too high.
	github.com/brianvoe/gofakeit/v6 => github.com/brianvoe/gofakeit/v6 v6.9.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ./../../core
	github.com/synapsecns/sanguine/ethergo => ./../../ethergo
	github.com/synapsecns/sanguine/services/omnirpc => ../omnirpc
	github.com/synapsecns/sanguine/services/scribe => ../scribe

)

require (
	github.com/99designs/gqlgen v0.17.36
	github.com/MichaelMure/go-term-markdown v0.1.4
	github.com/Yamashou/gqlgenc v0.10.0
	github.com/ethereum/go-ethereum v1.10.26
	github.com/friendsofgo/graphiql v0.2.2
	github.com/gin-gonic/gin v1.9.1
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/ipfs/go-log v1.0.5
	github.com/jftuga/ellipsis v1.0.0
	github.com/jftuga/termsize v1.0.2
	github.com/jpillora/backoff v1.0.0
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/ravilushqa/otelgqlgen v0.13.1
	github.com/richardwilkes/toolbox v1.74.0
	github.com/stretchr/testify v1.8.4
	github.com/synapsecns/sanguine/agents v0.0.243
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000
	github.com/synapsecns/sanguine/services/explorer v0.0.177
	github.com/synapsecns/sanguine/services/scribe v0.0.194
	github.com/urfave/cli/v2 v2.25.5
	github.com/vektah/gqlparser/v2 v2.5.8
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/sqlite v1.5.3
	gorm.io/gorm v1.25.2-0.20230530020048-26663ab9bf55
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/ImVexed/fasturl v0.0.0-20230304231329-4e41488060f3 // indirect
	github.com/MichaelMure/go-term-text v0.3.1 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230717121422-5aa5874ade95 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alecthomas/chroma v0.7.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/c-bata/go-prompt v0.2.6 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91 // indirect
	github.com/eliukblau/pixterm/pkg/ansimage v0.0.0-20191210081756-9fb6cf8c2f75 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/cors v1.4.0 // indirect
	github.com/gin-contrib/requestid v0.0.6 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-contrib/zap v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.4.1 // indirect
	github.com/go-git/go-git/v5 v5.8.1 // indirect
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gomarkdown/markdown v0.0.0-20191123064959-2c17d62f5098 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kyokomi/emoji/v2 v2.2.8 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/libs4go/crypto v0.0.1 // indirect
	github.com/libs4go/errors v0.0.3 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.15.1 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/pyroscope-io/client v0.7.2 // indirect
	github.com/pyroscope-io/godeltaprof v0.1.2 // indirect
	github.com/pyroscope-io/otel-profiling-go v0.4.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rung/go-safecast v1.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/skeema/knownhosts v1.2.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/synapsecns/sanguine/ethergo v0.0.2 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.21 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.2 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/contrib v1.16.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.42.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.42.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.39.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.25.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/exp v0.0.0-20230127193734-31bee513bff7 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/tools v0.9.3 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
