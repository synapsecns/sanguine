// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkDstTest} from "./FastBridgeV2.GasBench.Dst.t.sol";
import {RecipientMock} from "./mocks/RecipientMock.sol";

// solhint-disable func-name-mixedcase, no-empty-blocks
contract FastBridgeV2GasBenchmarkDstZapTest is FastBridgeV2GasBenchmarkDstTest {
    // To get an idea about how much overhead the Zap adds to the relaying process, we use a mock
    // recipient that has the hook function implemented as a no-op.
    // The mocked zapData are chosen to be similar to the real use cases:
    // - user address
    // - some kind of ID to decide what to do with the tokens next

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testFastBridgeV2GasBenchmarkDstZapTest() external {}

    function setUp() public virtual override {
        // In the inherited tests userB is always used as the recipient of the tokens.
        userB = address(new RecipientMock());
        vm.label(userB, "ContractRecipient");
        super.setUp();
    }

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        bytes memory mockZapData = abi.encode(userA, keccak256("Random ID"));
        setTokenTestZapData(mockZapData);
        setEthTestZapData(mockZapData);
    }
}
