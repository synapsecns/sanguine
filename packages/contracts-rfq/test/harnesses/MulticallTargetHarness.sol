// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {MulticallTarget} from "../../contracts/utils/MulticallTarget.sol";

contract MulticallTargetHarness is MulticallTarget {
    address public addressField;
    uint256 public uintField;

    string public constant REVERT_MESSAGE = "gm, this is a revert message";

    error CustomError();

    function setAddressField(address _addressField) external returns (address) {
        addressField = _addressField;
        return addressField;
    }

    function setUintField(uint256 _uintField) external returns (uint256) {
        uintField = _uintField;
        return uintField;
    }

    function customErrorRevert() external pure {
        revert CustomError();
    }

    function revertingFunction() external pure {
        revert(REVERT_MESSAGE);
    }

    function undeterminedRevert() external pure {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            revert(0, 0)
        }
    }
}
