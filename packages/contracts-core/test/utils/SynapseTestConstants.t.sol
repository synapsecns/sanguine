// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract SynapseTestConstants {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               DOMAINS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint32 internal constant DOMAIN_LOCAL = 1000;
    uint32 internal constant DOMAIN_REMOTE = 1500;
    // TODO: replace placeholder value
    uint32 internal constant DOMAIN_SYNAPSE = 4269;

    uint256 internal constant DOMAIN_AGENTS = 4;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   MASKS FOR TEST SUITE DEPLOYMENTS                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256 internal constant DEPLOY_MASK_DESTINATION = 0xF;
    // Default option for deploying Destination
    uint256 internal constant DEPLOY_MOCK_DESTINATION = 0x0;
    uint256 internal constant DEPLOY_PROD_DESTINATION = 0x1;

    uint256 internal constant DEPLOY_MASK_ORIGIN = 0xF0;
    // Default option for deploying Origin
    uint256 internal constant DEPLOY_MOCK_ORIGIN = 0x00;
    uint256 internal constant DEPLOY_PROD_ORIGIN = 0x10;

    uint256 internal constant DEPLOY_MASK_SUMMIT = 0xF00;
    // Default option for deploying Summit
    uint256 internal constant DEPLOY_MOCK_SUMMIT = 0x000;
    uint256 internal constant DEPLOY_PROD_SUMMIT = 0x100;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            MESSAGE TESTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint32 internal constant MESSAGES = 10;
    uint256 internal constant BLOCK_TIME = 12 seconds;
}
