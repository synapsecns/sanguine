# SYN OFT deployment and configuration

Run commands from `packages/contracts-adapter`. The scripts use `SynapseScript` and the existing deployment runners: `fsr` calls `run()`, and `fsr-str` calls `run(string)` with a token identifier such as `SYN`. Configure the wallet and chain RPC variables using `.env.example`. Run without `--broadcast` to simulate; append it to submit transactions.

`DEPLOY_ENVIRONMENT` in `.env` selects the configuration:

| Environment | Config | Deployment chains with SYN token addresses |
| --- | --- | --- |
| `production` (default) | `configs/global/SynapseOFTAdapter.json` | `ethereum`, `hyperevm` |
| `testnet` | `configs/global/testnet/SynapseOFTAdapter.json` | `ethereum_sepolia`, `hyperevm_testnet` |

Production also includes security routes for Arbitrum, Avalanche, Base, BNB, Optimism, and Polygon. Those route entries do not imply token or adapter deployments on those chains. Saved deployment artifacts live under `deployments/<chain>/`.

## Factory

```bash
npx fsr script/DeploySynapseOFTAdapterFactory.s.sol <chain> <walletName>
```

The script deploys `SynapseOFTAdapterFactory` through the repository's CREATE2 deployer, with zero salt and the broadcast wallet as initial owner. It then initializes the factory's endpoint from `endpoints[activeChain]`. A saved factory is reused and checked against the predicted address and owner; its endpoint is initialized if unset, otherwise checked against the config.

Matching factory addresses across chains require the same initial owner, CREATE2 deployer, and compiled creation bytecode. The endpoint can differ because initialization happens after deployment. The CREATE2 deployer must already exist on each chain; select it through a `Create2Factory` deployment artifact or `CREATE2_FACTORY`.

## Adapter

```bash
npx fsr-str script/DeploySynapseOFTAdapter.s.sol <chain> <walletName> SYN
```

Use the factory owner's wallet. The script reads `tokens.SYN.addresses[activeChain]` and `tokens.SYN.salt`, calls the saved factory's `deploy(token, salt)`, and saves `SynapseOFTAdapter.SYN` with empty constructor arguments. The adapter obtains its token, endpoint, and initial owner through the factory callback.

The adapter salt is configured separately from the factory's fixed zero salt. Use the checked-in salt for existing deployments. Matching adapter addresses require the same factory address, adapter salt, and compiled adapter creation bytecode; token and endpoint addresses can differ. Different tokens deployed through the same factory need distinct salts.

Reruns reuse the saved adapter and check its predicted address, token, endpoint, and ownership. The script requires both factory and adapter ownership to match the broadcast wallet, so it is not a post-handoff verification command once the adapter belongs to a multisig.

Grant the adapter permission to call the underlying token's `mint(address,uint256)`. Senders must approve the adapter for `burnFrom`. Configure the LayerZero routes before bridging.

For testnet, set `DEPLOY_ENVIRONMENT=testnet` and use the same commands with `ethereum_sepolia` and `hyperevm_testnet`.

## Wiring and security

Run on each chain where the local adapter is deployed and its token address is configured:

```bash
npx fsr-str script/WireSynapseOFTAdapter.s.sol ethereum <walletName> SYN
npx fsr-str script/WireSynapseOFTAdapter.s.sol hyperevm <walletName> SYN
```

The script selects routes from `wiring.chains`, independently of remote token addresses. It configures send and receive libraries and ULN security even when a remote adapter has not been deployed. Peer addresses come from saved `SynapseOFTAdapter.SYN` artifacts; missing remote artifacts skip only the peer step. Routes unsupported by either local message library are skipped entirely.

The config contains:

- `endpoints`: local LayerZero EndpointV2 addresses.
- `wiring.chains`: endpoint IDs (`eid`) and local `sendUln302` / `receiveUln302` library addresses.
- `wiring.blockConfirmations`: confirmation requirements for messages originating on each chain.
- `wiring.requiredDVNs`: local addresses of the required DVNs. The script sorts them and rejects empty, duplicate, or zero-address entries. All required DVNs must verify a message; optional DVNs are explicitly disabled.

The checked-in production policy requires **LayerZero Labs and Nethermind** on all eight chains:

| Source chain | Confirmations |
| --- | ---: |
| Ethereum | 64 |
| HyperEVM | 100 |
| Arbitrum | 100 |
| Avalanche | 100 |
| Base | 100 |
| BNB | 100 |
| Optimism | 100 |
| Polygon | 200 |

