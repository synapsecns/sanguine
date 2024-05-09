// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1DstTest, OptionsV1} from "./InterchainClientV1.Dst.t.sol";

contract InterchainClientV1DstNoOptPeriodWithGasDropTest is InterchainClientV1DstTest {
    function getOptions() internal pure override returns (OptionsV1 memory) {
        return OptionsV1({gasLimit: 100_000, gasAirdrop: 0.1 ether});
    }

    function getOptimisticPeriod() internal pure override returns (uint256) {
        return 0;
    }
}
