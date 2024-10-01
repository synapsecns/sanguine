// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

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

    function setUp() public virtual {
        srcToken = new MockERC20("SrcToken", 6);
        dstToken = new MockERC20("DstToken", 6);
        createFixtures();
        createFixturesV2();
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
        tokenParamsV2 =
            IFastBridgeV2.BridgeParamsV2({quoteRelayer: address(0), quoteExclusivitySeconds: 0, quoteId: ""});
        ethParamsV2 = IFastBridgeV2.BridgeParamsV2({quoteRelayer: address(0), quoteExclusivitySeconds: 0, quoteId: ""});

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
        txV2.sendChainGas = txV1.sendChainGas;
        txV2.deadline = txV1.deadline;
        txV2.nonce = txV1.nonce;
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
        txV1.sendChainGas = txV2.sendChainGas;
        txV1.deadline = txV2.deadline;
        txV1.nonce = txV2.nonce;
    }

    function getTxId(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public pure returns (bytes32) {
        return keccak256(abi.encode(bridgeTx));
    }

    function expectUnauthorized(address caller, bytes32 role) public {
        vm.expectRevert(abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role));
    }

    function cheatCollectedProtocolFees(address token, uint256 amount) public {
        stdstore.target(address(fastBridge)).sig("protocolFees(address)").with_key(token).checked_write(amount);
    }
}
