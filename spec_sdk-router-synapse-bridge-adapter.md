# SDK Router Support for SynapseBridgeAdapter

## Goal

Add `SynapseBridgeAdapter` support to `@synapsecns/sdk-router` as a `bridgeV2`-capable module for direct ERC20 bridging only: the origin token must already be SBA-supported on the origin chain, the bridge lands into the single remote token configured by SBA, `bridgeV2()` returns direct no-origin-swap routes, `intent({ allowMultipleTxs: true })` may append a destination-side swap after the SBA bridge step, and completion tracking follows the same LayerZero-style pattern already used by the OFT module.

## Context

- `bridgeV2` is already implemented in `packages/sdk-router/src/operations/bridge.ts`, but it only considers module sets where `isBridgeV2Supported === true`.
- Existing V2-capable modules implement the `SynapseModuleSet` contract in `packages/sdk-router/src/module/synapseModuleSet.ts` and hand off populated transactions to `SynapseIntentRouterSet.finalizeBridgeRouteV2()`.
- `SynapseBridgeAdapter` is not router-like. It exposes `bridgeERC20(dstEid, to, token, amount, gasLimit)`, `getNativeFee(dstEid, gasLimit)`, and point lookups for token mappings, but it does not expose route enumeration or swap-query construction.
- Under the accepted scope for this spec, SBA support is limited to non-swap origin routes. That means the SDK only needs to answer whether the user’s `fromToken` is directly bridgeable to the destination chain; the origin adapter can answer that via `getRemoteAddress(remoteEid, fromToken)`.
- The adapter contract has a hard minimum LayerZero gas limit of `200_000`, and the user chose to keep phase 1 on that hardcoded minimum instead of adding a new public SDK override.
- The user explicitly scoped out legacy bridgeV1 support. `allBridgeQuotes()`, `bridgeQuote()`, legacy `Query` objects, and legacy `bridge()` routing are not part of this feature.
- SBA deployment and chain metadata already exists elsewhere in the monorepo:
  - `packages/contracts-adapter/deployments/*/SynapseBridgeAdapter.json`
  - `packages/contracts-adapter/deployments/*/.chainId`
  - `packages/contracts-adapter/configs/global/chains.json`
- `sdk-router` is a published package and must remain self-contained at runtime. SBA metadata cannot be discovered by reading sibling workspace files after publish.
- Current repository constraints limit SBA `bridgeV2` origins to chains that have both SBA deployments and intent infrastructure. Based on the checked-in configs, those origin chains are `ETH`, `OPTIMISM`, `BSC`, `POLYGON`, `BASE`, `ARBITRUM`, `AVALANCHE`, and `BLAST`.
- Current repository constraints limit SBA destinations to active, bridge-supported chains that also have SBA deployments. Based on the checked-in configs, those destination chains are `ETH`, `OPTIMISM`, `BSC`, `POLYGON`, `FANTOM`, `METIS`, `CANTO`, `KLAYTN` (`kaia` in contracts-adapter configs), `BASE`, `ARBITRUM`, `AVALANCHE`, `DFK`, and `BLAST`.
- The OFT (`USDT0`) module already shows the preferred LayerZero integration pattern for a V2-only module: it is not part of legacy bridge routing, it returns `txHash` from `getSynapseTxId()`, and it checks completion through the LayerZero scan API.
- The existing V2 quote pipeline still starts from `getBridgeTokenCandidates()` and runs the swap engine set. Under this scope, SBA should only ever yield the direct input token as the bridge candidate, so the origin leg resolves through the existing `NoOpEngine` path with zero origin steps.

## Scope

- Add a new SBA-specific bridgeV2-only module set to `packages/sdk-router`.
- Bundle SBA chain, endpoint-ID, and deployment-address metadata inside `packages/sdk-router/src` as committed source.
- Make SBA routes available to `SynapseSDK.bridgeV2()` and to cross-chain `SynapseSDK.intent()`.
- Support direct SBA-backed quotes where `toToken` equals the mapped remote token.
- Support SBA as the bridge step inside `intent({ allowMultipleTxs: true })`, where the bridge lands into the mapped remote token and the existing destination-side intent flow may add a follow-up swap.
- Quote and surface the adapter’s LayerZero `nativeFee`.
- Populate origin-chain transactions through the existing SIR / TokenZap path when `fromSender` is provided.
- Add SBA tracking support to `getSynapseTxId()` and `getBridgeTxStatus()`, modeled after the OFT module’s LayerZero behavior.
- Add unit and integration-style test coverage for the new module and the new bridgeV2 path.
- Update SDK documentation for the new module and its constraints.

## Non-goals

