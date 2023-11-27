/*
Package fetchers provides fetchers to fetch data from Scribe, BridgeConfig, Swap, DeFiLlama and more.
/swap: The swap fetcher's purpose is to map swap ids to token addresses from swap events. With the token address,
we can get the token data from the erc20 contract. There is an LRU cache to store these mappings to reduce RPC calls.
/price: The price fetcher's purpose is to get the price of a token at a given timestamp from DeFiLlama.
There is an LRU cache to store these mappings to reduce RPC calls.
/scribe: The scribe fetcher is a wrapper around of the scribe gql client. All data from scribe is pulled using this
fetcher.
/token: The token fetcher's purpose is to get the token data from the bridge config, erc20 contract, etc. depending
on the type event. There is an LRU cache to store these mappings to reduce RPC calls.
*/
package fetchers
