# Remove Chain-Specific Pause Configs And Chain-Specific Warnings

## Goal

Remove the currently shipped chain-specific and global pause configuration data from the Synapse Interface maintenance artifacts, and remove the remaining hardcoded chain-specific bridge warnings, so the repo no longer surfaces obsolete chain-targeted warning states while the interface and widget remain operational.

## Context

`packages/synapse-interface` publishes maintenance data from `public/pauses/v1/paused-chains.json` and `public/pauses/v1/paused-bridge-modules.json`. The interface fetches those files in `components/Maintenance/hooks/useSynapsePauseData.ts`, stores them in Redux, and uses them for banner, warning, countdown, CTA-disabled, and bridge-quote filtering behavior.

`packages/widget` is coupled to the same artifacts. It fetches both files from the interface-owned URLs in `src/constants/index.ts` and `src/utils/getSynapsePauseData.ts`, then uses chain-pause data for warning/progress/CTA-disabled behavior and module-pause data for quote filtering.

Separately from the maintenance artifacts, `packages/synapse-interface/components/Warning.tsx` still renders hardcoded warnings for Harmony, Fantom, and Dogechain on the bridge page. Those warnings are not driven by the pause JSONs. Their copy is translated in every locale file under the `Warning` message namespace.

There is also an orphaned chain-warning definition in `packages/synapse-interface/constants/impairedChains.tsx` for Harmony. That file is not referenced elsewhere in the repository but still contains chain-specific warning content that should be cleaned up if the goal is to remove all repository-maintained chain-targeted warnings.

Current repository artifacts still ship:

- live chain-pause entries in `public/pauses/v1/paused-chains.json`
- a live global module pause in `public/pauses/v1/paused-bridge-modules.json`
- chain/global pause examples in `public/pauses/v1/examples/*`
- maintenance documentation in `README.md` that still instructs authors to use chain-pause and global-pause examples
- hardcoded bridge warnings for Harmony, Fantom, and Dogechain in `components/Warning.tsx`
- translated chain-warning copy in `messages/*.json`
- an unused Harmony impairment warning definition in `constants/impairedChains.tsx`

The package does not have comprehensive maintenance feature tests. Existing automated coverage is limited to the fetch helper in `components/Maintenance/functions/fetchJsonData.test.ts`, plus general package linting.

## Scope

This change removes shipped chain-specific and global pause configuration artifacts from `packages/synapse-interface` and removes repository-maintained hardcoded chain-specific warning content while preserving compatibility for existing consumers.

In scope:

- clear chain-pause entries from `public/pauses/v1/paused-chains.json`
- remove any shipped global pause entry from `public/pauses/v1/paused-bridge-modules.json`
- update or remove example artifacts that demonstrate chain-specific or global pause configs
- update `packages/synapse-interface/README.md` so the maintenance guide no longer documents the removed artifact content as the current repository state
- remove bridge-page hardcoded warnings for Harmony, Fantom, and Dogechain
- remove now-unused chain-warning translation keys from all supported locale files
- remove or neutralize the unused Harmony warning definition in `constants/impairedChains.tsx`
- document the widget audit result and keep widget code unchanged for this scoped cleanup

## Non-goals

- removing chain-pause runtime support from `packages/synapse-interface`
- removing maintenance UI or pause parsing from `packages/widget`
- changing pause JSON URLs, filenames, or fetch flow
- redesigning bridge-module pause support beyond removing shipped global config entries
- removing chain support, tokens, or chain metadata for Harmony, Fantom, or Dogechain
- adding a new maintenance configuration system

## Requirements

