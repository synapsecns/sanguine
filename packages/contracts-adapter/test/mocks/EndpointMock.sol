// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {
    MessagingFee,
    MessagingParams,
    MessagingReceipt
} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";

// solhint-disable no-empty-blocks
contract EndpointMock {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testEndpointMock() external {}

    function setDelegate(address delegate) external {}

    function send(
        MessagingParams calldata _params,
        address _refundAddress
    )
        external
        payable
        returns (MessagingReceipt memory)
    {}

    function quote(MessagingParams calldata _params, address _sender) external view returns (MessagingFee memory) {}

    function eid() external view returns (uint32) {}
}
