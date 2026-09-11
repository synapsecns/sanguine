// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IBurnableToken} from "./interfaces/IBurnableToken.sol";
import {IMintableToken} from "./interfaces/IMintableToken.sol";
import {ISynapseOFTAdapterFactory} from "./interfaces/ISynapseOFTAdapterFactory.sol";

import {OFTAdapter} from "@layerzerolabs/oft-evm/contracts/OFTAdapter.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @dev Initializes the LayerZero adapter and ownership from a single constructor parameter snapshot.
abstract contract SynapseOFTAdapterBase is OFTAdapter {
    constructor(ISynapseOFTAdapterFactory.ConstructorParams memory params)
        OFTAdapter(params.token, params.lzEndpoint, params.owner)
        Ownable(params.owner)
    {}
}

/// @notice OFT adapter for SynapseERC20 tokens.
/// @dev Uses burnFrom(address,uint256) and mint(address,uint256) directly on the token.
/// The adapter must have mint permission and senders must approve it for burnFrom.
contract SynapseOFTAdapter is SynapseOFTAdapterBase {
    /// @notice Initializes the adapter using the deploying factory's constructor parameters.
    /// @dev Omitting constructor arguments keeps CREATE2 addresses independent of the token address.
    constructor() SynapseOFTAdapterBase(ISynapseOFTAdapterFactory(msg.sender).getConstructorParams()) {}

    /// @notice Burns tokens from the sender for a cross-chain transfer.
    /// @dev Removes dust and checks slippage before calling burnFrom(from, amountSentLD).
    /// @param from The address whose tokens are burned.
    /// @param amountLD The requested amount in local decimals.
    /// @param minAmountLD The minimum amount to receive in local decimals.
    /// @param dstEid The destination endpoint ID.
    /// @return amountSentLD The amount burned in local decimals after removing dust.
    /// @return amountReceivedLD The amount to receive on the destination in local decimals.
    function _debit(
        address from,
        uint256 amountLD,
        uint256 minAmountLD,
        uint32 dstEid
    )
        internal
        virtual
        override
        returns (uint256 amountSentLD, uint256 amountReceivedLD)
    {
        (amountSentLD, amountReceivedLD) = _debitView(amountLD, minAmountLD, dstEid);
        IBurnableToken(address(innerToken)).burnFrom(from, amountSentLD);
    }

    /// @notice Mints tokens to the recipient of a cross-chain transfer.
    /// @dev Calls mint(to, amountLD), mapping a zero recipient to address(0xdead).
    /// The source endpoint ID is unused.
    /// @param to The recipient address.
    /// @param amountLD The amount to mint in local decimals.
    /// @return amountReceivedLD The amount minted in local decimals.
    function _credit(
        address to,
        uint256 amountLD,
        uint32 /* srcEid */
    )
        internal
        virtual
        override
        returns (uint256 amountReceivedLD)
    {
        // Minting to address(0) is unsupported by ERC20 tokens.
        if (to == address(0)) to = address(0xdead);
        IMintableToken(address(innerToken)).mint(to, amountLD);
        return amountLD;
    }
}
