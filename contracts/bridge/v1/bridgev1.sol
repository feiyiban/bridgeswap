// SPDX-License-Identifier: MIT

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";

import "./transferhelper.sol";
import "./role.sol";

pragma solidity ^0.8.0;

contract BridgeV1 is Initializable, Role, Pausable {
    uint public selfChainId;

    event MAPTransferOut(uint indexed fromChain, uint indexed toChain, string to, uint amount);
    event MAPTransferIn( uint indexed fromChain, uint indexed toChain, uint amount);

    constructor(uint _chainId) {
        selfChainId = _chainId;
    }

    function initialize(uint _chainId) public initializer {
        selfChainId = _chainId;
    }

     function setPause() external onlyManager{
        _pause();
    }

    function setUnpause() external onlyManager{
        _unpause();
    }

    function transferOut(address _token, string memory _receiver, uint _amount, uint _toChainId) external whenNotPaused {
        require(IERC20(_token).balanceOf(msg.sender) >= _amount, "balance too low");

        TransferHelper.safeTransferFrom(_token, msg.sender, address(this), _amount);

        emit MAPTransferOut(selfChainId, _toChainId, _receiver, _amount);
    }

    function transferIn(address _token, address payable _to, uint _amount, uint _fromChain, uint _toChain) external  whenNotPaused {
        TransferHelper.safeTransfer(_token, _to, _amount);
        emit MAPTransferIn(_fromChain, _toChain, _amount);
    }

    function withdraw(address _token, address payable _receiver, uint256 _amount) public onlyManager {
        IERC20(_token).transfer(_receiver, _amount);
    }
}