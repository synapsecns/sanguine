// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBurnableToken} from "./interfaces/IBurnableToken.sol";
import {ISynapseBridge} from "./interfaces/ISynapseBridge.sol";
import {ISynapseBridgeAdapter} from "./interfaces/ISynapseBridgeAdapter.sol";
import {ISynapseBridgeAdapterErrors} from "./interfaces/ISynapseBridgeAdapterErrors.sol";
import {ISynapseBridgeAdapterV2} from "./interfaces/ISynapseBridgeAdapterV2.sol";
import {ISynapseBridgeAdapterV2Errors} from "./interfaces/ISynapseBridgeAdapterV2Errors.sol";
import {ISynapseHyperCoreComposer} from "./interfaces/ISynapseHyperCoreComposer.sol";
import {BridgeMessage} from "./libs/BridgeMessage.sol";
import {HyperCoreMessage} from "./libs/HyperCoreMessage.sol";

import {MessagingFee, OApp, Origin} from "@layerzerolabs/oapp-evm/contracts/oapp/OApp.sol";
import {OptionsBuilder} from "@layerzerolabs/oapp-evm/contracts/oapp/libs/OptionsBuilder.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice A separately deployed SynapseBridgeAdapter version that can deliver bridged tokens to HyperCore.
/// @dev The original adapter and its deployed routes remain unchanged.
contract SynapseBridgeAdapterV2 is
    OApp,
    ISynapseBridgeAdapterV2,
    ISynapseBridgeAdapterErrors,
    ISynapseBridgeAdapterV2Errors
{
    using OptionsBuilder for bytes;
    using SafeERC20 for IERC20;

    uint64 public constant MIN_GAS_LIMIT = 200_000;
    uint8 public constant SYN_EVM_DECIMALS = 18;
    uint8 public constant SYN_CORE_DECIMALS = 8;
    uint256 public constant SYN_AMOUNT_SCALE = 1e10;
    uint256 private constant DIRECT_MESSAGE_LENGTH = 96;

    address public bridge;

    mapping(uint32 eid => mapping(address remoteAddr => address localAddr)) public getLocalAddress;
    mapping(uint32 eid => mapping(address localAddr => address remoteAddr)) public getRemoteAddress;
    mapping(address localAddr => TokenType tokenType) public getTokenType;
    mapping(uint32 eid => mapping(address token => uint256 amountScale)) public getHyperCoreAmountScale;
    mapping(address token => address composer) public getHyperCoreComposer;

    event BridgeSet(address bridge);
    event HyperCoreComposerSet(address indexed token, address indexed composer);
    event HyperCoreDecimalsSet(
        uint32 indexed dstEid, address indexed token, uint8 tokenDecimals, uint8 coreDecimals, uint256 amountScale
    );
    event TokenAdded(address token, TokenType tokenType, RemoteToken[] remoteTokens);
    event TokenSent(uint32 indexed dstEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);
    event TokenSentToHyperCore(
        uint32 indexed dstEid, address indexed to, address indexed token, uint256 amount, bytes32 guid
    );
    event TokenReceived(uint32 indexed srcEid, address indexed to, address indexed token, uint256 amount, bytes32 guid);
    event TokenReceivedOnHyperCore(
        uint32 indexed srcEid, address indexed to, address indexed token, uint256 amount, bytes32 guid
    );

    constructor(address endpoint_, address owner_) OApp(endpoint_, owner_) Ownable(owner_) {}

    // ════════════════════════════════════════════════ MANAGEMENT ═════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function addToken(address token, TokenType tokenType, RemoteToken[] memory remoteTokens) external onlyOwner {
        _checkAndSaveToken(token, tokenType);
        uint256 length = remoteTokens.length;
        if (length == 0) revert SBA__ZeroAmount();
        for (uint256 i = 0; i < length; ++i) {
            uint32 eid = remoteTokens[i].eid;
            if (getRemoteAddress[eid][token] != address(0)) revert SBA__RemotePairAlreadySet(eid, token);
            address remoteAddr = remoteTokens[i].addr;
            if (remoteAddr == address(0)) revert SBA__ZeroAddress();
            if (getLocalAddress[eid][remoteAddr] != address(0)) revert SBA__LocalPairAlreadyExists(eid, remoteAddr);
            getRemoteAddress[eid][token] = remoteAddr;
            getLocalAddress[eid][remoteAddr] = token;
        }
        emit TokenAdded(token, tokenType, remoteTokens);
    }

    /// @inheritdoc ISynapseBridgeAdapter
    function setBridge(address newBridge) external onlyOwner {
        if (newBridge == address(0)) revert SBA__ZeroAddress();
        if (bridge != address(0)) revert SBA__BridgeAlreadySet();
        bridge = newBridge;
        emit BridgeSet(newBridge);
    }

    /// @inheritdoc ISynapseBridgeAdapterV2
    function setHyperCoreDecimals(uint32 dstEid, address token, uint8 coreDecimals) external onlyOwner {
        _checkAndGetTokenType(token);
        if (getRemoteAddress[dstEid][token] == address(0)) revert SBA__RemotePairNotSet(dstEid, token);
        if (getHyperCoreAmountScale[dstEid][token] != 0) {
            revert SBAV2__HyperCoreDecimalsAlreadySet(dstEid, token);
        }

        uint8 tokenDecimals = IERC20Metadata(token).decimals();
        if (tokenDecimals != SYN_EVM_DECIMALS || coreDecimals != SYN_CORE_DECIMALS) {
            revert SBAV2__HyperCoreDecimalsInvalid(tokenDecimals, coreDecimals);
        }
        uint256 amountScale = SYN_AMOUNT_SCALE;
        getHyperCoreAmountScale[dstEid][token] = amountScale;
        emit HyperCoreDecimalsSet(dstEid, token, tokenDecimals, coreDecimals, amountScale);
    }

    /// @inheritdoc ISynapseBridgeAdapterV2
    function setHyperCoreComposer(address token, address composer) external onlyOwner {
        _checkAndGetTokenType(token);
        if (composer == address(0)) revert SBA__ZeroAddress();
        if (getHyperCoreComposer[token] != address(0)) revert SBAV2__HyperCoreComposerAlreadySet(token);
        if (composer.code.length == 0) revert SBAV2__HyperCoreComposerInvalid(composer);

        ISynapseHyperCoreComposer candidate = ISynapseHyperCoreComposer(composer);
        if (candidate.adapter() != address(this) || candidate.token() != token) {
            revert SBAV2__HyperCoreComposerInvalid(composer);
        }

        getHyperCoreComposer[token] = composer;
        emit HyperCoreComposerSet(token, composer);
    }

    // ════════════════════════════════════════════════ USER FACING ════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function bridgeERC20(uint32 dstEid, address to, address token, uint256 amount, uint64 gasLimit) external payable {
        bytes32 guid = _bridgeERC20({
            dstEid: dstEid,
            to: to,
            token: token,
            amount: amount,
            gasLimit: gasLimit,
            message: BridgeMessage.encodeBridgeMessage(to, token, amount)
        });
        // All custody and messaging calls have completed, and this function writes no state afterward.
        // slither-disable-next-line reentrancy-events
        emit TokenSent(dstEid, to, token, amount, guid);
    }

    /// @inheritdoc ISynapseBridgeAdapterV2
    function bridgeERC20ToHyperCore(
        uint32 dstEid,
        address to,
        address token,
        uint256 amount,
        uint64 gasLimit
    )
        external
        payable
    {
        uint256 amountScale = getHyperCoreAmountScale[dstEid][token];
        if (amountScale == 0) revert SBAV2__HyperCoreDecimalsNotSet(dstEid, token);
        if (amount % amountScale != 0) revert SBAV2__HyperCoreAmountNotRepresentable(amount, amountScale);
        uint256 coreAmount = amount / amountScale;
        if (coreAmount > type(uint64).max) revert SBAV2__HyperCoreAmountExceedsUint64(coreAmount);

        bytes32 guid = _bridgeERC20({
            dstEid: dstEid,
            to: to,
            token: token,
            amount: amount,
            gasLimit: gasLimit,
            message: HyperCoreMessage.encodeHyperCoreMessage(to, token, amount)
        });
        // All custody and messaging calls have completed, and this function writes no state afterward.
        // slither-disable-next-line reentrancy-events
        emit TokenSent(dstEid, to, token, amount, guid);
        // slither-disable-next-line reentrancy-events
        emit TokenSentToHyperCore(dstEid, to, token, amount, guid);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc ISynapseBridgeAdapter
    function getNativeFee(uint32 dstEid, uint64 gasLimit) external view returns (uint256 nativeFee) {
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        return _quote({
            _dstEid: dstEid,
            _message: BridgeMessage.encodeBridgeMessage(address(0), address(0), 0),
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _payInLzToken: false
        })
        .nativeFee;
    }

    /// @inheritdoc ISynapseBridgeAdapterV2
    function getNativeFeeToHyperCore(
        uint32 dstEid,
        address token,
        uint64 gasLimit
    )
        external
        view
        returns (uint256 nativeFee)
    {
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        if (getHyperCoreAmountScale[dstEid][token] == 0) {
            revert SBAV2__HyperCoreDecimalsNotSet(dstEid, token);
        }
        return _quote({
            _dstEid: dstEid,
            _message: HyperCoreMessage.encodeHyperCoreMessage(address(0), token, 0),
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _payInLzToken: false
        })
        .nativeFee;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _lzReceive(
        Origin calldata origin,
        bytes32 guid,
        bytes calldata message,
        address,
        bytes calldata
    )
        internal
        override
    {
        if (message.length == DIRECT_MESSAGE_LENGTH) {
            _receiveBridgeMessage(origin.srcEid, guid, message);
        } else {
            _receiveHyperCoreMessage(origin.srcEid, guid, message);
        }
    }

    function _receiveBridgeMessage(uint32 srcEid, bytes32 guid, bytes calldata message) internal {
        (address to, address srcToken, uint256 amount) = BridgeMessage.decodeBridgeMessage(message);
        address token = _mintOrWithdraw(srcEid, srcToken, to, amount, guid);
        // Only the authenticated LayerZero endpoint can enter this path, and no state is written after the bridge call.
        // slither-disable-next-line reentrancy-events
        emit TokenReceived(srcEid, to, token, amount, guid);
    }

    function _receiveHyperCoreMessage(uint32 srcEid, bytes32 guid, bytes calldata message) internal {
        (address to, address srcToken, uint256 amount) = HyperCoreMessage.decodeHyperCoreMessage(message);
        address token = _checkAndGetLocalAddress(srcEid, srcToken);
        address composer = getHyperCoreComposer[token];
        if (composer == address(0)) revert SBAV2__HyperCoreComposerNotSet(token);

        _mintOrWithdraw(token, composer, amount, guid);
        ISynapseHyperCoreComposer(composer).bridgeToHyperCore(to, amount);
        // Only the authenticated LayerZero endpoint can enter this path, and no state is written after these calls.
        // slither-disable-next-line reentrancy-events
        emit TokenReceivedOnHyperCore(srcEid, to, token, amount, guid);
        // slither-disable-next-line reentrancy-events
        emit TokenReceived(srcEid, to, token, amount, guid);
    }

    function _mintOrWithdraw(address token, address recipient, uint256 amount, bytes32 guid) internal {
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        _mintOrWithdraw(cachedBridge, token, recipient, amount, guid);
    }

    function _mintOrWithdraw(
        uint32 srcEid,
        address srcToken,
        address recipient,
        uint256 amount,
        bytes32 guid
    )
        internal
        returns (address token)
    {
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        token = _checkAndGetLocalAddress(srcEid, srcToken);
        _mintOrWithdraw(cachedBridge, token, recipient, amount, guid);
    }

    function _mintOrWithdraw(
        address cachedBridge,
        address token,
        address recipient,
        uint256 amount,
        bytes32 guid
    )
        internal
    {
        TokenType tokenType = _checkAndGetTokenType(token);
        if (tokenType == TokenType.MintBurn) {
            ISynapseBridge(cachedBridge).mint(recipient, token, amount, 0, guid);
        } else {
            ISynapseBridge(cachedBridge).withdraw(recipient, token, amount, 0, guid);
        }
    }

    function _bridgeERC20(
        uint32 dstEid,
        address to,
        address token,
        uint256 amount,
        uint64 gasLimit,
        bytes memory message
    )
        internal
        returns (bytes32 guid)
    {
        if (to == address(0)) revert SBA__ZeroAddress();
        if (amount == 0) revert SBA__ZeroAmount();
        if (gasLimit < MIN_GAS_LIMIT) revert SBA__GasLimitBelowMinimum();
        address cachedBridge = bridge;
        if (cachedBridge == address(0)) revert SBA__BridgeNotSet();
        TokenType tokenType = _checkAndGetTokenType(token);
        if (getRemoteAddress[dstEid][token] == address(0)) revert SBA__RemotePairNotSet(dstEid, token);

        if (tokenType == TokenType.MintBurn) {
            IBurnableToken(token).burnFrom(msg.sender, amount);
        } else {
            IERC20(token).safeTransferFrom(msg.sender, cachedBridge, amount);
        }

        return _lzSend({
            _dstEid: dstEid,
            _message: message,
            _options: OptionsBuilder.newOptions().addExecutorLzReceiveOption({_gas: gasLimit, _value: 0}),
            _fee: MessagingFee({nativeFee: msg.value, lzTokenFee: 0}),
            _refundAddress: msg.sender
        })
        .guid;
    }

    function _checkAndSaveToken(address token, TokenType tokenType) internal {
        if (token == address(0)) revert SBA__ZeroAddress();
        if (tokenType == TokenType.Unknown) revert SBA__TokenTypeUnknown();
        TokenType existingTokenType = getTokenType[token];
        if (existingTokenType == TokenType.Unknown) {
            getTokenType[token] = tokenType;
        } else if (existingTokenType != tokenType) {
            revert SBA__TokenAlreadyAdded(token);
        }
    }

    function _checkAndGetTokenType(address token) internal view returns (TokenType tokenType) {
        tokenType = getTokenType[token];
        if (tokenType == TokenType.Unknown) revert SBA__TokenUnknown(token);
    }

    function _checkAndGetLocalAddress(uint32 remoteEid, address remoteAddr) internal view returns (address localAddr) {
        localAddr = getLocalAddress[remoteEid][remoteAddr];
        if (localAddr == address(0)) revert SBA__LocalPairNotFound(remoteEid, remoteAddr);
    }
}
