// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionService, ExecutionServiceEvents, IExecutionService} from "../contracts/ExecutionService.sol";
import {Test} from "forge-std/Test.sol";
import {SynapseGasOracleMock} from "./mocks/SynapseGasOracleMock.sol";

contract ExecutionServiceTest is ExecutionServiceEvents, Test {
    ExecutionService public executionService;
    SynapseGasOracleMock public gasOracle;

    address icClient = address(0x123);
    address executorEOA = address(0x456);
    address owner = makeAddr("Owner");

    function setUp() public {
        gasOracle = new SynapseGasOracleMock();
        executionService = new ExecutionService(address(this));
        executionService.setInterchainClient(icClient);
        executionService.setExecutorEOA(executorEOA);
        executionService.setGasOracle(address(gasOracle));
        executionService.transferOwnership(owner);
    }
}
