// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ConfigureAppV1} from "./ConfigureAppV1.s.sol";

contract ConfigureExampleAppV1 is ConfigureAppV1 {
    constructor() ConfigureAppV1("ExampleAppV1") {}
}
