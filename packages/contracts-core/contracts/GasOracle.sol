// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MAX_SUMMIT_TIP} from "./libs/Constants.sol";
import {
    IncorrectDestinationDomain, LocalGasDataNotSet, RemoteGasDataNotSet, SummitTipTooHigh
} from "./libs/Errors.sol";
import {GasData, GasDataLib} from "./libs/stack/GasData.sol";
import {Number, NumberLib} from "./libs/stack/Number.sol";
import {Request, RequestLib} from "./libs/stack/Request.sol";
import {Tips, TipsLib} from "./libs/stack/Tips.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "./base/MessagingBase.sol";
import {GasOracleEvents} from "./events/GasOracleEvents.sol";
import {InterfaceDestination} from "./interfaces/InterfaceDestination.sol";
import {InterfaceGasOracle} from "./interfaces/InterfaceGasOracle.sol";

/**
 * @notice `GasOracle` contract is responsible for tracking the gas data for both local and remote chains.
 * ## Local gas data tracking
 * - `GasOracle` is using the available tools such as `tx.gasprice` to track the time-averaged values
 * for different "gas statistics" _(to be implemented in the future)_.
 * - These values are cached, so that the reported values are only changed when a big enough change is detected.
 * - In the MVP version the gas data is set manually by the owner of the contract.
 * - The reported values are included in Origin's State, whenever a new message is sent.
 * > This leads to cached "chain gas data" being included in the Guard and Notary snapshots.
 * ## Remote gas data tracking
 * - To track gas data for the remote chains, GasOracle relies on the Notaries to pass the gas data alongside
 * their attestations.
 * - As the gas data is cached, this leads to a storage write only when the gas data
 * for the remote chain changes significantly.
 * - GasOracle is in charge of enforcing the optimistic periods for the gas data it gets from `Destination`.
 * - The optimistic period is smaller when the "gas statistics" are increasing, and bigger when they are decreasing.
 * > Reason for that is that the decrease of the gas price leads to lower execution/delivery tips, and we want the
 * > Executors to be protected against that.
 */
