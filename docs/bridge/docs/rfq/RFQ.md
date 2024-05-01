# RFQ

RFQ is a bridge module supported by the Synapse Router that allows different market makers to post quotes on different bridge routes. Users can take these quotes by submitting an on-chain bridge request. In the event these requests are not fulfilled, users can request a refund after a set period of time.

### Actors

With the exception of the smart contract itself, RFQ is agnostic to how users receive quotes and where Solvers choose to post quotes. Below, we explain who the general actors interacting with the contract are and then explain the canonical RFQ implementation.

- **Solvers -** Solvers (also known as fillers) are market makers that provide liquidity & post quotes to the API. They are then in charge of fulfilling requests on-chain.
- **Users** - End users observe quotes from solvers and post requests on chain. In the event these requests cannot be fulfilled, the user can reclaim their funds after the optimistic window has passed.
- **Quoter -** The quoter runs a service ran by different interfaces to the Synapse Bridge that allows market makers to post quotes and users to read them. The spec of RFQ does not require this to be an “API” in the traditional sense. Interfaces can use protocols like libp2p, irc and dht’s to communicate quotes.

Right now, RFQ consists of three-different components, with each of the two off-chain components being ran by different actors:

- **API -** The RFQ api is an off-chain service ran by Quoters. user-interfaces that allows market makers/solvers to post quotes on different bridge routes. Solvers that have registered with the FastBridge contract can sign messages that post quotes signifying at what price they are willing to bridge tokens on a certain route.

  In the canonical implementation, users Solvers authenticated by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). The canonical implementation can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq).
- **Fast Bridge Contract -** The fast bridge contract is the core of the RFQ protocol and what allows solvers  to fulfill requests from users. A user deposits their funds into the FastBridge contract along with the lowest price they are willing to accept for a given route (a price they get by reading quotes from the Quoter).

  In the unlikely event no Solver is available to fulfill a users request, a user can permissionlessly  claim their funds back after waiting an optimistic period.

  Contract code level documentation can be found [here](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html).
- **Relayer** - The relayer is a service ran by the solvers. The relayer is responsible for posting quotes & fulfilling requests. While the relayer can be implemented in any way, the canonical implementation is a golang based relayer that provides a way to decide what chains/routes to quote on, how much to quote and which addresses not to relay for.

### Relayer

At a high level, the canonical implementation of the relayer has 3 different responsibilities.

- **Quoting** - Keep track of balances on each chain as well as in-flight funds and continuously post-quotes with these balances using the config to adjust quotes to the solvers specifications and posting to the API.
- **Relaying -** Fulfill users BridgeRequests [link to event] by relaying their funds on-chain. Once eligible, claim the users funds on the origin chain.
- **Rebalancing -** In order to handle the complexity of user flows, the Relayer provides an interface that allows funds to be rebalanced. This allows RFQ to be reflexive to cases where flows are mono-directional.

To facilitate this, the rfq relayer config is pretty complex. Here is a breakdown of an example config:

```yaml
submitter_config: # please see the more detailed submitter doc [here]
  chains:
    1:
      supports_eip_1559: true
      gas_estimate: 1000000
database:
  type: mysql # can be other mysql or sqlite
  dsn: root:password@hostname:3306)/database?parseTime=true # should be the dsn of your database. If using sqlite, this can be a path

signer: # please see more detailed signer config [here] (link)
  type: GCP
  file: /config/signer.txt

screener_api_url: 'http://screener-url'
rfq_url: 'http://rfq-api'
omnirpc_url: 'http://omnirpc'
rebalance_interval: 2m
relayer_api_port: '8081'

base_chain_config:
  confirmations: 0
  # Claim (72.5k) + Prove (57.5k) gas limits, rounded up
  origin_gas_estimate: 130_000
  # Relay gas limit, rounded up
  dest_gas_estimate: 110_000
  quote_offset_bps: 2
  native_token: ETH
  quote_pct: 90
  min_gas_token: 1000000000000000000
  fixed_fee_multiplier: 1.25

chains:
  1:
    rfq_address: "0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E" # rfq contract address on eth
    synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E" # ccctp contract address on eth
    token_messenger_address: "0xbd3fa81b58ba92a82136038b25adec7066af3155" # token messenger address on eth, note: only one of token_messenger_address or synapse_cctp_address actually needs to be present
    cctp_start_block: 19341000
    confirmations: 2
    tokens:
      USDC:
        address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
        decimals: 6
        price_usd: 1.0
        min_quote_amount: 10000
        rebalance_method: "circlecctp"
        maintenance_balance_pct: 20
        initial_balance_pct: 34
        max_rebalance_amount: 500000
      ETH:
        address: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
        decimals: 18
        price_usd: 2600
  10:
    rfq_address: "0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E"
    synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
    token_messenger_address: "0x2B4069517957735bE00ceE0fadAE88a26365528f"
    cctp_start_block: 116855000
    l1_fee_chain_id: 1
    # Prove + Claim L1 gas estimate
    l1_fee_origin_gas_estimate: 20
    # Relay L1 gas estimate
    l1_fee_dest_gas_estimate: 10
    tokens:
      USDC:
        address: "0x0b2c639c533813f4aa9d7837caf62653d097ff85"
        decimals: 6
        price_usd: 1.0
        min_quote_amount: 10000
        rebalance_method: "circlecctp"
        maintenance_balance_pct: 20
        initial_balance_pct: 34
        max_rebalance_amount: 500000
      ETH:
        address: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
        decimals: 18
        price_usd: 2600

quotable_tokens:
  10-0x0b2c639c533813f4aa9d7837caf62653d097ff85:
    - "1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
  1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48:
    - "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85"
  1-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE:
    - "10-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
  10-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE:
    - "1-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

fee_pricer:
  gas_price_cache_ttl: 60
  token_price_cache_ttl: 60

cctp_relayer_config:
  cctp_type: "circle"
  scribe_port: 80
  scribe_url: "scribe.scribe"
  circle_api_url: "https://iris-api.circle.com/v1/attestations"
  chains:
    - chain_id: 1
      synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      token_messenger_address: "0xbd3fa81b58ba92a82136038b25adec7066af3155"
    - chain_id: 10
      synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      token_messenger_address: "0x2B4069517957735bE00ceE0fadAE88a26365528f"
  base_omnirpc_url: "http://omnirpc.omnirpc.svc.cluster.local"
  unbonded_signer:
    type: GCP
    file: /config/signer.txt
  http_backoff_initial_interval_ms: 1000
  http_backoff_max_elapsed_time_ms: 300000

```

- **submitter_config** - This is covered [here]. At a high level this controls gas parameters used for on-chain transactions
- **database -** This is covered [here]. At a high level this is where the data is stored.
- **screener_api_url -** screener_api_url is an optional param for an api. While we recommend using the screener-api [link], this can be any api that returns `{”blocked”: bool}`
- **rfq_url -** the rfq_url is the url of the rfq api. Please see [API] page for details on running. For the Synapse interface this is [mainnet url] and [testnet_url]
- **omnirpc_url -** the omnirpc url to use, please see [here]
- **rebalance_interval:** 2m - how often to rebalance
- **relayer_api_port**: the relayer provides an api detailed [here], this is the endpoint for this api.Note this api should be secured/not public
- **base_chain_config:** Base chain config is the default config applied for each chain if the other chains do not override it. This is covered in the chains section.
- **chains**:  each chain has a different config that overrides base_chain_config. Here are the parameters for each chain
  - **rfq_address -** address of the rfq contract on chain. Current contract addresses can be found [here]
  - **synapse_cctp_address** (optional) - this is only applicable if **rebalance_method** is set to synapse.

      <aside>
      💡 The choice of wether to use synapse cctp or the circle token messenger is up to the user. Synapse will take a fee but unlike the token messenger, will not spend any of the users gas.

      </aside>

  - **token_messenger_address** (optional) - this is only applicable if **rebalance_method** is set to cctp. Tells our relayer to use the token messenger instead of synapse
  - **confirmations -** how many confirmations to wait before acting on an event. [should link to some docs on this]
  - **tokens -** this is a map of token symbol→token info for this chain. For example, token may be USDC, ETH, etc
    - **address** - address of the token on this chain id
    - **decimals**  - number of decimals this token uses. Please verify this against the token itself

        <aside>
        💡 RFQ does not yet support relays where tokens have different decimal counts

        </aside>

    - **min_quote_amount** - smallest amount to quote for a given chain. This should be balanced against expected gas spend.
    - **rebalance_method-** rebalance method for this particular kind of token. Some tokens may not have a rebalance method.
    - **maintenance_balance_pct -** percent of liquidity that should be maintained on the given chain for this token. If the balance is under this amount a rebalance is triggered
    - **initial_balance_pct -** percent of liquidity to maintain after a rebalance
    - **min_rebalance_amount -** amount of this token to try to rebalance
    - **max_rebalance_amount -** maximum amount of this token to try to rebalance at once
- **quotable_tokens:**

[//]: # (  - list of [chain-id]_[token_address]: []{ [chain-id]_[token_address]}. For example 1-0x00…. could be paired with 10-0x01)

    ```jsx
    "1-0x00":
    	- 1"-0x01"
    ```

- **cctp_relayer_config:** [should link out to this]

[API](https://www.notion.so/API-8d4fa8d6e36d4b98a6c801a1889f71ee?pvs=21)