1. `packages/synapse-interface/public/pauses/v1/paused-chains.json` must remain present at the current path and contain no shipped chain-specific or global pause entries after the change.
2. `packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json` must remain present at the current path and contain no shipped global pause entry after the change.
3. All maintained JSON artifacts must remain valid JSON arrays so existing interface and widget fetchers continue to parse them without code changes.
4. Example artifacts under `packages/synapse-interface/public/pauses/v1/examples/` must no longer demonstrate chain-specific pause configs or global pause configs.
5. `packages/synapse-interface/README.md` must stop presenting chain-specific and global pause configs as active maintenance artifacts that operators should author from repository examples.
6. The spec implementation must not delete the shared pause artifact files themselves, because the widget and interface both fetch them directly and treat missing files as failures.
7. `packages/synapse-interface/components/Warning.tsx` must stop rendering chain-specific warnings for Harmony, Fantom, and Dogechain.
8. Any chain-specific warning translation keys that become unused after removing the `Warning.tsx` logic must be removed from all locale files in `packages/synapse-interface/messages/`.
9. The shared `WarningMessage` presentation component must remain available to existing non-chain-warning consumers, either by keeping it exported from its current module or relocating it with import updates.
10. `packages/synapse-interface/constants/impairedChains.tsx` must no longer retain obsolete Harmony-specific warning content after the change.
11. No `packages/widget` source removal is required for this scoped cleanup. The widget audit should be reflected in the implementation notes or PR summary, not as code deletion.

## Technical approach

Keep the maintenance artifact contract stable while removing the obsolete data.

- Replace the contents of `public/pauses/v1/paused-chains.json` with an empty array.
- Remove any repository-shipped global pause record from `public/pauses/v1/paused-bridge-modules.json`. If no non-global module pauses remain, leave the file as an empty array.
- Update example JSON files so they no longer include chain-pause examples or global `ALL` module-pause examples. If an example file would otherwise become misleading, prefer an empty array over deleting the file unless the implementation also removes all references to that file.
- Rewrite the maintenance section of `packages/synapse-interface/README.md` to match the post-change artifact state. The README should no longer include chain-pause property documentation or chain/global example payloads as repository-maintained artifacts.
- Remove the bridge-page chain-warning decision logic from `components/Warning.tsx`. Preserve the reusable `WarningMessage` presentation component for other consumers such as maintenance warning rendering and portfolio unsupported-network notices.
- Remove the `Warning` component usage from `pages/state-managed-bridge/index.tsx` so the bridge page no longer shows Harmony, Fantom, or Dogechain-specific hardcoded notices.
- Delete the now-unused chain-warning translations from every locale file under `packages/synapse-interface/messages/`.
- Remove or simplify `constants/impairedChains.tsx` so it no longer contains Harmony-specific warning copy. If the file is truly unused, deleting it is acceptable.
- Do not change `components/Maintenance/Maintenance.tsx`, `components/Maintenance/hooks/useSynapsePauseData.ts`, `pages/swap/index.tsx`, or widget maintenance consumers for this task. With empty arrays or removed global entries, those maintenance code paths become dormant but still compatible.

Implementation invariants:

- `paused-chains.json` and `paused-bridge-modules.json` must keep their existing paths.
- Their top-level schema must remain an array.
- The change must be safe for both the interface local fallback imports and the widget remote fetch path.

## Affected areas

- `packages/synapse-interface/public/pauses/v1/paused-chains.json`
- `packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json`
- `packages/synapse-interface/public/pauses/v1/examples/paused-chains-example.json`
- `packages/synapse-interface/public/pauses/v1/examples/paused-bridge-modules-example.json`
- `packages/synapse-interface/README.md`
- `packages/synapse-interface/components/Warning.tsx`
- `packages/synapse-interface/pages/state-managed-bridge/index.tsx`
- `packages/synapse-interface/constants/impairedChains.tsx`
- `packages/synapse-interface/messages/en-US.json`
- `packages/synapse-interface/messages/ar.json`
- `packages/synapse-interface/messages/es.json`
- `packages/synapse-interface/messages/fr.json`
- `packages/synapse-interface/messages/jp.json`
- `packages/synapse-interface/messages/tr.json`
- `packages/synapse-interface/messages/zh-CN.json`

Audited but not expected to change in this scope:

- `packages/synapse-interface/components/Maintenance/hooks/useSynapsePauseData.ts`
- `packages/synapse-interface/components/Maintenance/Maintenance.tsx`
- `packages/synapse-interface/slices/bridgeQuote/thunks.ts`
- `packages/synapse-interface/components/Portfolio/components/SingleNetworkPortfolio.tsx`
- `packages/widget/src/constants/index.ts`
- `packages/widget/src/utils/getSynapsePauseData.ts`
- `packages/widget/src/components/Maintenance/Maintenance.tsx`
- `packages/widget/src/components/Widget.tsx`
- `packages/widget/src/state/slices/bridgeQuote/hooks.ts`

