// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;
pragma experimental ABIEncoderV2;

import {SynapseBaseScript} from "./SynapseBaseScript.sol";

// Imports for external consumption
import {StringUtils} from "./libs/StringUtils.sol";
import {stdJson} from "forge-std/Script.sol";

abstract contract SynapseScript is SynapseBaseScript {
    /// @notice Should return "💬"
    function getInfoEmoji() internal pure virtual override returns (string memory) {
        return unicode"💬";
    }

    /// @notice Should return "🟡"
    function getSkipEmoji() internal pure virtual override returns (string memory) {
        return unicode"🟡";
    }

    /// @notice Should return "❌"
    function getFailEmoji() internal pure virtual override returns (string memory) {
        return unicode"❌";
    }

    /// @notice Should return "✅"
    function getSuccessEmoji() internal pure virtual override returns (string memory) {
        return unicode"✅";
    }
}
