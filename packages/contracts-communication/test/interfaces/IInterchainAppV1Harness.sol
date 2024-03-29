// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IInterchainAppV1} from "../../contracts/interfaces/IInterchainAppV1.sol";
import {InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";

// solhint-disable func-name-mixedcase
interface IInterchainAppV1Harness is IInterchainAppV1 {
    event MessageReceived(uint64 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message);

    function exposed__sendInterchainMessageEVM(
        uint256 dstChainId,
        address receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc);

    function exposed__sendInterchainMessage(
        uint256 dstChainId,
        bytes32 receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc);

    function exposed__sendToLinkedApp(
        uint256 dstChainId,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc);

    function exposed__getInterchainFee(
        uint256 dstChainId,
        bytes memory options,
        uint256 messageLen
    )
        external
        view
        returns (uint256);

    function exposed__getMessageFee(
        uint256 dstChainId,
        OptionsV1 memory options,
        uint256 messageLen
    )
        external
        view
        returns (uint256);
}
