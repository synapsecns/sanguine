# Testing suite for the messaging contracts

TODO: This is an outdated README for a previous version of testing suite. Update this.

The suite contains a series of reusable tools. They are designed to be as composable as the tested contracts.

The testing suite is powered by <a href="https://github.com/foundry-rs/foundry">Foundry</a>.

## Directory structure

<pre>
test
├──<a href="./harnesses">harnesses</a> A collection of contracts, which are exposing the messaging contracts variables and functions for testing both in this suite and in the Go tests.
├──<a href="./suite">suite</a> The Foundry test files for the messaging contracts.
├──<a href="./tools">tools</a> Testing tools for the messaging contracts, that are used in <a href="./suite">suite</a>.
├──<a href="./utils">utils</a> Base test utilities to be used in <a href="./suite">suite</a>.
│   ├──<a href="./utils/proof">proof</a> Merkle proof generation.
</pre>

The directory structure for the messaging <a href="../contracts">contracts</a> is mirrored in <a href="./harnesses">harnesses</a>, <a href="./suite">suite</a> and <a href="./tools">tools</a>.

## Harnesses

For the majority of production contracts, there exists a corresponding harness. It exposes internal constants, functions and variables for testing. It also emits "logging" events to be used for testing. For instance:

```solidity
// From BasicClientHarness.t.sol
function _handleUnsafe(
  uint32 origin,
  uint32 nonce,
  uint256 rootSubmittedAt,
  bytes memory message
) internal override {
  emit LogBasicClientMessage(origin, nonce, rootSubmittedAt, message);
}

```

## Suite

Suite features a collection of testing contracts. For every production contract there is a corresponding testing one.

