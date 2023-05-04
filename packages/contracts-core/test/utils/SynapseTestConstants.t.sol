// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {TIPS_LENGTH, REQUEST_LENGTH} from "../../contracts/libs/Constants.sol";

abstract contract SynapseTestConstants {
    string internal constant LATEST_VERSION = "0.0.3";

    // ══════════════════════════════════════════════════ DOMAINS ══════════════════════════════════════════════════════

    uint32 internal constant DOMAIN_LOCAL = 1000;
    uint32 internal constant DOMAIN_REMOTE = 1500;
    uint32 internal constant DOMAIN_OTHER = 2000;
    // TODO: replace placeholder value
    uint32 internal constant DOMAIN_SYNAPSE = 10;

    uint256 internal constant DOMAIN_AGENTS = 4;

    // ═════════════════════════════════════ MASKS FOR TEST SUITE DEPLOYMENTS ══════════════════════════════════════════

    uint256 internal constant DEPLOY_MASK_DESTINATION = 0x0F;
    // Default option for deploying Destination
    uint256 internal constant DEPLOY_MOCK_DESTINATION = 0x00;
    uint256 internal constant DEPLOY_PROD_DESTINATION = 0x01;

    uint256 internal constant DEPLOY_MASK_ORIGIN = 0xF0;
    // Default option for deploying Origin
    uint256 internal constant DEPLOY_MOCK_ORIGIN = 0x00;
    uint256 internal constant DEPLOY_PROD_ORIGIN = 0x10;

    uint256 internal constant DEPLOY_MASK_DESTINATION_SYNAPSE = 0x0F_00;
    // Default option for deploying Synapse Chain Destination
    uint256 internal constant DEPLOY_MOCK_DESTINATION_SYNAPSE = 0x00_00;
    uint256 internal constant DEPLOY_PROD_DESTINATION_SYNAPSE = 0x01_00;

    uint256 internal constant DEPLOY_MASK_ORIGIN_SYNAPSE = 0xF0_00;
    // Default option for deploying Synapse Chain Origin
    uint256 internal constant DEPLOY_MOCK_ORIGIN_SYNAPSE = 0x00_00;
    uint256 internal constant DEPLOY_PROD_ORIGIN_SYNAPSE = 0x10_00;

    uint256 internal constant DEPLOY_MASK_SUMMIT = 0x0F_00_00;
    // Default option for deploying Summit (Synapse Chain)
    uint256 internal constant DEPLOY_MOCK_SUMMIT = 0x00_00_00;
    uint256 internal constant DEPLOY_PROD_SUMMIT = 0x01_00_00;

    uint256 internal constant DEPLOY_MASK_GAS_ORACLE = 0xF0_00_00;
    // Default option for deploying Gas Oracle
    uint256 internal constant DEPLOY_MOCK_GAS_ORACLE = 0x00_00_00;
    uint256 internal constant DEPLOY_PROD_GAS_ORACLE = 0x10_00_00;

    uint256 internal constant DEPLOY_MASK_GAS_ORACLE_SYNAPSE = 0x0F_00_00_00;
    // Default option for deploying Gas Oracle (Synapse Chain)
    uint256 internal constant DEPLOY_MOCK_GAS_ORACLE_SYNAPSE = 0x00_00_00_00;
    uint256 internal constant DEPLOY_PROD_GAS_ORACLE_SYNAPSE = 0x01_00_00_00;

    // ══════════════════════════════════════════════ ENCODING TESTS ═══════════════════════════════════════════════════

    uint256 internal constant MIN_BASE_MESSAGE_LENGTH = 32 + 32 + TIPS_LENGTH + REQUEST_LENGTH;

    // ═══════════════════════════════════════════════ MESSAGE TESTS ═══════════════════════════════════════════════════

    uint256 internal constant MAX_CONTENT_BYTES = 2 * 2 ** 10;
    uint256 internal constant MAX_SYSTEM_CALL_WORDS = MAX_CONTENT_BYTES / 32;

    uint32 internal constant MESSAGES = 10;
    uint256 internal constant BLOCK_TIME = 12 seconds;

    uint32 internal constant AGENT_ROOT_OPTIMISTIC_PERIOD = 1 days;
    uint32 internal constant BONDING_OPTIMISTIC_PERIOD = 1 days;
}