- No support for SBA in `bridgeQuote()`, `allBridgeQuotes()`, `applyBridgeDeadline()`, `applyBridgeSlippage()`, or the legacy `bridge()` path.
- No new public SDK parameter for overriding SBA gas limit.
- No origin-side swap support for SBA. If the input token is not directly bridgeable through SBA, no SBA quote should be returned.
- No native-token SBA entrypoint support. SBA itself only bridges ERC20 tokens.
- No dynamic on-chain route enumeration. The contract does not expose it.
- No widget or `synapse-interface` UI integration work.
- No automation requirement for regenerating SBA chain/deployment metadata from `contracts-adapter`; phase 1 may check in static metadata with comments naming the source files.

## Requirements

1. `SynapseSDK` must instantiate and register a new module set named `SynapseBridgeAdapter` in `allModuleSets`.
2. The SBA module set must report `isBridgeV2Supported = true`.
3. The SBA module set must expose `allEvents = ['TokenSent', 'TokenReceived']` so `getBridgeModuleName()` can identify SBA events.
4. SBA must not contribute any legacy `BridgeQuote` results. Its `getBridgeRoutes()` implementation must return `[]`, and any legacy `bridge()` implementation on the concrete module must throw a clear `bridge V1 not supported` error.
5. SBA chain/deployment discovery must use committed SDK-local chain metadata derived from the checked-in contracts-adapter deployment/config data, not runtime filesystem reads.
6. The chain metadata must normalize chain-name mismatches between packages, including `bnb -> SupportedChainId.BSC` and `kaia -> SupportedChainId.KLAYTN`.
7. `getBridgeTokenCandidates()` must:
   - Return `[]` when the SBA module is missing on either chain.
   - Resolve `remoteEid` from the destination chain and call `originAdapter.getRemoteAddress(remoteEid, fromToken)`.
   - Return `[]` when `getRemoteAddress(...)` returns `address(0)`.
   - Return `[]` when `toToken` is provided and does not equal the returned remote token.
   - Return a single direct candidate where `originToken = fromToken` and `destToken = remoteToken`.
8. `_collectV2Quotes()` must continue to use the existing V2 candidate flow, but SBA candidates must only produce no-op origin routes. Any origin route that requires an actual swap is out of scope for SBA.
9. `getBridgeRouteV2()` for SBA must:
    - Reject invalid candidate / chain / token combinations by returning `undefined`.
    - Reject any origin route with non-empty `steps`; SBA only supports direct no-op origin routes.
    - Treat SBA as a 1:1 bridge of the direct input token. `expectedToAmount` must equal the no-op origin route amount.
    - Set `minToAmount` equal to `expectedToAmount` for the SBA bridge step.
    - Set `gasDropAmount` to zero.
    - Set bridge fee amount to zero.
    - Quote `nativeFee` by calling `adapter.getNativeFee(dstEid, 200_000)`.
10. SBA must only produce single-transaction `bridgeV2()` quotes when the requested `toToken` is the exact remote SBA token for the chosen pair.
11. SBA must be eligible inside `intent({ allowMultipleTxs: true })` even when the requested final token differs from the SBA remote token. In that case the SBA step lands into the mapped remote token and the existing destination-side intent flow is responsible for the follow-up swap.
12. When `fromSender` and `toRecipient` are both present, SBA must return `zapData` that makes `SynapseIntentRouterSet.finalizeBridgeRouteV2()` build a valid SIR transaction.
13. The SBA zap payload must target the origin SBA contract and encode `bridgeERC20(dstEid, toRecipient, originBridgeToken, amount, 200_000)`.
14. The SBA zap payload must set `amountPosition` to the calldata slot for `bridgeERC20.amount` so TokenZap can patch the bridged amount at execution time.
15. When `fromSender` is absent, SBA quotes may still be returned, but `tx` must remain undefined just like other bridgeV2 modules.
16. `BridgeQuoteV2.moduleNames` must include `SynapseBridgeAdapter`, and SBA quotes must not include origin swap engine names because origin swaps are out of scope.
17. `BridgeQuoteV2.routerAddress` for SBA quotes must remain the SIR address for the origin chain, not the SBA contract address, because the public transaction still enters through SIR.
18. `getSynapseTxId()` for SBA must return the origin transaction hash unchanged, matching the OFT module’s LayerZero semantics.
19. `getBridgeTxStatus()` for SBA must use the LayerZero scan API and the same completion states already accepted by the OFT module (`CONFIRMING`, `DELIVERED`).
20. SBA `estimatedTime` must follow the OFT-style LayerZero model:
    - The concrete SBA module should be able to query pathway-specific confirmation counts from the LayerZero API when cache is cold.
    - The module set must expose a synchronous `getEstimatedTime()` by serving cached values and falling back to a static approximation when the cache is empty.
