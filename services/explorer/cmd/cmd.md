# Explorer

Explorer is an indexer and API serving platform analytics.

`./explorer placeholder`: A placeholder

## Directory Structure

<pre>
explorer
├── <a href="./cmd">cmd</a>: CLI
├── <a href="./config">config</a>: Config for the Explorer
├── <a href="./db">db</a>: Currently holds schema for db
├── <a href="./contracts">contracts</a>: Holds contracts generated via abigen
│   ├── <a href="./contracts/bridge">bridge</a>: SynapseBridge contract and more
│   ├── <a href="./contracts/bridgeconfig">bridgeconfig</a>: BridgeConfig contract and more
│   └── <a href="./contracts/swap">swap</a>: SwapFlashLoan contract and more
├── <a href="./internal">internal</a>: Workaround go-generate error, requiring file
└── <a href="./external">external</a>: External packages
└── └── <a href="./external/flattened-contracts">flattened-contracts</a>: Holds flattened contracts used by the /contracts generate.go files

</pre>
