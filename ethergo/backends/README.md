## Backends

Ethergo supports a wide variety of backends for testing contract and chain interactions. Which one you use depends on your needs. There is usually a trade-off between speed, requirements and visibility.

| Backend                  | Description                                                                                                                                    | Supports RPC Address | Supports Forking | Embedded (does not require Docker) | Speed                  |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|------------------|------------------------------------|------------------------|
| [Anvil](./anvil)         | Anvil is a dockerized-backend that supports a wide variety of options including forking, custom gas pricing, etc all configurable at boot time | ✅                    | ✅                | ❌                                  | Slow boot, fast run    |
| [Geth](./geth)           | Geth is an embedded [go-ethereum](https://github.com/ethereum/go-ethereum) node. This is the equivelant of the `geth --dev` command.           | ✅                    | ❌                | ✅                                  | Fastish boot, fast run |
| [Simulated](./simulated) | Geth [simulated backend](https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/bind/backends/simulated.go)                          | ✅                    | ❌                | ✅                                  | Practically instant    |
