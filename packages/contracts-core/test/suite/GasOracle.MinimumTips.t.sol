// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {NumberLib} from "../../contracts/libs/stack/Number.sol";
import {Request, RequestLib} from "../../contracts/libs/stack/Request.sol";
import {Tips, TipsLib} from "../../contracts/libs/stack/Tips.sol";
import {TIPS_MULTIPLIER} from "../../contracts/libs/Constants.sol";

import {Random, GasOracle, GasOracleTest} from "./GasOracle.t.sol";
import {RawGasData, RawGasData256} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleMinimumTipsTest is GasOracleTest {
    // We're using powers of two to avoid rounding errors (see NumberLib.compress()).
    // Tips are stored with TIPS_MULTIPLIER=2**32, so we use values above 2**32.
    uint256 public constant GAS_PRICE = 1 << 32;
    uint256 public constant DATA_PRICE = 1 << 36;
    uint256 public constant EXEC_BUFFER = 1 << 40;
    uint256 public constant AMORT_ATT_COST = 1 << 44;

    uint256 public constant HALF_BWAD = NumberLib.BWAD / 2;
    uint256 public constant TWO_BWAD = NumberLib.BWAD * 2;

    function setUp() public override {
        super.setUp();
        // Everything except etherPrice is not used and can be set to 0.
        // Ether price for local chain is set to 0.5 ETH.
        testedGO().setGasData({
            domain: DOMAIN_LOCAL,
            gasPrice: 0,
            dataPrice: 0,
            execBuffer: 0,
            amortAttCost: 0,
            etherPrice: HALF_BWAD,
            markup: 0
        });
        // All values are going to be used.
        // Ether price for remote chain is set to 2.0 ETH, markup is 50%.
        testedGO().setGasData({
            domain: DOMAIN_REMOTE,
            gasPrice: GAS_PRICE,
            dataPrice: DATA_PRICE,
            execBuffer: EXEC_BUFFER,
            amortAttCost: AMORT_ATT_COST,
            etherPrice: TWO_BWAD,
            markup: HALF_BWAD
        });
    }

    function test_getMinimumTips_summitTip(Random memory random) public {
        uint256 summitTip = 1 << 48;
        testedGO().setSummitTip(summitTip);
        uint256 paddedRequest = random.nextUint192();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Local chain Ether price is 0.5 ETH, so the summit tip is worth twice as much.
        assertEq(tips.summitTip() * TIPS_MULTIPLIER, summitTip * 2, "!summitTip");
    }

    function test_getMinimumTips_attestationTip(Random memory random) public {
        uint256 paddedRequest = random.nextUint192();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the attestation tip is worth 4 times more.
        assertEq(tips.attestationTip() * TIPS_MULTIPLIER, AMORT_ATT_COST * 4, "!attestationTip");
    }

    function test_getMinimumTips_executionTip(Random memory random) public {
        Request request = RequestLib.encodeRequest({
            gasDrop_: random.nextUint96(),
            gasLimit_: random.nextUint32(),
            version_: random.nextUint32()
        });
        uint256 paddedRequest = Request.unwrap(request);
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the execution tip is worth 4 times more.
        // Execution tip is execBuffer + gasPrice * gasLimit + dataPrice * contentLength.
        assertEq(
            tips.executionTip() * TIPS_MULTIPLIER,
            4 * (EXEC_BUFFER + GAS_PRICE * request.gasLimit() + DATA_PRICE * contentLength),
            "!executionTip"
        );
    }

    function test_getMinimumTips_deliveryTip(Random memory random) public {
        Request request = RequestLib.encodeRequest({
            gasDrop_: random.nextUint96(),
            gasLimit_: random.nextUint32(),
            version_: random.nextUint32()
        });
        uint256 paddedRequest = Request.unwrap(request);
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the execution tip is worth 4 times more.
        // Execution tip is execBuffer + gasPrice * gasLimit + dataPrice * contentLength.
        // Delivery tip is execution tip * markup (50%), therefore the final value is 2 times more.
        assertEq(
            tips.deliveryTip() * TIPS_MULTIPLIER,
            2 * (EXEC_BUFFER + GAS_PRICE * request.gasLimit() + DATA_PRICE * contentLength),
            "!deliveryTip"
        );
        // TODO: adjust test when delivery tip will also include the value of gas airdrop.
    }
}
