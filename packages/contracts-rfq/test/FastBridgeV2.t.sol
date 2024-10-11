// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib} from "../contracts/libs/BridgeTransactionV2.sol";

import {IFastBridge} from "../contracts/interfaces/IFastBridge.sol";

// solhint-disable-next-line no-unused-import
import {IFastBridgeV2} from "../contracts/interfaces/IFastBridgeV2.sol";

import {IFastBridgeV2Errors} from "../contracts/interfaces/IFastBridgeV2Errors.sol";
import {FastBridgeV2} from "../contracts/FastBridgeV2.sol";

import {MockERC20} from "./MockERC20.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Test} from "forge-std/Test.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";

// solhint-disable no-empty-blocks, max-states-count, ordering
abstract contract FastBridgeV2Test is Test, IFastBridgeV2Errors {
    using stdStorage for StdStorage;

    uint32 public constant SRC_CHAIN_ID = 1337;
    uint32 public constant DST_CHAIN_ID = 7331;
    uint256 public constant DEADLINE = 1 days;
    address public constant ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    FastBridgeV2 public fastBridge;
    MockERC20 public srcToken;
    MockERC20 public dstToken;

    address public relayerA = makeAddr("Relayer A");
    address public relayerB = makeAddr("Relayer B");
    address public guard = makeAddr("Guard");
    address public userA = makeAddr("User A");
    address public userB = makeAddr("User B");
    address public governor = makeAddr("Governor");
    address public refunder = makeAddr("Refunder");

    IFastBridgeV2.BridgeTransactionV2 internal tokenTx;
    IFastBridgeV2.BridgeTransactionV2 internal ethTx;
    IFastBridge.BridgeParams internal tokenParams;
    IFastBridge.BridgeParams internal ethParams;

    IFastBridgeV2.BridgeParamsV2 internal tokenParamsV2;
    IFastBridgeV2.BridgeParamsV2 internal ethParamsV2;

    bytes internal mockRequestV1;
    bytes internal invalidRequestV2;
    bytes internal mockRequestV3;

    function createInvalidRequestV2(bytes calldata requestV2) external pure returns (bytes memory) {
        // Cut the last 1 byte
        return requestV2[:requestV2.length - 1];
    }

    function createMockRequestV3(bytes calldata requestV2) external pure returns (bytes memory) {
        return abi.encodePacked(uint16(3), requestV2[2:]);
    }

    function setUp() public virtual {
        srcToken = new MockERC20("SrcToken", 6);
        dstToken = new MockERC20("DstToken", 6);
        createFixtures();
        mockRequestV1 = abi.encode(extractV1(tokenTx));
        // Invalid V2 request is formed before `createFixturesV2` to ensure it's not using callParams
        invalidRequestV2 = this.createInvalidRequestV2(BridgeTransactionV2Lib.encodeV2(tokenTx));
        createFixturesV2();
        // Mock V3 request is formed after `createFixturesV2` to ensure it's using callParams if needed
        mockRequestV3 = this.createMockRequestV3(BridgeTransactionV2Lib.encodeV2(ethTx));
        fastBridge = deployFastBridge();
        configureFastBridge();
        mintTokens();
    }

    function deployFastBridge() public virtual returns (FastBridgeV2);

    function configureFastBridge() public virtual {}

    function mintTokens() public virtual {}

    function createFixtures() public virtual {
        tokenParams = IFastBridge.BridgeParams({
            dstChainId: DST_CHAIN_ID,
            sender: userA,
            to: userB,
            originToken: address(srcToken),
            destToken: address(dstToken),
            originAmount: 1e6,
            destAmount: 0.99e6,
            sendChainGas: false,
            deadline: block.timestamp + DEADLINE
        });
        ethParams = IFastBridge.BridgeParams({
            dstChainId: DST_CHAIN_ID,
            sender: userA,
            to: userB,
            originToken: ETH_ADDRESS,
            destToken: ETH_ADDRESS,
            originAmount: 1 ether,
            destAmount: 0.99 ether,
            sendChainGas: false,
            deadline: block.timestamp + DEADLINE
        });

        setStorageBridgeTxV2(
            tokenTx,
            IFastBridge.BridgeTransaction({
                originChainId: SRC_CHAIN_ID,
                destChainId: DST_CHAIN_ID,
                originSender: userA,
                destRecipient: userB,
                originToken: address(srcToken),
                destToken: address(dstToken),
                originAmount: 1e6,
                destAmount: 0.99e6,
                // override this in tests with protocol fees
                originFeeAmount: 0,
                sendChainGas: false,
                deadline: block.timestamp + DEADLINE,
                nonce: 0
            })
        );
        setStorageBridgeTxV2(
            ethTx,
            IFastBridge.BridgeTransaction({
                originChainId: SRC_CHAIN_ID,
                destChainId: DST_CHAIN_ID,
                originSender: userA,
                destRecipient: userB,
                originToken: ETH_ADDRESS,
                destToken: ETH_ADDRESS,
                originAmount: 1 ether,
                destAmount: 0.99 ether,
                // override this in tests with protocol fees
                originFeeAmount: 0,
                sendChainGas: false,
                deadline: block.timestamp + DEADLINE,
                nonce: 1
            })
        );
    }

    function createFixturesV2() public virtual {
        // Override in tests with exclusivity params
        tokenParamsV2 = IFastBridgeV2.BridgeParamsV2({
            quoteRelayer: address(0),
            quoteExclusivitySeconds: 0,
            quoteId: bytes(""),
            callValue: 0,
            callParams: bytes("")
        });
        ethParamsV2 = IFastBridgeV2.BridgeParamsV2({
            quoteRelayer: address(0),
            quoteExclusivitySeconds: 0,
            quoteId: bytes(""),
            callValue: 0,
            callParams: bytes("")
        });

        tokenTx.exclusivityRelayer = address(0);
        tokenTx.exclusivityEndTime = block.timestamp;
        ethTx.exclusivityRelayer = address(0);
        ethTx.exclusivityEndTime = block.timestamp;
    }

    function setStorageBridgeTxV2(
        IFastBridgeV2.BridgeTransactionV2 storage txV2,
        IFastBridge.BridgeTransaction memory txV1
    )
        internal
    {
        txV2.originChainId = txV1.originChainId;
        txV2.destChainId = txV1.destChainId;
        txV2.originSender = txV1.originSender;
        txV2.destRecipient = txV1.destRecipient;
        txV2.originToken = txV1.originToken;
        txV2.destToken = txV1.destToken;
        txV2.originAmount = txV1.originAmount;
        txV2.destAmount = txV1.destAmount;
        txV2.originFeeAmount = txV1.originFeeAmount;
        txV2.deadline = txV1.deadline;
        txV2.nonce = txV1.nonce;
    }

    function setTokenTestCallParams(bytes memory callParams) public {
        tokenParamsV2.callParams = callParams;
        tokenTx.callParams = callParams;
    }

    function setTokenTestCallValue(uint256 callValue) public {
        tokenParamsV2.callValue = callValue;
        tokenTx.callValue = callValue;
    }

    function setTokenTestExclusivityParams(address relayer, uint256 exclusivitySeconds) public {
        tokenParamsV2.quoteRelayer = relayer;
        tokenParamsV2.quoteExclusivitySeconds = int256(exclusivitySeconds);
        tokenParamsV2.quoteId = bytes.concat("Token:", getMockQuoteId(relayer));

        tokenTx.exclusivityRelayer = relayer;
        tokenTx.exclusivityEndTime = block.timestamp + exclusivitySeconds;
    }

    function setEthTestCallParams(bytes memory callParams) public {
        ethParamsV2.callParams = callParams;
        ethTx.callParams = callParams;
    }

    function setEthTestCallValue(uint256 callValue) public {
        ethParamsV2.callValue = callValue;
        ethTx.callValue = callValue;
    }

    function setEthTestExclusivityParams(address relayer, uint256 exclusivitySeconds) public {
        ethParamsV2.quoteRelayer = relayer;
        ethParamsV2.quoteExclusivitySeconds = int256(exclusivitySeconds);
        ethParamsV2.quoteId = bytes.concat("ETH:", getMockQuoteId(relayer));

        ethTx.exclusivityRelayer = relayer;
        ethTx.exclusivityEndTime = block.timestamp + exclusivitySeconds;
    }

    function extractV1(IFastBridgeV2.BridgeTransactionV2 memory txV2)
        public
        pure
        returns (IFastBridge.BridgeTransaction memory txV1)
    {
        txV1.originChainId = txV2.originChainId;
        txV1.destChainId = txV2.destChainId;
        txV1.originSender = txV2.originSender;
        txV1.destRecipient = txV2.destRecipient;
        txV1.originToken = txV2.originToken;
        txV1.destToken = txV2.destToken;
        txV1.originAmount = txV2.originAmount;
        txV1.destAmount = txV2.destAmount;
        txV1.originFeeAmount = txV2.originFeeAmount;
        txV1.deadline = txV2.deadline;
        txV1.nonce = txV2.nonce;
    }

    function getMockQuoteId(address relayer) public view returns (bytes memory) {
        if (relayer == relayerA) {
            return bytes("created by Relayer A");
        } else if (relayer == relayerB) {
            return bytes("created by Relayer B");
        } else {
            return bytes("created by unknown relayer");
        }
    }

    function getTxId(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public pure returns (bytes32) {
        return keccak256(BridgeTransactionV2Lib.encodeV2(bridgeTx));
    }

    function expectUnauthorized(address caller, bytes32 role) public {
        vm.expectRevert(abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role));
    }

    function expectRevertInvalidEncodedTx() public {
        vm.expectRevert(BridgeTransactionV2Lib.BridgeTransactionV2__InvalidEncodedTx.selector);
    }

    function expectRevertUnsupportedVersion(uint16 version) public {
        vm.expectRevert(
            abi.encodeWithSelector(BridgeTransactionV2Lib.BridgeTransactionV2__UnsupportedVersion.selector, version)
        );
    }

    function cheatCollectedProtocolFees(address token, uint256 amount) public {
        stdstore.target(address(fastBridge)).sig("protocolFees(address)").with_key(token).checked_write(amount);
    }

    function cheatSenderNonce(address sender, uint256 nonce) public {
        stdstore.target(address(fastBridge)).sig("senderNonces(address)").with_key(sender).checked_write(nonce);
    }
}
