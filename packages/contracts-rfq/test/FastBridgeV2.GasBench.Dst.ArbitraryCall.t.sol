// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkDstTest} from "./FastBridgeV2.GasBench.Dst.t.sol";
import {RecipientMock} from "./mocks/RecipientMock.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2GasBenchmarkDstArbitraryCallTest is FastBridgeV2GasBenchmarkDstTest {
    // To get an idea about how much overhead the arbitrary call adds to the relaying process, we use a mock
    // recipient that has the hook function implemented as a no-op.
    // The mocked callParams are chosen to be similar to the real use cases:
    // - user address
    // - some kind of ID to decide what to do with the tokens next

    function setUp() public virtual override {
        // In the inherited tests userB is always used as the recipient of the tokens.
        userB = address(new RecipientMock());
        vm.label(userB, "ContractRecipient");
        super.setUp();
    }

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        bytes memory mockCallParams = abi.encode(userA, keccak256("Random ID"));
        setTokenTestCallParams(mockCallParams);
        setEthTestCallParams(mockCallParams);
    }
}
