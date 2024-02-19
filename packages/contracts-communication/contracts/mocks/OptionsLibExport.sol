pragma solidity 0.8.20;

import {OptionsLib, OptionsV1} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

contract OptionsLibMocks {
    function encodeOptions(OptionsV1 memory options) public view returns (bytes memory) {
        return OptionsLib.encodeOptionsV1(options);
    }

    function decodeOptions(bytes memory data) public view returns (OptionsV1 memory) {
        return OptionsLib.decodeOptionsV1(data);
    }

    function addressToBytes32(address convertable) public view returns (bytes32) {
        return TypeCasts.addressToBytes32(convertable);
    }
}
