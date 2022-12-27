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
}