21. The fallback SBA estimated-time approximation must use committed chain timing/security metadata derived from the contracts-adapter configs, not hardcoded one-off guesses per pair.
22. No public SDK type signature may change for this feature. `BridgeV2Parameters` and `IntentParameters` remain unchanged.

## Technical approach

- Add a new `packages/sdk-router/src/sba/` module namespace.
- Create a concrete `SynapseBridgeAdapterModule` class that implements `SynapseModule`.
- Create a `SynapseBridgeAdapterModuleSet` class that extends `SynapseModuleSet`.
- Add committed SBA chain metadata under `packages/sdk-router/src/constants/` or `packages/sdk-router/src/sba/` with:
  - SBA deployment address per chain ID.
  - LayerZero endpoint ID per chain ID.
  - Any chain-level timing/confirmation defaults needed for fallback ETA calculation.
- Source the first metadata snapshot from the contracts-adapter files listed in the Context section, but do not read them at runtime from published SDK code.
- Model the concrete module after the OFT module where behavior is LayerZero-specific:
  - `getSynapseTxId(txHash) => txHash`
  - `getBridgeTxStatus(txHash)` via LayerZero scan API
  - pathway-based ETA refresh plus cache
- Model the module set integration after other bridgeV2-only modules:
  - `getBridgeRoutes()` returns `[]`
  - `getFeeData()` returns zeroes
  - `getDefaultPeriods()` returns zeroes because legacy deadlines are out of scope
  - `applySlippage()` returns inputs unchanged because legacy slippage helpers are out of scope
- `getBridgeTokenCandidates()` should use the origin SBA contract as a point lookup only: resolve `remoteEid`, call `getRemoteAddress(remoteEid, fromToken)`, and derive the single direct candidate from the returned remote token.
- `getBridgeRouteV2()` should:
  - Validate `bridgeToken`, module existence, and requested-token compatibility through the existing `validateBridgeRouteV2Params()` path.
  - Require a zero-step origin route and return `undefined` if the route implies any origin swap.
  - Fetch `nativeFee` from the origin adapter using `getNativeFee(dstEid, 200_000)`.
  - Reuse the direct input amount because SBA does not apply an additional bridge fee or origin swap.
  - Build zap data only when `fromSender` and `toRecipient` are present.
- The concrete module should expose a helper such as `populateBridgeERC20(params, nativeFee)` that returns the adapter calldata and `msg.value`.
- The SBA zap-data encoding should mirror the existing OFT and FastBridge patterns: `target`, `payload`, and `amountPosition` are sufficient. `finalToken`, `forwardTo`, and `minFinalAmount` can remain at their existing defaults.
- Because `SynapseIntentRouterSet.finalizeBridgeRouteV2()` already appends the bridge step after the origin route, no SIR contract changes are required. For SBA that origin route should be the zero-step no-op route.
- No changes are needed in `_collectV1Quotes()` besides the new module appearing in `allModuleSets`; it will already be excluded from legacy quote collection because `isBridgeV2Supported` is true.
- The module set should cache pathway ETA refreshes with a TTL, similar to `UsdtModuleSet`, to avoid repeated external API calls.
- The module set should not attempt bridge-token discovery from static token registries and should not attempt destination-side amount-out simulation beyond 1:1 transfer semantics. Any post-bridge swap belongs to the existing `intent()` multi-step flow.

## Affected areas

- `packages/sdk-router/src/sdk.ts`
- `packages/sdk-router/src/constants/addresses.ts`
- `packages/sdk-router/src/constants/index.ts`
- `packages/sdk-router/src/constants/medianTime.ts`
- `packages/sdk-router/src/sba/index.ts`
- `packages/sdk-router/src/sba/synapseBridgeAdapterModule.ts`
- `packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.ts`
- `packages/sdk-router/src/operations/bridge.ts`
- `packages/sdk-router/src/operations/intent.ts`
- `packages/sdk-router/src/module/index.ts`
- `packages/sdk-router/src/types/index.ts`
- `packages/sdk-router/README.md`
- New tests under `packages/sdk-router/src/sba/`
- Existing SDK-level tests for bridge flows

## Edge cases and failure handling

