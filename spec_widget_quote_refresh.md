# Widget Quote Refresh Fix

## Goal

Fix the widget's stale-quote refresh bug so a fresh quote is not re-fetched on every mouse movement, while preserving the intended widget UX:

- fetch a new quote when inputs change
- treat quotes as stale after the configured timeout
- refresh stale quotes only in response to real user activity when the wallet is not busy

This spec is intentionally TDD-first. The implementation must follow red-green-refactor order, with automated tests defining the behavior before the production hook changes land.

## Context

The widget quote flow lives in:

- `packages/widget/src/components/Widget.tsx`
- `packages/widget/src/hooks/useBridgeQuoteUpdater.ts`
- `packages/widget/src/hooks/useIntervalTimer.ts`
- `packages/widget/src/state/slices/bridgeQuote/hooks.ts`

Current behavior:

- `Widget.tsx` fetches quotes through `fetchAndStoreBridgeQuote()` and passes the current quote plus the refresh callback into `useBridgeQuoteUpdater`.
- `useBridgeQuoteUpdater` determines staleness from `quote.timestamp`, a timer value from `useIntervalTimer`, and `calculateTimeBetween`.
- Once a quote is considered stale, the hook adds a one-shot `mousemove` listener and refreshes the quote on the next mouse move.
- The timestamp stored on the quote is created before `synapseSDK.bridgeV2(...)` resolves, not when the quote is received.
- The stale comparison uses `Math.abs(currentTime - quoteTime)`, so a newly fetched quote can be treated as stale if the timer value is older than the quote by at least the stale timeout.

Observed failure mode:

- If the first valid quote arrives more than 15 seconds after the timer's initial baseline, or if a later quote arrives after a long gap, the new quote can be flagged stale immediately.
- The next mouse move refreshes the quote, but the timer baseline may still be older than the new quote, so the quote is still considered stale.
- This creates repeated quote refreshes on subsequent mouse movement, which is not the intended behavior.

Relevant comparison point:

- `packages/synapse-interface` uses a different stale-refresh strategy in `components/StateManagedBridge/hooks/useStaleQuoteUpdater.ts`.
- The interface implementation does not compute staleness from `quote.timestamp`; it uses local timers, an optional short auto-refresh window, and then falls back to a one-shot `mousemove` refresh.

## Scope

This spec covers fixes inside the `packages/widget` quote refresh flow only.

In scope:

- Stale-quote detection logic
- Quote lifecycle cleanup and stale scheduling
- Mousemove refresh arming behavior
- Widget-local test harness setup so `packages/widget` has a runnable `Jest` + `jsdom` test path
- Focused automated tests for the stale refresh hook and related timing logic
- One widget-level smoke test that verifies the hook is wired into the widget refresh flow correctly

## Non-goals

- Changing the public `Bridge` component API
- Changing normal quote fetch triggers based on amount, chain, token, wallet, or destination state
- Adding background polling or any automatic quote refresh window
- Porting or partially porting the `synapse-interface` bridge-page stale-refresh model
- Adding new stale-quote UI, countdown UI, or quote-confirmation UI
- Changing how quotes are selected from SDK results
- Changing request de-duping behavior or replacing `currentSDKRequestID`
- Refactoring quote fetch thunks, wallet state, allowance fetching, or maintenance logic beyond what is required by the stale-refresh hook
- Changing `packages/widget-example` or requiring consumer integration changes
- Broad browser E2E coverage for this fix
- Cleaning up the existing quote `timestamp` field as a prerequisite for this fix
- Broad dead-code cleanup outside the immediate quote-refresh path
- Refactoring unrelated wallet, allowance, or maintenance logic
- Removing the stale quote feature entirely

## Requirements

1. A quote fetched successfully by the widget must be treated as fresh for a full stale timeout starting from quote arrival, not from request start.
2. Idle mouse movement must not refresh a fresh quote.
3. After a quote becomes stale, at most one mousemove-driven refresh may fire for that stale cycle.
4. A new quote received after a mousemove-driven refresh must not remain stale immediately unless the stale timeout has actually elapsed since that quote arrived.
5. If quote fetching is in progress or the wallet is pending, stale-refresh listeners must not remain armed.
6. Existing user-facing widget behavior must be preserved:
   - input changes still fetch quotes
   - stale quotes still refresh only after user activity
   - there is no background polling for the widget
7. The fix must not require consumer changes in `packages/widget-example` or other widget integrations.
8. The implementation must keep request de-duping semantics based on `currentSDKRequestID`.
9. The implementation must be delivered with strict red-green-refactor TDD:
   - write or update failing automated tests first
   - make the minimum production change needed to pass them
   - refactor only after the tests are green
10. `packages/widget` must gain a runnable local automated test path based on `Jest` and `jsdom`.

## TDD Implementation Contract

The implementation sequence is part of the scope, not a suggestion.

Required delivery order:

