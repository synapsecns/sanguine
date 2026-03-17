# Synapse Interface Bridge ETA Formatting And Start-Time Alignment

## Goal

Make bridge ETA text in `packages/synapse-interface` compact and consistent by displaying durations as `XmYs` or `Xs`, and make pending bridge countdowns start from wallet confirmation time instead of the earlier button-click time.

## Context

- The bridge quote ETA shown in `components/StateManagedBridge/BridgeExchangeRateInfo.tsx` currently reads `bridgeQuote.estimatedTime` from Redux and renders either raw seconds or `estimatedTime / 60` minutes with no seconds component and no compact formatting.
- The bridge quote fetch path in `slices/bridgeQuote/thunks.ts` passes `estimatedTime` through from `synapseSDK.bridgeV2(...)` unchanged. In `@synapsecns/sdk-router`, `estimatedTime` is defined in seconds.
- Pending bridge transactions are staged in `pages/state-managed-bridge/index.tsx` before wallet confirmation. After `wallet.sendTransaction(...)` resolves, the code updates the pending record with the tx hash but still leaves `timestamp` unset.
- The `_transactions` bridge activity flow in `utils/hooks/use_TransactionsListener.ts` promotes a pending bridge into `_transactions` using `timestamp: tx.id`, where `id` is the pre-wallet-click timestamp, not the wallet-confirmation timestamp.
- `_Transaction.tsx`, `TimeRemaining.tsx`, `calculateEstimatedTimeStatus.ts`, and `AnimatedProgressBar.tsx` all treat `timestamp` as the countdown start time.
- `synapse-interface` uses `next-intl` translation namespaces and keeps locale JSON files in `packages/synapse-interface/messages`. There is no shared bridge ETA formatting helper today.
- The package has Jest config and existing unit tests, but `package.json` scripts do not currently expose a working test runner command. Validation should therefore call out the concrete command needed to run specific Jest files locally.

## Scope

- Add a shared compact duration formatter for bridge ETA display in `synapse-interface`.
- Use the shared formatter in the state-managed bridge quote ETA UI.
- Use the shared formatter in the `_Transaction` countdown UI.
- Start `_Transaction` countdown timing from wallet confirmation time, using the timestamp captured immediately after `wallet.sendTransaction(...)` resolves.
- Keep pending bridge records separate from displayed `_transactions` until both tx hash and confirmed start timestamp are available.
- Update locale messages only as needed to support compact countdown copy.

## Non-goals

- Do not change ETA generation logic in `@synapsecns/sdk-router`.
- Do not change widget ETA formatting in `packages/widget`.
- Do not refactor the legacy `components/Activity/Transaction` flow unless required for compilation.
- Do not replace the existing pending bridge transaction architecture with a new store or async flow.
- Do not change bridge quote polling, bridge module selection, or transaction status polling thresholds.

## Requirements

1. Bridge quote ETA in `BridgeExchangeRateInfo` must render compact durations:
   - If `estimatedTime < 60`, show seconds only, for example `45s`.
   - If `estimatedTime >= 60`, show minutes and seconds, for example `1m0s`, `1m30s`, `21m30s`.
   - Do not render decimal minutes.
   - Do not zero-pad the seconds field beyond normal integer rendering. Example: `1m5s`, not `1m05s`.
2. `_Transaction` countdown text in `TimeRemaining` must use the same compact duration rules while the transaction is still within the estimated window.
3. The delayed state in `TimeRemaining` must continue to show `Waiting...`; if a delayed-duration parenthetical is shown, it must use the same compact formatter on the absolute delayed duration.
4. Countdown start time for newly submitted bridge transactions must be captured immediately after `wallet.sendTransaction(...)` resolves and returns the transaction hash.
5. Time spent with the wallet modal open before confirmation must not reduce the displayed remaining ETA.
6. The `_transactions` store must use the wallet-confirmation timestamp as its `timestamp` field for bridge countdown calculations, “Began” display, sorting, and progress-bar start time.
7. Quote ETA placeholder behavior must remain unchanged when no valid quote ETA is available.
8. Existing non-bridge or non-ETA behaviors must remain unchanged.

## Technical approach

- Add a small shared formatter in `packages/synapse-interface/utils/time.ts` or a nearby dedicated utility file used by both quote display and `_Transaction` UI.
- The formatter should accept a non-negative duration in seconds and return:
  - `Xs` for values under 60
  - `XmYs` for values at or above 60
- The formatter should normalize input to whole seconds before formatting.
- `BridgeExchangeRateInfo` should stop doing inline `> 60` division and instead render the formatter output directly.
- `TimeRemaining` should stop using `Math.ceil(remainingTime / 60)` and raw second strings, and should render the shared formatter for non-delayed pending states.
- Keep `Waiting...` behavior for delayed transactions. If a delayed parenthetical continues to be shown after the existing threshold, format the absolute delayed seconds through the same formatter.
- In `pages/state-managed-bridge/index.tsx`, separate the pending record identifier from the displayed transaction start time:
  - Keep the existing pending record `id` as the stable local identifier.
  - Capture a new `submittedAt` or equivalent unix-seconds value immediately after `wallet.sendTransaction(...)` resolves.
  - Pass that timestamp through `updatePendingBridgeTransaction(...)`.
  - Set `isSubmitted: true` at the same point so the stored pending state reflects reality even if that flag remains otherwise unused.
- In `slices/transactions/actions.ts`, `slices/transactions/reducer.ts`, and related call sites, make the pending-transaction timestamp semantics explicit:
  - Pending bridge records may exist before wallet confirmation.
  - A pending bridge record should only be promoted into `_transactions` once both `transactionHash` and the post-confirmation timestamp are present.
