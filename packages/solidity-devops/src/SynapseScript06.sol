// SPDX-License-Identifier: MIT
pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;

import {SynapseBaseScript} from "./base/SynapseBaseScript.sol";

// Imports for external consumption
import {StringUtils} from "./libs/StringUtils.sol";
import {stdJson} from "forge-std/Script.sol";

abstract contract SynapseScript06 is SynapseBaseScript {
    /// @notice Should return "💬"
    function getInfoEmoji() internal pure virtual override returns (string memory) {
        return "💬";
    }

    /// @notice Should return "🟡"
    function getSkipEmoji() internal pure virtual override returns (string memory) {
        return "🟡";
    }

    /// @notice Should return "❌"
    function getFailEmoji() internal pure virtual override returns (string memory) {
        return "❌";
    }

    /// @notice Should return "✅"
    function getSuccessEmoji() internal pure virtual override returns (string memory) {
        return "✅";
    }
}
