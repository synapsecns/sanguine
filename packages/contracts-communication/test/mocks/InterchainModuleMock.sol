// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, IInterchainModule} from "../../contracts/interfaces/IInterchainModule.sol";

contract InterchainModuleMock is IInterchainModule {
    function requestVerification(uint256 destChainId, IInterchainDB.InterchainEntry memory entry) external payable {}

    function getModuleFee(uint256 destChainId) external view returns (uint256) {}

    function mockVerifyEntry(address interchainDB, IInterchainDB.InterchainEntry memory entry) external {
        IInterchainDB(interchainDB).verifyEntry(entry);
    }
}