- If the origin chain is not intents-supported, SBA must produce no V2 quote even if the adapter is deployed there.
- If the destination chain is paused or otherwise unsupported by `isChainIdSupported()`, SBA must produce no V2 quote even if the adapter is deployed there.
- If the origin token has no SBA mapping for the destination chain, SBA must produce no quote.
- If the requested `toToken` does not equal the SBA remote token and `allowMultipleTxs` is false, SBA must produce no quote.
- If the V2 pipeline does not yield a zero-step origin route for the direct input token, SBA must produce no quote.
- If `getNativeFee()` reverts or the adapter lookup data is missing, SBA must produce no quote rather than a partially populated quote.
- If the LayerZero status API is unavailable, `getBridgeTxStatus()` must return `false` rather than throwing for routine polling flows, matching existing LayerZero-style modules.
- If the LayerZero ETA refresh fails, SBA must continue to return the static fallback ETA.
- If the checked-in chain/deployment metadata diverges from on-chain deployments, quote generation may silently under-report or omit routes. This is acceptable for phase 1 but must be documented as an operational risk.

## Phase plan

1. Add committed SBA metadata to `sdk-router`.
2. Add `SynapseBridgeAdapterModule` with LayerZero tracking helpers, calldata population, and ETA-refresh support.
3. Add `SynapseBridgeAdapterModuleSet` with `getRemoteAddress(...)`-based direct candidate discovery and bridgeV2 route generation.
4. Register the SBA module set in `SynapseSDK` and exports.
5. Add SBA-specific unit tests for direct candidate discovery, native-fee quoting, zap-data encoding, tx-id/status behavior, and ETA fallback/cache behavior.
6. Add SDK-level tests covering `bridgeV2()` direct SBA quotes and `intent({ allowMultipleTxs: true })` SBA bridge-step participation.
7. Update `README.md` to document the new module, its bridgeV2-only scope, hardcoded gas-limit policy, and LayerZero-tracking behavior.

## Acceptance criteria

- `SynapseSDK.bridgeV2()` can return SBA-backed quotes for supported direct token pairs on supported origin/destination chains.
- SBA quotes report `moduleNames` containing `SynapseBridgeAdapter`.
- SBA quotes report a non-zero `nativeFee` when the adapter quotes one.
- SBA quotes return `tx` only when `fromSender` is provided.
- SBA quotes use SIR as `routerAddress`.
- SBA quotes never include origin swap steps or origin swap module names.
- `SynapseSDK.intent()` can include SBA as the bridge step when `allowMultipleTxs` is enabled and a destination-side swap is required.
- `SynapseSDK.bridgeQuote()` and `SynapseSDK.allBridgeQuotes()` behavior is unchanged.
- `SynapseSDK.getSynapseTxId(originChainId, 'SynapseBridgeAdapter', txHash)` returns `txHash`.
- `SynapseSDK.getBridgeTxStatus(destChainId, 'SynapseBridgeAdapter', txHash)` queries LayerZero-style status and returns a boolean without requiring new public parameters.
- Unsupported chains, unsupported tokens, and mismatched final tokens return no SBA quote instead of malformed results.
- The SDK package builds and tests pass with the new SBA code included.

## Validation plan

- Add unit tests for SBA chain metadata normalization and direct candidate filtering rules.
- Add unit tests that mock `getRemoteAddress(...)` to prove SBA only returns direct input-token candidates and rejects unsupported or mismatched routes.
- Add unit tests for SBA calldata generation and verify the `amountPosition` used in zap data points to `bridgeERC20.amount`.
- Add unit tests that mock `getNativeFee()` and confirm the returned `nativeFee` is forwarded into both the quote and the populated transaction step.
- Add unit tests that mock LayerZero status responses and verify SBA completion logic matches the OFT module behavior.
- Add unit tests for ETA cache behavior and static fallback behavior.
- Add SDK-level tests that mock or stub SBA module behavior inside `bridgeV2()` and `intent()` flows so the new V2 path is covered directly.
- Run `yarn test` or the package-local SDK test command in `packages/sdk-router`.
- Run the package build for `packages/sdk-router`.
- Manually verify at least one real supported pair on a live provider by calling `bridgeV2()` with `fromSender` set and inspecting the populated SIR transaction for SBA calldata and non-zero `nativeFee`.

## Risks and assumptions

- Assumption: the current SBA destination handler remains compatible with the minimum LayerZero gas limit of `200_000`. If the adapter receive logic grows, this hardcoded default will need to be revisited.
- Assumption: SBA `getRemoteAddress(...)` remains the authoritative source for whether a direct input token is bridgeable to a specific destination chain.
- Assumption: using the OFT module’s LayerZero tx-hash tracking model is acceptable for SBA even though SBA also emits its own `TokenSent` / `TokenReceived` events.
- Risk: checked-in SBA chain/deployment metadata can drift from on-chain deployments even though token support itself is discovered live from the adapter.
- Risk: ETA and completion tracking depend on the external LayerZero scan API.
- Risk: bridgeV2 coverage in this package is currently thin. This feature should add direct tests instead of relying only on existing legacy bridge tests.
