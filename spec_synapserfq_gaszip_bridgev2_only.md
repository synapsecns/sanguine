# SynapseRFQ and Gas.zip BridgeV2-Only Workflows

## Goal
Make `SynapseRFQ` and `Gas.zip` available only through the `bridgeV2` and intent-style workflows, while leaving the generic legacy bridge orchestration intact for other modules. After this change, legacy quote APIs must stop surfacing these two modules, but their V2 candidate discovery, route generation, and transaction construction must continue to work.

## Context
- `SynapseSDK` constructs `synapseRouterSet`, `synapseCCTPRouterSet`, `fastBridgeRouterSet`, `gasZipModuleSet`, `relayModuleSet`, and `usdtModuleSet`, but only `fastBridgeRouterSet`, `gasZipModuleSet`, `relayModuleSet`, and `usdtModuleSet` are placed in `allModuleSets`.
- `allBridgeQuotes()` iterates every module in `allModuleSets` and asks each module for legacy `BridgeRoute` entries via `getBridgeRoutes()`. This means `SynapseRFQ` and `Gas.zip` currently participate in legacy `bridgeQuote()` and `allBridgeQuotes()` because both module sets implement `getBridgeRoutes()`.
- `bridgeV2()` follows a separate path. It filters `allModuleSets` by `isBridgeV2Supported`, collects candidates through `getBridgeTokenCandidates()`, and builds V2 routes through `getBridgeRouteV2()`.
- `RelayModuleSet` and `UsdtModuleSet` are the existing V2-only convention in `sdk-router`: they keep `isBridgeV2Supported = true`, implement V2 candidate and route generation, and return `[]` from `getBridgeRoutes()`.
- `FastBridgeRouterSet` is mixed-mode today. It has V2 entry points, but it also still implements legacy V1 routing, legacy fee calculation, legacy slippage logic, and legacy destination-query building.
- `GasZipModuleSet` is also mixed-mode today, although its auxiliary helper methods already look like a V2-only module contract.
- `FastBridgeRouterSet` has one important shared helper: `getBridgeZapData()` uses `getDefaultPeriods().destPeriod` to compute the FastBridge deadline embedded in V2 zap calldata. That deadline behavior must not change as part of this refactor.
- Because the REST `/bridge` controller, widget quote flows, and any direct SDK consumers of `allBridgeQuotes()` rely on the legacy path, they will stop receiving RFQ and Gas.zip routes automatically once these module sets stop returning legacy routes. No generic orchestration changes are required for that behavior shift.

## Scope
- Convert `FastBridgeRouterSet` and `GasZipModuleSet` to V2-only routing behavior.
- Remove or neutralize their legacy route-generation behavior.
- Preserve generic bridge orchestration, public bridge operations, and `SynapseSDK` wiring unchanged.
- Update `sdk-router` tests so the new module contract is enforced.

## Non-goals
- Do not change `SynapseSDK.allModuleSets`, `operations/bridge.ts`, `SynapseModuleSet`, or the generic legacy bridge APIs.
- Do not change `SynapseBridge` or `SynapseCCTP`.
- Do not change intent orchestration, `SynapseIntentRouterSet`, or module selection rules.
- Do not change `FastBridgeRouter`, `GasZipModule`, status lookup helpers, or event-to-module mapping behavior unless a direct module cleanup requires it.
- Do not update public docs, README files, REST API docs, or widget code in this scope.

