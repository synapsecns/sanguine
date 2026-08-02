# SYN delivery to HyperCore

<!-- markdownlint-disable MD013 -->

Tracking: [CALL-2439](https://linear.app/protochain/issue/CALL-2439/implement-native-synapsebridge-delivery-of-syn-to-hypercore)

This package contains the contract boundary for delivering SYN through the existing
[SynapseBridge](https://github.com/synapsecns/synapse-contracts) supply model to a linked
[HIP-1 asset](https://hyperliquid.gitbook.io/hyperliquid-docs/hyperliquid-improvement-proposals-hips/hip-1-native-token-standard).
SYN does not become an OFT.

## Scope of this change

- **`SynapseBridgeAdapterV2` is a separate deployment.** The deployed `SynapseBridgeAdapter` contract and its
  CREATE2-derived addresses are unchanged.
- **Direct EVM delivery stays supported.** The original 96-byte message is decoded exactly as before.
- **HyperCore delivery uses a versioned message.** The source adapter rejects dust and `uint64` overflow before it
  burns or escrows SYN.
- **`SynapseHyperCoreComposer` is SYN-specific.** Both contracts enforce 18 HyperEVM decimals, 8 HyperCore wei
  decimals, and an amount scale of `10^10`.
- **No deployment or production configuration is included.** Token, bridge, Core index, peer, DVN, executor, and
  owner addresses must not be guessed before audited artifacts exist.

## Intended flow

1. `SynapseBridgeAdapterV2.bridgeERC20ToHyperCore` verifies the configured SYN route and amount.
2. The source adapter burns mint-and-burn SYN, or escrows an underlying token in `SynapseBridge`.
3. LayerZero delivers the versioned message to the authenticated destination adapter peer.
4. The HyperEVM `SynapseBridge` mints or withdraws SYN to `SynapseHyperCoreComposer`.
5. The Composer checks its immutable adapter and token bindings, both required HyperCore accounts, exact decimal
   representation, `uint64` bounds, and current asset-bridge capacity.
6. The Composer transfers HyperEVM SYN to the token system address, then submits the versioned spot-send action to
   [CoreWriter](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/hyperevm/interacting-with-hypercore).

## Safety invariants

- **No dust is burned.** Every source amount must be divisible by `10^10` before source custody changes.
- **No silent fallback exists.** A failed HyperCore delivery never changes the recipient to a HyperEVM address.
- **Immediate destination failures roll back the mint.** Composer validation, ERC20 transfer, and CoreWriter call
  occur inside the LayerZero receive transaction. A revert rolls back the destination `SynapseBridge` mint or
  withdrawal, leaving the LayerZero message available for retry.
- **Capacity is checked before transfer.** The Composer reads the Core-side system balance and rejects an amount that
  is not currently backed by available HIP-1 bridge capacity. It also reserves capacity across all deliveries in the
  same HyperEVM block because those transactions read the same pre-transfer HyperCore state.
- **Bindings are immutable.** Decimal routes and Composer bindings cannot be overwritten after configuration.
- **Existing bridge behavior is isolated.** The v1 adapter source and deployed bytecode remain unchanged.

## Atomicity boundary

An accepted HyperEVM transaction is not proof that the spot-send action executed on HyperCore. Hyperliquid processes
EVM-to-Core transfers and CoreWriter actions after the EVM block. The Composer checks the conditions that would make
the action fail, including account activation and bridge capacity, before it emits the transfer and CoreWriter call.
Production monitoring must still confirm all of the following for every canary:

- the LayerZero GUID was delivered once;
- the HyperEVM transaction succeeded;
- the asset-bridge `Transfer` event credited the Composer on HyperCore;
- the CoreWriter action executed successfully;
- the final recipient received the exact 8-decimal Core amount;
- the Composer has no stranded HyperCore or HyperEVM SYN balance.

See [Hyperliquid interaction timings](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/hyperevm/interaction-timings)
and the [LayerZero Composer implementation](https://github.com/LayerZero-Labs/devtools/blob/main/packages/hyperliquid-composer/contracts/HyperLiquidComposer.sol).

## Deployment gates

Do not deploy or link the HIP-1 token until each gate is resolved:

1. **HyperEVM SYN contract:** use the existing `SynapseERC20` authority model with `SynapseBridge` as its only minter.
2. **HyperCore finalizer:** a factory-deployed ERC20 must expose the authorized finalizer in the first storage slot or
   the `keccak256("HyperCore deployer")` slot. Otherwise deploy from the finalizer EOA and retain the exact creation
   nonce required by `finalizeEvmContract`.
3. **Full bridge capacity:** fund the HyperCore asset bridge with the complete intended circulatable supply during
   genesis. Do not partially top it up later.
4. **Composer activation:** activate the deployed Composer account on HyperCore before enabling any source route.
5. **Peer isolation:** V2 peers used for HyperCore delivery must point only to the reviewed V2 deployments.
6. **Independent review:** audit the token authority, source custody action, message codec, destination mint,
   precompile reads, system-address transfer, and CoreWriter encoding as one supply path.

The HyperCore token index, HyperEVM SYN address, SynapseBridge address, Composer address, and all LayerZero security
configuration remain intentionally unset in this change.
