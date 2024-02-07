// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════
import {NumberLib} from "../contracts/libs/stack/Number.sol";
import {GasOracle} from "../contracts/GasOracle.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DeployerUtils} from "./utils/DeployerUtils.sol";
import {RawGasData256} from "../test/utils/libs/SynapseStructs.t.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {console, stdJson} from "forge-std/Script.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-console
// solhint-disable ordering
contract SetupGasOracle003Script is DeployerUtils {
    using stdJson for string;
    using Strings for uint256;

    // Need to follow alphabetic order
    struct GasDataJSON {
        uint256 amortAttCost;
        uint256 dataPrice;
        uint256 etherPriceUSDCent;
        uint256 execBuffer;
        uint256 gasPrice;
        uint256 markupPercent;
        uint256 summitTip;
    }

    string public constant GAS_ORACLE = "GasOracle";

    string public gasDataConfig;

    GasOracle public gasOracle;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
    }

    /// @dev Function to exclude script from coverage report
    function testScript() external {}

    /// @notice Main function with the setup logic.
    /// @dev To setup GasOracle on $chainName
    /// forge script script/SetupGasOracle003.s.sol -f chainName --broadcast
    /// @dev To simulate setup on $chainName
    /// forge script script/SetupGasOracle003.s.sol -f chainName
    function run() external {
        setupDevnetIfEnabled();
        startBroadcast(true);
        gasDataConfig = loadGlobalDeployConfig("Messaging003GasData");
        gasOracle = GasOracle(loadDeployment(GAS_ORACLE));
        uint256[] memory domains = gasDataConfig.readUintArray(".domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            setupDomainGasData(domains[i]);
        }
        stopBroadcast();
    }

    function setupDomainGasData(uint256 domain) internal {
        console.log("Checking new gas data for domain: %s", domain);
        GasDataJSON memory gasDataJSON =
            abi.decode(gasDataConfig.parseRaw(string.concat(".gasData.", domain.toString())), (GasDataJSON));
        uint256 ethPriceUSDCent = gasDataConfig.readUint(".ethPriceUSDCent");
        RawGasData256 memory gasData256;
        gasData256.amortAttCost = gasDataJSON.amortAttCost;
        gasData256.dataPrice = gasDataJSON.dataPrice;
        gasData256.etherPrice = (gasDataJSON.etherPriceUSDCent << NumberLib.BWAD_SHIFT) / ethPriceUSDCent;
        gasData256.execBuffer = gasDataJSON.execBuffer;
        gasData256.gasPrice = gasDataJSON.gasPrice;
        gasData256.markup = (gasDataJSON.markupPercent << NumberLib.BWAD_SHIFT) / 100;
        // Round by keeping only the highest 9 bits
        gasData256.round();
        console.log("   amortAttCost: %s", gasData256.amortAttCost);
        console.log("   dataPrice: %s", gasData256.dataPrice);
        console.log("   etherPrice: %s", gasData256.etherPrice);
        console.log("       etherPriceUSDCent: %s", gasDataJSON.etherPriceUSDCent);
        console.log("   execBuffer: %s", gasData256.execBuffer);
        console.log("   gasPrice: %s", gasData256.gasPrice);
        console.log("   markup: %s", gasData256.markup);
        console.log("      markupPercent: %s", gasDataJSON.markupPercent);
        console.log("   summitTip: %s", gasDataJSON.summitTip);
        if (isDataDifferent(domain, gasData256)) {
            console.log("Setting up new gas data for domain: %s", domain);
            gasOracle.setGasData({
                domain: uint32(domain),
                gasPrice: gasData256.gasPrice,
                dataPrice: gasData256.dataPrice,
                execBuffer: gasData256.execBuffer,
                amortAttCost: gasData256.amortAttCost,
                etherPrice: gasData256.etherPrice,
                markup: gasData256.markup
            });
            require(!isDataDifferent(domain, gasData256), "Failed to set gas data");
            console.log("   Gas data successfully set");
        } else {
            console.log("Gas data is up to date for domain: %s", domain);
        }
        console.log();
    }

    function isDataDifferent(uint256 domain, RawGasData256 memory gasData256) internal view returns (bool) {
        (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        ) = gasOracle.getDecodedGasData(uint32(domain));
        return gasData256.gasPrice != gasPrice || gasData256.dataPrice != dataPrice
            || gasData256.execBuffer != execBuffer || gasData256.amortAttCost != amortAttCost
            || gasData256.etherPrice != etherPrice || gasData256.markup != markup;
    }
}
