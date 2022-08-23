// Package simulated represents a simulated backend. It complies to the standard chain
// backend and is used for any tests that don't require a websocket (see https://github.com/ethereum/go-ethereum/issues/21457)
// It's relatively useful for any tests that don't require a full backend (often integration tests). Beyond the very
// real limitations fo this backend (json rpc, different mempool mechanics, etc) that make it differ from geth, geth also
// introduces some superficial limitations on us that we want to get around. Geth isn't made for multichain mocking so
// they require the use of chain id 1337. ToAddress get around this we use go:generate to copy the backend using multichainsimulation.
//
// We don't make any changes to the fundamental backend simulation, but we do use the
package simulated
