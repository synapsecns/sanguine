// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract AgentsTestContract {
  event AgentsEventA(
    address indexed sender,
    uint256 indexed valueA,
    uint256 indexed valueB,
    uint256 valueC
  );
  event AgentsEventB(
    address indexed sender,
    bytes valueA,
    uint256 valueB,
    uint256 valueC
  );

  event IMessageReceipientHandleEvent(
    uint32 indexed _origin,
    uint32 indexed _nonce,
    bytes32 _sender,
    uint256 _rootSubmittedAt,
    bytes _message
  );

  function emitAgentsEventA(
    uint256 valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emit AgentsEventA(msg.sender, valueA, valueB, valueC);
  }

  function emitAgentsEventB(
    bytes memory valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emit AgentsEventB(msg.sender, valueA, valueB, valueC);
  }

  function emitAgentsEventAandB(
    uint256 valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emitAgentsEventA(valueA, valueB, valueC);
    emitAgentsEventB(abi.encodePacked(valueA), valueB, valueC);
  }

  function handle(
    uint32 _origin,
    uint32 _nonce,
    bytes32 _sender,
    uint256 _rootSubmittedAt,
    bytes memory _message
  ) external {
    emit IMessageReceipientHandleEvent(
      _origin,
      _nonce,
      _sender,
      _rootSubmittedAt,
      _message
    );
  }
}
