// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {BurningProcessor} from "../../src/processors/BurningProcessor.sol";
import {LockingProcessor} from "../../src/processors/LockingProcessor.sol";
import {IInterchainFactory} from "../../src/interfaces/IInterchainFactory.sol";

contract MockInterchainFactory is IInterchainFactory {
    address internal _initialAdmin;
    address internal _processor;

    address internal _interchainToken;
    address internal _underlyingToken;

    function deployBurningProcessor(
        address interchainToken,
        address underlyingToken
    )
        external
        returns (address deployedProcessor)
    {
        _interchainToken = interchainToken;
        _underlyingToken = underlyingToken;
        deployedProcessor = address(new BurningProcessor());
        delete _interchainToken;
        delete _underlyingToken;
    }

    function deployLockingProcessor(
        address interchainToken,
        address underlyingToken
    )
        external
        returns (address deployedProcessor)
    {
        _interchainToken = interchainToken;
        _underlyingToken = underlyingToken;
        deployedProcessor = address(new LockingProcessor());
        delete _interchainToken;
        delete _underlyingToken;
    }

    function getInterchainTokenDeployParameters()
        external
        view
        override
        returns (address initialAdmin, address processor)
    {
        initialAdmin = _initialAdmin;
        processor = _processor;
    }

    function getProcessorDeployParameters()
        external
        view
        override
        returns (address interchainToken, address underlyingToken)
    {
        interchainToken = _interchainToken;
        underlyingToken = _underlyingToken;
    }
}
