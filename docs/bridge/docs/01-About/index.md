---
title: About
---

import AnimatedLogo from '@site/src/components/AnimatedLogo'
import SVGBridge from '@site/src/components/SVGBridge'
import { BridgeFlow } from '@site/src/components/BridgeFlow'
import { CCTPFlow } from '@site/src/components/CCTPFlow'
import { RFQFlow } from '@site/src/components/RFQFlow'

<AnimatedLogo />

# Synapse

:::note
Synapse Protocol lives under the Cortex umbrella, Convert SYN to CX before the end of the migration period. Learn more about Cortex in the [Cortex Docs](https://docs.cortexprotocol.com/docs/StartHere)
:::

Synapse is an Interchain Programming Interface. Developers read and write interchain data with Synapse, which has settled $50B in transactions between 2M+ users, and generated $30M+ in fees. Eventually, all Synapse Protocol features will fold into Cortex.



Source: [Synapse Explorer analytics](https://explorer.synapseprotocol.com).

## Interchain Bridge

<figure>
    <SVGBridge />
    <figcaption>Synapse Bridge</figcaption>
</figure>

* [Overview](/docs/Bridge)
* [Bridge guide](/docs/Bridge#how-to-bridge)
<!-- * [Supported routes](/docs/Bridge#how-to-bridge) -->

## Developers

Embed or build a custom Bridge application.

* **[SDK](/docs/Bridge/SDK)** – Call Synapse Router functions from your frontend or backend application.
* **[REST API](/docs/Bridge/REST-API)** – Endpoints and example code
* **[Widget](/docs/Bridge/Widget)** – Embed a customized Synapse Bridge in your application.

## Synapse Routers

Synapse Router automatically determines the appropriate router for each Bridge transaction.

* **[Synapse Router](/docs/Routers/Synapse-Router)** – Returns and executes quotes for supported interchain transactions.
* **[CCTP](/docs/Routers/CCTP)** – Native router for USDC transactions.
* **[RFQ](/docs/RFQ)** – Relayers bid for the right to provide immediate delivery.

<figure id="flowGroup">
    <figure>
        <BridgeFlow />
        <figcaption>Synapse Router &ndash; Mint and burn any token between chains</figcaption>
    </figure>
    <figure>
        <CCTPFlow />
        <figcaption>CCTP &ndash; Use Circle contracts to mint and burn native USDC</figcaption>
    </figure>
    <figure>
        <RFQFlow />
        <figcaption>RFQ &ndash; Take Immediate delivery from a destination relayer, who receives your origin chain assets on confirmation.</figcaption>
    </figure>
</figure>

## Community & Support

Connect with other developers and the Synapse team

* **[Discord](https://discord.gg/4rMzuEnKqe)**
* **[Twitter](https://twitter.com/SynapseProtocol)**
* **[Telegram](https://t.me/synapseprotocol)**
* **[Forum](https://common.xyz/cortex-dao)**

## Additional Links

Synapse transactions can be observed confirmed via the following methods:

* **[Synapse Bridge](https://synapseprotocol.com)** – Bridge, Swap, and Stake via Synapse's cross-chain pools.
* **[Synapse Explorer](https://explorer.synapseprotocol.com)** – Public explorer for Synapse Bridge transactions.