## Requirements
- `SynapseRFQ` must no longer return legacy `BridgeRoute` results from `getBridgeRoutes()`.
- `Gas.zip` must no longer return legacy `BridgeRoute` results from `getBridgeRoutes()`.
- `SynapseRFQ` and `Gas.zip` must remain `bridgeV2`-capable by keeping `isBridgeV2Supported = true`.
- `SynapseRFQ.getBridgeTokenCandidates()` and `SynapseRFQ.getBridgeRouteV2()` must keep working for currently supported V2 routes.
- `GasZipModuleSet.getBridgeTokenCandidates()` and `GasZipModuleSet.getBridgeRouteV2()` must keep working for currently supported V2 routes.
- `allBridgeQuotes()` and `bridgeQuote()` must stop surfacing RFQ and Gas.zip solely because those modules stop producing legacy routes. No generic orchestration code should be edited to special-case them.
- `bridgeV2()` and `intent()` must continue to be able to return RFQ and Gas.zip quotes when their V2 module logic succeeds.
- `GasZipModuleSet` must follow the same V2-only convention already used by Relay and USDT0: `getBridgeRoutes()` returns `[]`, and its non-V2 helper methods remain zero-value or no-op implementations.
- `FastBridgeRouterSet` must follow the same V2-only routing convention, but it must preserve any helper that is still required by V2 execution. In particular, the FastBridge deadline embedded in V2 zap data must remain unchanged.
- `getBridgeModuleName()`, `getSynapseTxId()`, and `getBridgeTxStatus()` behavior for `SynapseRFQ` must remain intact.
- Existing or new tests must reflect that RFQ- and Gas.zip-specific legacy routing is no longer supported.

## Technical approach
- In `packages/sdk-router/src/rfq/fastBridgeRouterSet.ts`:
  - Change `getBridgeRoutes()` to return `[]` immediately with a `Bridge V1 is not supported` comment, matching the V2-only pattern already used by Relay and USDT0.
  - Convert V1-only helper behavior that is no longer part of supported flows:
    - `getFeeData()` returns a zero fee/config stub.
    - `applySlippage()` becomes an identity implementation, since RFQ will no longer be a supported legacy quote source.
  - Preserve `getDefaultPeriods()` at its current effective RFQ values, or move the V2 deadline constant into a private helper that keeps the value unchanged. Do not zero out RFQ deadlines unless zap-data deadline generation is first moved off `getDefaultPeriods()`.
  - Remove dead V1-only helpers and imports once `getBridgeRoutes()` no longer uses them, including `filterOriginQuotes()`, `createRFQDestQuery()`, and any imports that become unused.
  - Leave V2-specific logic intact: quote caching, `getBridgeTokenCandidates()`, `getBridgeRouteV2()`, `getBridgeZapData()`, protocol-fee handling, and FastBridge address lookup remain the active implementation surface.
- In `packages/sdk-router/src/gaszip/gasZipModuleSet.ts`:
  - Change `getBridgeRoutes()` to return `[]` immediately with a `Bridge V1 is not supported` comment.
  - Keep existing V2 entry points unchanged: `getBridgeTokenCandidates()`, `getBridgeRouteV2()`, block-height freshness checks, gas.zip quote retrieval, and zap-data creation.
  - Keep existing zero/no-op helper behavior (`getFeeData()`, `getDefaultPeriods()`, `applySlippage()`) as the explicit V2-only contract.
  - Remove any V1-only imports or unreachable code that becomes dead after `getBridgeRoutes()` is simplified.
- Testing strategy:
  - Rewrite RFQ tests away from legacy slippage math and legacy destination-query generation.
  - Add direct module tests asserting RFQ and Gas.zip `getBridgeRoutes()` return empty arrays.
  - Add targeted RFQ V2 coverage to protect the deadline behavior used in zap-data construction.
  - If a lightweight SDK-level assertion is practical with current test infrastructure, add one to verify that legacy quote aggregation no longer includes RFQ or Gas.zip while the generic `bridgeV2` path itself remains untouched. If not, keep validation at module level plus existing SDK coverage.

## Affected areas
- `packages/sdk-router/src/rfq/fastBridgeRouterSet.ts`
- `packages/sdk-router/src/gaszip/gasZipModuleSet.ts`
- `packages/sdk-router/src/rfq/fastBridgeRouterSet.test.ts`
- `packages/sdk-router/src/gaszip/gasZipModuleSet.test.ts` if a new test file is added
- `packages/sdk-router/src/sdk.test.ts` or another targeted SDK/operations test file only if an aggregate assertion is introduced

