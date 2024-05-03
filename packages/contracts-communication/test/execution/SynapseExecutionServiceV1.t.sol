// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ClaimableFeesEvents} from "../../contracts/events/ClaimableFeesEvents.sol";
import {SynapseExecutionServiceEvents} from "../../contracts/events/SynapseExecutionServiceEvents.sol";
import {SynapseExecutionServiceV1} from "../../contracts/execution/SynapseExecutionServiceV1.sol";
import {IClaimableFees} from "../../contracts/interfaces/IClaimableFees.sol";
import {ISynapseExecutionServiceV1} from "../../contracts/interfaces/ISynapseExecutionServiceV1.sol";

import {ProxyTest} from "../proxy/ProxyTest.t.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseExecutionServiceV1Test is ProxyTest, ClaimableFeesEvents, SynapseExecutionServiceEvents {
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

    SynapseExecutionServiceV1 public implementation;
    SynapseExecutionServiceV1 public service;

    address public admin = makeAddr("Admin");
    address public governor = makeAddr("Governor");

    function setUp() public virtual {
        implementation = new SynapseExecutionServiceV1();
        service = SynapseExecutionServiceV1(deployProxy(address(implementation)));
    }

    function expectEventClaimerFractionSet(uint256 claimerFraction) internal {
        vm.expectEmit(address(service));
        emit ClaimerFractionSet(claimerFraction);
    }

    function expectEventFeeRecipientSet(address feeRecipient) internal {
        vm.expectEmit(address(service));
        emit FeeRecipientSet(feeRecipient);
    }

    function expectEventExecutorEOASet(address executor) internal {
        vm.expectEmit(address(service));
        emit ExecutorEOASet(executor);
    }

    function expectEventGasOracleSet(address gasOracle) internal {
        vm.expectEmit(address(service));
        emit GasOracleSet(gasOracle);
    }

    function expectEventGlobalMarkupSet(uint256 globalMarkup) internal {
        vm.expectEmit(address(service));
        emit GlobalMarkupSet(globalMarkup);
    }

    function expectEventExecutionRequested(bytes32 transactionId, address client, uint256 executionFee) internal {
        vm.expectEmit(address(service));
        emit ExecutionRequested(transactionId, client, executionFee);
    }

    function expectRevertClaimerFractionAboveMax(uint256 claimerFraction, uint256 maxAllowed) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IClaimableFees.ClaimableFees__ClaimerFractionAboveMax.selector, claimerFraction, maxAllowed
            )
        );
    }

    function expectRevertGasOracleZeroAddress() internal {
        vm.expectRevert(ISynapseExecutionServiceV1.SynapseExecutionService__GasOracleZeroAddress.selector);
    }

    function expectRevertFeeAmountBelowMin(uint256 feeAmount, uint256 minRequired) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                ISynapseExecutionServiceV1.SynapseExecutionService__FeeAmountBelowMin.selector, feeAmount, minRequired
            )
        );
    }

    function expectRevertFeeRecipientZeroAddress() internal {
        vm.expectRevert(IClaimableFees.ClaimableFees__FeeRecipientZeroAddress.selector);
    }

    function expectRevertOptionsVersionNotSupported(uint256 version) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                ISynapseExecutionServiceV1.SynapseExecutionService__OptionsVersionNotSupported.selector, version
            )
        );
    }

    function expectRevertExecutorZeroAddress() internal {
        vm.expectRevert(ISynapseExecutionServiceV1.SynapseExecutionService__ExecutorZeroAddress.selector);
    }

    function expectRevertZeroAmount() internal {
        vm.expectRevert(IClaimableFees.ClaimableFees__FeeAmountZero.selector);
    }

    function expectRevertNotGovernor(address caller) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, GOVERNOR_ROLE)
        );
    }

    function expectRevertNotInterchainClient(address caller) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, IC_CLIENT_ROLE)
        );
    }
}
