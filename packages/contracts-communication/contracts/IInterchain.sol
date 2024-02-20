pragma solidity 0.8.20;

interface IInterchain {
    function interchainSend(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        address[] calldata modules
    )
        external
        payable;

    function interchainReceive(bytes calldata transaction) external;
}
