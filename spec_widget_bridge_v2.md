<!-- markdownlint-disable MD013 -->

# Widget BridgeV2 Migration

## Goal

Update `@synapsecns/widget` to source bridge quotes and executable bridge
transaction payloads from `sdk-router` `bridgeV2`, while preserving the widget's
existing UI, Redux flow, approval flow, transaction tracking flow, and public
package API.

## Context

- The widget currently fetches quotes with `synapseSDK.allBridgeQuotes(...)`,
  applies slippage locally, and submits transactions through
  `synapseSDK.bridge(...)`.
- The widget tracks submitted bridges separately through `getSynapseTxId(...)`
  and `getBridgeTxStatus(...)`, and that same tracking pattern is still used by
  `packages/synapse-interface`.
- `bridgeV2` returns a `BridgeQuoteV2` object with `expectedToAmount`,
  `minToAmount`, `routerAddress`, `moduleNames`, and an optional populated `tx`
  when `fromSender` is provided; it does not return `originQuery` / `destQuery`.
- In the widget, the connected wallet address is currently both the sender and
  the destination address for executed bridge transactions.
- The widget currently allows quote display before a wallet is connected. That
  behavior should remain unless `bridgeV2` makes it impossible.
- Maintenance filtering in the widget is module-based today. With `bridgeV2`,
  quotes can include multiple module names, so filtering must evaluate the
  returned `moduleNames` array.
- There is no established internal spec directory in this repository; this spec
  is intentionally stored as a root-level `spec_*.md` document.

## Scope

- Migrate widget quote fetching from legacy `allBridgeQuotes` output to
  `bridgeV2` output.
- Migrate widget bridge execution from `synapseSDK.bridge(...)` to the populated
  `tx` returned by `bridgeV2`.
- Preserve the existing widget quote refresh, approval, bridge submit, and
  transaction status polling flows unless a change is required by `bridgeV2`.
- Ensure quotes are invalidated and refreshed when the connected wallet address
  changes.
- Keep transaction status tracking functional for bridges initiated from
  `bridgeV2` quotes.

## Non-goals

- Do not redesign the widget UI, props, theming, or consumer integration API.
- Do not add destination-address input support to the widget.
- Do not replace the existing status-tracking mechanism with a new protocol or
  service.
- Do not port all `synapse-interface` bridge behaviors wholesale when they would
  change current widget behavior unnecessarily.
- Do not introduce new test infrastructure for `packages/widget`.

## Requirements

- The widget quote thunk must call `synapseSDK.bridgeV2(...)` instead of
  `synapseSDK.allBridgeQuotes(...)`.
- Quote requests must preserve current pre-connect behavior:
  - When no wallet is connected, the widget may still fetch and display quotes.
  - When a wallet is connected, the quote request must include `fromSender` and
    `toRecipient` set to the connected address so the returned quote contains
    address-specific transaction data.
- Quote requests must pass `slippagePercentage: 0.1` to preserve the current
  effective slippage behavior that was previously applied through
  `applyBridgeSlippage(...)`.
- Quote selection must preserve the widget's current route preference:
  - Filter out paused quotes first.
  - Prefer an active quote whose `moduleNames` includes `SynapseRFQ`.
  - If no active RFQ quote exists, use the first remaining active quote, relying
    on `bridgeV2` sort order by best output.
- Paused-module filtering must treat a quote as paused when any returned
  `moduleNames` entry matches a paused bridge module for the origin chain.
- The widget quote state must continue to expose the fields the current UI and
  bridge flow depend on:
  - displayed output amount
  - display string
  - router address
  - exchange rate
  - estimated time
  - bridge module name
  - request id
  - quote timestamp
- The widget quote state must additionally retain the `bridgeV2` populated
  transaction payload required for execution when a connected address is
  available.
- The widget must derive the stored `bridgeModuleName` from the returned
  `moduleNames` array using the final bridge module entry, consistent with
  `synapse-interface`.
- The widget must keep its separate allowance flow:
  - continue fetching allowance outside the quote thunk
  - continue using `bridgeQuote.routerAddress` as the spender
  - continue using the current wallet slice and approval validation flow
