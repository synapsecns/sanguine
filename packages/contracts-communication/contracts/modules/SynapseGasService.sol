pragma solidity 0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

abstract contract SynapseGasService is Ownable {
    address payable public executor;
    address public gasOracle;

    function getModuleFee(uint256 dstChainId) public pure returns (uint256) {
        // Get Latest Posted Destination Gas Price from oracle
        // Requires: Access to origin USD Gas Price / Destination USD Gas PRice
        // Get current price of origin chain assets
        // Get current price of destination chain assets
        // Calculate the estiamted fee based on preset gas limit
        // return

        // TODO: Right now, we don't have all of the info needed to provide a real fee estimation - we will provide 1 wei to enable other functionality to be built.
        return 1;
    }

    function setExecutor(address _executor) public onlyOwner {
        executor = payable(_executor);
    }

    function _payFeesForExecution(uint256 feeAmount) internal {
        // Transfer fee to executor
        executor.transfer(feeAmount);
    }
}
