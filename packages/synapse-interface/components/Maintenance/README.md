# Maintenance Instructions

This explains how to utilize the Maintenance feature the Synapse Interface Webapp in order to pause a chain or bridge module.

## How it works

There are a few maintenance components implemented around the app:
1. Banner - located at the top of the page.
2. Countdown Progress Bar - located at the top of Bridge / Swap cards.
3. Warning Message - located below the input UI in Bridge / Swap cards.

These components ingest data fetched from the following JSON files:

- [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json)
- [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json)

To control when the Banner, Countdown Progress Bar, and Warning Message components are displayed, update the [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json).

To specify which bridge modules (SynapseRFQ, SynapseBridge, or SynapseCCTP) are paused, update the [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json).

After updating the proper JSON files, the following steps must be taken to ensure the production webapp reflects the changes made:
1. Merge branch in `master`
2. Merge `master` branch into `fe-release` branch

Although Step 1 will already reflect changes in the webapp, Step 2 is required in the slim chance that the github API is down, so that the production webapp can use the local JSON files as a reliable backup data source.

## Chain Pause

You can pause the Bridge and Swap functionalities on specific chains using their chainIds. Pauses can be applied independently to Bridge or Swap functions, or to both simultaneously.

For Bridge functionality, you can specify the origin and destination chainIds to pause. You can set a start and end time for the pause, or leave it indefinite if the duration is uncertain.

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
Boolean indicating whhether to hide Warning Message in Bridge or Swap card.

`disableCountdown`
Boolean indicating whether to hide Countdown Progress Bar.

### Example

`paused-chains.json`
```tsx
  [
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
      "inputWarningMessage": "Base bridging is paused until maintenance is complete.",
      "bannerMessage": "Optimism + Base Bridging will be paused 10 minutes ahead of Ecotone (March 14 00:00 UTC, 20:00 EST). Will be back online shortly following the network upgrade.",
      "progressBarMessage": "Base maintenance in progress",
      "disableBanner": false,
      "disableWarning": false,
      "disableCountdown": false
    }
  ]
```


## Bridge Module Pause

You are able to pause a specific bridge module on a given chain. Currently, there are the following bridge modules:
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
  },
]
```