- In `utils/hooks/use_TransactionsListener.ts`, use `tx.timestamp` when creating `_transactions` entries instead of `tx.id`.
- Keep `_TransactionDetails.timestamp` as the single source of truth for countdown start time.
- For localization:
  - Prefer compact unit suffixes from the `Time` namespace rather than hard-coded English literals if new strings are needed.
  - Add any required keys to all locale files and keep `checkTranslationJsons.js` compatibility intact.
  - Preserve existing `Waiting`, `Began`, `Complete`, `Reverted`, and `Refunded` strings.

## Affected areas

- `packages/synapse-interface/components/StateManagedBridge/BridgeExchangeRateInfo.tsx`
- `packages/synapse-interface/components/_Transaction/components/TimeRemaining.tsx`
- `packages/synapse-interface/components/_Transaction/_Transaction.tsx`
- `packages/synapse-interface/components/_Transaction/helpers/calculateEstimatedTimeStatus.ts`
- `packages/synapse-interface/components/_Transaction/components/AnimatedProgressBar.tsx`
- `packages/synapse-interface/pages/state-managed-bridge/index.tsx`
- `packages/synapse-interface/utils/hooks/use_TransactionsListener.ts`
- `packages/synapse-interface/slices/transactions/actions.ts`
- `packages/synapse-interface/slices/transactions/reducer.ts`
- `packages/synapse-interface/utils/time.ts` or a new colocated time-formatting utility
- `packages/synapse-interface/messages/*.json`
- `packages/synapse-interface/__tests__` or a colocated `*.test.ts` file for the new formatter and timestamp behavior

## Edge cases and failure handling

- `estimatedTime = 0` must render as `0s`.
- `estimatedTime = 60` must render as `1m0s`.
- Large values must continue to render in minutes and seconds only; hours formatting is out of scope for this change.
- If `estimatedTime` is missing, non-finite, or the quote is still loading, the quote panel must continue to show the existing placeholder.
- If wallet confirmation is rejected or `sendTransaction` throws before returning a hash, no `_transactions` entry should be created and the existing pending record cleanup behavior must remain intact.
- If a tx hash exists but the post-confirmation timestamp is unexpectedly missing, `use_TransactionsListener` must not promote the record into `_transactions`.
- Hyperliquid deposit code uses the same pending bridge action types. Type or action-shape changes must not break that path; if it does not participate in `_Transaction` countdowns, it may continue to omit ETA-specific values.

## Phase plan

1. Add a shared compact duration formatter and unit tests for representative second values.
2. Replace inline ETA formatting in `BridgeExchangeRateInfo` with the shared formatter.
3. Replace inline countdown formatting in `TimeRemaining` with the shared formatter while preserving delayed-state copy.
4. Update the bridge submission flow to capture the countdown start timestamp after `wallet.sendTransaction(...)` resolves and store it in pending bridge state.
5. Update pending bridge action/reducer semantics so promotion into `_transactions` requires both tx hash and confirmed timestamp.
6. Update `_TransactionsListener` to use the stored post-confirmation timestamp instead of the pending record id.
7. Add or update locale keys needed for compact countdown rendering and run translation consistency checks.
8. Manually verify quote ETA display, countdown start timing, delayed-state behavior, and progress bar timing on at least one under-a-minute route and one multi-minute route.

## Acceptance criteria

- A bridge quote with `estimatedTime = 45` renders `45s`.
- A bridge quote with `estimatedTime = 60` renders `1m0s`.
- A bridge quote with `estimatedTime = 90` renders `1m30s`.
- A bridge quote with `estimatedTime = 1290` renders `21m30s`.
- A pending `_Transaction` that has 90 seconds remaining renders a compact countdown using the same formatter, not `2m remaining`.
- If the user waits 20 seconds before confirming in the wallet, the newly created `_Transaction` still starts near the full quoted ETA after wallet confirmation instead of starting 20 seconds short.
- The “Began” timestamp shown in the transaction menu reflects the wallet-confirmation time, not the earlier bridge-button click time.
- The animated progress bar starts from the wallet-confirmation timestamp.
- Delayed transactions still show `Waiting...`, and any displayed delayed-duration parenthetical uses compact formatting.
- Placeholder rendering for missing or invalid quote ETA is unchanged.

## Validation plan

- Add unit tests for the compact duration formatter covering `0`, `1`, `59`, `60`, `61`, `90`, and `1290` seconds.
- Add a targeted test for the bridge pending-transaction promotion logic if practical, or otherwise add a reducer-level assertion that the stored timestamp survives the pending-to-transaction handoff unchanged.
- Run the relevant Jest command directly for the new or updated test files using the package-local Jest config.
- Run `npm run lint:check` in `packages/synapse-interface`.
- Run `node scripts/checkTranslationJsons.js` in `packages/synapse-interface`.
- Manual QA:
  - Quote under 60 seconds shows only seconds.
  - Quote at or above 60 seconds shows `XmYs`.
  - Countdown does not start until wallet confirmation completes.
  - Delayed state still transitions to `Waiting...`.

## Risks and assumptions

- Assumption: the timestamp captured immediately after `wallet.sendTransaction(...)` resolves is an acceptable proxy for “user confirmed in wallet” and is the easiest repo-consistent implementation point.
- Assumption: compact formatting should remain minutes-and-seconds only, with no hour/day expansion, because all current bridge ETA displays are already minute-scale or smaller in the interface.
- Risk: adding new translation keys requires updating every locale file and keeping translation consistency scripts green.
- Risk: `PendingBridgeTransaction.timestamp` is currently treated as present in TypeScript but populated with `undefined` at runtime. The implementation should tighten this contract enough to prevent future regressions without broad store refactors.
- Risk: Hyperliquid deposit code shares the pending bridge action types and must be preserved during any timestamp/type cleanup.