- Bridge execution must submit the populated transaction returned by the
  selected `bridgeV2` quote and must no longer construct bridge calldata through
  `synapseSDK.bridge(...)`.
- If the wallet is connected but the currently stored quote does not contain an
  executable `tx`, the bridge action must not proceed until a connected-address
  quote refresh completes.
- When the connected wallet address changes, the widget must:
  - clear the current quote immediately
  - issue a fresh quote request when the current selections and amount remain
    valid
  - ensure the new quote is keyed to the new address before it becomes
    actionable
- Existing stale-quote refresh behavior must continue to work with `bridgeV2`
  quotes.
- Existing transaction tracking must continue to work:
  - keep storing transaction metadata after submit
  - keep resolving `kappa` / Synapse tx id through `getSynapseTxId(...)`
  - keep polling completion through `getBridgeTxStatus(...)`
- Native-token routes must continue to work by submitting the `value` already
  encoded in the `bridgeV2` populated transaction.

## Technical approach

- Update `packages/widget/src/state/slices/bridgeQuote/hooks.ts` to map
  `BridgeQuoteV2` into the widget's existing quote shape instead of mapping
  legacy `BridgeQuote`.
- The quote-mapping logic should:
  - call `bridgeV2` with `fromChainId`, `toChainId`, token addresses,
    stringified input amount, and `slippagePercentage: 0.1`
  - include `fromSender` and `toRecipient` only when `connectedAddress` is
    available
  - filter paused routes by `quote.moduleNames`
  - preserve widget RFQ-first selection by checking
    `moduleNames.includes('SynapseRFQ')`
  - compute display values from `expectedToAmount`
  - store the final bridge module as `moduleNames[moduleNames.length - 1]`
  - store the returned `tx` payload for later submission
- Replace quote fields that are only meaningful for the legacy flow:
  - `originQuery` and `destQuery` should no longer drive execution
  - any reducer types that still include them should be removed or retired from
    the active flow
- Update `packages/widget/src/state/slices/bridgeQuote/reducer.ts` so the stored
  quote type reflects the `bridgeV2` execution payload while keeping existing
  fields required by `Receipt`, `BridgeButton`, and stale quote refresh.
- Update `packages/widget/src/components/Widget.tsx` to:
  - pass `connectedAddress` into quote fetching
  - add `connectedAddress` to the quote-fetch trigger dependencies
  - continue resetting the quote before issuing a fresh request
  - continue refreshing allowance off `bridgeQuote.routerAddress`
  - pass the stored quote transaction payload into bridge execution
- Update `packages/widget/src/state/slices/bridgeTransaction/hooks.ts` to submit
  the stored `bridgeV2` populated transaction directly through the signer
  instead of calling `synapseSDK.bridge(...)`.
- Keep bridge transaction state output unchanged enough for
  `useTransactionListener`, `Transactions`, and `Transaction` to continue
  working without a behavioral rewrite.
- Keep `packages/widget/src/hooks/useBridgeTxStatus.tsx` on the existing
  `getSynapseTxId(...)` plus `getBridgeTxStatus(...)` path. No new status API
  should be introduced in this scope because the repository reference
  implementation still uses the same tracking contract.
- If necessary, add a minimal validity guard where bridge execution currently
  assumes all valid quotes are executable so that a disconnected quote cannot be
  used after a wallet connects and before the refreshed address-bound quote
  arrives.

## Affected areas

- `packages/widget/src/state/slices/bridgeQuote/hooks.ts`
- `packages/widget/src/state/slices/bridgeQuote/reducer.ts`
- `packages/widget/src/components/Widget.tsx`
- `packages/widget/src/state/slices/bridgeTransaction/hooks.ts`
- `packages/widget/src/components/BridgeButton.tsx` if an explicit
  executable-quote guard is needed
- `packages/widget/src/components/Receipt.tsx` only if quote typing changes
  require small updates
- `packages/widget/src/hooks/useBridgeTxStatus.tsx` only for type alignment, not
  for a flow rewrite
- `packages/widget/src/state/slices/transactions/*` only if bridge-module
  metadata typing requires adjustment

