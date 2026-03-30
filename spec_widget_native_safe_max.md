# Widget Native Asset Safe Max

## Goal

Prevent the widget's native-asset balance control from pre-filling an amount that leaves no room for origin-chain gas or quote-level native fees, while preserving the widget's current manual-input and send-validation behavior.

## Context

- `packages/widget` currently has no amount-related public props; the bridge input is local component state inside the widget.
- The widget's balance control in `packages/widget/src/components/AvailableBalance.tsx` currently fills the full token balance on click.
- The widget's validation in `packages/widget/src/hooks/useValidations.tsx` only checks `rawBalance >= inputAmount` and does not reserve gas or `nativeFee`.
- The bridge quote flow in `packages/widget/src/state/slices/bridgeQuote/hooks.ts` already normalizes and stores `quote.nativeFee` plus the populated `quote.tx`.
- The widget does not currently estimate transaction gas for bridging.
- `packages/synapse-interface` has a separate gas-aware max flow, but the user requested widget-only scope.
- The visible destination quote and receipt in the widget are driven by Redux quote state, so safe-max bootstrap work must not overwrite the user-visible quote state just to compute a bridgeable max.

## Scope

- Apply the feature only to `packages/widget`.
- Apply the feature only when the origin asset is native on the selected origin chain.
- Change the widget's balance control so native assets display and apply a safe max instead of raw wallet balance.
- Compute native safe max from buffered estimated gas cost plus `quote.nativeFee`.
- Support empty-input bootstrap for native safe max with a two-step quote calculation.
- Keep the native safe-max flow isolated from the visible quote/receipt state used by the rest of the widget.
- Update widget tests and package README to describe the changed native balance-control behavior.

## Non-goals

- No changes to `packages/synapse-interface`.
- No new public widget props or exported types.
- No hard enforcement for manual input above the safe max.
- No change to non-native asset behavior.
- No change to the widget's existing send-button validation beyond the native balance-control prefill behavior.
- No attempt to guarantee success for users who manually override the suggested native safe max.

## Requirements

- For non-native origin assets, the balance control must keep its current behavior:
  - display the parsed wallet balance
  - fill the raw wallet balance on click
- For native origin assets, the balance control must display the bridgeable safe max once it has been computed for the current selections and wallet state.
- For native origin assets, clicking the balance control must set the input to the computed safe max, not the raw wallet balance.
- Native safe max must be calculated as:
  - `safeMaxWei = max(0, rawBalanceWei - quote.nativeFeeWei - bufferedGasFeeWei)`
  - `bufferedGasFeeWei = ceil(estimatedGasFeeWei * 1.5)`
  - `estimatedGasFeeWei = estimatedGasLimit * selectedFeePerGas`
- `selectedFeePerGas` must come from the origin-chain provider used by the widget and prefer `maxFeePerGas`; if `maxFeePerGas` is unavailable, use `gasPrice`; if neither is available, safe max is unavailable.
- Native safe max must use bigint arithmetic end-to-end; no floating-point math may be used for wei calculations.
- The native safe-max control must not fall back to filling raw native balance when the required quote or gas estimate is unavailable.
- While native safe max is not ready for the current inputs, the control must mirror the existing non-native balance-display behavior for the current wallet-balance state, but remain non-clickable.
- Because the user chose prefill-only scope, manual input above the safe max must remain allowed. Existing raw-balance validation and `Send` enablement must remain unchanged.
- The widget must not dispatch provisional safe-max quote requests through the Redux quote slice if doing so would replace the user-visible destination quote or receipt data.
- Native safe-max calculation must use the same active-quote selection rules as the widget's visible quote flow, including paused-module filtering and RFQ preference.
- Safe-max results must be invalidated when any of the following change:
  - connected address
  - origin chain
  - destination chain
  - origin token
  - destination token
  - origin wallet balance
  - paused-module inputs that affect quote selection
