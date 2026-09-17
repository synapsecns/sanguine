// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {SynapseOFTAdapterFactory} from "../src/SynapseOFTAdapterFactory.sol";

import {EndpointMock} from "./mocks/EndpointMock.sol";
import {MintableTestToken} from "./mocks/MintableTestToken.sol";

import {
    ILayerZeroEndpointV2,
    MessagingFee,
    MessagingReceipt,
    Origin
} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {OAppReceiver} from "@layerzerolabs/oapp-evm/contracts/oapp/OAppReceiver.sol";
import {IOAppCore} from "@layerzerolabs/oapp-evm/contracts/oapp/interfaces/IOAppCore.sol";
import {IOFT, OFTReceipt, SendParam} from "@layerzerolabs/oft-evm/contracts/interfaces/IOFT.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";
import {Test} from "forge-std/Test.sol";

contract SynapseOFTAdapterTest is Test {
    uint32 internal constant SRC_EID = 1;
    uint32 internal constant DST_EID = 2;
    uint256 internal constant CONVERSION_RATE = 1e12;
    bytes32 internal constant PEER = keccak256("Peer");
    bytes32 internal constant GUID = keccak256("Guid");

    SynapseOFTAdapter internal adapter;
    MintableTestToken internal token;
    EndpointMock internal endpoint;

    address internal delegate = makeAddr("Delegate");
    address internal user = makeAddr("User");
    address internal recipient = makeAddr("Recipient");

    function setUp() public {
        token = new MintableTestToken();
        endpoint = new EndpointMock();
        SynapseOFTAdapterFactory factory = new SynapseOFTAdapterFactory(delegate);
        vm.prank(delegate);
        factory.initialize(address(endpoint));
        vm.prank(delegate);
        adapter = SynapseOFTAdapter(factory.deploy(address(token), bytes32(0)));

        vm.prank(delegate);
        adapter.setPeer(DST_EID, PEER);
    }

    function testSetPeerRevertsForNonOwner() public {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user));
        vm.prank(user);
        adapter.setPeer(SRC_EID, PEER);
    }

    function testSendBurnsApprovedAmountAndRemovesDust() public {
        uint256 amountLD = 1 ether + 123;
        uint256 expectedAmountLD = 1 ether;
        token.mint(user, amountLD);

        vm.prank(user);
        token.approve(address(adapter), expectedAmountLD);
        _mockSend();

        vm.prank(user);
        (MessagingReceipt memory messagingReceipt, OFTReceipt memory oftReceipt) =
            adapter.send(_sendParam(recipient, amountLD, expectedAmountLD), MessagingFee(0, 0), user);

        assertEq(messagingReceipt.guid, GUID);
        assertEq(oftReceipt.amountSentLD, expectedAmountLD);
        assertEq(oftReceipt.amountReceivedLD, expectedAmountLD);
        assertEq(token.balanceOf(user), 123);
        assertEq(token.totalSupply(), 123);
        assertEq(token.allowance(user, address(adapter)), 0);
        assertEq(token.balanceOf(address(adapter)), 0);
    }

    function testSendRevertsWithoutBurnAllowance() public {
        uint256 amountLD = 1 ether;
        token.mint(user, amountLD);

        vm.expectRevert(
            abi.encodeWithSelector(IERC20Errors.ERC20InsufficientAllowance.selector, address(adapter), 0, amountLD)
        );
        vm.prank(user);
        adapter.send(_sendParam(recipient, amountLD, amountLD), MessagingFee(0, 0), user);

        assertEq(token.balanceOf(user), amountLD);
        assertEq(token.totalSupply(), amountLD);
    }

    function testSendRevertsWhenDustCausesSlippage() public {
        uint256 amountLD = 1 ether + 123;
        token.mint(user, amountLD);
        vm.prank(user);
        token.approve(address(adapter), amountLD);

        vm.expectRevert(abi.encodeWithSelector(IOFT.SlippageExceeded.selector, 1 ether, amountLD));
        vm.prank(user);
        adapter.send(_sendParam(recipient, amountLD, amountLD), MessagingFee(0, 0), user);

        assertEq(token.balanceOf(user), amountLD);
        assertEq(token.totalSupply(), amountLD);
    }

    function testLzReceiveMintsTokensAndIncreasesSupply() public {
        uint64 amountSD = 123_456;
        uint256 amountLD = uint256(amountSD) * CONVERSION_RATE;
        _setSourcePeer();

        vm.prank(address(endpoint));
        adapter.lzReceive(_origin(PEER), GUID, _message(recipient, amountSD), address(0), "");

        assertEq(token.balanceOf(recipient), amountLD);
        assertEq(token.totalSupply(), amountLD);
        assertEq(token.balanceOf(address(adapter)), 0);
    }

    function testLzReceiveMapsZeroRecipientToDeadAddress() public {
        uint64 amountSD = 7;
        uint256 amountLD = uint256(amountSD) * CONVERSION_RATE;
        _setSourcePeer();

        vm.prank(address(endpoint));
        adapter.lzReceive(_origin(PEER), GUID, _message(address(0), amountSD), address(0), "");

        assertEq(token.balanceOf(address(0xdead)), amountLD);
        assertEq(token.balanceOf(address(0)), 0);
        assertEq(token.totalSupply(), amountLD);
    }

    function testLzReceiveRevertsForNonEndpoint() public {
        _setSourcePeer();

        vm.expectRevert(abi.encodeWithSelector(OAppReceiver.OnlyEndpoint.selector, user));
        vm.prank(user);
        adapter.lzReceive(_origin(PEER), GUID, _message(recipient, 1), address(0), "");
    }

    function testLzReceiveRevertsForUntrustedPeer() public {
        bytes32 unknownPeer = keccak256("Unknown Peer");
        _setSourcePeer();

        vm.expectRevert(abi.encodeWithSelector(IOAppCore.OnlyPeer.selector, SRC_EID, unknownPeer));
        vm.prank(address(endpoint));
        adapter.lzReceive(_origin(unknownPeer), GUID, _message(recipient, 1), address(0), "");
    }

    function testConstructorSetsMetadataAndOwner() public view {
        assertEq(adapter.token(), address(token));
        assertEq(adapter.owner(), delegate);
        assertEq(address(adapter.endpoint()), address(endpoint));
        assertEq(adapter.sharedDecimals(), 6);
        assertEq(adapter.decimalConversionRate(), CONVERSION_RATE);
        assertTrue(adapter.approvalRequired());
    }

    function _setSourcePeer() internal {
        vm.prank(delegate);
        adapter.setPeer(SRC_EID, PEER);
    }

    function _mockSend() internal {
        MessagingReceipt memory receipt = MessagingReceipt({guid: GUID, nonce: 1, fee: MessagingFee(0, 0)});
        vm.mockCall(address(endpoint), abi.encodeWithSelector(ILayerZeroEndpointV2.send.selector), abi.encode(receipt));
    }

    function _sendParam(address to, uint256 amountLD, uint256 minAmountLD) internal pure returns (SendParam memory) {
        return SendParam({
            dstEid: DST_EID,
            to: bytes32(uint256(uint160(to))),
            amountLD: amountLD,
            minAmountLD: minAmountLD,
            extraOptions: "",
            composeMsg: "",
            oftCmd: ""
        });
    }

    function _origin(bytes32 sender) internal pure returns (Origin memory) {
        return Origin({srcEid: SRC_EID, sender: sender, nonce: 1});
    }

    function _message(address to, uint64 amountSD) internal pure returns (bytes memory) {
        return abi.encodePacked(bytes32(uint256(uint160(to))), amountSD);
    }
}
