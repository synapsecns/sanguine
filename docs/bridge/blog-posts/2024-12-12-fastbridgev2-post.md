---
slug: synapse-intent-network-launch
title: The Synapse Intent Network
# authors: [synapse]
tags: [update, fastbridgev2, intent-network]
---

import { RFQFlow } from '@site/src/components/RFQFlow'

We are excited to announce the **Synapse Intent Network**.

<!--truncate-->

# Summary

The Synapse Intent Network is a cross-chain communication protocol that enables seamless asset transfers and message passing between different blockchain networks. At its core, it provides a robust infrastructure for executing complex cross-chain operations while maintaining security and efficiency.

<figure>
	<RFQFlow />
	<figcaption>RFQ flow: get Quote, get txData, sign transaction</figcaption>
</figure>

<br/>

This major protocol upgrade represents a fundamental shift in the architecture, introducing intent-based routing, advanced bridging capabilities, and significant gas optimizations. The changes reflect a deeper focus on user experience while substantially improving the protocol's security, efficiency, and scalability.
## Key Improvements At-A-Glance

* **Gas Optimization Revolution**
  * Achieved 30-50% reduction in transaction costs
  * Single-slot storage optimizations save up to 20,000 gas per operation

* **Powerful Intent-Based Routing**
  * New architecture enables complex multi-step transactions to be executed atomically
  * Reduces cross-contract calls
  * Allows for sophisticated bridging scenarios

* **Advanced Batching Capabilities**
  * Multicall support with selective result processing
  * Reduces base transaction costs by 21,000 gas per operation
  * Enables atomic execution of complex operations

## Synapse Intent Network: Technical Improvements

### Quoting and API Improvements

The transition to the Synapse Intent Network brings significant changes to our quoting infrastructure, focusing on real-time price discovery and efficient market making. The most notable addition is active quoting alongside our existing passive quoting system.

#### Active Quoting

Traditional passive quoting works like an order book - relayers post standing offers that users can take. While this model is simple and efficient for stable market conditions, it can lag during volatile periods. Active quoting addresses this limitation by enabling relayers to respond to quote requests in real-time:

```typescript
// Example WebSocket quote request format
interface QuoteRequest {
  data: {
    origin_chain_id: number;
    dest_chain_id: number;
    origin_token_addr: string;
    dest_token_addr: string;
    origin_amount_exact: string;
    expiration_window: number;
  }
}
```

This hybrid approach gives relayers flexibility in their market-making strategies. Simple integrations can stick with passive quoting, while sophisticated relayers can implement dynamic pricing models that account for immediate market conditions, liquidity depth, and cross-chain gas costs.

### WebSocket API Evolution

Supporting this new quoting model required rethinking our API infrastructure. The new WebSocket layer eliminates the need for polling and complex state management:

```typescript
const operations = {
  subscribe: "Subscribe to specific chains",
  send_quote: "Respond to quote request",
  request_quote: "New quote request notification"
}
```

The real-time nature of WebSockets dramatically reduces quote latency. Rather than repeatedly querying for updates, relayers receive instant notifications about new opportunities. This improved efficiency translates directly to better pricing for end users as relayers can operate with tighter spreads.

## Contract Improvements

:::info

The Synapse Intent Network is backwards combatible with the original Fastbridge Contracts.

:::

### Gas Optimizations

A core focus of V2 was reducing transaction costs without sacrificing code clarity. Through careful struct packing and custom error implementations, we've achieved significant gas savings:

```solidity
struct BridgeTxDetails {
    BridgeStatus status;           // 1 byte
    uint32 destChainId;           // 4 bytes
    uint16 proverID;              // 2 bytes
    uint40 proofBlockTimestamp;   // 5 bytes
    address proofRelayer;         // 20 bytes
} // Total: 32 bytes (1 slot)
```

These optimizations balance efficiency with maintainability. While more aggressive packing was possible, it would have made the code significantly harder to reason about. The current implementation provides meaningful gas savings while keeping the codebase approachable for new developers.

### Atomic Operations with Multicall and Zaps

Cross-chain operations often require multiple transactions. V2 introduces multicall support and a new Zap interface to address this friction:

```solidity
// Multicall enables efficient batching
fastBridge.multicallNoResults([
    abi.encodeCall(IFastBridge.prove, (request1)),
    abi.encodeCall(IFastBridge.claim, (request1))
], false);

// Zaps enable complex atomic operations
interface IZapRecipient {
    function zap(
        address token,
        uint256 amount,
        bytes memory zapData
    ) external payable returns (bytes4);
}
```

The Zap interface is intentionally minimal, maximizing composability while maintaining security. It enables powerful workflows like "bridge and stake" or "bridge and provide liquidity" in a single atomic transaction. These compound operations significantly improve the user experience by reducing transaction overhead and eliminating partial execution risks.

## Final Thoughts

The Synapse Intent Network (FastBridgeV2) update represents a thoughtful evolution of our protocol. Each change was evaluated not just for its immediate benefits, but for its long-term impact on protocol composability and user experience. The result is a more efficient, developer-friendly, and user-centric cross-chain bridging system. Please reach out to us if you have any questions or feedback.
