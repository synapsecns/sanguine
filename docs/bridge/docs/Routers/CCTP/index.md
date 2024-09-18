# Overview

Synapse CCTP is a custom module built on top of Circle's [Cross-Chain Transfer Protocol](https://www.circle.com/en/cross-chain-transfer-protocol) that allows for bridge requests to natively mint & burn USDC on [supported chains](https://developers.circle.com/stablecoins/docs/cctp-getting-started#supported-blockchains).

Synapse's CCTP implementation consists of two main components:

- [CCTP Relayer](./Relayer.md): An off-chain service that fulfills transactions requested through the CCTP contracts. The relayer is responsible for fetching attestations from the [Circle API](https://developers.circle.com/stablecoin/reference) and submitting them to the CCTP contracts. Anyone can run a relayer.
- [CCTP Contracts](./Contracts.md): A set of smart contracts that allow for the minting and burning of USDC on supported chains, and instant swaps to/from any supported asset. These contracts are deployed on each supported chain and are responsible for minting and burning USDC.

As a modular component of Synapse's router system, CCTP can be configured to bridge through any supported liquidity source, such as [Curve](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/curve/CurveV1Module.sol), [Algebra](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/algebra/AlgebraModule.sol), [DAI PSM](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/dss/DssPsmModule.sol),  and others.
