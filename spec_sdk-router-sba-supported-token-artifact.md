# SDK Router SBA Supported-Token Artifact

## Goal

Add a committed hardcoded SBA supported-token artifact to `@synapsecns/sdk-router` and make it the deterministic source of truth for SBA bridge candidates, while removing SBA-local wrapped-native handling and relying on the existing swap engine path to reach SBA bridge tokens on the origin chain.

## Context

- `SynapseBridgeAdapterModuleSet.getBridgeTokenCandidates()` currently discovers SBA support by calling `adapter.getRemoteAddress(dstEid, token)` at runtime in [packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.ts#L82).
- The current SBA path also special-cases native input by asking `SynapseBridgeAdapterModule` for `getWrappedNativeToken()`, which is backed by `SwapQuoterV2.weth()` and an SBA-local cache in [packages/sdk-router/src/sba/synapseBridgeAdapterModule.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModule.ts#L53).
- `_collectV2Quotes()` already has the generic shape needed for origin-side conversion: it asks each bridgeV2 module for candidates, dedupes `originToken`, and uses `swapEngineSet.getBestQuote()` plus `generateRoute()` to route the requested `fromToken` into each candidate origin token in [packages/sdk-router/src/operations/bridge.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/operations/bridge.ts#L145).
- `NoOpEngine` already covers direct origin-token matches, and `DefaultPoolsEngine` already participates in the same path for intent-preview-based swaps, so origin-side native wrapping does not need SBA-specific logic when the candidate list already includes the wrapped token.
- SBA chain metadata is already shipped as committed SDK-local source in [packages/sdk-router/src/sba/metadata.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/metadata.ts). The published package cannot read sibling workspace files at runtime.
- The contracts-side SBA configuration is derived from `packages/contracts-adapter/configs/global/tokens.json`, `packages/contracts-adapter/configs/global/chains.json`, and the skip rules in `packages/contracts-adapter/script/DeploySBA.s.sol`. `DeploySBA.s.sol` currently skips tokens that are not deployed on multiple chains and explicitly excludes `GMX`.
- On March 18, 2026, a live verification sweep across the current SBA-supported chains using RPC endpoints selected from `https://chainlist.org/rpcs.json` found no mismatches for the supported-token snapshot implied by the repo configs:
  - all 149 included chain-token entries returned `getTokenType(token) ∈ {1,2}`
  - excluded `BUSD` on BNB returned `0`
  - excluded `GMX` on Arbitrum and Avalanche returned `0`
- That March 18, 2026 sweep is the manual verification pass for the snapshot to be hardcoded in `sdk-router`; this scope does not require SDK-side refresh or validation tooling.
- Current tests and README still describe the narrower phase-1 behavior where SBA normally returns direct origin routes and only has one origin-side exception for native gas token -> wrapped native.

## Scope

- Add a committed hardcoded SBA supported-token artifact under `packages/sdk-router/src/sba/`.
- Use that artifact for SBA bridge candidate discovery and SBA pair validation.
- Remove SBA-local wrapped-native discovery logic and its tests.
- Keep the existing swap engine pipeline as the only mechanism for origin-side wrapping or swapping into SBA bridge tokens.
- Update SDK tests and README to reflect the new candidate model and the removal of SBA-native special casing.

## Non-goals

- No changes to legacy `bridgeQuote()`, `allBridgeQuotes()`, `bridge()`, or other bridge V1 behavior.
- No new public SDK parameters or public API surface.
- No runtime reads from `contracts-adapter` or other sibling workspaces after publish.
- No new SBA-specific origin-swap engine, router contract, or swap-module path.
- No requirement to implement live RPC verification, artifact refresh tooling, or automated artifact generation in this phase.
- No widget or `synapse-interface` integration work.
- No requirement for runtime or build-time derivation from `contracts-adapter`; the SDK may consume a committed manually maintained snapshot directly.

## Requirements

1. Add a committed hardcoded SBA supported-token artifact in `sdk-router` that can answer, for a given origin chain and destination chain, which local SBA tokens are bridgeable and which remote token each maps to.
2. The artifact must encode the manually verified SBA token snapshot for the current supported SDK chains, and runtime SBA routing in `sdk-router` must consume only that committed data.
3. The artifact must normalize chain naming differences the same way `metadata.ts` already does, including `ethereum -> SupportedChainId.ETH`, `bnb -> SupportedChainId.BSC`, and `kaia -> SupportedChainId.KLAYTN`.
4. The artifact must only include tokens that belong in the current verified supported-chain snapshot:
   - token exists on at least two supported SBA chains
   - token is not explicitly excluded, such as `GMX`
   - both the origin and destination chain exist in `SBA_CHAIN_METADATA`
5. The single-chain exclusion must be reflected in the hardcoded artifact. Under the current repo snapshot, that exclusion removes `BUSD` on BNB from the SBA artifact.
6. The artifact shape must support efficient lookup by:
   - origin chain
   - origin token address
   - destination chain
   - optional exact destination token filter
7. `SynapseBridgeAdapterModuleSet.getBridgeTokenCandidates()` must stop calling `adapter.getRemoteAddress(...)`.
8. `SynapseBridgeAdapterModuleSet.getBridgeTokenCandidates()` must stop special-casing native input and must not call `getWrappedNativeToken()` or any equivalent SBA-local helper.
9. SBA candidate discovery must return every artifact-defined SBA bridge token for the requested `fromChainId -> toChainId` pair, optionally filtered by exact `toToken` when `toToken` is supplied.
10. SBA candidate discovery must no longer require `fromToken` itself to already be an SBA token. The existing origin-side swap pipeline is responsible for reaching any returned SBA origin token.
11. `SynapseBridgeAdapterModuleSet.getBridgeRouteV2()` must validate the chosen candidate against the artifact instead of re-querying `getRemoteAddress(...)`.
12. `getBridgeRouteV2()` must preserve the current SBA route semantics:
    - bridge amount is 1:1 with the amount that reaches the SBA origin token
    - `expectedToAmount = originSwapRoute.expectedToAmount`
    - `minToAmount = expectedToAmount` for no-op origin routes
    - `minToAmount = originSwapRoute.minToAmount` when an origin-side swap route exists
    - `nativeFee` still comes from `adapter.getNativeFee(dstEid, 200_000)`
    - zap-data encoding still uses `bridgeERC20(dstEid, to, token, amount, gasLimit)`
13. Single-transaction `bridgeV2()` behavior must remain strict: SBA quotes are only returned when the artifact-mapped destination token exactly matches the requested `toToken`.
14. `intent({ allowMultipleTxs: true })` behavior must remain unchanged: SBA may bridge into the artifact-mapped destination token first, then the existing destination-side swap flow may append a follow-up step.
15. Remove the SBA module code that exists only for wrapped-native discovery:
    - `swapQuoterContract`
    - wrapped-native cache
    - `getWrappedNativeToken()`
    - any SBA-only `SwapQuoterV2.weth()` dependency
16. If no origin-side swap route exists from the requested `fromToken` into any artifact-defined SBA origin token, SBA must return no quote rather than introducing a new SBA-local fallback path.
17. The README section for SBA must no longer describe a special native-wrap exception. It must describe artifact-backed SBA candidates plus the existing origin swap engine path instead.

## Technical approach

- Add a new SBA artifact module, for example `packages/sdk-router/src/sba/supportedTokens.ts`.
- Store the snapshot as committed TypeScript data and expose helper functions rather than reading sibling-workspace JSON at runtime.
- Treat the March 18, 2026 manually verified snapshot as the input to encode in that committed data; no SDK-side refresh or validation workflow is required in this scope.
- Use a data shape that is keyed by origin chain and origin token, with per-destination remote token mappings. The artifact may also store a stable token ID or symbol for readability, but runtime routing only needs addresses and chain IDs.
- Build helper APIs around the artifact, for example:
  - `getSbaSupportedTokens(fromChainId, toChainId, toToken?)`
  - `getSbaRemoteToken(fromChainId, originToken, toChainId)`
  - optional assertions or duplicate-detection helpers during module initialization or test-time validation
- Refactor `SynapseBridgeAdapterModuleSet.getBridgeTokenCandidates()` to:
  - require SBA modules/metadata on both chains, as today
  - read supported origin tokens from the artifact for the requested chain pair
  - emit `BridgeTokenCandidate` values directly from artifact mappings
  - ignore `fromToken` as a candidate-discovery filter
- Keep `_collectV2Quotes()` as the origin-side route planner. This already dedupes candidate origin tokens and routes the requested input token into them through the shared swap engine set. No new SBA-local wrapping/swapping code is required there.
- Refactor `SynapseBridgeAdapterModuleSet.getBridgeRouteV2()` to use the artifact lookup as the pair validator. The route should trust the selected `bridgeToken` plus the artifact mapping and continue to call the live adapter only for `getNativeFee(...)` and transaction execution parameters.
- Remove the now-unused SBA module methods and fields that only supported the wrapped-native branch. If `getRemoteAddress()` is no longer used anywhere in `sdk-router`, remove that SBA module helper as well.
- Keep transaction population, ETA refresh/cache, `getSynapseTxId()`, and LayerZero status polling unchanged.
- Update README language from "direct route with one native-wrap exception" to "artifact-defined SBA bridge tokens that may be reached via the existing origin swap engine path."

## Affected areas

- New file under `packages/sdk-router/src/sba/` for the committed supported-token artifact and its helpers.
- [packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.ts)
- [packages/sdk-router/src/sba/synapseBridgeAdapterModule.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModule.ts)
- [packages/sdk-router/src/sba/index.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/index.ts)
- [packages/sdk-router/src/operations/bridge.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/operations/bridge.ts) for comments or small supporting adjustments, if needed
- [packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.test.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModuleSet.test.ts)
- [packages/sdk-router/src/sba/synapseBridgeAdapterModule.test.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sba/synapseBridgeAdapterModule.test.ts)
- [packages/sdk-router/src/sdk.test.ts](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/src/sdk.test.ts)
- [packages/sdk-router/README.md](/workspace/.worktrees/feat/sba-sdk/packages/sdk-router/README.md)

## Edge cases and failure handling

- If either chain lacks an SBA module instance or SBA metadata, SBA candidate discovery must return `[]`.
- If the artifact has no entry for the requested chain pair, SBA candidate discovery must return `[]`.
- If `toToken` is provided and no artifact mapping for the pair matches it exactly, SBA candidate discovery must return `[]`.
- If multiple token IDs would create conflicting `(originChainId, originToken, destChainId)` mappings, artifact construction or validation must fail fast rather than silently choosing one.
- If the origin swap engine set cannot route the requested `fromToken` into any SBA artifact token, no SBA quote should be returned.
- If `getNativeFee()` reverts, the SBA route should be dropped, matching current "no partial quote" behavior.
- If the artifact drifts from deployed SBA configuration, the SDK may omit valid routes or surface routes that later fail at execution time. This is an accepted operational risk of the committed hardcoded-snapshot approach and must be documented.
- Removing SBA-local wrapped-native discovery means native-origin SBA quotes now depend entirely on swap-engine coverage. If the origin swap engines cannot express that wrap path, the SDK must return no SBA quote.

## Phase plan

1. Add the committed hardcoded SBA supported-token artifact and helper lookups under `packages/sdk-router/src/sba/`.
2. Refactor `SynapseBridgeAdapterModuleSet` to build candidates and validate routes from the artifact instead of `getRemoteAddress(...)`.
3. Remove wrapped-native-specific SBA module code and clean up now-unused imports, fields, and tests.
4. Update SDK-level tests to prove direct SBA routes, artifact-backed multi-candidate routing, and native-origin routing through the existing swap engine path.
5. Update `packages/sdk-router/README.md` to document the new artifact-backed candidate model and the removal of SBA-native special casing.

## Acceptance criteria

- SBA candidate discovery is driven by a committed SDK-local artifact rather than live `getRemoteAddress(...)` calls.
- The committed artifact excludes tokens that only exist on a single supported SBA chain and excludes `GMX`.
- `SynapseSDK.bridgeV2()` still returns direct SBA quotes when the requested `fromToken` is already an artifact-defined SBA origin token and the requested `toToken` matches the mapped remote token.
- `SynapseSDK.bridgeV2()` can return SBA quotes for native-origin requests only when the existing swap engine path can route native input into an artifact-defined SBA origin token; no SBA-specific wrapped-native helper is involved.
- `SynapseSDK.intent({ allowMultipleTxs: true })` still uses SBA as the bridge step and may append a destination-side swap when the requested final token differs from the SBA mapped token.
- No SBA codepath in `sdk-router` calls `SwapQuoterV2.weth()` for wrapped-native discovery.
- No SBA codepath in `sdk-router` relies on live `getRemoteAddress(...)` for bridge candidate discovery.
- Existing SBA route construction still produces the same `bridgeERC20(...)` zap payload, `nativeFee` handling, ETA behavior, and tx-status behavior.
- README text matches the new implementation shape.

## Validation plan

- Add unit tests for artifact lookup helpers, including:
  - chain-name normalization coverage
  - exclusion of unsupported tokens such as `GMX`
  - representative known mappings for Harmony, DFK, and Klaytn/Kaia-related SBA routes
- Update `SynapseBridgeAdapterModuleSet` tests to assert:
  - candidates come from artifact entries
  - `toToken` filtering works against artifact data
  - no-candidate behavior for unsupported chain pairs still works
  - route validation no longer depends on `getRemoteAddress()`
- Remove or rewrite `SynapseBridgeAdapterModule` tests that only exist for `getWrappedNativeToken()`.
- Update SDK-level tests to verify:
  - direct SBA `bridgeV2()` quotes still work
  - native-origin SBA `bridgeV2()` quotes work through the generic swap-engine path without SBA-specific wrapped-native mocks
  - `intent({ allowMultipleTxs: true })` still uses SBA then destination swap
- Run:
  - `yarn --cwd packages/sdk-router test`
  - `yarn --cwd packages/sdk-router lint:check`
  - `yarn --cwd packages/sdk-router build`

## Risks and assumptions

- Assumption: the hardcoded artifact values being committed here already match the intended SBA token universe, based on the repo snapshot and the March 18, 2026 manual verification pass.
- Assumption: the existing origin swap engine set, especially `DefaultPoolsEngine` for native wrapping where supported, is sufficient to reach artifact-defined SBA origin tokens without SBA-local helpers.
- Assumption: no consumer depends on SBA-specific wrapped-native discovery APIs, since they are internal-only today.
- Risk: artifact drift can cause false positives or false negatives in quote discovery until the committed hardcoded snapshot is updated manually.
- Risk: artifact-backed candidate enumeration may increase the number of origin swap quotes attempted for SBA compared with the current direct-token lookup path, though `_collectV2Quotes()` already dedupes origin tokens before route generation.
