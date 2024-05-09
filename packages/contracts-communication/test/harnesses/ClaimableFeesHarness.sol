// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ClaimableFees} from "../../contracts/fees/ClaimableFees.sol";

contract ClaimableFeesHarness is ClaimableFees {
    uint256 internal _claimableAmount;
    address internal _feeRecipient;
    uint256 internal _claimerFraction;

    event BeforeFeesClaimed(uint256 amount, uint256 reward);

    function setup(uint256 claimableAmount, address feeRecipient, uint256 claimerFraction) external {
        _claimableAmount = claimableAmount;
        _feeRecipient = feeRecipient;
        _claimerFraction = claimerFraction;
    }

    function getClaimerFraction() public view override returns (uint256) {
        return _claimerFraction;
    }

    function getClaimableAmount() public view override returns (uint256) {
        return _claimableAmount;
    }

    function getFeeRecipient() public view override returns (address) {
        return _feeRecipient;
    }

    function _beforeFeesClaimed(uint256 fullAmount, uint256 reward) internal override {
        _claimableAmount -= fullAmount;
        emit BeforeFeesClaimed(fullAmount, reward);
    }
}