## Edge cases and failure handling

- Do not delete `paused-chains.json`. If that file is removed instead of emptied, widget fetches will fail before fresh module-pause data can be stored.
- Do not change the JSON schema shape. Both interface and widget parse these artifacts as arrays of known object shapes and do not have a migration layer.
- If example artifacts are retained, they must not accidentally preserve a removed global `ALL` example or a chain-pause example that contradicts the README.
- Do not remove the generic `WarningMessage` export in a way that breaks existing imports from maintenance components or portfolio components.
- Removing translation keys requires updating all locale files consistently; leaving stale keys behind is acceptable at runtime but does not satisfy this cleanup scope.
- If documentation is reduced, it must not imply that widget code was cleaned up in the same change. Widget cleanup is only required if chain-pause support is intentionally retired later.

## Phase plan

1. Update live maintenance artifacts.
   Replace `paused-chains.json` with `[]`.
   Remove shipped global pause records from `paused-bridge-modules.json`, leaving `[]` if nothing else remains.
2. Update example artifacts.
   Remove chain-pause example content and any global-pause example content from the example JSON files while keeping the files valid and repo-consistent.
3. Remove hardcoded chain-specific warnings.
   Delete Harmony, Fantom, and Dogechain warning behavior from `components/Warning.tsx` and remove its usage from the bridge page while preserving the reusable `WarningMessage` UI primitive.
4. Clean up message and legacy warning definitions.
   Remove the now-unused chain-warning translation keys from all locale files and delete or neutralize the unused Harmony warning content in `constants/impairedChains.tsx`.
5. Update package documentation.
   Rewrite the maintenance guide in `packages/synapse-interface/README.md` so it matches the new artifact set and no longer advertises removed chain/global configs.
6. Validate compatibility.
   Confirm both JSON files still parse, the interface can ingest empty arrays, the bridge page no longer renders chain-specific hardcoded warnings, and no widget source change is necessary for this scoped cleanup.

## Acceptance criteria

- `paused-chains.json` exists and contains `[]`.
- `paused-bridge-modules.json` exists and contains no global `ALL` pause entry.
- No example JSON file in `public/pauses/v1/examples/` contains a chain-pause example or a global pause example.
- `packages/synapse-interface/README.md` no longer documents chain-specific or global pause configs as maintained repository artifacts.
- The bridge page no longer renders Harmony, Fantom, or Dogechain-specific hardcoded warnings.
- `constants/impairedChains.tsx` no longer contains Harmony-specific warning copy.
- The chain-warning translation keys have been removed from every locale file that previously defined them.
- `packages/widget` source files are unchanged by this task.
- The resulting artifact set remains consumable by the existing interface and widget maintenance fetchers without path or schema changes.

## Validation plan

- Parse all modified JSON files with a JSON validator or a simple Node read/parse check.
- Run `yarn --cwd packages/synapse-interface lint:check` after documentation and JSON updates.
- Manually verify that the interface no longer surfaces a maintenance banner, warning, countdown, or disabled CTA solely because of repository-shipped chain/global pause artifact data.
- Manually verify that selecting Harmony, Fantom, or Dogechain on the bridge page no longer renders the old hardcoded warning block.
- Manually verify that the widget still loads maintenance data successfully when the shared artifact files exist but are empty.

## Risks and assumptions

- Assumption: the requested change includes removing all chain-specific warning copy and UI in `packages/synapse-interface`, but does not require chain support removal.
- Assumption: keeping empty files at the existing paths is acceptable and preferred over deleting the artifacts.
- Risk: if product intent is actually to retire chain-pause support entirely, this scoped cleanup will be incomplete because both interface and widget still contain dormant chain-pause code paths.
- Risk: the widget currently fetches chain and module pause data together. Any future attempt to remove only `paused-chains.json` without refactoring widget fetch behavior would also break fresh module-pause ingestion.
