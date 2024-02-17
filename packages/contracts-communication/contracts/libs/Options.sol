// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library OptionsLib {
    struct Options {
        uint8 version;
        uint256 gasLimit;
        // uint256 msgValue;
        uint256 gasAirdrop;
    }

    function encodeOptions(Options memory options) internal pure returns (bytes memory) {
        return abi.encode(options.version, options.gasLimit, options.gasAirdrop);
    }

    function decodeOptions(bytes memory data) internal pure returns (Options memory) {
        (uint8 version, uint256 gasLimit, uint256 gasAirdrop) = abi.decode(data, (uint8, uint256, uint256));
        return Options(version, gasLimit, gasAirdrop);
    }
}
