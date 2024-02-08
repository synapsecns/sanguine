// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {IInterchainFactory} from "../../src/interfaces/IInterchainFactory.sol";

import {BurningProcessor} from "../../src/processors/BurningProcessor.sol";
import {LockingProcessor} from "../../src/processors/LockingProcessor.sol";
import {InterchainERC20} from "../../src/tokens/InterchainERC20.sol";

import {InterchainERC20Harness} from "../harnesses/InterchainERC20Harness.sol";

contract MockInterchainFactory is IInterchainFactory {
    address internal _firstArg;
    address internal _secondArg;

    modifier withArgs(address firstArg, address secondArg) {
        _firstArg = firstArg;
        _secondArg = secondArg;
        _;
        delete _firstArg;
        delete _secondArg;
    }

    function deployInterchainToken(
        string memory name,
        string memory symbol,
        uint8 decimals,
        address initialAdmin,
        address processor
    )
        external
        withArgs(initialAdmin, processor)
        returns (address deployedToken)
    {
        deployedToken = address(new InterchainERC20(name, symbol, decimals));
    }

    function deployInterchainTokenHarness(
        string memory name,
        string memory symbol,
        uint8 decimals,
        address initialAdmin,
        address processor
    )
        external
        withArgs(initialAdmin, processor)
        returns (address deployedToken)
    {
        deployedToken = address(new InterchainERC20Harness(name, symbol, decimals));
    }

    function deployBurningProcessor(
        address interchainToken,
        address underlyingToken
    )
        external
        withArgs(interchainToken, underlyingToken)
        returns (address deployedProcessor)
    {
        deployedProcessor = address(new BurningProcessor());
    }

    function deployLockingProcessor(
        address interchainToken,
        address underlyingToken
    )
        external
        withArgs(interchainToken, underlyingToken)
        returns (address deployedProcessor)
    {
        deployedProcessor = address(new LockingProcessor());
    }

    function getInterchainTokenDeployParameters()
        external
        view
        override
        returns (address initialAdmin, address processor)
    {
        initialAdmin = _firstArg;
        processor = _secondArg;
    }

    function getProcessorDeployParameters()
        external
        view
        override
        returns (address interchainToken, address underlyingToken)
    {
        interchainToken = _firstArg;
        underlyingToken = _secondArg;
    }
}
