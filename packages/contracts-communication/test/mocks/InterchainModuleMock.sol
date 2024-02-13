// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";

contract InterchainModuleMock is IInterchainModule {
    function verifyEntry(uint256 destChainId, IInterchainDB.InterchainEntry memory entry) external payable {}

    function getModuleFee(uint256 destChainId) external view returns (uint256) {}

    function mockConfirmEntry(address interchainDB, IInterchainDB.InterchainEntry memory entry) external {
        IInterchainDB(interchainDB).confirmEntry(entry);
    }
}
