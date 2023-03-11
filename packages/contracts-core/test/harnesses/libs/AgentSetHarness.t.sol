// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { AgentSet } from "../../../contracts/libs/AgentSet.sol";

contract AgentSetHarness {
    using AgentSet for AgentSet.DomainAddressSet;

    AgentSet.DomainAddressSet internal set;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function add(uint32 domain, address account) external returns (bool) {
        bool value = AgentSet.add(set, domain, account);
        return value;
    }

    function remove(uint32 domain, address account) external returns (bool) {
        bool value = AgentSet.remove(set, domain, account);
        return value;
    }

    function contains(address account) external view returns (bool, uint32) {
        (bool isActive, uint32 domain) = AgentSet.contains(set, account);
        return (isActive, domain);
    }

    function contains(uint32 domain, address account) external view returns (bool) {
        bool value = AgentSet.contains(set, domain, account);
        return value;
    }

    function length(uint32 domain) external view returns (uint256) {
        uint256 value = AgentSet.length(set, domain);
        return value;
    }

    function at(uint32 domain, uint256 index) external view returns (address) {
        address value = AgentSet.at(set, domain, index);
        return value;
    }

    function values(uint32 domain) external view returns (address[] memory) {
        address[] memory value = AgentSet.values(set, domain);
        return value;
    }
}
