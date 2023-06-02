---
sidebar_position: 4
---

# Running an Executor

Because the Executor is not in a position to commit fraud or perform a Denial of Service attack, there is no need for an Executor to post any stake. The Executor just needs to keep track of all the messages in the probation state and as soon as the optimistic period has lapsed, it can move them from the probationary state to the accepted state on the Destination chains.

The Executor does not have a bonded Signing key, but it does need an address with Gas on each of the chains it supports.
Running an Executor requires more in terms of the deployment, because it uses a database as well as other microservices developed by the Synapse team. The reference implementation uses these other Microservices, so these need to also be deployed in the Environment.

The reference implementation instructions for the Executor can be found here:

https://github.com/synapsecns/sanguine/blob/master/agents/agents/executor/cmd/cmd.md

The executor configuration file could look something like this:

```yaml
    execute_interval: 5
    summit_chain_id: 10
    summit_address: 0xSummit
    chains:
      - chain_id: 123
        temp_rpc: "https://chain123.rpc
        origin_address: 0xabc
        destination_address: 0xdef
      - <another chain config>
    unbonded_signer:
      type: "File"
      file: "/config/executor_signer.txt"
```
The Executor requires running an instance of scribe. The instructions for that is here:

https://github.com/synapsecns/sanguine/blob/master/services/scribe/cmd/cmd.md

And this is an example scrib config file:

```yaml
    embedded_scribe_config:
      chains:
        - chain_id: 137
          required_confirmations: 0
          contract_sub_chunk_size: 1000
          contract_chunk_size: 1000
          store_concurrency: 1
          store_concurrency_threshold: 500
          contracts:
            - address: 0xF3773BE7cb59235Ced272cF324aaeb0A4115280f
              start_block: 40189736
            - address: 0xde5BB62aBCF588EC200674757EDB2f6889aCd065
              start_block: 40189736
        - chain_id: 10
          required_confirmations: 0
          contract_sub_chunk_size: 500
          contract_chunk_size: 500
          store_concurrency: 1
          store_concurrency_threshold: 500
          contracts:
            - address: 0xF3773BE7cb59235Ced272cF324aaeb0A4115280f
              start_block: 79864523
            - address: 0xde5BB62aBCF588EC200674757EDB2f6889aCd065
              start_block: 79864305
            - address: 0x128fF47f1a614c61beC9935898C33B91486aA04e
              start_block: 79864192
        - chain_id: 43114
          required_confirmations: 0
          contract_sub_chunk_size: 2000
          contract_chunk_size: 10000
          store_concurrency: 1
          store_concurrency_threshold: 500
          contracts:
            - address: 0xF3773BE7cb59235Ced272cF324aaeb0A4115280f
              start_block: 27262747
            - address: 0xde5BB62aBCF588EC200674757EDB2f6889aCd065
              start_block: 27262744
      rpc_url: "https://rpc.interoperability.institute/confirmations"
      refresh_rate: 0
```
