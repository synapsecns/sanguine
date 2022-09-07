module github.com/synapsecns/sanguine/services/explorer

go 1.19

replace github.com/synapsecns/synapse-contracts v0.0.0-20220822000752-397774c0ecad => ./external/synapse-contracts

require github.com/ethereum/go-ethereum v1.10.23

require (
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220906165534-d0df966e6959 // indirect
	golang.org/x/tools v0.1.12 // indirect
)
