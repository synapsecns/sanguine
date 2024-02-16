pragma solidity 0.8.20;

interface IInterchainModuleV1 {
    function sendModuleMessage(bytes calldata transaction) external payable;

    function receiveModuleMessage() external;

    function estimateFee(uint256 dstChainId) external view returns (uint256);
}
