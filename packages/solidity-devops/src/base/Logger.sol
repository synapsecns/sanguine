// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {StringUtils} from "../libs/StringUtils.sol";

import {console2} from "forge-std/Script.sol";

abstract contract Logger {
    using StringUtils for *;

    string private constant TAB = "    ";

    /// @dev Current indent level for all log messages
    uint256 private currentIndentLevel;

    modifier withIndent() {
        increaseIndent();
        _;
        decreaseIndent();
    }

    /// @notice Returns the current indent string.
    /// @dev Handy if the log message contains arguments other than strings and printLog() is not used.
    function currentIndent() internal view returns (string memory) {
        return TAB.duplicate(currentIndentLevel);
    }

    /// @notice Increases the indent level for all log messages.
    function increaseIndent() internal {
        ++currentIndentLevel;
    }

    /// @notice Decreases the indent level for all log messages.
    function decreaseIndent() internal {
        require(currentIndentLevel > 0, "Indent level cannot be negative");
        --currentIndentLevel;
    }

    /// @notice Prints the log message with the current indent level.
    function printLog(string memory logString) internal view {
        console2.log("%s%s", currentIndent(), logString);
    }

    /// @notice Prints the informational log message with the current indent level.
    function printInfo(string memory logString) internal view {
        printLog(getInfoEmoji().concat(" ", logString));
    }

    /// @notice Prints the log message with the current indent level plus one.
    function printLogWithIndent(string memory logString) internal view {
        printLog(TAB.concat(logString));
    }

    /// @notice Prints the "skipping" message with the current indent level plus one.
    function printSkipWithIndent(string memory reason) internal view {
        printLogWithIndent(getSkipEmoji().concat(" Skipping: ", reason));
    }

    /// @notice Prints the "fail" message with the current indent level plus one.
    function printFailWithIndent(string memory logString) internal view {
        printLogWithIndent(getFailEmoji().concat(" ", logString));
    }

    /// @notice Prints the "success" message with the current indent level plus one.
    function printSuccessWithIndent(string memory logString) internal view {
        printLogWithIndent(getSuccessEmoji().concat(" ", logString));
    }

    /// @notice Should return "💬"
    function getInfoEmoji() internal pure virtual returns (string memory);

    /// @notice Should return "🟡"
    function getSkipEmoji() internal pure virtual returns (string memory);

    /// @notice Should return "❌"
    function getFailEmoji() internal pure virtual returns (string memory);

    /// @notice Should return "✅"
    function getSuccessEmoji() internal pure virtual returns (string memory);
}