## Edge cases and failure handling

- If `bridgeV2` returns no active quotes after paused-module filtering, the
  widget must continue to surface an invalid quote state and disable sending.
- If the selected quote does not include `tx` because no `fromSender` was
  supplied, the widget may still display the quote but must not submit it until
  the connected-address refresh completes.
- If the connected address changes during an in-flight quote request, the stale
  quote must not remain actionable after the address change.
- If an approval succeeds, the existing allowance refresh behavior must continue
  to update approval state without forcing unrelated UI changes.
- If bridge submission fails after a `bridgeV2` quote was selected, the widget
  must preserve its current error handling and wallet-pending cleanup behavior.
- If a route is native-token based, the submitted signer payload must honor the
  `value` returned by the quote instead of recomputing it locally.
- Transaction status polling must continue to tolerate temporary failures by
  keeping the transaction pending until a positive completion result is observed
  or the user clears the item.

## Phase plan

1. Update quote-state types and thunk inputs in
   `packages/widget/src/state/slices/bridgeQuote/*` to accept
   `connectedAddress`, call `bridgeV2`, and map `BridgeQuoteV2` into widget
   state.
2. Update `packages/widget/src/components/Widget.tsx` to trigger quote fetches
   on connected-address changes, keep the current reset-and-refetch flow, and
   continue allowance refresh from `routerAddress`.
3. Update `packages/widget/src/state/slices/bridgeTransaction/hooks.ts` so
   bridge submit uses the stored `bridgeV2` populated transaction.
4. Add any minimal button or execution guard needed so only executable
   connected-address quotes can be submitted.
5. Verify that transaction tracking still records `bridgeModuleName`,
   `estimatedTime`, hash, and completion status without changing the existing
   tracking UX.
6. Run package validation and manual bridge-flow checks.

## Acceptance criteria

- With valid selections and amount, the widget fetches quotes from `bridgeV2`
  rather than `allBridgeQuotes`.
- The widget still displays quotes before wallet connection.
- After wallet connection, the widget refreshes the quote so the stored quote is
  tied to the connected address and contains a usable transaction payload.
- Switching the connected account while the widget has a valid quote clears the
  old quote and refreshes it for the new address.
- The widget still prefers RFQ routes when an active RFQ quote is available,
  matching current widget behavior.
- Approval still targets the selected quote's `routerAddress` and the button
  state still updates after allowance refresh.
- Executing a bridge sends the `bridgeV2` populated transaction instead of
  generating calldata through `synapseSDK.bridge(...)`.
- Submitted transactions still appear in the widget transaction list and still
  resolve to complete through the existing status polling flow.
- Native-token routes continue to submit successfully with the correct
  transaction value.
- No public widget props or documented integration entry points change as part
  of this work.

## Validation plan

- Run `yarn --cwd packages/widget lint:check`.
- Run `yarn --cwd packages/widget build`.
- Manually verify in an existing widget example app that:
  - a quote appears while disconnected
  - connecting a wallet refreshes that quote
  - changing accounts refreshes that quote again
  - an ERC20 route still follows approve then bridge
  - a native-token route still bridges successfully
  - a submitted transaction continues to move from pending to complete in the
    widget UI

## Risks and assumptions

- Assumption: `bridgeV2` remains the supported quote and populated-transaction
  entry point, while `getSynapseTxId(...)` and `getBridgeTxStatus(...)` remain
  the supported status-tracking path. This matches the current `sdk-router` and
  `synapse-interface` implementation.
- Assumption: the last entry in `moduleNames` is the bridge module that should
  continue to be stored as `bridgeModuleName` for tracking and UI labels.
- Assumption: preserving current widget behavior means keeping RFQ-first route
  preference instead of adopting `synapse-interface`'s separate RFQ-vs-non-RFQ
  threshold heuristic.
- Risk: a disconnected quote can be displayed without `tx`; the implementation
  must explicitly prevent that quote from being used after connect until the
  refreshed quote arrives.
- Risk: quote requests are asynchronous, so address changes can race with stale
  responses; the implementation should rely on the existing reset/request-id
  pattern and not leave old quotes actionable after address changes.
