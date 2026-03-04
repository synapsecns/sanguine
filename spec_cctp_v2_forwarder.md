# CCTP V2 Forwarding Service Support in `sdk-router`

## Problem
`packages/sdk-router` currently exposes CCTP through legacy `SynapseCCTPRouterSet` behavior:
- Not included in `bridgeV2` (`isBridgeV2Supported = false`).
- Built around Synapse CCTP V1-style request/fulfillment assumptions.
- Does not support Circle Forwarding Service execution via SIR.

We need a CCTP V2 path that:
- Uses Circle infrastructure for destination execution (no self-relay).
- Enforces forwarding-service semantics.
- Integrates through `bridgeV2` and SynapseIntentRouter (SIR), not a new standalone router flow.

## Goals
1. Add a new bridge-V2 module named `CCTPv2`.
2. Include `CCTPv2` by default in `bridgeV2` quote discovery.
3. Execute via SIR (`BridgeRouteV2.zapData`) only.
4. Use Circle Forwarding Service hook format with no extra custom hook payload.
5. Keep legacy `SynapseCCTP` behavior unchanged for existing V1 quote APIs.
6. Define accurate `estimatedTime` policy for CCTP V2 using Circle finality settings.

## Non-Goals
1. Replacing/removing `SynapseCCTP` in this phase.
2. Adding widget/interface/rest-api integration in this spec.
3. Building a custom attestation + destination relay flow.
4. Supporting arbitrary custom forwarding hook payloads.

## Scope
- Package: `packages/sdk-router` only.
- Primary APIs impacted: `bridgeV2` path and module composition in `sdk.ts`.
- Legacy APIs `bridgeQuote/allBridgeQuotes` continue using existing module behavior.
- This document reflects current implemented progress for the `sdk-router` CCTP V2 module.

## External Constraints (Validated March 4, 2026)
1. Forwarding service is enabled through burn hook data with reserved `cctp-forward` format.
2. Forwarding service does not support wrapper-contract forwarding when `destinationCaller` is set.
3. Forwarding-service fixed service fees documented by Circle:
   - Ethereum destination: `1.25` USDC.
   - All other destinations: `0.20` USDC.
4. CCTP fee/finality data is available via Circle API `GET /v2/burn/USDC/fees/{sourceDomainId}/{destDomainId}` (use `?forward=true` for forwarding-aware responses where available).
5. Message/forwarding status is available via `GET /v2/messages/{sourceDomainId}?transactionHash=...`.
6. Circle finality semantics:
   - `minFinalityThreshold = 2000`: Standard Transfer attestation.
   - `1000 <= minFinalityThreshold < 2000`: Fast Transfer attestation.
7. Circle required-block-confirmation timing (as source chain) for Standard Transfer:
   - Ethereum, Arbitrum, Base, Optimism: `~15-19 minutes`.
   - Avalanche, Polygon PoS: `~8 seconds`.
8. Spot checks against live Iris fee responses on March 4, 2026 show fee tiers include `finalityThreshold` values `1000` and `2000`, consistent with Circle finality modes.

## Requirements

### Functional
1. Add `CCTPv2ModuleSet` (bridge-V2 only):
   - `moduleName = 'CCTPv2'`
   - `isBridgeV2Supported = true`
   - `getBridgeRoutes()` returns `[]`
   - `getBridgeTokenCandidates()` returns supported native-USDC CCTP V2 pairs
   - `getBridgeRouteV2()` returns `BridgeRouteV2` with `zapData`
2. Add `CCTPv2Module` for low-level calldata/status helpers.
3. Register module in `sdk.ts` and `allModuleSets` so it is active by default in `bridgeV2`.
4. Keep `SynapseCCTPRouterSet` unchanged and still available for legacy quote APIs.

### Forwarding-Service Semantics
1. Build burns using `TokenMessengerV2.depositForBurnWithHook` (or equivalent V2 hook-enabled burn call).
2. Use forwarding hook payload with no custom append data:
   - `0x636374702d666f72776172640000000000000000000000000000000000000000`
3. Set `destinationCaller = bytes32(0)` for forwarding-service compatibility.
4. If forwarding prerequisites are not met (unsupported chain/config/API data), do not return a `CCTPv2` quote.