1. Add widget-local test infrastructure in `packages/widget` so hook and widget behavior can run under `Jest` with `jsdom`.
2. Add the hook-level failing tests that capture the stale-refresh bug and the required cleanup behavior.
3. Add one widget-level failing smoke test that proves the hook remains correctly wired through `Widget.tsx`.
4. Update the production implementation until those tests pass.
5. Perform only scoped cleanup that is justified by the passing tests.

TDD evidence expectations:

- The PR should make it obvious which tests define the behavior under change.
- The new tests must fail against the pre-fix implementation and pass with the fix.
- Manual verification remains useful, but it is secondary confirmation and not a substitute for the automated coverage in this spec.

## Technical approach

The widget should preserve its current stale-on-user-activity UX, but the implementation should move away from the current interval-baseline math because that is the source of the bug and is harder to reason about than a quote-scoped timeout.

Implementation shape:

- Rework `packages/widget/src/hooks/useBridgeQuoteUpdater.ts` so each valid quote owns its own stale timeout lifecycle.
- Remove quote staleness dependence on `useIntervalTimer`, `quote.timestamp`, and `calculateTimeBetween`.
- When a valid quote arrives and refresh is allowed:
  - clear any previously armed stale timeout
  - remove any previously armed mousemove listener
  - start a `setTimeout(staleTimeout)` for the active quote
- When that timeout fires:
  - mark the active quote as stale inside the hook's local control flow
  - install a one-shot `mousemove` listener for the active stale cycle
- When the mousemove listener fires:
  - call the refresh callback once
  - clear the listener reference immediately
  - treat the current stale cycle as closed so later mouse movement does not retrigger until a new quote becomes stale
- When the quote changes, becomes invalid, starts loading, or wallet pending begins:
  - clear the stale timeout
  - remove any mousemove listener
  - reset any local stale-cycle bookkeeping

Invariants:

- A quote's stale timeout starts when that quote becomes the active valid quote in the hook, not from widget mount time and not from request-start time.
- Only the currently active quote may arm or own a stale cycle.
- Loading and wallet-pending states are hard stops for stale-refresh scheduling.
- The hook must remain self-contained so `Widget.tsx` does not need consumer-visible behavior changes.

Rationale:

- Preserves the current widget UX: no background polling, stale quotes still refresh only after user activity.
- Eliminates the stale-baseline race entirely.
- Reduces long-term timing ambiguity by expressing stale behavior directly in the updater hook instead of coordinating multiple helpers.

## Affected areas

- `packages/widget/src/hooks/useBridgeQuoteUpdater.ts`
- `packages/widget/src/hooks/useIntervalTimer.ts` only if dead code cleanup or decoupling is performed as part of the refactor
- `packages/widget/src/components/Widget.tsx` only if hook wiring or imports need adjustment during the refactor
- `packages/widget/src/state/slices/bridgeQuote/hooks.ts` only if the implementation removes now-unused timestamp plumbing as a scoped follow-up
- `packages/widget/src/utils/calculateTimeBetween.ts` only if no longer needed after the refactor
- `packages/widget/src/state/slices/bridgeQuote/reducer.ts` if any state shape clarification is needed
- `packages/widget/package.json`
- A widget-local Jest config, test setup file, and new test files near the hook or widget layer
- `packages/widget/CHANGELOG.md` only if this repository expects changelog updates for behavior fixes

## Test scenarios

The following scenarios are required. They are the behavioral contract for the implementation.

### Hook-level scenarios

1. Fresh quote stays fresh for a full timeout from arrival
   - Start with a valid quote becoming active now.
   - Assert that mouse movement before `staleTimeout` does not call the refresh callback.
   - Assert that the quote only becomes refreshable after the timeout elapses.

2. Long idle before first quote does not make the first quote stale
   - Advance fake time well past `staleTimeout` before the first valid quote appears.
   - Activate the first valid quote.
   - Assert that immediate mouse movement still does not refresh it.
   - Assert that the quote gets its own full fresh window from arrival.

3. One stale cycle produces one mousemove refresh
   - Let a valid quote age past `staleTimeout`.
   - Trigger one mousemove event and assert exactly one refresh call.
   - Trigger additional mousemove events without a new quote and assert there are no extra refreshes.

4. Fresh quote after stale refresh is not immediately stale again
   - Age quote A until stale.
   - Fire the mousemove refresh for quote A.
   - Replace it with quote B.
   - Assert quote B does not refresh on immediate mouse movement and receives a full fresh timeout of its own.

5. Loading clears pending stale work
   - Arm a stale timeout or stale mousemove listener.
   - Switch `isQuoteLoading` to `true`.
   - Assert timers and listeners are cleared and mouse movement no longer triggers refresh while loading is active.

6. Wallet pending clears pending stale work
   - Arm a stale timeout or stale mousemove listener.
   - Switch `isWalletPending` to `true`.
   - Assert timers and listeners are cleared and mouse movement no longer triggers refresh while wallet pending is active.

