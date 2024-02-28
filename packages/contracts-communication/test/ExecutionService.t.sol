// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionService, ExecutionServiceEvents, IExecutionService} from "../contracts/ExecutionService.sol";
import {Test} from "forge-std/Test.sol";
import {GasOracleMock} from "./mocks/GasOracleMock.sol";

contract ExecutionServiceTest is ExecutionServiceEvents, Test {
    ExecutionService public executionService;
    GasOracleMock public gasOracle;

    address icClient = address(0x123);
    address executorEOA = address(0x456);

    function setUp() public {
        gasOracle = new GasOracleMock();
        executionService = new ExecutionService();
        executionService.setInterchainClient(icClient);
        executionService.setExecutorEOA(executorEOA);
        executionService.setGasOracle(address(gasOracle));
    }
}
