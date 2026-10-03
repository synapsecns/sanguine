// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseHyperCoreComposer} from "../src/SynapseHyperCoreComposer.sol";
import {ISynapseBridgeAdapter} from "../src/interfaces/ISynapseBridgeAdapter.sol";
import {SynapseBridgeAdapterTest, SynapseBridgeAdapterV2} from "./SBA.t.sol";

import {TestToken} from "./mocks/TestToken.sol";
import {TestTokenDecimals} from "./mocks/TestTokenDecimals.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseBridgeAdapterManagementTest is SynapseBridgeAdapterTest {
    address internal bridge = makeAddr("Bridge");
    address internal owner = makeAddr("Owner");
    address internal token = makeAddr("Token");
    address internal anotherToken = makeAddr("AnotherToken");
    bytes31 internal symbol = "SYMBOL";
    bytes31 internal anotherSymbol = "ANOTHERSYMBOL";
    string internal readableSymbol = "SYMBOL";
    TestToken internal hyperCoreToken;

    mapping(uint32 eid => address token) internal mockRemoteAddressMap;
    ISynapseBridgeAdapter.RemoteToken[] internal allRemoteTokens;
    ISynapseBridgeAdapter.RemoteToken[] internal firstRemoteToken;
    ISynapseBridgeAdapter.RemoteToken[] internal secondRemoteToken;

    function afterAdapterDeployed() internal virtual override {
        mockRemoteAddressMap[DST_EID] = makeAddr("DST Token");
        mockRemoteAddressMap[OTHER_DST_EID] = makeAddr("Other DST Token");
        allRemoteTokens.push(ISynapseBridgeAdapter.RemoteToken(DST_EID, mockRemoteAddressMap[DST_EID]));
        allRemoteTokens.push(ISynapseBridgeAdapter.RemoteToken(OTHER_DST_EID, mockRemoteAddressMap[OTHER_DST_EID]));
        firstRemoteToken.push(allRemoteTokens[0]);
        secondRemoteToken.push(allRemoteTokens[1]);
        hyperCoreToken = new TestToken();
    }

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapterV2) {
        return new SynapseBridgeAdapterV2(endpoint, owner);
    }

    function checkTokenAdded(
        address token_,
        ISynapseBridgeAdapter.TokenType tokenType_,
        ISynapseBridgeAdapter.RemoteToken[] memory remoteTokens_
    )
        internal
        view
    {
        // Check token type by address
        ISynapseBridgeAdapter.TokenType adapterTokenType = adapter.getTokenType(token_);
        assertEq(uint8(adapterTokenType), uint8(tokenType_));
        for (uint256 i = 0; i < remoteTokens_.length; i++) {
            ISynapseBridgeAdapter.RemoteToken memory remoteToken = remoteTokens_[i];
            // Check remote address by local address
            address adapterRemoteAddress = adapter.getRemoteAddress(remoteToken.eid, token_);
            assertEq(adapterRemoteAddress, remoteToken.addr);
            // Check local address by remote address
            address adapterLocalAddress = adapter.getLocalAddress(remoteToken.eid, remoteToken.addr);
            assertEq(adapterLocalAddress, token_);
        }
    }

    function test_constructor() public view {
        assertEq(address(adapter.endpoint()), endpoint);
        assertEq(adapter.owner(), owner);
        assertEq(adapter.bridge(), address(0));
        assertEq(adapter.MIN_GAS_LIMIT(), 200_000);
        assertEq(adapter.SYN_EVM_DECIMALS(), 18);
        assertEq(adapter.SYN_CORE_DECIMALS(), 8);
        assertEq(adapter.SYN_AMOUNT_SCALE(), 1e10);
    }

    // ═════════════════════════════════════════════════ ADD TOKEN ═════════════════════════════════════════════════════

    function addToken(
        address token_,
        ISynapseBridgeAdapter.TokenType tokenType_,
        ISynapseBridgeAdapter.RemoteToken[] memory remoteTokens_
    )
        internal
    {
        vm.prank(owner);
        adapter.addToken(token_, tokenType_, remoteTokens_);
    }

    function test_addToken_mintBurn_oneByOne() public {
        // Add first token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add second token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_allAtOnce() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenAlreadyAddedAsWithdrawDeposit() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
    }

    function test_addToken_mintBurn_revert_remotePairAlreadySet_sameRemoteAddress() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_remotePairAlreadySet_diffRemoteAddress() public {
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, anotherToken))
        );
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_localPairAlreadyExists() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as single entry
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add as a batch
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenTypeUnknown() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_remoteTokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(ISynapseBridgeAdapter.RemoteToken(firstRemoteToken[0].eid, address(0)))
        );
        // Add as a batch
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(firstRemoteToken[0], ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, address(0)))
        );
    }

    function test_addToken_mintBurn_revert_existingTokenEmptyArray() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_mintBurn_revert_newTokenEmptyArray() public {
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_withdrawDeposit_oneByOne() public {
        // Add first token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add second token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_allAtOnce() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenAlreadyAddedAsMint() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
    }

    function test_addToken_withdrawDeposit_revert_remotePairAlreadySet_sameRemoteAddress() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_remotePairAlreadySet_diffRemoteAddress() public {
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, anotherToken))
        );
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_localPairAlreadyExists() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as single entry
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add as a batch
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenTypeUnknown() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_remoteTokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(ISynapseBridgeAdapter.RemoteToken(firstRemoteToken[0].eid, address(0)))
        );
        // Add as a batch
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(firstRemoteToken[0], ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, address(0)))
        );
    }

    function test_addToken_withdrawDeposit_revert_existingTokenEmptyArray() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_withdrawDeposit_revert_newTokenEmptyArray() public {
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_revert_tokenTypeUnknown() public {
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    // ════════════════════════════════════════════════ SET BRIDGE ═════════════════════════════════════════════════════

    function test_setBridge() public {
        expectEventBridgeSet(bridge);
        vm.prank(owner);
        adapter.setBridge(bridge);
        assertEq(adapter.bridge(), bridge);
    }

    function test_setBridge_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_bridgeAlreadySet() public {
        vm.prank(owner);
        adapter.setBridge(bridge);
        expectRevertBridgeAlreadySet();
        vm.prank(owner);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_zeroAddress() public {
        expectRevertZeroAddress();
        vm.prank(owner);
        adapter.setBridge(address(0));
    }

    // ══════════════════════════════════════════ SET HYPERCORE DECIMALS ══════════════════════════════════════════════

    function test_setHyperCoreDecimals() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        vm.expectEmit(address(adapter));
        emit HyperCoreDecimalsSet(DST_EID, address(hyperCoreToken), 18, 8, 1e10);
        vm.prank(owner);
        adapter.setHyperCoreDecimals(DST_EID, address(hyperCoreToken), 8);
        assertEq(adapter.getHyperCoreAmountScale(DST_EID, address(hyperCoreToken)), 1e10);
    }

    function test_setHyperCoreDecimals_revert_invalidDecimals() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreDecimalsInvalid.selector, 18, 19));
        vm.prank(owner);
        adapter.setHyperCoreDecimals(DST_EID, address(hyperCoreToken), 19);
    }

    function test_setHyperCoreDecimals_revert_invalidTokenDecimals() public {
        TestTokenDecimals invalidToken = new TestTokenDecimals(6);
        addToken(address(invalidToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreDecimalsInvalid.selector, 6, 8));
        vm.prank(owner);
        adapter.setHyperCoreDecimals(DST_EID, address(invalidToken), 8);
    }

    function test_setHyperCoreDecimals_revert_alreadySet() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        vm.startPrank(owner);
        adapter.setHyperCoreDecimals(DST_EID, address(hyperCoreToken), 8);
        vm.expectRevert(
            abi.encodeWithSelector(SBAV2__HyperCoreDecimalsAlreadySet.selector, DST_EID, address(hyperCoreToken))
        );
        adapter.setHyperCoreDecimals(DST_EID, address(hyperCoreToken), 8);
        vm.stopPrank();
    }

    function test_setHyperCoreDecimals_revert_remotePairNotSet() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertRemotePairNotSet(OTHER_DST_EID, address(hyperCoreToken));
        vm.prank(owner);
        adapter.setHyperCoreDecimals(OTHER_DST_EID, address(hyperCoreToken), 8);
    }

    function test_setHyperCoreDecimals_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.setHyperCoreDecimals(DST_EID, address(hyperCoreToken), 8);
    }

    // ══════════════════════════════════════════ SET HYPERCORE COMPOSER ══════════════════════════════════════════════

    function test_setHyperCoreComposer() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        SynapseHyperCoreComposer composer =
            new SynapseHyperCoreComposer(address(adapter), address(hyperCoreToken), 1337);
        vm.expectEmit(address(adapter));
        emit HyperCoreComposerSet(address(hyperCoreToken), address(composer));
        vm.prank(owner);
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
        assertEq(adapter.getHyperCoreComposer(address(hyperCoreToken)), address(composer));
    }

    function test_setHyperCoreComposer_revert_invalidBinding() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        SynapseHyperCoreComposer composer = new SynapseHyperCoreComposer(address(this), address(hyperCoreToken), 1337);
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreComposerInvalid.selector, address(composer)));
        vm.prank(owner);
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
    }

    function test_setHyperCoreComposer_revert_invalidTokenBinding() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        TestToken differentToken = new TestToken();
        SynapseHyperCoreComposer composer =
            new SynapseHyperCoreComposer(address(adapter), address(differentToken), 1337);
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreComposerInvalid.selector, address(composer)));
        vm.prank(owner);
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
    }

    function test_setHyperCoreComposer_revert_notContract() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        address composer = makeAddr("Composer");
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreComposerInvalid.selector, composer));
        vm.prank(owner);
        adapter.setHyperCoreComposer(address(hyperCoreToken), composer);
    }

    function test_setHyperCoreComposer_revert_alreadySet() public {
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        SynapseHyperCoreComposer composer =
            new SynapseHyperCoreComposer(address(adapter), address(hyperCoreToken), 1337);
        vm.startPrank(owner);
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
        vm.expectRevert(abi.encodeWithSelector(SBAV2__HyperCoreComposerAlreadySet.selector, address(hyperCoreToken)));
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
        vm.stopPrank();
    }

    function test_setHyperCoreComposer_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        addToken(address(hyperCoreToken), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        SynapseHyperCoreComposer composer =
            new SynapseHyperCoreComposer(address(adapter), address(hyperCoreToken), 1337);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.setHyperCoreComposer(address(hyperCoreToken), address(composer));
    }
}
