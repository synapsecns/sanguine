# OFT deployment scripts

Run these commands from `packages/contracts-adapter`. Both scripts use `SynapseScript` and the existing `fsr` deployment runner, which selects the chain and wallet and saves deployment artifacts. Configure the wallet and `<CHAIN>_RPC` variables using `.env.example`; optional verifier variables follow the same conventions. Run without `--broadcast` to simulate first.

## Factory

Both scripts read `DEPLOY_ENVIRONMENT` from `.env`, defaulting to `production`. Production uses `configs/global/SynapseOFTAdapter.json`; `testnet` uses `configs/global/testnet/SynapseOFTAdapter.json`. Its `endpoints` map contains the LayerZero EndpointV2 address for each chain. Factory deployment uses `endpoints[activeChain]`.

```bash
npx fsr script/DeploySynapseOFTAdapterFactory.s.sol <chain> <walletName>
npx fsr script/DeploySynapseOFTAdapterFactory.s.sol <chain> <walletName> --broadcast
```

This deploys the factory with zero salt and `msg.sender` (the broadcast wallet) as its initial owner via the repository's CREATE2 deployer, saves it as `SynapseOFTAdapterFactory`, and initializes its endpoint. A saved deployment is reused, with its address and owner checked against the deployment parameters and its endpoint checked against the configuration. An uninitialized saved factory is initialized on rerun.

For matching factory addresses across chains, use the same broadcast wallet, CREATE2 deployer address, and compiled creation bytecode. The factory salt is always zero. The endpoint can differ. The CREATE2 deployer must already exist on each chain; the existing `Create2Factory` deployment artifact or `CREATE2_FACTORY` environment variable can select it.

## Adapter

Add each token to the corresponding environment's global config under `tokens`. The salt is shared across chains; `addresses` maps each chain to its token address. For example:

```json
{
  "endpoints": {
    "ethereum": "0x1a44076050125825900e736c501f859c50fe728c"
  },
  "tokens": {
    "SYN": {
      "salt": "0x0000000000000000000000000000000000000000000000000000000000000000",
      "addresses": {
        "ethereum": "<SynapseERC20 token address>"
      }
    }
  }
}
```

The mainnet config includes Ethereum and HyperEVM endpoints; the testnet config includes Ethereum Sepolia and HyperEVM testnet endpoints. Token entries must be populated before running the adapter script.

```bash
npx fsr-str script/DeploySynapseOFTAdapter.s.sol <chain> <walletName> SYN
npx fsr-str script/DeploySynapseOFTAdapter.s.sol <chain> <walletName> SYN --broadcast
```

Use the factory owner's wallet. The script loads the saved `SynapseOFTAdapterFactory`, calls its owner-only `deploy(token, salt)`, and saves the adapter as `SynapseOFTAdapter.SYN` with empty constructor arguments. Reruns reuse the saved adapter and check its predicted address, token, endpoint, and owner.

Use the same adapter salt for the same token across chains, and a distinct salt for each token served by a factory. Matching addresses require the same factory address and adapter creation bytecode; token and endpoint addresses can differ.

After deployment, grant the adapter the token's mint permission and configure its LayerZero peers and messaging settings before bridging.

## Testnet commands

Set `DEPLOY_ENVIRONMENT=testnet` in `.env`, then use the same entry points:

```bash
npx fsr script/DeploySynapseOFTAdapterFactory.s.sol ethereum_sepolia <walletName>
npx fsr-str script/DeploySynapseOFTAdapter.s.sol ethereum_sepolia <walletName> SYN
```

Use `hyperevm_testnet` for the other testnet. Append `--broadcast` to submit transactions after simulation.

## Wiring

After deploying the adapters on each chain and saving their artifacts, run:

```bash
npx fsr-str script/WireSynapseOFTAdapter.s.sol <chain> <walletName> SYN
```

`DEPLOY_ENVIRONMENT` selects the same global config used for deployment. The script wires only chains in `tokens.SYN.addresses`, resolving peers from their saved `SynapseOFTAdapter.SYN` artifacts. Peer changes are skipped when a remote deployment is missing; unsupported LayerZero routes are skipped entirely. Append `--broadcast` to submit transactions after simulation.

The global config's `wiring` object contains:

- `chains`: LayerZero endpoint IDs (`eid`) and `sendUln302` / `receiveUln302` library addresses, keyed by chain.
- `blockConfirmations`: source-chain confirmation counts, keyed by chain. A send configuration uses the local count; a receive configuration uses the remote chain's count.
- `requiredDVNs`: arrays of required DVN addresses on each chain. The script sorts these addresses before configuring ULN.

The checked-in `blockConfirmations` and `requiredDVNs` maps are intentionally empty in both environments. Fill in the chosen policy before wiring: confirmation counts for every selected chain and required DVNs for the active chain. OFT wiring explicitly disables optional DVNs instead of inheriting the library defaults. A chain with no configured required DVNs or confirmation count is rejected. Library addresses and endpoint IDs are taken from LayerZero's [deployment metadata](https://metadata.layerzero-api.com/v1/metadata/deployments).

Matching configuration is skipped on reruns. When the wallet is not the app owner, peer changes are printed as calldata; when it is not the endpoint delegate, library and security changes are printed as calldata. Submit that calldata through the corresponding owner or delegate account.

The wiring script configures peers, libraries, and ULN security. OFT send callers still supply their LayerZero execution options.