- If the user changes bridge selections or amount while a native safe-max calculation is in flight, stale async results must be ignored.
- If the safe max computes to zero, the balance control must display `0.0` for the bridgeable amount and remain non-clickable.

## Technical approach

- Implement native safe-max logic as widget-local state and async behavior, not as new public API.
- Add a widget-local helper or hook, for example `useNativeSafeMax`, owned by `packages/widget`, rather than reusing `packages/synapse-interface`.
- Refactor the quote-selection logic currently embedded in `packages/widget/src/state/slices/bridgeQuote/hooks.ts` into a reusable widget-local helper so both the Redux quote thunk and the native safe-max flow choose quotes the same way.
- The native safe-max helper must accept or derive:
  - current origin and destination selections
  - connected address
  - raw origin balance
  - origin-chain provider or signer/provider pair
  - paused-module state
  - `synapseSDK`
- Detect native origin assets using the same widget-native convention already used elsewhere in the package.
- Do not mutate the existing Redux bridge quote state to bootstrap safe max. Instead, perform native safe-max quote calls directly through `synapseSDK.bridgeV2(...)` from the widget-local flow.
- The native safe-max flow must run in two quote-based passes:
  1. Bootstrap pass:
     - use the full raw native balance as the provisional amount
     - request a quote directly from the SDK with sender and recipient set to the connected address
     - select the same active quote the widget would use
     - require an executable populated transaction with `to` and `data`
     - estimate gas and compute a candidate safe max
  2. Refinement pass:
     - request a second quote using the candidate safe-max amount
     - estimate gas again using the refined populated transaction
     - recompute safe max using the refined quote's `nativeFee`
     - use the refined value as the final displayed and applied safe max
- Stop after the refinement pass even if the value changes again. This keeps the flow bounded and avoids user-visible loops.
- Estimate gas from the populated transaction returned by the selected quote using the origin-chain provider.
- Include the connected address as `from` when estimating gas if the provider requires it.
- Derive gas cost from the estimated gas limit and current fee data from the same origin-chain provider.
- Apply the 150% gas buffer with integer rounding up in wei.
- Format the displayed native safe max with the same visible precision convention used by the current balance label.
- Format the applied input amount from the final safe max with high precision, matching the widget's current max-fill behavior rather than the rounded label text.
- Extend `AvailableBalance` so it can render:
  - the unchanged non-native balance path
  - a native not-ready state that mirrors the non-native balance-display behavior but stays non-clickable
  - a native ready state that displays bridgeable amount and applies the precomputed safe max on click
- Keep `useValidations.tsx` unchanged for safe-max enforcement. The chosen scope is prefill-only, so no new validation branch should be introduced for manual native overage.
- Reset or invalidate native safe-max state when wallet pending state or selection changes make the cached value unsafe to reuse.

## Affected areas

- `packages/widget/src/components/Widget.tsx`
- `packages/widget/src/components/AvailableBalance.tsx`
- `packages/widget/src/hooks/`
- `packages/widget/src/state/slices/bridgeQuote/hooks.ts`
- `packages/widget/src/utils/`
- `packages/widget/src/components/Widget.test.tsx`
- `packages/widget/src/components/Receipt.test.tsx` only if shared quote-selection refactors affect existing receipt coverage
- `packages/widget/README.md`
- A new repo-root spec file is this document; implementation should not add new public docs outside widget scope unless needed for package README clarity

## Edge cases and failure handling

- Empty input with native origin asset:
  - safe max must still be computable through the bootstrap pass
  - the visible quote/output area must not be overwritten by bootstrap-only quote requests
- Missing connected address, missing selections, or missing balance:
  - keep existing placeholder behavior
  - native safe max remains unavailable
- Quote returns no active routes after paused-module filtering:
  - native safe max remains unavailable
  - do not replace the current input
- Quote returns no executable `tx`:
  - native safe max remains unavailable
  - do not replace the current input
- Provider cannot return fee data or gas estimate:
  - native safe max remains unavailable
  - do not replace the current input