### Fee and Finality Policy
1. Use live Circle fee API (with cache) for every route evaluation.
2. Select the slowest available route by always choosing the maximum available `finalityThreshold` from API response.
3. Compute protocol fee budget from API `minimumFee` (bps) and transfer amount.
4. Include forwarding fee budget in `maxFee`:
   - Prefer API `forwardFee` data when present.
   - If `forwardFee` is not present in fee response, use fixed service-fee constants:
     - `1.25` USDC for Ethereum destination.
     - `0.20` USDC for all other destinations.
5. If fee API request fails or returns unusable data for a route, return no `CCTPv2` quote for that route (fail closed for this module only).

### Estimated Time Policy
1. `estimatedTime` must be derived from selected `finalityThreshold` and source-chain finality timing from Circle docs.
2. Standard mode (`finalityThreshold >= 2000`) timing estimates by source chain:
   - `ETH`, `ARBITRUM`, `BASE`, `OPTIMISM`: `1020` seconds (midpoint of `15-19 minutes`).
   - `AVALANCHE`, `POLYGON`: `8` seconds (rounded to `10` by SDK quote time precision).
3. Fast mode (`1000 <= finalityThreshold < 2000`) timing estimates by source chain:
   - `ETH`, `ARBITRUM`, `BASE`, `OPTIMISM`: `600` seconds.
   - `AVALANCHE`, `POLYGON`: `6` seconds (rounded to `10` by SDK quote time precision).
4. If `finalityThreshold` is unmappable for policy (for example `< 1000`), return no `CCTPv2` quote (fail closed).
5. Module-level `getEstimatedTime(chainId)` should return non-zero fallback values for supported CCTP V2 source chains.

### API Strategy
1. Runtime source of truth is Circle API (not static quote constants).
2. Use in-memory TTL caching for fee/status calls to limit request volume:
   - Fee API TTL: 15 seconds.
   - Status API TTL: 5 seconds.
3. Do not use stale-cache-beyond-TTL fallback for quote generation.

### Transaction Lifecycle
1. `getSynapseTxId(txHash)` returns `txHash:originChainId` for `CCTPv2`.
2. `getBridgeTxStatus()` queries Circle `v2/messages` and returns `true` only when:
   - Message `status` is terminal success, and
   - `forwardState` is terminal success for forwarding-enabled transfers.
3. Pending/unknown/API-error states return `false` (no throw for transient fetch failures).

## Proposed Design

### Files
Add new module namespace:
1. `packages/sdk-router/src/cctpV2/cctpV2Module.ts`
2. `packages/sdk-router/src/cctpV2/cctpV2ModuleSet.ts`
3. `packages/sdk-router/src/cctpV2/api.ts`
4. `packages/sdk-router/src/cctpV2/index.ts`

Update:
1. `packages/sdk-router/src/sdk.ts` (module registration)
2. `packages/sdk-router/src/constants/*` (CCTP V2 maps)
3. `packages/sdk-router/README.md`
4. `packages/sdk-router/CHANGELOG.md`

### `CCTPv2Module` Responsibilities
1. Encode burn calldata for hook-enabled CCTP V2 forwarding path.
2. Provide amount-position detection for zap payload substitution.
3. Provide status lookup helper over `/v2/messages`.

### `CCTPv2ModuleSet` Responsibilities
1. Candidate discovery for supported source/destination USDC pairs.
2. Live fee fetch + route finality selection (max threshold).
3. `maxFee` composition (protocol + forwarding fee budget).
4. Zap data generation (`encodeZapData`) for SIR finalization.

### Constants
Add explicit maps for:
1. `CCTP_V2_SUPPORTED_CHAIN_IDS`
2. `CCTP_V2_DOMAIN_MAP` (chainId -> domainId)
3. `CCTP_V2_USDC_ADDRESS_MAP` (chainId -> native USDC)
4. `CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP`
5. `CCTP_V2_FORWARD_SERVICE_FEE_USDC` fallback constants (ETH vs non-ETH)
6. API hosts (`iris-api.circle.com`, `iris-api-sandbox.circle.com` if test coverage needs it)

## Quote Behavior
1. Single-tx mode (`allowMultipleTxs = false`): return route only when destination token equals destination native USDC for the candidate path.
2. Multi-tx mode (`allowMultipleTxs = true`): allow destination fallback to bridged USDC if final requested token differs.
3. No-route conditions:
   - Unsupported chain/domain/token messenger mapping
   - Missing/invalid live fee data
   - Non-positive resulting output after fees

