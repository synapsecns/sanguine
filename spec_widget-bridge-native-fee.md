# Widget Bridge Native Fee In Quote Receipt

## Goal

Show a non-zero bridge quote `nativeFee` in the widget so users can see the additional origin-chain native asset cost before submitting, and make the adjacent `Send` and `Receive` receipt rows self-describing by including their token symbols, without changing the current bridge form flow or CTA behavior.

## Context

- The widget package lives in `packages/widget` and fetches quotes through `synapseSDK.bridgeV2(...)` in `packages/widget/src/state/slices/bridgeQuote/hooks.ts`.
- The SDK already exposes `nativeFee` on `BridgeQuoteV2` in `packages/sdk-router/src/types/bridge.ts`, but the widget currently discards that field when normalizing quote state.
- The current quote details surface is the expandable receipt in `packages/widget/src/components/Receipt.tsx`. It currently renders `Router`, `Origin`, `Destination`, `Send`, and `Receive`.
- The `Send` and `Receive` values are currently passed from `packages/widget/src/components/Widget.tsx` as formatted numeric strings without token symbols, so those rows are less explicit than the new `Bridge fee` row will be.
- The collapsed receipt summary only shows estimated time and router; it does not currently show fee details.
- The widget package is English-only today and does not currently validate whether the connected wallet has enough origin native balance to cover quote-side native fees.
- The bridge execution path already uses `quote.tx.value` when present in `packages/widget/src/state/slices/bridgeTransaction/hooks.ts`, so the native fee is operationally relevant even though it is not shown in the UI.

## Scope

- Persist `nativeFee` from the selected `bridgeV2` quote into widget bridge quote state.
- Render a new `Bridge fee` row in the expanded receipt when `nativeFee` is non-zero.
- Format the fee as `<amount> <origin native symbol>`, for example `0.00042 ETH`.
- Update the expanded receipt so `Send` shows `<amount> <origin token symbol>` and `Receive` shows `<amount> <destination token symbol>`.
- Keep the fee as a separate line item from `Send`.
- Keep the fee display limited to the expanded receipt details and do not alter the collapsed summary.

## Non-goals

- Changing the collapsed receipt summary content.
- Changing the meaning of the existing `Send` row.
- Adding a `Total wallet spend` row.
- Adding native-balance validation, warnings, or CTA disabling based on `nativeFee`.
- Adding fiat conversion for the fee.
- Adding localization or translating the new label in this change.
- Changing the token selector dropdown UI.

## Requirements

1. When the selected bridge quote contains a `nativeFee` greater than zero, the widget must show a `Bridge fee` row in the expanded receipt.
2. The `Bridge fee` row must appear only in the expanded receipt details, not in the collapsed summary line.
3. The `Bridge fee` row must be rendered as a separate line item and must not modify the displayed value of `Send`.
4. The `Send` row must display the formatted amount plus the selected origin token symbol, for example `100 USDC`.
5. The `Receive` row must display the formatted amount plus the selected destination token symbol, for example `99.8 USDT`.
6. The `Bridge fee` row value must use the origin chain native currency symbol from widget chain metadata and display as `<amount> <symbol>`.
7. The `Bridge fee` row must be hidden when:
   - there is no valid quote,
   - the quote is loading,
   - `nativeFee` is absent, invalid, or equal to zero.
8. The `Bridge fee` row must be positioned with the other quote details in the expanded receipt and should appear after `Send` and before `Receive`.
9. The displayed `Bridge fee` amount must not collapse a non-zero fee into a misleading `0 <symbol>` or `0.0000 <symbol>` value.
10. Existing bridge execution, approval flow, and button enabled/disabled behavior must remain unchanged.

## Technical approach

- Extend the widget-local bridge quote state in `packages/widget/src/state/slices/bridgeQuote/reducer.ts` to include a `nativeFee: bigint` field on `BridgeQuote`, with `0n` in `EMPTY_BRIDGE_QUOTE`.
- In `packages/widget/src/state/slices/bridgeQuote/hooks.ts`, read `quote.nativeFee` from the selected SDK quote and normalize it into widget state as `BigInt(quote.nativeFee)`.
- Keep `nativeFee` as raw bigint state rather than pre-formatting it in the thunk. Formatting should remain a presentation concern in the receipt component.
- Update `packages/widget/src/components/Widget.tsx` so the `send` and `receive` values passed into `Receipt` include the selected origin and destination token symbols.
- Update `packages/widget/src/components/Receipt.tsx` to:
  - render the `Send` and `Receive` strings as amount-plus-symbol values,
  - derive the origin chain native symbol from `CHAINS_BY_ID[originChainId]?.nativeCurrency.symbol`,
  - compute a formatted fee string from `quote.nativeFee`,
  - conditionally render a `Bridge fee` row only when the quote is valid and the fee is greater than zero.