Testnet requires **LayerZero Labs and Mantle01**, with **5 confirmations** from either Ethereum Sepolia or HyperEVM testnet. Library addresses, endpoint IDs, and DVN addresses are based on [LayerZero deployment metadata](https://metadata.layerzero-api.com/v1/metadata/deployments).

Send configuration uses the local chain's confirmation count. Receive configuration uses the remote source chain's count. For example, HyperEVM's receive configuration for Ethereum requires 64 confirmations, while Ethereum's receive configuration for HyperEVM requires 100.

Matching settings are skipped on reruns. If the wallet is not the adapter owner, peer changes are printed as multisig calldata. If it is not the endpoint delegate, library and security changes are printed as multisig calldata. Submit that calldata through the corresponding authority.

Wiring configures execution budgets from `wiring.enforcedOptions`, keyed by destination chain. Both `SEND` and `SEND_AND_CALL` enforce 200,000 `lzReceive` gas. Sends with composition to HyperEVM also enforce 250,000 `lzCompose` gas at index 0; other destinations have no enforced compose budget. Native value is zero. The testnet config uses the same budgets for Sepolia and HyperEVM testnet. These updates require the adapter owner; the script prints multisig calldata when run by another wallet and skips settings that already match.

Caller-supplied `extraOptions` add to these enforced budgets. Use empty extra options when the configured budget is sufficient, and supply additional compose gas for composers on other destinations.

## HyperCore composer

`SynapseComposer` extends LayerZero's recovery-enabled Hyperliquid composer. It uses the underlying SynapseERC20 token, its linked HyperCore index, the difference between ERC20 decimals and Core `weiDecimals`, and an authorized recovery address.

These parameters are already populated under `tokens.SYN.composer` in each environment's config:

| Environment | Core token index | Decimal difference |
| --- | ---: | ---: |
| Production | 873 | 10 |
| Testnet | 3071 | 10 |

The decimal difference is 18 ERC20 decimals minus 8 Core wei decimals. `recoveryAddress` is configured separately in each environment and controls recovery of composer-held assets.

```bash
npx fsr-str script/DeploySynapseComposer.s.sol hyperevm <walletName> SYN
```

For testnet, set `DEPLOY_ENVIRONMENT=testnet` and select `hyperevm_testnet`. The script uses regular CREATE, loads the saved `SynapseOFTAdapter.SYN`, and saves `SynapseComposer.SYN` with constructor arguments for verification. It rejects chains other than HyperEVM mainnet and testnet, verifies the adapter's token and endpoint, and checks the composer settings when reusing an existing deployment.

Before use, link the underlying token to HyperCore, fund its asset bridge, and activate the composer and recipient accounts on HyperCore. Deployment does not perform these steps. See the [LayerZero Hyperliquid guide](https://docs.layerzero.network/v2/developers/hyperliquid/hyperliquid-oft-deployment) for token setup and linking.

For inbound transfers, use the composer as the OFT destination and encode `composeMsg = abi.encode(uint256(0), recipient)`. Supply options funding both `lzReceive` and `lzCompose` at index 0. A synchronous Core transfer failure can refund the tokens to the recipient on HyperEVM. Outbound HyperCore transfers require a native transfer to HyperEVM followed by an OFT send.

### Source refunds and manual recovery

If the compose payload cannot be decoded, the composer records `failedMessages[guid]` and keeps the tokens on HyperEVM. Its `refundToSrc(guid)` method sends them back through the adapter, using the constructor's token allowance. This path requires a wired reverse route, sufficient native messaging fees, and enforced `SEND` receive-gas options on the HyperEVM adapter for the source endpoint. The composer supplies empty extra options, and adding native funds alone cannot replace the missing options.

Without enforced options, the recovery address can call `recoverEvmERC20(amount)` to withdraw the affected tokens, then transfer them to the user on HyperEVM or approve the adapter and bridge them back with explicit execution options. Manual recovery does not clear `failedMessages[guid]`: track settled GUIDs to prevent a later source-refund call from consuming other composer-held funds. Core asset retrieval is also available through the inherited recovery functions.

## Ownership and endpoint delegate

After deployment and configuration, transfer both authorities to the chain's multisig:

```bash
npx fsr-str script/TransferOwnershipSynapseOFTAdapter.s.sol <chain> <currentOwnerWallet> SYN
```

The script reads `multisig[activeChain]` from the selected environment's `SynapseOFTAdapter.json`. Production has entries for Ethereum and HyperEVM. Testnet multisig entries must be added before using this script there. The target must be a nonzero contract address.

The script verifies the adapter's token and endpoint, sets the endpoint delegate through `adapter.setDelegate(multisig)`, then transfers adapter ownership. Delegate transfer comes first because `setDelegate` is owner-only. These are separate transactions when both changes are needed; reruns skip completed settings and verify the final state. If both authorities already match, the script exits without changes.

Factory ownership and the composer's immutable recovery address are separate from this handoff. Future peer and enforced-option updates require the adapter owner; the wiring script’s endpoint library and ULN updates require the endpoint delegate.
