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
1. Add a new bridge-V2 module named `CircleCCTPV2`.
2. Include `CircleCCTPV2` by default in `bridgeV2` quote discovery.
3. Execute via SIR (`BridgeRouteV2.zapData`) only.
4. Use Circle Forwarding Service hook format with no extra custom hook payload.
5. Keep legacy `SynapseCCTP` behavior unchanged for existing V1 quote APIs.

## Non-Goals
1. Replacing/removing `SynapseCCTP` in this phase.
2. Adding widget/interface/rest-api integration in this spec.
3. Building a custom attestation + destination relay flow.
4. Supporting arbitrary custom forwarding hook payloads.

## Scope
- Package: `packages/sdk-router` only.
- Primary APIs impacted: `bridgeV2` path and module composition in `sdk.ts`.
- Legacy APIs `bridgeQuote/allBridgeQuotes` continue using existing module behavior.

## External Constraints (Validated March 3, 2026)
1. Forwarding service is enabled through burn hook data with reserved `cctp-forward` format.
2. Forwarding service does not support wrapper-contract forwarding when `destinationCaller` is set.
3. Forwarding-service fixed service fees documented by Circle:
   - Ethereum destination: `1.25` USDC.
   - All other destinations: `0.20` USDC.
4. CCTP fee/finality data is available via Circle API `GET /v2/burn/USDC/fees/{sourceDomainId}/{destDomainId}` (use `?forward=true` for forwarding-aware responses where available).
5. Message/forwarding status is available via `GET /v2/messages/{sourceDomainId}?transactionHash=...`.

## Requirements

### Functional
1. Add `CircleCCTPV2ModuleSet` (bridge-V2 only):
   - `moduleName = 'CircleCCTPV2'`
   - `isBridgeV2Supported = true`
   - `getBridgeRoutes()` returns `[]`
   - `getBridgeTokenCandidates()` returns supported native-USDC CCTP V2 pairs
   - `getBridgeRouteV2()` returns `BridgeRouteV2` with `zapData`
2. Add `CircleCCTPV2Module` for low-level calldata/status helpers.
3. Register module in `sdk.ts` and `allModuleSets` so it is active by default in `bridgeV2`.
4. Keep `SynapseCCTPRouterSet` unchanged and still available for legacy quote APIs.

### Forwarding-Service Semantics
1. Build burns using `TokenMessengerV2.depositForBurnWithHook` (or equivalent V2 hook-enabled burn call).
2. Use forwarding hook payload with no custom append data:
   - `0x636374702d666f72776172640000000000000000000000000000000000000000`
3. Set `destinationCaller = bytes32(0)` for forwarding-service compatibility.
4. If forwarding prerequisites are not met (unsupported chain/config/API data), do not return a `CircleCCTPV2` quote.

### Fee and Finality Policy
1. Use live Circle fee API (with cache) for every route evaluation.
2. Select the slowest available route by always choosing the maximum available `finalityThreshold` from API response.
3. Compute protocol fee budget from API `minimumFee` (bps) and transfer amount.
4. Include forwarding fee budget in `maxFee`:
   - Prefer API `forwardFee` data when present.
   - If `forwardFee` is not present in fee response, use fixed service-fee constants:
     - `1.25` USDC for Ethereum destination.
     - `0.20` USDC for all other destinations.
5. If fee API request fails or returns unusable data for a route, return no `CircleCCTPV2` quote for that route (fail closed for this module only).

### API Strategy
1. Runtime source of truth is Circle API (not static quote constants).
2. Use in-memory TTL caching for fee/status calls to limit request volume:
   - Fee API TTL: 15 seconds.
   - Status API TTL: 5 seconds.
3. Do not use stale-cache-beyond-TTL fallback for quote generation.

### Transaction Lifecycle
1. `getSynapseTxId(txHash)` returns `txHash`.
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

### `CCTPV2Module` Responsibilities
1. Encode burn calldata for hook-enabled CCTP V2 forwarding path.
2. Provide amount-position detection for zap payload substitution.
3. Provide status lookup helper over `/v2/messages`.

### `CCTPV2ModuleSet` Responsibilities
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

## Testing Plan

### Unit
1. `cctpV2Module.test.ts`
   - Calldata encoding for hook-enabled burn.
   - Amount position detection.
   - Forwarding hook constant usage.
2. `cctpV2ModuleSet.test.ts`
   - Candidate filtering.
   - Max-finality route selection from mocked fee response.
   - `maxFee` calculation from `minimumFee` + forwarding fee.
   - API error path returns no quote.
   - Forwarding fee fallback (only when `forwardFee` missing).

### SDK Integration
1. `sdk.test.ts`
   - `bridgeV2` includes `CircleCCTPV2` by default when supported.
   - `moduleNames` includes `CircleCCTPV2`.
   - Legacy `SynapseCCTP` V1 quote APIs remain unchanged.

### Regression
1. Existing bridge-V2 modules unaffected.
2. Existing `SynapseCCTPRouterSet` tests unaffected.

## Rollout
1. Implement module + API client + constants.
2. Wire into SDK and tests.
3. Update docs/changelog.
4. Validate supported chain mappings against Circle supported-blockchain docs before release.

## Acceptance Criteria
1. `bridgeV2` returns `CircleCCTPV2` quotes by default for supported CCTP V2 USDC routes.
2. Generated tx path is SIR-based and includes forwarding hook payload.
3. Finality policy always chooses the slowest available API finality threshold.
4. Module fails closed (no quote) when live API data is unavailable or unusable.
5. Legacy `SynapseCCTP` quote behavior remains unchanged.

## Open Questions
None.

## References
1. Forwarding service concept: https://developers.circle.com/cctp/concepts/forwarding-service
2. CCTP technical guide (finality + fees): https://developers.circle.com/cctp/technical-guide
3. Fee API: https://developers.circle.com/api-reference/cctp/all/get-burn-usdc-fees
4. Messages/status API: https://developers.circle.com/api-reference/cctp/all/get-messages-v2
5. V1->V2 migration: https://developers.circle.com/cctp/migration-from-v1-to-v2
