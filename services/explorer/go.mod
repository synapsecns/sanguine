module github.com/synapsecns/sanguine/services/explorer

go 1.19

replace (
	github.com/Yamashou/gqlgenc => github.com/synapsecns/gqlgenc v0.10.0-hotfix
	// later versions give errors on uint64 being too high.
	github.com/brianvoe/gofakeit/v6 => github.com/brianvoe/gofakeit/v6 v6.9.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/synapsecns/sanguine/core => ./../../core
	github.com/synapsecns/sanguine/ethergo => ./../../ethergo
	github.com/synapsecns/sanguine/tools => ./../../tools
)

require (
	github.com/MichaelMure/go-term-markdown v0.1.4
	github.com/brianvoe/gofakeit/v6 v6.18.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.10.23
	github.com/hashicorp/consul/sdk v0.11.0
	github.com/jftuga/ellipsis v1.0.0
	github.com/jftuga/termsize v1.0.2
	github.com/stretchr/testify v1.8.0
	github.com/synapsecns/sanguine/tools v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.14.1
	golang.org/x/exp v0.0.0-20220827204233-334a2380cb91
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/MichaelMure/go-term-text v0.3.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/alecthomas/chroma v0.7.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/c-bata/go-prompt v0.2.6 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91 // indirect
	github.com/eliukblau/pixterm/pkg/ansimage v0.0.0-20191210081756-9fb6cf8c2f75 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/gen2brain/beeep v0.0.0-20210529141713-5586760f0cc1 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-toast/toast v0.0.0-20190211030409-01e6764cf0a4 // indirect
	github.com/godbus/dbus/v5 v5.0.6 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/gomarkdown/markdown v0.0.0-20191123064959-2c17d62f5098 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/gopherjs/gopherwasm v1.1.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kyokomi/emoji/v2 v2.2.8 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/nanmu42/etherscan-api v1.4.0 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/synapsecns/sanguine/core v0.0.0-00010101000000-000000000000 // indirect
	github.com/tadvi/systray v0.0.0-20190226123456-11a2b8fa57af // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/image v0.0.0-20220722155232-062f8c9fd539 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220906165534-d0df966e6959 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)
