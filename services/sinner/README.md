# Sinner
[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/services/sinner.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/services/sinner)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/services/sinner)](https://goreportcard.com/report/github.com/synapsecns/sanguine/services/sinner)

Sinner tracks the message lifecycle for the synapse interchain network.

## Scope

While this applications goal is to encompass the full range of events sent through SIN, the initial scope is much smaller & pragmatically focused on developers building on SIN & the developent of the off-chain protocols itself. To that end, at a high level the goal of the application is to track the message through the following stages. Completed types are contained in the table below:

## Config:

Example Indexer Config:

```yaml
default_refresh_rate: 1
scribe_url: "http://scribe.com/graphql"
db_path: "/tmp/a.db"
db_type: sqlite
skip_migrations: false
chains:
  - chain_id: 444
    contracts:
      - address: "0x537ab51470984D6D9aDF8953C0D2ed8eDA4050ED"
        start_block: 1
        contract_type: origin
      - address: "0xA944636Ac279e0346AF96Ef7e236025C6cBFE609"
        start_block: 1
        contract_type: execution_hub
  - chain_id: 421614
    contracts:
      - address: "0x537ab51470984D6D9aDF8953C0D2ed8eDA4050ED"
        start_block: 1
        contract_type: origin
      - address: "0xA944636Ac279e0346AF96Ef7e236025C6cBFE609"
        start_block: 1
        contract_type: execution_hub
  - chain_id: 11155111
    contracts:
      - address: "0x537ab51470984D6D9aDF8953C0D2ed8eDA4050ED"
        start_block: 1
        contract_type: origin
      - address: "0xA944636Ac279e0346AF96Ef7e236025C6cBFE609"
        start_block: 1
        contract_type: execution_hub
```

Example Server Config

  ```yaml
http_port: 8080
db_path: "/tmp/a.db"
db_type: sqlite
skip_migrations: true
  ```

If running the `unified` command, the config will be a combination of the above two configs (simply add the `http_port` to the top of your indexer config).
### **Message Stages**

1. [x] Message sent on origin chain.
2. [ ] At least one Guard submitted the snapshot for the origin chain taken after the message was sent.
3. [ ] At least one Notary submitted the snapshot for the origin chain taken after the message was sent.
4. [ ] Attestation created from any of the (3) snapshots was submitted to the destination chain.
5. [ ] Optimistic period for the message has passed (awaiting the execution).
6. [ ] An Executor tried to execute the message, but the destination app reverted (awaiting the retry).
7. [x] An Executor successfully executed the message.


### **Tracked Events & States**

_Note: the above only reflects progress in tracking message states in Sinner, not the agents_

1. [x] `Origin` on origin chain emits an event for every sent message.
  * `event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message);`
2. [ ] `Summit` on Synapse Chain emits an event for every state in the Guard snapshot that was submitted.
  * `event StateSaved(bytes state);`
  * We are interested in the states that have:
    * Origin domain matching the origin domain of the message.
    * Nonce higher or equal as the nonce of the message.
3. [ ] `Inbox` on Synapse Chain emits an event for every snapshot that was submitted (both Guard and Notary).
  * `event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapPayload, bytes snapSignature)`
  * We are only interested in Notary snapshots, so `domain != 0`.
    * Notary domain could be different from the origin domain of the message, that's fine.
  * [ ] `snapPayload` is a concatenation of encoded states. We are interested in the states that have:
    * Origin domain matching the origin domain of the message.
    * Nonce higher or equal as the nonce of the message.
  * [ ] In the same transaction <code>Summit</code> emits an event with the created attestation: - <code>event AttestationSaved(bytes attestation);</code>

   In other words, the submitted Guard/Notary snapshot "enables" messages coming from the set of origin domains (matching the domains of the states) with nonce less or equal to the nonce of the submitted state for the origin domain.

4. `LightInbox` on the destination chain emits an event for every submitted attestation.
  * [ ] `event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature);`
  * [ ] We are only interested in the attestation with payload matching any of the attestations from (3).
  * [ ] Optimistic period for the message starts when the first suitable attestation is submitted.
5. This could be tracked by checking the timestamp of the last mined destination block, no events needed.
6. `Destination` emits the event whenever an Executor tries to execute the message.
  * [ ] `event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool success);`
  * [ ] We are only interested in the events that have:
    * `remoteDomain` matching the origin domain of the message.
    * `messageHash` matching the hash of the message (emitted in (1)).
    * `success` equal to `false`.
  * [ ] These "failed" messages could be later retried by any Executor.
7. [x] `Destination` emits the event whenever an Executor successfully executes the message.
  * `event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool success);`
  * We are only interested in the events that have:
    * `remoteDomain` matching the origin domain of the message.
    * `messageHash` matching the hash of the message (emitted in (1)).
    * `success` equal to `true`.
  * [ ] The executed message could not be retried anymore.
