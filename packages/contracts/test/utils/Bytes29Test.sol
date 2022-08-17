// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

abstract contract Bytes29Test is Test {
    using TypedMemView for bytes29;

    uint40 internal constant WRONG_TYPE = 0xFF_00_FF_00_FF;

    function _createTestView() internal virtual returns (bytes29 _view);

    function _prepareMistypedTest(uint40 _correctType) internal returns (bytes29 _view) {
        assert(_correctType != WRONG_TYPE);
        bytes29 _correctView = _createTestView();
        // Change type of memory view w/o changing the data
        _view = _correctView.castTo(WRONG_TYPE);
        // Expect "wrong type" revert
        vm.expectRevert(_expectedRevertMessage(_correctType));
    }

    function _modifyTestView(uint256 _from, bytes memory _newData) internal returns (bytes memory) {
        return _modifyView(_createTestView(), _from, _newData);
    }

    function _modifyTestView(
        uint256 _from,
        uint256 _len,
        bytes memory _newData
    ) internal returns (bytes memory) {
        return _modifyView(_createTestView(), _from, _len, _newData);
    }

    /// @dev replace [_from, _from + _newData.length) with _newData
    function _modifyView(
        bytes29 _view,
        uint256 _from,
        bytes memory _newData
    ) internal view returns (bytes memory newPayload) {
        return _modifyView(_view, _from, _newData.length, _newData);
    }

    /// @dev replace [_from, _from + len) with _newData
    function _modifyView(
        bytes29 _view,
        uint256 _from,
        uint256 _len,
        bytes memory _newData
    ) internal view returns (bytes memory newPayload) {
        assert(_from + _len <= _view.len());
        newPayload = abi.encodePacked(
            _view.prefix(_from, 0).clone(),
            _newData,
            _view.sliceFrom(_from + _len, 0).clone()
        );
        assert(newPayload.length == _view.len() + _newData.length - _len);
    }

    function _expectedRevertMessage(uint40 _correctType) internal pure returns (bytes memory) {
        (, uint256 g) = TypedMemView.encodeHex(WRONG_TYPE);
        (, uint256 e) = TypedMemView.encodeHex(_correctType);
        return
            abi.encodePacked(
                "Type assertion failed. Got 0x",
                uint80(g),
                ". Expected 0x",
                uint80(e)
            );
    }
}
