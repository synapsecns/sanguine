// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";

import {ChainGas, GasData, GasDataHarness, Number} from "../../harnesses/libs/GasDataHarness.t.sol";

import {RawChainGas, RawGasData} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract GasDataLibraryTest is SynapseLibraryTest {
    GasDataHarness internal libHarness;

    function setUp() public {
        libHarness = new GasDataHarness();
    }

    function test_encodeGasData(RawGasData memory rgd) public {
        GasData gd = libHarness.encodeGasData({
            gasPrice_: rgd.gasPrice.encodeNumber(),
            dataPrice_: rgd.dataPrice.encodeNumber(),
            execBuffer_: rgd.execBuffer.encodeNumber(),
            amortAttCost_: rgd.amortAttCost.encodeNumber(),
            etherPrice_: rgd.etherPrice.encodeNumber(),
            markup_: rgd.markup.encodeNumber()
        });
        uint256 expected = uint256(rgd.gasPrice.number) * 2 ** 80 + uint256(rgd.dataPrice.number) * 2 ** 64
            + uint256(rgd.execBuffer.number) * 2 ** 48 + uint256(rgd.amortAttCost.number) * 2 ** 32
            + uint256(rgd.etherPrice.number) * 2 ** 16 + uint256(rgd.markup.number);
        assertEq(GasData.unwrap(gd), expected, "!encodeGasData");
        assertEq(GasData.unwrap(libHarness.wrapGasData(expected)), expected, "!wrapGasData");
        assertEq(Number.unwrap(libHarness.gasPrice(gd)), rgd.gasPrice.number, "!gasPrice");
        assertEq(Number.unwrap(libHarness.dataPrice(gd)), rgd.dataPrice.number, "!dataPrice");
        assertEq(Number.unwrap(libHarness.execBuffer(gd)), rgd.execBuffer.number, "!execBuffer");
        assertEq(Number.unwrap(libHarness.amortAttCost(gd)), rgd.amortAttCost.number, "!amortAttCost");
        assertEq(Number.unwrap(libHarness.etherPrice(gd)), rgd.etherPrice.number, "!etherPrice");
        assertEq(Number.unwrap(libHarness.markup(gd)), rgd.markup.number, "!markup");
    }

    function test_encodeChainGas(RawChainGas memory rcg) public {
        ChainGas cd = libHarness.encodeChainGas({gasData_: rcg.gasData.castToGasData(), domain_: rcg.domain});
        uint256 expected = uint256(rcg.gasData.encodeGasData()) * 2 ** 32 + uint256(rcg.domain);
        assertEq(ChainGas.unwrap(cd), expected, "!encodeChainGas");
        assertEq(ChainGas.unwrap(libHarness.wrapChainGas(expected)), expected, "!wrapChainGas");
        assertEq(libHarness.domain(cd), rcg.domain, "!domain");
        assertEq(GasData.unwrap(libHarness.gasData(cd)), rcg.gasData.encodeGasData(), "!gasData");
    }
}