## Edge cases and failure handling
- Legacy callers using `bridgeQuote()` or `allBridgeQuotes()` for routes that previously resolved through RFQ or Gas.zip may now receive fewer quotes or `No route found`. This is expected behavior.
- `bridgeV2()` already separates V1 and V2 collection paths. This refactor must not alter that orchestration or add module-specific conditions there.
- RFQ deadlines are a shared implementation detail for V2 zap construction. A naive Relay-style `getDefaultPeriods()` stub would break that V2 path; the implementation must avoid that regression.
- Direct callers of `applyBridgeSlippage('SynapseRFQ', ...)` will no longer receive RFQ-specific legacy slippage behavior once RFQ becomes V2-only. This is acceptable under the new module contract.
- RFQ status and tx-id lookup must continue to work for existing and future RFQ-backed transfers, even though legacy quote generation is removed.

## Phase plan
1. Update `FastBridgeRouterSet` so legacy route generation is disabled, unreachable V1 helper logic is stubbed or removed, and RFQ V2 deadline handling is explicitly preserved.
2. Update `GasZipModuleSet` so legacy route generation is disabled and any resulting dead code or imports are cleaned up without changing V2 behavior.
3. Rewrite or add `sdk-router` tests for the new V2-only contract, including RFQ and Gas.zip `getBridgeRoutes()` coverage and an RFQ V2 deadline invariant.
4. Run the `sdk-router` test suite and fix any remaining assertions that still assume RFQ or Gas.zip participate in legacy bridge aggregation.

## Acceptance criteria
- `FastBridgeRouterSet.getBridgeRoutes()` returns `[]` for any input.
- `GasZipModuleSet.getBridgeRoutes()` returns `[]` for any input.
- `SynapseRFQ` no longer appears in results produced through legacy quote flows that depend on `getBridgeRoutes()`.
- `Gas.zip` no longer appears in results produced through legacy quote flows that depend on `getBridgeRoutes()`.
- RFQ V2 candidate discovery and route generation still work for supported routes.
- Gas.zip V2 candidate discovery and route generation still work for supported routes.
- RFQ V2 zap-data deadline behavior is unchanged from the pre-refactor implementation.
- No generic bridge orchestration files need to be modified to obtain the new behavior.
- `sdk-router` tests pass after the module changes.

## Validation plan
- Run `yarn --cwd packages/sdk-router test`.
- Confirm RFQ unit tests no longer assert legacy slippage math or legacy destination-query construction.
- Add and run tests that assert RFQ and Gas.zip `getBridgeRoutes()` return empty arrays.
- Add and run a targeted RFQ V2 test that exercises zap-data generation and verifies the embedded deadline input is unchanged by the refactor.
- Sanity-check that `fastBridgeRouterSet.ts` and `gasZipModuleSet.ts` no longer retain unused V1-only imports or helpers.

## Risks and assumptions
- Assumption: the intended product behavior is that legacy `/bridge`, widget V1 quote flows, and any direct `allBridgeQuotes()` consumers should stop surfacing RFQ and Gas.zip routes without adding a compatibility shim.
- Assumption: preserving RFQ’s current deadline value for V2 zap-data generation is more important than making every helper look identical to Relay or USDT0.
- Risk: public docs in `packages/sdk-router/README.md` and `docs/bridge/docs/03-Bridge/01-SDK.md` still describe RFQ as a legacy `bridgeQuote()` module. This spec intentionally leaves documentation untouched, so those references will become stale until a follow-up update is made.
- Risk: downstream packages may have tests or UI assumptions that expect RFQ or Gas.zip in legacy quote responses. Those failures are expected fallout from the intended behavior change, not a reason to modify generic bridge orchestration in this scope.
