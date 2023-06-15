// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {GasData, ChainGas, GasDataLib, Number} from "../../../../contracts/libs/stack/GasData.sol";

contract GasDataHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function encodeGasData(
        Number gasPrice_,
        Number dataPrice_,
        Number execBuffer_,
        Number amortAttCost_,
        Number etherPrice_,
        Number markup_
    ) public pure returns (GasData) {
        GasData result =
            GasDataLib.encodeGasData(gasPrice_, dataPrice_, execBuffer_, amortAttCost_, etherPrice_, markup_);
        return result;
    }

    function wrapGasData(uint256 paddedGasData) public pure returns (GasData) {
        return GasDataLib.wrapGasData(paddedGasData);
    }

    function encodeChainGas(GasData gasData_, uint32 domain_) public pure returns (ChainGas) {
        ChainGas result = GasDataLib.encodeChainGas(gasData_, domain_);
        return result;
    }

    function wrapChainGas(uint256 paddedChainGas) public pure returns (ChainGas) {
        return GasDataLib.wrapChainGas(paddedChainGas);
    }

    function gasPrice(GasData gasData_) public pure returns (Number) {
        return gasData_.gasPrice();
    }

    function dataPrice(GasData gasData_) public pure returns (Number) {
        return gasData_.dataPrice();
    }

    function execBuffer(GasData gasData_) public pure returns (Number) {
        return gasData_.execBuffer();
    }

    function amortAttCost(GasData gasData_) public pure returns (Number) {
        return gasData_.amortAttCost();
    }

    function etherPrice(GasData gasData_) public pure returns (Number) {
        return gasData_.etherPrice();
    }

    function markup(GasData gasData_) public pure returns (Number) {
        return gasData_.markup();
    }

    function domain(ChainGas chainData_) public pure returns (uint32) {
        return chainData_.domain();
    }

    function gasData(ChainGas chainData_) public pure returns (GasData) {
        return chainData_.gasData();
    }

    function snapGasHash(ChainGas[] memory snapGas) public pure returns (bytes32) {
        return GasDataLib.snapGasHash(snapGas);
    }
}
