---
title: About
---

import AnimatedLogo from '@site/src/components/AnimatedLogo'
import SVGBridge from '@site/src/components/SVGBridge'
import { BridgeFlow } from '@site/src/components/BridgeFlow'
import { CCTPFlow } from '@site/src/components/CCTPFlow'
import { RFQFlow } from '@site/src/components/RFQFlow'

<AnimatedLogo />

# Use Synapse

Synapse is an Interchain Programming Interface. Developers read and write interchain data with Synapse, which has settled $50B in transactions between 2M+ users, and generated $30M+ in fees [[Explorer](https://explorer.synapseprotocol.com)].

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

* **[SDK](/docs/Developers/Bridge-SDK)** – Call Synapse Router functions from your frontend or backend application.
* **[REST API](/docs/Developers/REST-API)** – Endpoints and example code
* **[Widget](/docs/Developers/Widget)** – Embed a customized Synapse Bridge in your application.

## Synapse Routers

Synapse Router automatically determines the appropriate router type to use for each Bridge transaction.

<figure id="flowGroup">
    <figure>
        <BridgeFlow />
        <figcaption>Synapse Router</figcaption>
    </figure>
    <figure>
        <CCTPFlow />
        <figcaption>Synapse CCTP</figcaption>
    </figure>
    <figure>
        <RFQFlow />
        <figcaption>Synapse RFQ</figcaption>
    </figure>
</figure>

* **[Synapse Router](/docs/Routers/Synapse-Router)** – Executable quotes for arbitrary blockchain transactions.
* **[CCTP](/docs/Routers/CCTP)** – Native router for USDC transactions.
* **[RFQ](/docs/Routers/RFQ)** – Fast router that allows on-chain agents to bid on interchain delivery.

## Essential Services

Bolt-on services for reliability and ease-of-use:

* **[Scribe](/docs/Services/Scribe)** – Index logs, receipts and transactions across multiple chains
* **[OmniRPC](/docs/Services/Omnirpc)** – Interchain RPC load balancer and verifier
* **[Signer](/docs/Services/Signer)** – Support for the AWS Key Management System (KMS)
* **[Submitter](/docs/Services/Submitter)** – Gas management service to ensure transaction confirmation
* **[Observability](/docs/Services/Observability)** – Open telemetry system for Synapse SDK


## Community & Support

Connect with other developers and the Synapse team

* **[Discord](https://discord.gg/synapseprotocol)**
* **[Twitter](https://twitter.com/SynapseProtocol)**
* **[Telegram](https://t.me/synapseprotocol)**
* **[Forum](https://forum.synapseprotocol.com/)**

## Additional Links

Synapse transactions can be observed confirmed via the following methods:

* **[Synapse Bridge](https://synapseprotocol.com)** – Bridge, Swap, and Stake via Synapse's cross-chain pools.
* **[Synapse Explorer](https://explorer.synapseprotocol.com)** – Public explorer for Synapse Bridge transactions.
