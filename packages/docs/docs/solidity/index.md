---
sidebar_position: 1
---

# Integrating Messaging

For Smart Contract developers who need to send Cross-Chain Messages, you need to call the “dispatch” method on the Synapse Carbon Origin contract that has already been deployed on the chain that you are sending from. The dispatch function has the following interface:

```solidity
/**
 * @notice Dispatch the message to the recipient located on destination domain.
 * @param _destination      	   Domain of destination chain
 * @param _recipient        	   Address of recipient on destination chain as bytes32
 * @param _optimisticSeconds Optimistic period for message execution on destination chain
 * @param _tips             	   Payload with information about paid tips
 * @param _messageBody       Raw bytes content of message
 * @return messageNonce        Nonce of the dispatched message
 * @return messageHash      	   Hash of the dispatched message
 */
  function dispatch(
    uint32 _destination,
    bytes32 _recipient,
    uint32 _optimisticSeconds,
    bytes memory _tips,
    bytes memory _messageBody
  ) external payable returns (uint32 messageNonce, bytes32 messageHash);
```

The current version does not require anything for tips. Below is an example implementation of a test smart contract that calls the dispatch:

```solidity
function sendMessage(
    address _origin
    uint32 _destination,
    address _recipient,
    uint32 _optimisticSeconds,
    bytes memory _message
  ) external {
    bytes32 recipient = TypeCasts.addressToBytes32(_recipient);
    bytes memory tips = TipsLib.emptyTips();
    (uint32 nonce, ) = InterfaceOrigin(_origin).dispatch(
    _destination,
    recipient,
    _optimisticSeconds,
    tips,
    _message
  );
  emit MessageSent(
    _destination,
    nonce,
    TypeCasts.addressToBytes32(address(this)),
    recipient,
    _message);
}
```

Then, on the receiving chain, there will be a Smart Contract that acts as the recipient of the message. This Smart Contract needs to implement the following interface:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IMessageRecipient {
  function handle(
    uint32 _origin,
    uint32 _nonce,
    bytes32 _sender,
    uint256 _rootTimestamp,
    bytes memory _message
  ) external;
}
```

Notice that the “message” is of type “bytes” in Solidity, so it is up to the Smart Contract developer to interpret what that message is. The Synapse Carbon messaging system serves merely as a fundamental building block that allows developers to treat sending a message as a black box. Simply send some bytes on the Origin chain, and those bytes will be received on the Destination chain.