7. Quote invalidation clears timers and listeners
   - Start with a valid quote and an active timeout or stale listener.
   - Replace the quote with an invalid quote state.
   - Assert all stale-cycle state is cleared immediately.

8. Quote replacement closes the old stale cycle
   - Start quote A and let it arm a stale cycle.
   - Replace it with quote B before or after A becomes stale.
   - Assert any timer or listener owned by quote A is removed.
   - Assert only quote B can later arm the next stale cycle.

9. Background rerenders do not re-arm fresh quote listeners
   - Rerender with the same active valid quote and unchanged stale inputs.
   - Assert no early refresh listener is attached and no extra callback fires.

### Widget-level smoke scenario

1. Widget wiring preserves normal fetch behavior and stale refresh behavior
   - Render `Widget.tsx` with controlled quote state and mocked quote-fetch flow.
   - Assert the normal input-driven quote fetch path still works.
   - After a valid quote arrives, assert immediate mouse movement does not refetch.
   - After the quote becomes stale, assert one mousemove triggers exactly one refresh through the widget's existing refresh path.
   - Assert no consumer-visible API changes are required to exercise the behavior.

## Edge cases and failure handling

- First quote arrives long after the widget mounted: the quote must still get a full fresh window before any stale listener is armed.
- Quote request latency exceeds the stale timeout: the quote must still be considered fresh when it arrives.
- Quote is refreshed while a previous stale listener is armed: old listeners must be removed before the new quote becomes active.
- Quote becomes invalid because inputs are cleared: any stale timers or listeners must be removed immediately.
- Wallet prompt begins while a stale listener is armed: the listener must be removed and not re-armed while wallet pending remains active.
- Rapid sequence of input changes and request IDs: only the latest quote may become active, and stale-refresh logic must follow the active quote only.
- Background rerenders unrelated to quote state must not re-arm stale listeners for a fresh quote.

## Phase plan

1. Add widget-local `Jest` + `jsdom` test infrastructure and a real test command in `packages/widget`.
2. Add failing hook tests for quote freshness, stale transition, one-shot mousemove refresh, and cleanup behavior.
3. Add one failing widget-level smoke test that proves the widget still uses the same refresh flow and does not refetch on fresh-quote mouse movement.
4. Rework `packages/widget/src/hooks/useBridgeQuoteUpdater.ts` so each valid quote owns its own stale timeout and one-shot mousemove refresh cycle.
5. Remove quote-staleness dependence on `useIntervalTimer` and `calculateTimeBetween`; retain or delete those helpers only if they still serve other code paths.
6. Make only the minimal call-site or state-shape changes needed to get the new tests green.
7. Manually verify the widget example with a long-lived page session and a slow quote scenario.

## Acceptance criteria

- A newly received quote is not treated as stale before `staleTimeout` has elapsed since the quote arrived.
- Moving the mouse repeatedly over a fresh quote does not trigger quote fetching.
- Once a quote is stale, exactly one refresh occurs on the next mouse move.
- After that refresh, additional mouse movement does not trigger more refreshes until the new quote becomes stale.
- Input changes, chain changes, token changes, and wallet changes still trigger quotes exactly as before.
- `packages/widget` has a runnable local `Jest` + `jsdom` test path that covers the required scenarios in this spec.
- No consumer code changes are required in `packages/widget-example`.

## Validation plan

Automated validation:

- Run the new widget-local test command and keep it green.
- Cover, at minimum:
  - quote becomes stale only after timeout from quote arrival
  - a quote fetched after a long initial idle period is still fresh
  - mousemove triggers one refresh only when stale
  - a newly refreshed quote does not remain stale immediately
  - loading cancels stale timers and listeners
  - wallet pending cancels stale timers and listeners
  - quote invalidation clears timers and listeners
  - a widget-level smoke path confirms the hook is wired correctly through `Widget.tsx`

Manual verification in `packages/widget-example`:

- load the page and wait more than 15s before entering an amount
- confirm the first quote does not refetch on immediate mouse movement
- wait until the quote is actually stale, then move the mouse once and confirm one refresh
- continue moving the mouse and confirm no repeated refreshes until the quote becomes stale again
- simulate a slow quote by throttling the network and confirm arrival time, not request start time, controls staleness

## Risks and assumptions

- Assumption: the intended widget UX is the current documented contract in `useBridgeQuoteUpdater` comments and `packages/widget/CHANGELOG.md`, not the richer `synapse-interface` auto-refresh behavior.
- Assumption: widget consumers expect no new background polling and no API changes.
- Accepted tradeoff: this is a slightly larger internal change than a timestamp-only patch, because it includes local test harness setup and TDD scaffolding to make the timing behavior durable.
- Risk: the hook refactor must be careful not to leave behind orphaned timeouts or listeners across rapid quote changes.
- Risk: if `useIntervalTimer`, `calculateTimeBetween`, or quote timestamp plumbing are still used elsewhere, cleanup should stay scoped and not broaden this bug fix unnecessarily.
