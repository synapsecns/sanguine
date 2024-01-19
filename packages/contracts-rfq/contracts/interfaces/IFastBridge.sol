// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridge {
    struct BridgeTransaction {
        uint32 originChainId;
        uint32 destChainId;
        address originSender; // user (origin)
        address destRecipient; // user (dest)
        address originToken;
        address destToken;
        uint256 originAmount; // amount in on origin bridge less originFeeAmount
        uint256 destAmount;
        uint256 originFeeAmount;
        bool sendChainGas;
        uint256 deadline; // user specified deadline for destination relay
        uint256 nonce;
    }

    struct BridgeProof {
        uint96 timestamp;
        address relayer;
    }

    // ============ Events ============

    event BridgeRequested(
        bytes32 indexed transactionId,
        address indexed sender,
        bytes request,
        uint32 destChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        bool sendChainGas
    );
    event BridgeRelayed(
        bytes32 indexed transactionId,
        address indexed relayer,
        address indexed to,
        uint32 originChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        uint256 chainGasAmount
    );
    event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash);
    event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer);
    event BridgeDepositClaimed(
        bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount
    );
    event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount);

    // ============ Methods ============

    struct BridgeParams {
        uint32 dstChainId;
        address sender;
        address to;
        address originToken;
        address destToken;
        uint256 originAmount; // should include protocol fee (if any)
        uint256 destAmount; // should include relayer fee
        bool sendChainGas;
        uint256 deadline;
    }

    /// @notice Initiates bridge on origin chain to be relayed by off-chain relayer
    /// @param params The parameters required to bridge
    function bridge(BridgeParams memory params) external payable;

    /// @notice Relays destination side of bridge transaction by off-chain relayer
    /// @param request The encoded bridge transaction to relay on destination chain
    function relay(bytes memory request) external payable;

    /// @notice Provides proof on origin side that relayer provided funds on destination side of bridge transaction
    /// @param request The encoded bridge transaction to prove on origin chain
    /// @param destTxHash The destination tx hash proving bridge transaction was relayed
    function prove(bytes memory request, bytes32 destTxHash) external;

    /// @notice Completes bridge transaction on origin chain by claiming originally deposited capital
    /// @param request The encoded bridge transaction to claim on origin chain
    /// @param to The recipient address of the funds
    function claim(bytes memory request, address to) external;

    /// @notice Disputes an outstanding proof in case relayer provided dest chain tx is invalid
    /// @param transactionId The transaction id associated with the encoded bridge transaction to dispute
    function dispute(bytes32 transactionId) external;

    /// @notice Refunds an outstanding bridge transaction in case optimistic bridging failed
    /// @param request The encoded bridge transaction to refund
    function refund(bytes memory request) external;

    // ============ Views ============

    /// @notice Decodes bridge request into a bridge transaction
    /// @param request The bridge request to decode
    function getBridgeTransaction(bytes memory request) external pure returns (BridgeTransaction memory);

    /// @notice Checks if the dispute period has passed so bridge deposit can be claimed
    /// @param transactionId The transaction id associated with the encoded bridge transaction to check
    /// @param relayer The address of the relayer attempting to claim
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool);
}