The underlying testing logic is usually implemented in the corresponding [Tools](#tools) contract. The tests itself are implemented in the testing contract.

```solidity
// From BasicClient.t.sol
contract BasicClientTest is BasicClientTools {
  function setUp() public override {
    super.setUp();
    setupBasicClients();
  }

  function test_setup_constructor() public {
    // code for testing state after setUp and constructor
  }

  function test_sendMessage_noTips() public {
    // code for testing "send a message with no tips"
    // should succeed
  }

  function test_sendMessage_withTips() public {
    // code for testing "send a message with tips"
    // should succeed
  }

  function test_sendMessage_revert_noRecipient() public {
    // code for testing "send a message without recipient"
    // should revert
  }
}

```

### Test function naming convention

A mix of snake_case and camelCase is used.

- `test_someFunction_whenCondition()` is used for tests, when no reverts are supposed to happen:
  - `someFunction` refers to a function that is being tested
  - `whenCondition` (optional) refers to special condition for a test. A minimal yet explicit name should be used. E.g. `noTips` instead of `doesNotUseAnyTipsWhatsoever`.
- `test_someFunction_revert_whenCondition` is used for tests, when reverts are supposed to happen:
  - `someFunction` refers to a function that is being tested
  - `whenCondition` refers to a condition when the function is supposed to revert. By default, functions are not supposed to revert, so the revert condition should always be mentioned in the function name. See above for picking a minimalistic condition name.

### Test function workflow

Usual workflow is:

- Set up conditions for the test.
- Trigger the function that is being tested.
- Check any of the following things:
  - Events emitted
  - State after the function call
  - Revert that should have happened

Some of the tests do that a couple of times. Example being executing messages on `Destination`:

- Set up conditions:
  - Dispatch a few messages from the origin (remote) chain to destination (local) chain.
  - Prepare merkle proofs for all these messages.
  - Create and sign attestation for the remote chain.
  - Submit it to `Destination` contract on local chain.
  - Wait until optimistic period is over.
- Trigger `Destination.execute()` on local chain, providing a valid merkle proof.
- Check things:
  - Destination's Event was emitted, as well as the "logging" event was emitted by the message recipient.
  - Or, check that call was reverted, if _(insert condition)_

### Example of test function: `submitAttestation()`

```solidity
// From Destination.t.sol
function test_submitAttestation() public {
  // Create messages sent from remote domain and prepare attestation
  createMessages({
    context: userRemoteToLocal,
    recipient: address(suiteApp(DOMAIN_LOCAL))
  });
  createSuggestedAttestation(DOMAIN_REMOTE);
  expectAttestationAccepted();
  // Should emit corresponding event and mark root submission time
  destinationSubmitAttestation({ domain: DOMAIN_LOCAL, returnValue: true });
  assertEq(
    destinationSubmittedAt(DOMAIN_LOCAL),
    block.timestamp,
    '!rootSubmittedAt'
  );
}

```

```solidity
// From DestinationTools.t.sol
// Creates test messages and prepares their merkle proofs for future execution
function createMessages(MessageContext memory context, address recipient)
  public
{
  bytes32 recipientBytes32 = addressToBytes32(recipient);
  rawMessages = new bytes[](MESSAGES);
  messageHashes = new bytes32[](MESSAGES);
  for (uint32 index = 0; index < MESSAGES; ++index) {
    // Construct a dispatched message
    createDispatchedMessage({
      context: context,
      mockTips: true,
      body: MOCK_BODY,
      recipient: recipientBytes32,
      optimisticSeconds: APP_OPTIMISTIC_SECONDS
    });
    // Save raw message and its hash for later use
    rawMessages[index] = messageRaw;
    messageHashes[index] = keccak256(messageRaw);
    // Dispatch message on remote Origin
    originDispatch();
  }
  // Create merkle proofs for dispatched messages
  proofGen.createTree(messageHashes);
}

```

Following steps are taken:

1. A collection of messages sharing the same context (user sends a message from `DOMAIN_REMOTE` to `DOMAIN_LOCAL`) is created.

- Each dispatched messages is constructed first.
- Its payload and hash is saved.
- It is then dispatched from corresponding `Origin` (on `DOMAIN_REMOTE` in that example)
- Finally, a collection of merkle proofs is generated.

2. A suggested attestation (i.e. referencing the latest state) is created for `DOMAIN_REMOTE`. Attestation's root could be later used for executing the dispatched messages.
3. We expect `AttestationAccepted` event to be emitted.
4. We submit attestation to `Destination` and expect return value of `true`.
5. We check the submission time for the attestation root.

### Example of test function: `execute()`

```solidity
// From Destination.t.sol
function test_execute() public {
  AppHarness app = suiteApp(DOMAIN_LOCAL);
  test_submitAttestation();
  skip(APP_OPTIMISTIC_SECONDS);
  // Should be able to execute all messages once optimistic period is over
  for (uint32 i = 0; i < MESSAGES; ++i) {
    checkMessageExecution({ context: userRemoteToLocal, app: app, index: i });
  }
}

```

```solidity
// From DestinationTools.t.sol
// Prepare app to receive a message from Destination
function prepareApp(
  MessageContext memory context,
  AppHarness app,
  uint32 nonce
) public {
  // App will revert if any of values passed over by Destination will differ (see AppHarness)
  app.prepare({
    origin: context.origin,
    nonce: nonce,
    sender: addressToBytes32(context.sender),
    message: _createMockBody(context.origin, context.destination, nonce)
  });
}

// Check given message execution
function checkMessageExecution(
  MessageContext memory context,
  AppHarness app,
  uint32 index
) public {
  uint32 nonce = index + 1;
  // Save mock data in app to check against data passed by Destination
  prepareApp(context, app, nonce);
  // Recreate tips used for that message
  createMockTips(nonce);
  expectLogTips();
  expectExecuted({ domain: context.origin, index: index });
  // Trigger Destination.execute() on destination chain
  destinationExecute({ domain: context.destination, index: index });
  // Check executed message status
  assertEq(
    destinationMessageStatus(context, index),
    attestationRoot,
    '!messageStatus'
  );
}

```

Following steps are taken:

1. We reuse `test_submitAttestation()` to get us to the state where messages are dispatched, and attestation is submitted to `Destination`
2. We execute messages one by one, and check that everything went fine:

- That recipient is passed the correct `(origin, nonce, sender, message)` data
- That `Destination` has the same `tips` payload which was used for sending a message.
- That `Executed` event is emitted.
- That message is marked as executed on `Destination` (preventing another execution)

Everything is reusable!

## Tools

For every tested contract, a corresponding `<...>Tools` contract is used for basic testing logic. In this contract, testing data is created and saved for later verification. The `Tools` contract reuse functions from one another to make the testing easier.

```solidity
abstract contract DestinationTools is OriginTools {
  // Here we define constants and state variables used for testing
  bytes[] internal rawMessages;

  // ...

  // Here we define functions to create test data

  // Creates test messages and prepares their merkle proofs for future execution
  function createMessages(MessageContext memory context, address recipient)
    public
  {
    // ...
  }

  /*╔══════════════════════════════════════════════════════════════════════╗*\
  ▏*║                            EXPECT EVENTS                             ║*▕
  \*╚══════════════════════════════════════════════════════════════════════╝*/

  // Here we define wrappers for expecting a given event

  function expectAttestationAccepted() public {
    vm.expectEmit(true, true, true, true);
    emit AttestationAccepted(
      attestationDomain,
      attestationNonce,
      attestationRoot,
      signatureNotary
    );
  }

  // ...

  /*╔══════════════════════════════════════════════════════════════════════╗*\
  ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
  \*╚══════════════════════════════════════════════════════════════════════╝*/

  // Here we define wrappers for triggering a Destination function and expecting a revert

  // Trigger destination.submitAttestation() with saved data and expect a revert
  function destinationSubmitAttestation(
    uint32 domain,
    bytes memory revertMessage
  ) public {
    DestinationHarness destination = suiteDestination(domain);
    vm.expectRevert(revertMessage);
    vm.prank(broadcaster);
    destination.submitAttestation(attestationRaw);
  }

  // ...

  /*╔══════════════════════════════════════════════════════════════════════╗*\
  ▏*║                          TRIGGER FUNCTIONS                           ║*▕
  \*╚══════════════════════════════════════════════════════════════════════╝*/

  // Here we define wrappers for triggering a Destination function and expecting a given return value

  // Trigger destination.submitAttestation() with saved data and check the return value
  function destinationSubmitAttestation(uint32 domain, bool returnValue)
    public
  {
    DestinationHarness destination = suiteDestination(domain);
    vm.prank(broadcaster);
    assertEq(
      destination.submitAttestation(attestationRaw),
      returnValue,
      '!returnValue'
    );
    if (returnValue) {
      rootSubmittedAt = block.timestamp;
    }
  }
}

```

### Test data

Test data needs to be unique for every test, while also being easily reconstructible for checking. Data is usually constructed from the given parameters, the remaining ones are mocked. Functions for data construction are designed to be composable:

```solidity
abstract contract OriginTools is MessageTools {
  // Create a dispatched message: given {context, body, recipient, optimistic period}
  // pass MOCK_X constant to mock field X instead
  function createDispatchedMessage(
    MessageContext memory context,
    bool mockTips,
    bytes memory body,
    bytes32 recipient,
    uint32 optimisticSeconds
  ) public {
    createMessage({
      origin: context.origin,
      sender: _getSender(context, recipient),
      nonce: _nextOriginNonce(context.origin),
      destination: context.destination,
      mockTips: mockTips,
      body: body,
      recipient: recipient,
      optimisticSeconds: optimisticSeconds
    });
  }
}
```

Here `OriginTools` implements a function to construct a payload for a dispatched message. It reuses the generic `createMessage()` from `MessageTools` instead of manually encoding the payload from scratch.

## Utils

A collection of the base contracts used for Foundry tests.

### `SynapseTestSuite`

Inherits from [SynapseUtilities](#synapseutilities), [SynapseTestStorage](#synapseteststorage).

The base contract to be used in tests. `SynapseTestSuite.setUp()` deploys all messaging contracts for three chains:

- `DOMAIN_SYNAPSE`: to be used as the "master chain" for messaging contracts. Attestations about all chains `Origin` state are posted here. Also here happens bonding and unbonding of the off-chain actors.
- `DOMAIN_LOCAL`: to be used as the "main testing chain". Tests where messages are sent, will usually feature messages **from "local chain" to "remote chain"**.
- `DOMAIN_REMOTE`: to be used as the "auxiliary testing chain". Tests where messages are received, will usually feature messages **from "remote chain" to "local chain"**.

In the same `setUp` functions, a collection of off-chain actors are created. For each actor the private key is saved for later signing.

- Notary. A predetermined amount of notaries is created for each chain.
- Guard. A predetermined amount of guards is created.
- Owner. A single account is created. For every `Ownable` contract, the deployer contract transfers the ownership to `owner`. This emulates real world behavior, and helps with testing: calls without `vm.prank(address)` are made from the testing contract, which is also the deployer, so it makes sense to give away the ownership to prevent false negatives.
- Proxy admin. Some of the contracts are supposed to be deployed as upgradeable proxies. Proxy admin address is helpful for:
  - Fuzzing, to remove it from the list of callers. Calls from proxy admin are not forwarded, causing the test to fail, when it's not supposed to.
  - Testing upgradeability (to be implemented when needed)
- Attacker. A single account is created for various tests of unauthorized access (including signatures of unauthorized agents).
- User. A single account is created to serve as `msg.sender` for legit tests like "sending a message".
- Broadcaster. A single account is created to serve as `msg.sender` for tests, where signature of a off-chain actor is submitted on chain.

### `SynapseUtilities`

Inherits from `Test`, a default Foundry testing contract.

Features some useful utilities that don't require access to state variables, like typecasts, key generation, string formatting.

### `SynapseTestStorage`

Inherits from [SynapseConstants](#synapseconstants), [SynapseEvents](#synapseevents).

Features all storage variables used for testing (like saved deployments, actors, etc), as well as a collection of handy getters for them.
Also features a tool to generate Merkle proofs, and the preset context variables for the messaging tests.

### `SynapseConstants`

Features a collection of constants used for the messaging tests.

### `SynapseEvents`

Features all events from the production contracts, as well all the "logging" events from all the harnesses. This way all events are accessable in the testing contracts without the need to redefine them.
