---
sidebar_position: 4
---

# Running an Executor

Because the Executor is not in a position to commit fraud or perform a Denial of Service attack, there is no need for an Executor to post any stake. The Executor just needs to keep track of all the messages in the probation state and as soon as the optimistic period has lapsed, it can move them from the probationary state to the accepted state on the Destination chains.
<br/>
The Executor does not have a bonded Signing key, but it does need an address with Gas on each of the chains it supports.
Running an Executor requires more in terms of the deployment, because it uses a database as well as other microservices developed by the Synapse Carbon. The reference implementation uses these other Microservices, so these need to also be deployed in the Environment.
