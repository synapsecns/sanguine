// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBurnableToken} from "./interfaces/IBurnableToken.sol";
import {ISynapseBridge} from "./interfaces/ISynapseBridge.sol";
import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "./interfaces/ISynapseBridgeAdapterErrors.sol";
import {BridgeMessage} from "./libs/BridgeMessage.sol";

import {MessagingFee, OApp, Origin} from "@layerzerolabs/oapp-evm/contracts/oapp/OApp.sol";
import {OptionsBuilder} from "@layerzerolabs/oapp-evm/contracts/oapp/libs/OptionsBuilder.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract SynapseBridgeAdapter is OApp, ISynapseBridgeAdapter, ISynapseBridgeAdapterErrors {
    using OptionsBuilder for bytes;
    using SafeERC20 for IERC20;

    uint64 public constant MIN_GAS_LIMIT = 200_000;

    address public bridge;

    mapping(uint32 eid => mapping(address remoteAddr => address localAddr)) public getLocalAddress;
    mapping(uint32 eid => mapping(address localAddr => address remoteAddr)) public getRemoteAddress;
    mapping(address localAddr => TokenType tokenType) public getTokenType;

    event BridgeSet(address bridge);
    event TokenAdded(address token, TokenType tokenType, RemoteToken[] remoteTokens);
    event TokenSent(uint32 indexed dstEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);
    event TokenReceived(uint32 indexed srcEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);

    constructor(address endpoint_, address owner_) OApp(endpoint_, owner_) Ownable(owner_) {}

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function addToken(address token, TokenType tokenType, RemoteToken[] memory remoteTokens) external onlyOwner {
        // Check local token and its type, store if it's the first addition
        _checkAndSaveToken(token, tokenType);
        uint256 length = remoteTokens.length;
        if (length == 0) revert SBA__ZeroAmount();
        for (uint256 i = 0; i < length; ++i) {
            // Check that a remote token pair has not been set for the local token and eid
            uint32 eid = remoteTokens[i].eid;
            if (getRemoteAddress[eid][token] != address(0)) revert SBA__RemotePairAlreadySet(eid, token);
            // Check that a remote address is not zero and have not been used for any other local token and given eid
            address remoteAddr = remoteTokens[i].addr;
            if (remoteAddr == address(0)) revert SBA__ZeroAddress();
            if (getLocalAddress[eid][remoteAddr] != address(0)) revert SBA__LocalPairAlreadyExists(eid, remoteAddr);
            // Store remote <> local address mappings
            getRemoteAddress[eid][token] = remoteAddr;
            getLocalAddress[eid][remoteAddr] = token;
        }
        emit TokenAdded(token, tokenType, remoteTokens);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function setBridge(address newBridge) external onlyOwner {
        // Check: new parameters
        if (newBridge == address(0)) revert SBA__ZeroAddress();
        // Check: existing state
        if (bridge != address(0)) revert SBA__BridgeAlreadySet();
        // Store
        bridge = newBridge;
        emit BridgeSet(newBridge);
    }

    // ════════════════════════════════════════════════ USER FACING ════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function bridgeERC20(uint32 dstEid, address to, address token, uint256 amount, uint64 gasLimit) external payable {
        // Verify the parameters
        if (to == address(0)) revert SBA__ZeroAddress();
        if (amount == 0) revert SBA__ZeroAmount();
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        // Cache bridge address
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        // Check token's type (note: this reverts if token is unknown)
        TokenType tokenType = _checkAndGetTokenType(token);
        // Check that a remote token pair exists for the given local token and eid
        if (getRemoteAddress[dstEid][token] == address(0)) revert SBA__RemotePairNotSet(dstEid, token);
        // Burn tokens from sender or deposit them into the bridge as prerequisite
        if (tokenType == TokenType.MintBurn) {
            IBurnableToken(token).burnFrom(msg.sender, amount);
        } else {
            IERC20(token).safeTransferFrom(msg.sender, cachedBridge, amount);
        }
        // Send the bridge message (note: we use the source token address for the message)
        bytes32 guid = _lzSend({
            _dstEid: dstEid,
            _message: BridgeMessage.encodeBridgeMessage(to, token, amount),
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _fee: MessagingFee({nativeFee: msg.value, lzTokenFee: 0}),
            _refundAddress: msg.sender
        }).guid;
        emit TokenSent(dstEid, to, token, amount, guid);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function getNativeFee(uint32 dstEid, uint64 gasLimit) external view returns (uint256 nativeFee) {
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        // Since all the messages have the same length, we can use arbitrary data for fee estimation
        bytes memory message = BridgeMessage.encodeBridgeMessage(address(0), address(0), 0);
        return _quote({
            _dstEid: dstEid,
            _message: message,
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _payInLzToken: false
        }).nativeFee;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Handles the received message from origin's Adapter counterpart.
    /// All validation checks for integrity of the message have been performed at this point.
    function _lzReceive(
        Origin calldata _origin,
        bytes32 _guid,
        bytes calldata _message,
        address, // _executor,
        bytes calldata // _extraData
    )
        internal
        override
    {
        (address to, address srcToken, uint256 amount) = BridgeMessage.decodeBridgeMessage(_message);
        // Cache guid to avoid stack too deep error
        bytes32 guid = _guid;
        // Cache bridge address
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        // Cache token type and address (note: this reverts if either source or local token is unknown)
        address token = _checkAndGetLocalAddress(_origin.srcEid, srcToken);
        TokenType tokenType = _checkAndGetTokenType(token);
        // Mint or withdraw tokens from the bridge as the result of the message
        // Note: the fees are set to 0 to enable 1:1 bridging
        // Note: guid is used as "kappa" (which is SynapseBridge's own global unique identifier),
        // clashes between two different systems are not possible due to being different keccak256 hashes.
        if (tokenType == TokenType.MintBurn) {
            ISynapseBridge(cachedBridge).mint(to, token, amount, 0, guid);
        } else {
            ISynapseBridge(cachedBridge).withdraw(to, token, amount, 0, guid);
        }
        emit TokenReceived(_origin.srcEid, to, token, amount, guid);
    }

    /// @dev Checks that the token is not already added with a different token type, stores if not added already.
    function _checkAndSaveToken(address token, TokenType tokenType) internal {
        if (token == address(0)) revert SBA__ZeroAddress();
        if (tokenType == TokenType.Unknown) revert SBA__TokenTypeUnknown();
        TokenType existingTokenType = getTokenType[token];
        if (existingTokenType == TokenType.Unknown) {
            // Save token type on the first addition
            getTokenType[token] = tokenType;
        } else if (existingTokenType != tokenType) {
            // Validate that the token type is the same on later additions
            revert SBA__TokenAlreadyAdded(token);
        }
    }

    /// @dev Checks that the local token has been added and returns its type.
    function _checkAndGetTokenType(address token) internal view returns (TokenType tokenType) {
        tokenType = getTokenType[token];
        if (tokenType == TokenType.Unknown) revert SBA__TokenUnknown(token);
    }

    /// @dev Checks that the local address pair exists for a given remoteEid:remoteAddr and returns it.
    function _checkAndGetLocalAddress(uint32 remoteEid, address remoteAddr) internal view returns (address localAddr) {
        localAddr = getLocalAddress[remoteEid][remoteAddr];
        if (localAddr == address(0)) revert SBA__LocalPairNotFound(remoteEid, remoteAddr);
    }
}