contract GasOracle is MessagingBase, GasOracleEvents, InterfaceGasOracle {
    // ══════════════════════════════════════════ IMMUTABLES & CONSTANTS ═══════════════════════════════════════════════

    address public immutable destination;

    // TODO: come up with refined values for the optimistic periods
    uint256 public constant GAS_DATA_INCREASED_OPTIMISTIC_PERIOD = 5 minutes;
    uint256 public constant GAS_DATA_DECREASED_OPTIMISTIC_PERIOD = 1 hours;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    mapping(uint32 => GasData) internal _gasData;

    // Fixed value for the summit tip, denominated in Ethereum Mainnet Wei.
    uint256 public summitTipWei;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_, address destination_) MessagingBase("0.0.3", synapseDomain_) {
        destination = destination_;
    }

    /// @notice Initializes GasOracle contract:
    /// - msg.sender is set as contract owner
    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable2Step_init();
    }

    /// @notice MVP function to set the gas data for the given domain.
    function setGasData(
        uint32 domain,
        uint256 gasPrice,
        uint256 dataPrice,
        uint256 execBuffer,
        uint256 amortAttCost,
        uint256 etherPrice,
        uint256 markup
    ) external onlyOwner {
        GasData updatedGasData = GasDataLib.encodeGasData({
            gasPrice_: NumberLib.compress(gasPrice),
            dataPrice_: NumberLib.compress(dataPrice),
            execBuffer_: NumberLib.compress(execBuffer),
            amortAttCost_: NumberLib.compress(amortAttCost),
            etherPrice_: NumberLib.compress(etherPrice),
            markup_: NumberLib.compress(markup)
        });
        if (GasData.unwrap(updatedGasData) != GasData.unwrap(_gasData[domain])) {
            _setGasData(domain, updatedGasData);
        }
    }

    /// @notice MVP function to set the summit tip.
    function setSummitTip(uint256 summitTipWei_) external onlyOwner {
        if (summitTipWei_ > MAX_SUMMIT_TIP) revert SummitTipTooHigh();
        summitTipWei = summitTipWei_;
        emit SummitTipUpdated(summitTipWei_);
    }

    /// @inheritdoc InterfaceGasOracle
    function updateGasData(uint32 domain) external {
        (bool wasUpdated, GasData updatedGasData) = _fetchGasData(domain);
        if (wasUpdated) {
            _setGasData(domain, updatedGasData);
        }
    }

    /// @inheritdoc InterfaceGasOracle
    function getDecodedGasData(uint32 domain)
        external
        view
        returns (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        )
    {
        GasData gasData = _gasData[domain];
        gasPrice = NumberLib.decompress(gasData.gasPrice());
        dataPrice = NumberLib.decompress(gasData.dataPrice());
        execBuffer = NumberLib.decompress(gasData.execBuffer());
        amortAttCost = NumberLib.decompress(gasData.amortAttCost());
        etherPrice = NumberLib.decompress(gasData.etherPrice());
        markup = NumberLib.decompress(gasData.markup());
    }

    /// @inheritdoc InterfaceGasOracle
    function getGasData() external view returns (uint256 paddedGasData) {
        return GasData.unwrap(_gasData[localDomain]);
    }

    /// @inheritdoc InterfaceGasOracle
    function getMinimumTips(uint32 destination_, uint256 paddedRequest, uint256 contentLength)
        external
        view
        returns (uint256 paddedTips)
    {
        if (destination_ == localDomain) revert IncorrectDestinationDomain();
        GasData localGasData = _gasData[localDomain];
        uint256 localEtherPrice = localGasData.etherPrice().decompress();
        if (localEtherPrice == 0) revert LocalGasDataNotSet();
        GasData remoteGasData = _gasData[destination_];
        uint256 remoteEtherPrice = remoteGasData.etherPrice().decompress();
        if (remoteEtherPrice == 0) revert RemoteGasDataNotSet();
        Request request = RequestLib.wrapPadded(paddedRequest);
        // TODO: figure out unchecked math
        // We store the fixed value of the summit tip in Ethereum Mainnet Wei already.
        // To convert it to local Ether, we need to divide by the local Ether price (using BWAD math).
        uint256 summitTip = (summitTipWei << NumberLib.BWAD_SHIFT) / localEtherPrice;
        // To convert the cost from remote Ether to local Ether, we need to multiply by the ratio of the Ether prices.
        uint256 attestationTip = remoteGasData.amortAttCost().decompress() * remoteEtherPrice / localEtherPrice;
        // Total cost for Executor to execute a message on the remote chain has three components:
        // - Execution: gas price * requested gas limit
        // - Calldata: data price * content length
        // - Buffer: additional fee to account for computations before and after the actual execution
        // Same logic for converting the cost from remote Ether to local Ether applies here.
        // forgefmt: disable-next-item
        uint256 executionTip = (
            remoteGasData.gasPrice().decompress() * request.gasLimit() +
            remoteGasData.dataPrice().decompress() * contentLength +
            remoteGasData.execBuffer().decompress()
        ) * remoteEtherPrice / localEtherPrice;
        // Markup for executionTip is assigned to the Delivery tip. Markup is denominated in BWAD units.
        // Execution tip is already denominated in local Ether units.
        uint256 deliveryTip = (executionTip * remoteGasData.markup().decompress()) >> NumberLib.BWAD_SHIFT;
        // The price of the gas airdrop is also included in the Delivery tip.
        // TODO: enable when gasDrop is implemented
        // deliveryTip += request.gasDrop() * remoteEtherPrice / localEtherPrice;
        // Use calculated values to encode the tips.
        return Tips.unwrap(
            TipsLib.encodeTips256({
                summitTip_: summitTip,
                attestationTip_: attestationTip,
                executionTip_: executionTip,
                deliveryTip_: deliveryTip
            })
        );
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Sets the gas data for the given domain, and emits a corresponding event.
    function _setGasData(uint32 domain, GasData updatedGasData) internal {
        _gasData[domain] = updatedGasData;
        emit GasDataUpdated(domain, GasData.unwrap(updatedGasData));
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the updated gas data for the given domain by
    /// optimistically consuming the data from the `Destination` contract.
    function _fetchGasData(uint32 domain) internal view returns (bool wasUpdated, GasData updatedGasData) {
        GasData current = _gasData[domain];
        // Destination only has the gas data for the remote domains.
        if (domain == localDomain) return (false, current);
        (GasData incoming, uint256 dataMaturity) = InterfaceDestination(destination).getGasData(domain);
        // Zero maturity means that either there is no data for the domain, or it was just updated.
        // In both cases, we don't want to update the local data.
        if (dataMaturity == 0) return (false, current);
        // Update each gas parameter separately.
        updatedGasData = GasDataLib.encodeGasData({
            gasPrice_: _updateGasParameter(current.gasPrice(), incoming.gasPrice(), dataMaturity),
            dataPrice_: _updateGasParameter(current.dataPrice(), incoming.dataPrice(), dataMaturity),
            execBuffer_: _updateGasParameter(current.execBuffer(), incoming.execBuffer(), dataMaturity),
            amortAttCost_: _updateGasParameter(current.amortAttCost(), incoming.amortAttCost(), dataMaturity),
            etherPrice_: _updateGasParameter(current.etherPrice(), incoming.etherPrice(), dataMaturity),
            markup_: _updateGasParameter(current.markup(), incoming.markup(), dataMaturity)
        });
        wasUpdated = GasData.unwrap(updatedGasData) != GasData.unwrap(current);
    }

    /// @dev Returns the updated value for the gas parameter, given the maturity of the incoming data.
    function _updateGasParameter(Number current, Number incoming, uint256 dataMaturity)
        internal
        pure
        returns (Number updatedParameter)
    {
        // We apply the incoming value only if its optimistic period has passed.
        // The optimistic period is smaller when the the value is increasing, and bigger when it is decreasing.
        if (incoming.decompress() > current.decompress()) {
            return dataMaturity < GAS_DATA_INCREASED_OPTIMISTIC_PERIOD ? current : incoming;
        } else {
            return dataMaturity < GAS_DATA_DECREASED_OPTIMISTIC_PERIOD ? current : incoming;
        }
    }
}
