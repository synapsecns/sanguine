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

// solhint-disable no-empty-blocks, ordering
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

    IFastBridge.BridgeTransaction public tokenTx;
    IFastBridge.BridgeTransaction public ethTx;
    IFastBridge.BridgeParams public tokenParams;
    IFastBridge.BridgeParams public ethParams;

    function setUp() public virtual {
        srcToken = new MockERC20("SrcToken", 6);
        dstToken = new MockERC20("DstToken", 6);
        createFixtures();
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

        tokenTx = IFastBridge.BridgeTransaction({
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
        });
        ethTx = IFastBridge.BridgeTransaction({
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
        });
    }

    function getTxId(IFastBridge.BridgeTransaction memory bridgeTx) public pure returns (bytes32) {
        return keccak256(abi.encode(bridgeTx));
    }

    function expectUnauthorized(address caller, bytes32 role) public {
        vm.expectRevert(abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role));
    }

    function cheatCollectedProtocolFees(address token, uint256 amount) public {
        stdstore.target(address(fastBridge)).sig("protocolFees(address)").with_key(token).checked_write(amount);
    }
}