- Use existing widget formatting utilities where possible. If the existing `formatBigIntToString(..., 4)` style would produce an unreadable zero-like display for tiny but non-zero fees, add a small widget-local helper or logic to preserve enough significant decimals for visibility.
- Do not add any new props to the public widget API. This is an internal quote-state and receipt rendering change.
- Do not add new validation paths to `useValidations.tsx` or `BridgeButton.tsx`.

## Affected areas

- `packages/widget/src/state/slices/bridgeQuote/reducer.ts`
- `packages/widget/src/state/slices/bridgeQuote/hooks.ts`
- `packages/widget/src/components/Widget.tsx`
- `packages/widget/src/components/Receipt.tsx`
- `packages/widget/src/utils/formatBigIntToString.ts` or a new nearby widget-local formatting helper if additional precision logic is required
- `packages/widget/src/constants/chains.ts` only if a chain metadata gap is discovered during implementation

## Edge cases and failure handling

- If the SDK returns `nativeFee` as `'0'`, the receipt must not show the `Bridge fee` row.
- If the SDK returns a malformed or missing `nativeFee`, the widget should treat it as zero for rendering purposes rather than crash the receipt.
- If origin chain metadata is unexpectedly missing a native symbol, the implementation may fall back to showing the formatted amount without a symbol, but this should be treated as defensive fallback rather than expected behavior.
- If origin or destination token metadata is unexpectedly missing a token symbol, the implementation may fall back to the amount-only `Send` or `Receive` string rather than blocking receipt rendering.
- For very small non-zero fees, the formatting must preserve enough precision to communicate that a fee exists.
- Routes with zero native fee must continue to render exactly as they do today.
- The widget must continue to submit the quote transaction exactly as returned by the SDK; this feature is informational only.

## Phase plan

1. Extend widget bridge quote state to persist `nativeFee` from SDK quote normalization.
2. Update the receipt input props so `Send` and `Receive` render as amount plus token symbol.
3. Implement receipt-side formatting for native fee display using origin chain native symbol.
4. Add the conditional `Bridge fee` row to the expanded receipt between `Send` and `Receive`.
5. Verify that zero-fee routes do not show the row and non-zero-fee routes do.
6. Run widget package lint/build validation and perform manual UI verification with at least one zero-fee and one non-zero-fee route.

## Acceptance criteria

- A quote with `nativeFee = 0` shows no new fee row in the expanded receipt.
- A quote with `nativeFee > 0` shows a `Bridge fee` row in the expanded receipt.
- The `Send` row shows the origin token symbol alongside the amount.
- The `Receive` row shows the destination token symbol alongside the amount.
- The `Bridge fee` row label is exactly `Bridge fee`.
- The `Bridge fee` value is shown as amount plus origin native symbol, such as `0.00042 ETH`.
- The collapsed summary remains unchanged.
- The `Send` row remains semantically the input token amount only and is not redefined to include the native fee.
- No passive warning, validation error, or CTA disabled state is introduced for insufficient native balance as part of this change.
- Existing bridge and approval flows continue to function without API changes to widget consumers.

## Validation plan

- Run `npm --prefix packages/widget run lint:check`.
- Run `npm --prefix packages/widget run build`.
- Manually verify in the widget UI:
  - one route whose selected quote has `nativeFee = 0`,
  - one route whose selected quote has `nativeFee > 0`,
  - a route with a very small non-zero native fee to confirm the display is not rounded to zero,
  - a native-origin and ERC20-origin route if available,
  - that `Send` and `Receive` both show the correct token symbols after changing token selections.
- Confirm the collapsed receipt summary text is unchanged before and after the feature.
- Confirm no CTA state changes occur when the fee is present.

## Risks and assumptions

- Assumption: `synapseSDK.bridgeV2(...)` continues returning `nativeFee` as a decimal string compatible with `BigInt(...)`.
- Assumption: widget-supported chains in `CHAINS_BY_ID` provide native currency symbols suitable for display.
- Assumption: selected bridgeable tokens always expose stable `symbol` values suitable for receipt display.
- Risk: naive formatting could make small non-zero fees appear as zero, which would undermine the UX goal; implementation should explicitly guard against this.
- Risk: some routes may encode total native spend in `quote.tx.value` beyond user expectations, but this spec intentionally limits scope to displaying the quoted `nativeFee`, not reconciling all wallet outflow semantics.
