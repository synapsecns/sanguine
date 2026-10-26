---
title: Blast Wind-Down and Recovery
---

# Blast wind-down and recovery

Blast (chain ID 81457) is retired from supported Synapse frontend bridge and swap selections and pool listings. Do not initiate a new Blast transfer using an old link. Frontend removal does not itself pause contracts, settle pending transfers, withdraw liquidity, burn tokens, or refund holders.

## Existing transactions and positions

Keep your wallet address, transaction hashes, token balances, LP balances, and staking records. Historical chain and token metadata and pool identifiers remain available for lookup. Contact [Synapse Support](https://discord.com/invite/synapseprotocol) about pending transfers or recovery of an existing position. The retained pool and staking identifiers are not a guarantee that a retired network can be managed through the frontend wallet connection.

The following addresses are retained for recovery and reconciliation:

| Item | Blast address |
| --- | --- |
| nUSD/USDB pool (`blast2pool`) | `0xa4bd1AAD7cF04567c10f38FC4355E91bba32aC9c` |
| nUSD/USDB LP token | `0xb758FF41E0A811b9eB0Ff973D373c354588524b6` |
| nETH/WETH pool (`blastethpool`) | `0x999fcd13C54B26E02a6Ccd185f71550b3a4641c0` |
| nETH/WETH LP token | `0x1600084daCcd3b58ACeFcdb0F3D7c9649359116b` |
| MiniChef (pool IDs 0 and 1, respectively) | `0x3100dC8464A8523306c3C5034de24a8927d6E590` |
| nUSD | `0x3194B0A295D87fDAA54DF852c248F7a6BAF6c6e0` |
| nETH | `0xce971282fAAc9faBcF121944956da7142cccC855` |

See also [liquidity pools](../06-Contracts/04-Liquidity-Pools.md) and [intermediate tokens](../06-Contracts/03-Intermediate-Tokens.md).

## Requirements before any pool retirement or refund

Any on-chain wind-down requires a separate reviewed plan. Preserve a verified withdrawal or claim path for LP holders, including staked LP tokens; collecting admin fees does not settle LP claims. Confirm deployed pool controls before changing them: a pool pause can also disable single-token withdrawals, and a messaging endpoint pause is a separate control that can leave transfers pending.

Before considering refunds on another chain, reconcile an entitlement snapshot against wallet LP balances, staked positions, token liabilities, and outstanding cross-chain messages. Specify the snapshot block and how later withdrawals, transfers, and settlements affect entitlements. Verify and reserve sufficient assets for refunds and nUSD/nETH backing, and reconcile each payment against the original claim to prevent duplicate recovery across chains. These are planning requirements; this notice does not announce an approved refund program, snapshot, funding, or completed payments.
