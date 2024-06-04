// SPDX-License-Identifier: MIT or Apache-2.0
pragma solidity ^0.8.21;

import {IAstriaMintableERC20} from "./IAstriaMintableERC20.sol";
import {ERC20} from "lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract AstriaMintableERC20 is IAstriaMintableERC20, ERC20 {
    address public immutable BRIDGE;

    address public immutable WITHDRAWER;

    event Mint(address indexed account, uint256 amount);

    event Burn(address indexed account, uint256 amount);

    modifier onlyBridge() {
        require(msg.sender == BRIDGE, "AstriaMintableERC20: only bridge can mint");
        _;
    }

    modifier onlyWithdrawer() {
        require(msg.sender == BRIDGE, "AstriaMintableERC20: only withdrawer can burn");
        _;
    }

    constructor(
        address _bridge,
        address _withdrawer,
        string memory _name,
        string memory _symbol
    ) ERC20(_name, _symbol) {
        BRIDGE = _bridge;
        WITHDRAWER = _withdrawer;
    }

    function mint(address _to, uint256 _amount)
        external
        virtual
        onlyBridge
    {
        _mint(_to, _amount);
        emit Mint(_to, _amount);
    }

    function burn(address _from, uint256 _amount)
        external
        virtual
        onlyWithdrawer
    {
        _burn(_from, _amount);
        emit Burn(_from, _amount);
    }
}
