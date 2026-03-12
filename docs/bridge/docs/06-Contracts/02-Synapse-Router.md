---
title: Synapse Router
---

:::note This list may be incomplete

The canonical list is hosted within the SynapseCNS on [Github](https://github.com/synapsecns/synapse-contracts).

:::

# Synapse Router

Synapse Router contracts route through the [SynapseBridge](https://github.com/synapsecns/synapse-contracts/blob/ed93453430635e2d43704d5599d3318c43a23033/contracts/bridge/SynapseBridge.sol#L63-L118) contract, which is used for event indexing.

:::tip Events

Contracts in the `deployments` folder of each chain's `SynapseBridge.json` file emit `TokenMint` or `TokenWithdraw` events when a transaction completes on the destination chain.

:::

**Router address**: `0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a`

| Chain     | Address                                      |
|-----------|----------------------------------------------|
| Arbitrum  | `0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9` [↗](https://arbiscan.io/address/0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9) |
| Aurora    | `0xaeD5b25BE1c3163c907a471082640450F928DDFE` [↗](https://explorer.mainnet.aurora.dev/address/0xaeD5b25BE1c3163c907a471082640450F928DDFE/transactions) |
| Avalanche | `0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE` [↗](https://snowtrace.io/address/0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE) |
| Base      | `0xf07d1C752fAb503E47FEF309bf14fbDD3E867089` [↗](https://basescan.org/address/0xf07d1C752fAb503E47FEF309bf14fbDD3E867089) |
| Blast     | `0x55769baf6ec39b3bf4aae948eb890ea33307ef3c` [↗](https://blastscan.io/address/0x55769baf6ec39b3bf4aae948eb890ea33307ef3c) |
| Boba      | `0x432036208d2717394d2614d6697c46DF3Ed69540` [↗](https://blockexplorer.boba.network/address/0x432036208d2717394d2614d6697c46DF3Ed69540/transactions) |
| BSC       | `0xd123f70AE324d34A9E76b67a27bf77593bA8749f` [↗](https://bscscan.com/address/0xd123f70AE324d34A9E76b67a27bf77593bA8749f) |
| Canto     | `0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0` [↗](https://evm.explorer.canto.io/address/0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0) |
| Cronos    | `0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9` [↗](https://cronoscan.com/address/0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9) |
| DFK       | `0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A` [↗](https://subnets.avax.network/defi-kingdoms/dfk-chain/explorer/address/0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A) |
| Dogechain | `0x9508BF380c1e6f751D97604732eF1Bae6673f299` [↗](https://explorer.dogechain.dog/address/0x9508BF380c1e6f751D97604732eF1Bae6673f299) |
| Ethereum  | `0x2796317b0fF8538F253012862c06787Adfb8cEb6` [↗](https://etherscan.io/address/0x2796317b0fF8538F253012862c06787Adfb8cEb6) |
| Fantom    | `0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b` [↗](https://ftmscan.com/address/0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b) |
| Harmony   | `0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b` [↗](https://explorer.harmony.one/address/0xaf41a65f786339e7911f4acdad6bd49426f2dc6b) |
| Klaytn    | `0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b` [↗](https://scope.klaytn.com/account/0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b?tabId=txList) |
| Metis     | `0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c` [↗](https://andromeda-explorer.metis.io/address/0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c) |
| Moonbeam  | `0x84A420459cd31C3c34583F67E0f0fB191067D32f` [↗](https://moonscan.io/address/0x84A420459cd31C3c34583F67E0f0fB191067D32f) |
| Moonriver | `0xaeD5b25BE1c3163c907a471082640450F928DDFE` [↗](https://moonriver.moonscan.io/address/0xaeD5b25BE1c3163c907a471082640450F928DDFE) |
| Optimism  | `0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b` [↗](https://optimistic.etherscan.io/address/0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b) |
| Polygon   | `0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280` [↗](https://polygonscan.com/address/0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280) |
