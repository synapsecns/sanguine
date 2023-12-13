module github.com/synapsecns/sanguine/services/rfq

go 1.20


replace (
	github.com/Yamashou/gqlgenc => github.com/synapsecns/gqlgenc v0.10.0-hotfix
	// later versions give erros on uint64 being too high.
	github.com/brianvoe/gofakeit/v6 => github.com/brianvoe/gofakeit/v6 v6.9.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/soheilhy/cmux => github.com/lepsta/cmux v0.0.0-20221204213707-47c4a1bf4a43
	github.com/synapsecns/sanguine/core => ./../../core
	github.com/synapsecns/sanguine/ethergo => ./../../ethergo
	github.com/synapsecns/sanguine/services/omnirpc => ../omnirpc
	github.com/synapsecns/sanguine/tools => ./../../tools

)
