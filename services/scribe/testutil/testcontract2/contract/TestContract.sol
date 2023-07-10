// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract TestContract {
  event EventA(
    address indexed sender,
    uint256 indexed valueA,
    uint256 indexed valueB,
    uint256 valueC
  );
  event EventB(
    address indexed sender,
    bytes valueA,
    uint256 valueB,
    uint256 valueC
  );

  function emitEventA(
    uint256 valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emit EventA(msg.sender, valueA, valueB, valueC);
  }

  function emitEventB(
    bytes memory valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emit EventB(msg.sender, valueA, valueB, valueC);
  }

  function emitEventAandB(
    uint256 valueA,
    uint256 valueB,
    uint256 valueC
  ) public {
    emitEventA(valueA, valueB, valueC);
    emitEventB(abi.encodePacked(valueA), valueB, valueC);
  }
}
