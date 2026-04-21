This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app)

## Getting Started

First, run development server:

```bash
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.tsx`. The page auto-updates as you edit and save the file.

---

# Local Development Setup Guide

This guide is for running `@synapsecns/synapse-interface` and `@synapsecns/sdk-router` simultaneously for local development, with continuous recompilation on changes.

## Prerequisites

Ensure you've installed Node.js (version 18.17.0) and Yarn on your machine. This setup assumes you're using Yarn Workspaces and Lerna to manage your project, with `@synapsecns/sdk-router` and `@synapsecns/synapse-interface` as part of the same workspace.

## Steps

1. **Install dependencies**
   From the root directory of your workspace, run:

```shell
yarn install
```

This will handle dependency installation and local package linking.

2. **Watch for changes in `@synapsecns/sdk-router`**
   Open a terminal, navigate to the workspace root, and run:

```shell
lerna run --scope @synapsecns/sdk-router start --stream
```

This triggers TSDX in watch mode for `@synapsecns/sdk-router`, triggering rebuilds on file changes.

3. **Run the Next.js application in development mode**
   In a separate terminal window, navigate to the `synapse-interface` directory and start the dev server:

```shell
yarn dev
```

This command watches for file changes and automatically rebuilds the application, including updated dependencies.

After completing these steps, any changes to `@synapsecns/sdk-router` will be automatically detected and rebuilt. The `@synapsecns/synapse-interface` application will then pick up and incorporate these updates.

Make sure the `@synapsecns/sdk-router` dependency in `synapse-interface`'s `package.json` is declared by name and version (like `"@synapsecns/sdk-router": "0.1.0"`), matching `sdk-router`'s `package.json` version.

---

# Maintenance Guide

This guide covers the shared maintenance data consumed by `@synapsecns/synapse-interface` and `@synapsecns/widget`.

## How it works

The interface renders maintenance UI from shared pause artifacts:

1. Banner at the top of the page.
2. Countdown Progress Bar at the top of Bridge / Swap cards.
3. Warning Message below the input UI in Bridge / Swap cards.

The widget still depends on these same pause artifact files and repository paths for its maintenance state, even though it fetches them through different remote URLs. Keep both files present at their current paths even when they are empty, because the widget is coupled to those artifact locations.

The shared artifacts are:

- [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json)
- [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json)

Current repository state:

- `paused-chains.json` is intentionally shipped as an empty array. This repository no longer maintains chain-specific pause records in the shared artifacts.
- `paused-bridge-modules.json` remains the maintained pause artifact for bridge module filtering. Records must stay chain-scoped; do not add a global `ALL` pause record.
- Example files under `public/pauses/v1/examples/` are intentionally empty arrays so they do not advertise removed chain-specific or global configs.

When a shared pause artifact changes, the production webapp picks it up after:

1. Merging the branch into `master`.
2. Merging `master` into `fe-release`.

This deployment flow applies to the interface only. After Step 1, the [Github Pages](https://github.com/synapsecns/sanguine/deployments/github-pages) deployment for the branch must finish before the production webapp reflects the new data. Step 2 preserves the interface's local fallback path if the GitHub API is unavailable.

## Bridge Module Pause

Use `paused-bridge-modules.json` to pause a specific bridge module on a specific chain. Supported bridge modules are:

- `SynapseRFQ`
- `SynapseCCTP`
- `SynapseBridge`

### Bridge Module Pause Props

`chainId`
Chain ID of the chain where the bridge module should be paused.

`bridgeModuleName`
Accepts `SynapseRFQ`, `SynapseBridge`, or `SynapseCCTP`.

### Example

```json
[]
```

## Adding a New Language

1. Add the new locale code to `next.config.js`.
2. Include the new locale in the `LanguageSelector`.
3. Populate `/messages/{locale.json}` translations for the new locale. The keys in this should match `en-US.json`. See some of the other language files for reference.
