// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../utils/SynapseTestSuite.t.sol";

abstract contract NotaryManagerTools is SynapseTestSuite {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectNewOrigin(address origin) public {
        vm.expectEmit(true, true, true, true);
        emit NewOrigin(origin);
    }

    function expectNewNotary(address notary) public {
        vm.expectEmit(true, true, true, true);
        emit NewNotary(notary);
    }

    function expectFakeSlashed(address reporter) public {
        vm.expectEmit(true, true, true, true);
        emit FakeSlashed(reporter);
    }
}
