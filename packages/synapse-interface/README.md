This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app)

## Getting Started

First, run the development server:

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

This guide explains how to use the Maintenance feature to pause a chain or bridge module for `@synapsecns/synapse-interface` and `@synapsecns/widget` packages.

## How it works

There are a few maintenance components implemented around the Synapse Interface Webapp:
1. Banner - located at the top of the page.
2. Countdown Progress Bar - located at the top of Bridge / Swap cards.
3. Warning Message - located below the input UI in Bridge / Swap cards.

NOTE: Currently, the Synapse Widget implements the Countdown Progress Bar and Warning Message displayed in the Bridge Widget, but not the Banner.

These components ingest data fetched from the following JSON files:
- [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json)
- [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json)


To control when the Banner, Countdown Progress Bar, and Warning Message components are displayed, update the [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json).

To specify which bridge modules (SynapseRFQ, SynapseBridge, or SynapseCCTP) are paused, update the [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json).


After updating the proper JSON files, the following steps must be taken to ensure the production webapp reflects the changes made:
1. Merge the new branch into `master`
2. Merge `master` branch into `fe-release` branch

After Step 1 is completed, the [Github Pages](https://github.com/synapsecns/sanguine/deployments/github-pages) must finish building for the respective branch to take effect on the production webapp.

Although completing Step 1 will already reflect changes in the webapp, Step 2 is required in the slim chance that the github API is down, so that the production webapp can use the local JSON files as a reliable backup data source.


## Chain Pause

You can pause the Bridge and Swap functionalities on specific chains using their chainIds. Pauses can be applied independently to Bridge or Swap functions, or to both simultaneously.

For Bridge functionality, you can specify the origin and destination chainIds to pause. For Swap functionality, you can pause a chain either by including the specific chainId in either the origin or destination. You can set a start and end time for the pause, or leave it indefinite if the duration is uncertain.

Additionally, you can control which components are displayed during the pause event.


### Chain Pause Props

`id`
Unique ID used to distinguish maintenance component instances. Use 'EVENT_NAME-pause' format. (e.g: arbitrum-chain-pause)

`pausedFromChains`
An array containing origin chainIds to pause.

`pausedToChains`
An array containing destination chainIds to pause.

`pauseBridge`
Boolean indicating whether to pause Bridge functionality.

`pauseSwap`
Boolean indicating whether to pause Swap functionality.

`startTimePauseChain`
UTC time of when to start chain(s) pause.

`endTimePauseChain`
UTC time of when to end chain(s) pause. If null, chain(s) pause will continue indefinitely.

`startTimeBanner`
UTC time of when to start displaying Banner.

`endTimeBanner`
UTC time of when to end displaying Banner. If null, Banner will display indefinitely.

`inputWarningMessage`
String to display in Warning Message shown in Bridge or Swap card.

`bannerMessage`
String to display in Banner.

`progressBarMessage`
String to display in Countdown Progress Bar.

`disableBanner`
Boolean indicating whether to hide Banner.

`disableWarning`
Boolean indicating whether to hide Warning Message in Bridge or Swap card.

`disableCountdown`
Boolean indicating whether to hide Countdown Progress Bar.


### Example

`paused-chains.json`
```tsx
  [
    // Bridge Pause
    {
      "id": "base-chain-pause",
      "pausedFromChains": [8453],
      "pausedToChains": [8453],
      "pauseBridge": true,
      "pauseSwap": false,
      "startTimePauseChain": "2024-04-12T17:41:00Z",
      "endTimePauseChain": null,
      "startTimeBanner": "2024-04-12T04:40:00Z",
      "endTimeBanner": null,
      "inputWarningMessage": "Base bridging is paused until maintenance is complete.",
      "bannerMessage": "Base bridging is paused until maintenance is complete.",
      "progressBarMessage": "Base maintenance in progress",
      "disableBanner": false,
      "disableWarning": false,
      "disableCountdown": false
    },
    {
      "id": "ecotone-fork-pause",
      "pausedFromChains": [10, 8453],
      "pausedToChains": [10, 8453],
      "pauseBridge": true,
      "pauseSwap": false,
      "startTimePauseChain": "2024-03-13T23:35:00Z",
      "endTimePauseChain": "2024-03-14T00:25:00Z",
      "startTimeBanner": "2024-03-13T23:20:00Z",
      "endTimeBanner": "2024-03-14T00:25:00Z",
      "inputWarningMessage": "",
      "bannerMessage": "Optimism + Base Bridging will be paused 10 minutes ahead of Ecotone (March 14 00:00 UTC, 20:00 EST). Will be back online shortly following the network upgrade.",
      "progressBarMessage": "Ecotone Fork maintenance in progress",
      "disableBanner": false,
      "disableWarning": true,
      "disableCountdown": false
    },
    // Swap Pause
    {
      "id": "arbitrum-swap-pause",
      "pausedFromChains": [42161],
      "pausedToChains": [42161],
      "pauseBridge": false,
      "pauseSwap": true,
      "startTimePauseChain": "2024-03-13T23:35:00Z",
      "endTimePauseChain": "2024-03-14T00:25:00Z",
      "startTimeBanner": "2024-03-13T23:20:00Z",
      "endTimeBanner": "2024-03-14T00:25:00Z",
      "inputWarningMessage": "Swapping on Arbitrum is paused until maintenance is complete.",
      "bannerMessage": "Swapping on Arbitrum is paused until maintenance is complete.",
      "progressBarMessage": "Arbitrum maintenance in progress",
      "disableBanner": false,
      "disableWarning": false,
      "disableCountdown": false
    }
  ]
```


## Bridge Module Pause

You can pause a specific bridge module on a given chain. Currently, there are the following bridge modules:
- SynapseRFQ
- SynapseCCTP
- SynapseBridge


### Bridge Module Pause Props

`chainId`
Chain ID of Chain to pause specific bridge module.

`bridgeModuleName`
Accepts 'SynapseRFQ', 'SynapseBridge', 'SynapseCCTP', or 'ALL'. If selecting 'ALL', all bridge modules will be paused for respective chainId.


### Example

`paused-bridge-modules.json`
```tsx
  [
    {
      "chainId": 42161,
      "bridgeModuleName": "ALL"
    },
    {
      "chainId": 10,
      "bridgeModuleName": "SynapseRFQ"
    },
    {
      "chainId": 10,
      "bridgeModuleName": "SynapseCCTP"
    },
    {
      "chainId": 8453,
      "bridgeModuleName": "SynapseBridge"
    }
  ]
```
