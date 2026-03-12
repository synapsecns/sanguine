# Transaction Support FAQ

## What does a Bridge transaction look like?

After submitting and signing a transaction from your wallet, gas fees are collected, and the transaction is sent to the origin chain router. Once accepted, the bridged asset is removed from your portfolio, and a progress bar shows the estimated confirmation time.

Once confirmed on the destination chain, the asset is added to your portfolio, and destination hash is available from the progress menu dropdown. The transaction appears as part of your history in the Activity tab once it is index by the Synapse Explorer.

Gas token airdrops and rebates are delivered to your wallet automatically. However, only bridgeable assets are shown in your Synapse portfolio.

## Did my transaction initiate?

Transactions that do not initiate on the origin chain return an error message. Your assets will remain in your portfolio, under your control.

In the event that your transaction does not initiate, double check that you have sufficient funds to send, and to cover gas fees, and you can safely try again.

## My transaction failed to initiate after several tries

Occasionally, technical issues or a high volume of activity on the origin chain may prevent new transactions from being accepted.

In most cases, these issues safely resolve within 30-60 minutes. Activity levels can be found on native block explorers (e.g the [Etherscan gas tracker](https://etherscan.io/gastracker)).

You can also adjust your wallet’s gas settings to make transactions more likely to be accepted during times of peak activity.

## Why is my transaction taking so long?
Synapse time estimates are based on destination block times. Occasionally, a transaction may post to a later block than expected.

Block explorer links in the progress dropdown menu can confirm whether a confirmation on-chain but not yet received by Synapse.

## My transaction failed to complete

Transactions that fail to complete are not lost, and are manually addressed by the Synapse support team. You can reach Support via the [Synapse Discord channel](https://discord.com/invite/synapseprotocol) at any time.

:::note For DeFi Kingdoms

NFT transactions can take twice as long as tokens. Contact Support if your transaction has been pending for two hours or more.

:::

## I received a different asset than expected
In the event of an sudden increase in slippage, Synapse will deliver the intermediate asset sent to the destination chain instead of swapping it for an unexpected loss.

This asset appears in your portfolio and can be safely [swapped](https://synapseprotocol.com/swap) for the asset of your choice on the destination chain.

## Did I receive my rebate or gas airdrop?
While rebates and airdrops appear in your wallet automatically, only bridgeable assets are shown in your Synapse portfolio.

If you don’t see an asset you should have received, first check your wallet while connected to the destination chain for your bridge transaction.

## Help!
Don’t panic! Contact Synapse Support on Discord to answer any other questions you might have.
