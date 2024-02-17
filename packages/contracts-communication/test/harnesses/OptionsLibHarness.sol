// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsLib} from "../../contracts/libs/Options.sol";

contract OptionsLibHarness {
    function encodeOptions(
        uint8 version,
        uint256 gasLimit,
        //msgValue
        // uint256 gasAirdrop,
        OptionsLib.NativeDrop[] memory nativeDrops
    )
        external
        pure
        returns (bytes memory)
    {
        OptionsLib.Options memory options = OptionsLib.Options(version, gasLimit, nativeDrops);
        return OptionsLib.encodeOptions(options);
    }

    function decodeOptions(bytes calldata data) external pure returns (uint8, uint256, OptionsLib.NativeDrop[] memory) {
        OptionsLib.Options memory options = OptionsLib.decodeOptions(data);
        return (options.version, options.gasLimit, options.nativeDrops);
    }
}
