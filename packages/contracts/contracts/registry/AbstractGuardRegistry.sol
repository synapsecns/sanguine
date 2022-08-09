// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

abstract contract AbstractGuardRegistry {
    function _checkGuardAuth(bytes memory _report)
        internal
        view
        returns (address _guard, bytes29 _data)
    {
        // TODO: check if _report is valid, once guard message standard is finalized
    }

    function _isGuard(address _guard) internal view virtual returns (bool);
}
