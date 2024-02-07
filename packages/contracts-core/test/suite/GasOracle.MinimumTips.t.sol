// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IncorrectDestinationDomain, LocalGasDataNotSet, RemoteGasDataNotSet} from "../../contracts/libs/Errors.sol";
import {NumberLib} from "../../contracts/libs/stack/Number.sol";
import {Request, RequestLib} from "../../contracts/libs/stack/Request.sol";
import {Tips, TipsLib} from "../../contracts/libs/stack/Tips.sol";
import {TIPS_MULTIPLIER} from "../../contracts/libs/Constants.sol";

import {Random, GasOracle, GasOracleTest} from "./GasOracle.t.sol";
import {RawGasData, RawGasData256, RawRequest} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleMinimumTipsTest is GasOracleTest {
    // We're using powers of two to avoid rounding errors (see NumberLib.compress()).
    // Tips are stored with TIPS_MULTIPLIER=2**32, so we use values above 2**32.
    uint256 public constant GAS_PRICE = 1 << 40;
    uint256 public constant DATA_PRICE = 1 << 42;
    uint256 public constant EXEC_BUFFER = 1 << 44;
    uint256 public constant AMORT_ATT_COST = 1 << 46;

    uint256 public constant HALF_BWAD = NumberLib.BWAD / 2;
    uint256 public constant TWO_BWAD = NumberLib.BWAD * 2;
    uint256 public constant FOUR_BWAD = NumberLib.BWAD * 4;

    function setUp() public override {
        super.setUp();
        // Everything except etherPrice is not used and can be set to 0.
        // Ether price for local chain is set to 0.5 ETH.
        setLocalEtherPrice(HALF_BWAD);
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

    // ══════════════════════════════════════════════ TESTS: REVERTS ═══════════════════════════════════════════════════

    function test_getMinimumTips_revert_incorrectDestination() public {
        vm.expectRevert(IncorrectDestinationDomain.selector);
        testedGO().getMinimumTips(DOMAIN_LOCAL, 0, 0);
    }

    function test_getMinimumTips_revert_localGasDataNotSet() public {
        setLocalEtherPrice(0);
        vm.expectRevert(LocalGasDataNotSet.selector);
        testedGO().getMinimumTips(DOMAIN_REMOTE, 0, 0);
    }

    function test_getMinimumTips_revert_remoteGasDataNotSet() public {
        testedGO().setGasData({
            domain: DOMAIN_REMOTE,
            gasPrice: 0,
            dataPrice: 0,
            execBuffer: 0,
            amortAttCost: 0,
            etherPrice: 0,
            markup: 0
        });
        vm.expectRevert(RemoteGasDataNotSet.selector);
        testedGO().getMinimumTips(DOMAIN_REMOTE, 0, 0);
    }

    // ════════════════════════════════════════════ TESTS: MINIMUM TIPS ════════════════════════════════════════════════

    function test_getMinimumTips_summitTip(Random memory random) public {
        uint256 summitTip = 1 << 48;
        testedGO().setSummitTip(summitTip);
        uint256 paddedRequest = rawTestRequest(random).encodeRequest();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Local chain Ether price is 0.5 ETH, so the summit tip is worth 2x.
        assertEq(tips.summitTip() * TIPS_MULTIPLIER, summitTip * 2, "!summitTip: 0.5ETH");
        // Local chain Ether price is 2.0 ETH, so the summit tip is worth 0.5x.
        setLocalEtherPrice(TWO_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(tips.summitTip() * TIPS_MULTIPLIER, summitTip / 2, "!summitTip: 2.0ETH");
        // Local chain Ether price is 4.0 ETH, so the summit tip is worth 0.25x.
        setLocalEtherPrice(FOUR_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(tips.summitTip() * TIPS_MULTIPLIER, summitTip / 4, "!summitTip: 4.0ETH");
    }

    function test_getMinimumTips_attestationTip(Random memory random) public {
        uint256 paddedRequest = rawTestRequest(random).encodeRequest();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the attestation tip is worth 4x.
        assertEq(tips.attestationTip() * TIPS_MULTIPLIER, AMORT_ATT_COST * 4, "!attestationTip: 0.5ETH");
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 2.0 ETH.
        // So the attestation tip is worth 1x.
        setLocalEtherPrice(TWO_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(tips.attestationTip() * TIPS_MULTIPLIER, AMORT_ATT_COST, "!attestationTip: 2.0ETH");
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 4.0 ETH.
        // So the attestation tip is worth 0.5x.
        setLocalEtherPrice(FOUR_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(tips.attestationTip() * TIPS_MULTIPLIER, AMORT_ATT_COST / 2, "!attestationTip: 4.0ETH");
    }

    function test_getMinimumTips_executionTip(Random memory random) public {
        RawRequest memory request = rawTestRequest(random);
        uint256 paddedRequest = request.encodeRequest();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the execution tip is worth 4x.
        // Execution tip is execBuffer + gasPrice * gasLimit + dataPrice * contentLength.
        assertEq(
            tips.executionTip() * TIPS_MULTIPLIER,
            4 * (EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength),
            "!executionTip: 0.5ETH"
        );
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 2.0 ETH.
        // So the execution tip is worth 1x.
        setLocalEtherPrice(TWO_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(
            tips.executionTip() * TIPS_MULTIPLIER,
            EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength,
            "!executionTip: 2.0ETH"
        );
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 4.0 ETH.
        // So the execution tip is worth 0.5x.
        setLocalEtherPrice(FOUR_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(
            tips.executionTip() * TIPS_MULTIPLIER,
            (EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength) / 2,
            "!executionTip: 4.0ETH"
        );
    }

    function test_getMinimumTips_deliveryTip(Random memory random) public {
        RawRequest memory request = rawTestRequest(random);
        uint256 paddedRequest = request.encodeRequest();
        uint256 contentLength = random.nextUint16();
        Tips tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 0.5 ETH.
        // So the execution tip is worth 4x.
        // Execution tip is execBuffer + gasPrice * gasLimit + dataPrice * contentLength.
        // Delivery tip is execution tip * markup (50%), therefore the final value is worth 2x.
        assertEq(
            tips.deliveryTip() * TIPS_MULTIPLIER,
            2 * (EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength),
            "!deliveryTip: 0.5ETH"
        );
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 2.0 ETH.
        // So the execution tip is worth 1x.
        // Delivery tip is execution tip * markup (50%), therefore the final value is worth 0.5x.
        setLocalEtherPrice(TWO_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(
            tips.deliveryTip() * TIPS_MULTIPLIER,
            (EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength) / 2,
            "!deliveryTip: 2.0ETH"
        );
        // Remote chain Ether price is 2.0 ETH, local chain Ether price is 4.0 ETH.
        // So the execution tip is worth 0.5x.
        // Delivery tip is execution tip * markup (50%), therefore the final value is worth 0.25x.
        setLocalEtherPrice(FOUR_BWAD);
        tips = TipsLib.wrapPadded(testedGO().getMinimumTips(DOMAIN_REMOTE, paddedRequest, contentLength));
        assertEq(
            tips.deliveryTip() * TIPS_MULTIPLIER,
            (EXEC_BUFFER + GAS_PRICE * request.gasLimit + DATA_PRICE * contentLength) / 4,
            "!deliveryTip: 4.0ETH"
        );
        // TODO: adjust test when delivery tip will also include the value of gas airdrop.
    }

    function rawTestRequest(Random memory random) internal pure returns (RawRequest memory rr) {
        rr = random.nextRequest();
        // Set sensible max values for gasDrop and gasLimit.
        rr.boundRequest({maxGasDrop: 10 ** 20, maxGasLimit: 10 ** 8});
    }

    function setLocalEtherPrice(uint256 etherPrice) public {
        testedGO().setGasData({
            domain: DOMAIN_LOCAL,
            gasPrice: 0,
            dataPrice: 0,
            execBuffer: 0,
            amortAttCost: 0,
            etherPrice: etherPrice,
            markup: 0
        });
    }
}
