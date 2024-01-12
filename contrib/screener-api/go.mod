module github.com/synapsecns/sanguine/contrib/screener-api

go 1.21

replace github.com/synapsecns/sanguine/core => ../../core

require (
	github.com/go-resty/resty/v2 v2.11.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
)

require golang.org/x/net v0.17.0 // indirect