## Edge Cases
1. API returns fee entries but no usable finality/fee combination -> no quote.
2. API provides protocol fee but omits forwarding fee -> use fixed service fee constants only.
3. `fromSender` missing -> quote can be returned, tx omitted (existing bridge-V2 pattern).
4. `toRecipient` missing and cannot be defaulted -> no zapData/tx.
5. Circle status available but non-terminal -> `getBridgeTxStatus = false`.
6. Scientific-notation `minimumFee` values are treated as unusable input in current implementation -> no quote (fail closed).

## Testing Plan

### Unit
1. `cctpV2Module.test.ts`
   - Calldata encoding for hook-enabled burn.
   - Amount position detection.
   - Forwarding hook constant usage.
2. `cctpV2ModuleSet.test.ts`
   - Candidate filtering.
   - Max-finality route selection from mocked fee response.
   - Finality-derived `estimatedTime` selection and chain mapping (with SDK rounding behavior).
   - Unmappable finality threshold returns no quote (fail closed).
   - Scientific-notation `minimumFee` values are treated as unusable and return no quote.
   - `maxFee` calculation from `minimumFee` + forwarding fee.
   - API error path returns no quote.
   - Forwarding fee fallback (only when `forwardFee` missing).

### SDK Integration
1. `sdk.test.ts`
   - `bridgeV2` includes `CCTPv2` by default when supported.
   - `moduleNames` includes `CCTPv2`.
   - `estimatedTime` for `CCTPv2` is non-zero and matches finality-derived policy.
   - `getEstimatedTime(..., 'CCTPv2')` returns non-zero chain-aware values.
   - Legacy `SynapseCCTP` V1 quote APIs remain unchanged.

### Regression
1. Existing bridge-V2 modules unaffected.
2. Existing `SynapseCCTPRouterSet` tests unaffected.

## Rollout
1. Implement module + API client + constants.
2. Wire into SDK and tests.
3. Update docs/changelog.
4. Validate supported chain mappings against Circle supported-blockchain docs before release.

## Implementation Status (As of March 4, 2026)
Implemented:
1. `CCTPv2` bridge-v2 module, forwarding hook, fee API integration, and max-finality (`max finalityThreshold`) route selection.
2. Route-level `estimatedTime` emission based on selected `finalityThreshold` mode (standard vs fast) and source-chain mapping.
3. Fail-closed behavior for unmappable finality thresholds (`no quote`).
4. Module-level `getEstimatedTime()` non-zero fallback for supported CCTP V2 source chains.
5. Unit and SDK tests covering finality-derived estimated-time behavior.

Not Implemented Yet (for this spec scope):
1. `packages/sdk-router/README.md` updates listed in proposed design are not yet landed.
2. `packages/sdk-router/CHANGELOG.md` updates listed in proposed design are not yet landed.

## Acceptance Criteria
1. `bridgeV2` returns `CCTPv2` quotes by default for supported CCTP V2 USDC routes.
2. Generated tx path is SIR-based and includes forwarding hook payload.
3. Finality policy always chooses the slowest available API finality threshold.
4. Module fails closed (no quote) when live API data is unavailable or unusable.
5. `CCTPv2` quotes include non-zero `estimatedTime` derived from finality mode and source-chain mapping.
6. Unmappable `finalityThreshold` values fail closed (`no quote`).
7. Legacy `SynapseCCTP` quote behavior remains unchanged.

## Open Questions
None.

## References
1. Forwarding service concept: https://developers.circle.com/cctp/concepts/forwarding-service
2. CCTP technical guide (finality + fees): https://developers.circle.com/cctp/technical-guide
3. Fee API: https://developers.circle.com/api-reference/cctp/all/get-burn-usdc-fees
4. Messages/status API: https://developers.circle.com/api-reference/cctp/all/get-messages-v2
5. V1->V2 migration: https://developers.circle.com/cctp/migration-from-v1-to-v2
6. Required block confirmations (finality timing): https://developers.circle.com/cctp/required-block-confirmations
7. Supported blockchains/domains: https://developers.circle.com/cctp/supported-blockchains
