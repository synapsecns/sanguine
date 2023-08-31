module github.com/synapsecns/sanguine/contrib/devnet

go 1.19

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/ethereum/go-ethereum v1.10.26
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	golang.org/x/sync v0.3.0
)

replace (
	github.com/synapsecns/sanguine/core latest => ../../core
	github.com/synapsecns/sanguine/ethergo latest => ../../ethergo
)

require github.com/go-stack/stack v1.8.1 // indirect
