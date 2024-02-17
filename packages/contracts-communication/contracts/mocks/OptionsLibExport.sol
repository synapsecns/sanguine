pragma solidity 0.8.20;

import {OptionsLib} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

contract OptionsLibMocks {
    function encodeOptions(OptionsLib.Options memory options) public view returns (bytes memory) {
        return OptionsLib.encodeOptions(options);
    }

    function decodeOptions(bytes memory data) public view returns (OptionsLib.Options memory) {
        return OptionsLib.decodeOptions(data);
    }

    function addressToBytes32(address convertable) public view returns (bytes32) {
        return TypeCasts.addressToBytes32(convertable);
    }
}
