---
title: Bridge
---

import { BridgeFlow } from '@site/src/components/BridgeFlow'
import SVGBridge from '@site/src/components/SVGBridge'

# Synapse Bridge

The [Synapse Bridge](https://synapseprotocol.com) and [Solana Bridge](https://solana.synapseprotocol.com/) seamlessly swap on-chain assets between [20+ EVM and non-EVM blockchains](./Supported-Routes) in a safe and secure manner.

<br />

<figure>
    <BridgeFlow />
    <figcaption>User assets are sent to a bridge contract, moved to the `destChain`, and returned to the user.</figcaption>
</figure>

## Developers

Add the [Custom Widget](#) to your DeFi application, or build your own DeFi applications using the [Synapse SDK](#).

## Bridge Functions

The [Synapse Router](#) will return an appropriate bridge function based on the chains and tokens involved.

* **Canonical** – Assets are wrapped and then bridged.
* **[Liquidity Pools](#)** – Assets are swapped via Synapse liquidity pools.
* **[CCTP](#)** – Native router for USDC

## Pool Liquidity

![liquidity pool tokens](lp-tokens.svg)\
Synapse liquidity pools use the nexus USD (nUSD) and nexus ETH (nETH) interchain stablecoins. nUSD and nETH are fully backed by the nexus USD and nexus ETH liquidity pools on Ethereum.

When a token is bridged using a Synapse Liquidity Pool, it is converted to a nexus token on the source chain, which is then bridged to the destination chain, before being converted back into a native token.

## How to Bridge

<figure>
    <SVGBridge />
    <figcaption>Synapse Bridge</figcaption>
</figure>

### Instructions

1. Connect your wallet
2. Select your origin and destination chains from the dropdown menus
3. Select your origin token from the portfolio view, or dropdown menu
4. Enter the amount you wish to send
5. If you wish to send assets to a different wallet address, enable `Show withdrawal address` from the Settings menu (optional).
5. Connect your wallet to the origin chain, if necessary
6. Click `Bridge` to send a quote to your wallet for confirmation
7. Sign and Confirm the Bridge action from your wallet

## Bridge Contracts

Synapse Bridge contracts are available [here](https://docs.synapseprotocol.com/synapse-bridge/contract-addresses).
