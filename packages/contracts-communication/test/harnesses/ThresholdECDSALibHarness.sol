// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ThresholdECDSA, ThresholdECDSALib} from "../../contracts/libs/ThresholdECDSA.sol";

contract ThresholdECDSALibHarness {
    ThresholdECDSA internal thresholdECDSA;

    function addSigner(address account) external {
        ThresholdECDSALib.addSigner(thresholdECDSA, account);
    }

    function removeSigner(address account) external {
        ThresholdECDSALib.removeSigner(thresholdECDSA, account);
    }

    function modifyThreshold(uint256 threshold) external {
        ThresholdECDSALib.modifyThreshold(thresholdECDSA, threshold);
    }

    function isSigner(address account) external view returns (bool) {
        return ThresholdECDSALib.isSigner(thresholdECDSA, account);
    }

    function getSigners() external view returns (address[] memory) {
        return ThresholdECDSALib.getSigners(thresholdECDSA);
    }

    function getThreshold() external view returns (uint256) {
        return ThresholdECDSALib.getThreshold(thresholdECDSA);
    }

    function verifySignedHash(bytes32 hash, bytes calldata signatures) external view {
        ThresholdECDSALib.verifySignedHash(thresholdECDSA, hash, signatures);
    }
}
