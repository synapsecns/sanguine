# Maintenance Instructions

This explains how to utilize the Maintenance feature on Synapse Protocol's Webapp in order to pause a chain or bridge module.

## How it works

There are a few maintenance components we utilized around the app:
1. Banner - located at the top of the page.
2. Countdown Progress Bar - located at the top of Bridge / Swap cards.
3. Warning Message - located below the input UI in Bridge / Swap cards.

These components ingest data fetched from the following JSON files:

Pause Chains - [JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json)

Pause Bridge Modules - [JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json)

In order to update when / if the Banner, Countdown Progress Bar, and Warning Message components are shown, you will need to update the [Pause Chains JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-chains.json).

In order to update which bridge modules are paused (SynapseRFQ, SynapseBridge, or SynapseCCTP), you will need to update the [Pause Bridge Modules JSON](https://github.com/synapsecns/sanguine/blob/master/packages/synapse-interface/public/pauses/v1/paused-bridge-modules.json)


## Chain Pause Props
`id`
Unique ID used to distinguish maintenance component instances. Use 'EVENT_NAME-pause' format. (e.g arbitrum-chain-pause)

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

## Bridge Module Pause Props

`chainId`
Chain ID of Chain to pause specific bridge module.

`bridgeModuleName`
Accepts 'SynapseRFQ', 'SynapseBridge', 'SynapseCCTP', or 'ALL'. If selecting 'ALL', all bridge modules will be paused for respective chainId.

