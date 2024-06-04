// SPDX-License-Identifier: MIT or Apache-2.0
pragma solidity ^0.8.21;

import {IAstriaMintableERC20} from "./IAstriaMintableERC20.sol";
import {ERC20} from "lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract AstriaMintableERC20 is IAstriaMintableERC20, ERC20 {
    // the `astriaBridgeSenderAddress` built into the astria-geth node
    address public immutable BRIDGE;

    // the contract address of the `AstriaWithdrawerERC20` contract which 
    // is authorized to burn tokens
    address public immutable WITHDRAWER;

    // the 32-byte asset ID of the token represented on the sequencer chain
    uint256 public immutable SEQUENCER_ASSET_ID;

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
        uint256 _sequencerAssetId,
        string memory _name,
        string memory _symbol
    ) ERC20(_name, _symbol) {
        BRIDGE = _bridge;
        WITHDRAWER = _withdrawer;
        SEQUENCER_ASSET_ID = _sequencerAssetId;
    }

    function getSequencerAssetId() external view returns (uint256) {
        return SEQUENCER_ASSET_ID;
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