- Safe max computes to a negative value:
  - clamp to zero
  - show `0.0`
  - keep the control non-clickable
- User edits the amount or changes selections during calculation:
  - ignore stale completions
  - do not overwrite the newer user input with late async results
- Because manual input stays allowed by scope, users can still type a native amount above the suggested safe max and reach the wallet flow. This is accepted behavior for this feature.

## Phase plan

1. Extract widget-local quote-selection logic so the Redux quote thunk and the safe-max flow can share RFQ preference and paused-module filtering.
2. Add a widget-local native safe-max helper or hook that performs direct SDK quote calls, gas estimation, fee-data lookup, buffering, and stale-result cancellation.
3. Update `Widget.tsx` to own the native safe-max state machine and pass the right props into `AvailableBalance`.
4. Update `AvailableBalance.tsx` to render native loading, unavailable, and ready states while preserving non-native behavior.
5. Keep existing validation and send-button behavior unchanged, except for wiring the native balance control to the precomputed safe max.
6. Add or expand tests for bootstrap flow, refinement flow, stale-result suppression, non-native regression coverage, zero-safe-max behavior, and the prefill-only boundary.
7. Update `packages/widget/README.md` to note that native balance control shows bridgeable balance rather than raw wallet balance.

## Acceptance criteria

- Selecting a non-native origin token preserves the current balance-label and max-fill behavior.
- Selecting a native origin token causes the balance control to compute and display bridgeable amount rather than raw wallet balance.
- Clicking the native balance control fills the final refined safe max, not the raw native balance.
- The native safe-max calculation reserves both:
  - buffered estimated gas fee using a 150% multiplier
  - the selected quote's `nativeFee`
- Native safe-max bootstrap does not overwrite the visible destination quote or receipt while computing a prefill value.
- Manual native input above the safe max remains possible, and the widget's current raw-balance validation behavior remains unchanged.
- Selection or balance changes during safe-max computation do not allow stale async completions to overwrite the current state.
- When safe max cannot be computed, the widget does not fall back to filling raw native balance through the balance control.
- When safe max is not ready, the native balance control mirrors the non-native balance label behavior but remains non-clickable.

## Validation plan

- Add widget tests that mock SDK quote responses, fee data, and gas estimation to cover:
  - native empty-input bootstrap
  - native refinement pass
  - native zero-safe-max result
  - native unavailable state on quote failure
  - native unavailable state on gas-estimation failure
  - stale async result suppression after selection or input changes
  - non-native regression behavior
  - prefill-only behavior, proving manual over-safe-max input is still allowed
- Run widget tests with `cd packages/widget && npm test`.
- Run widget lint with `cd packages/widget && npm run lint:check`.
- Run widget build with `cd packages/widget && npm run build`.
- Manually verify in the widget example app that:
  - native balance control shows a loading state before safe max is ready
  - once wallet balance is available but safe max is not ready, native balance control shows the same `Available {parsedBalance}` label as non-native tokens without becoming clickable
  - native balance control shows bridgeable amount once ready
  - clicking it fills the reduced amount
  - the visible quote/receipt does not flicker to bootstrap-only data
  - non-native tokens still fill full balance

## Risks and assumptions

- Assumption: using current provider fee data with a 150% buffer is an acceptable approximation for origin-chain gas reservation in the widget.
- Assumption: one refinement pass is sufficient to align the final safe max with the quote used for the resulting prefill without introducing repeated auto-adjustment loops.
- Risk: some quote routes may change between provisional full-balance quoting and refined safe-max quoting. The bounded two-pass flow intentionally favors deterministic UI over repeated convergence attempts.
- Risk: some providers may fail gas estimation for specific populated transactions. In those cases the native balance control will remain unavailable rather than falling back to raw balance.
- Risk: because enforcement is intentionally out of scope, users who manually enter a larger native amount can still hit wallet-level insufficient-funds failures.
