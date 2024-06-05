// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {AstriaMintableERC20} from "../src/AstriaMintableERC20.sol";

contract AstriaMintableERC20Script is Script {
    function setUp() public {}

    function deploy() public {
        string memory name = vm.envString("NAME");
        string memory symbol = vm.envString("SYMBOL");
        address bridge = vm.envAddress("BRIDGE");
        uint32 assetWithdrawalDecimals = uint32(vm.envUint("ASSET_WITHDRAWAL_DECIMALS"));

        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        new AstriaMintableERC20(bridge, assetWithdrawalDecimals, name, symbol);
        vm.stopBroadcast();
    }

    function getBalance() public view {
        AstriaMintableERC20 astriaMintableERC20 = AstriaMintableERC20(vm.envAddress("ASTRIA_MINTABLE_ERC20_ADDRESS"));
        address account = vm.envAddress("USER_ADDRESS");
        uint256 balance = astriaMintableERC20.balanceOf(account);
        console.logUint(balance);
    }
}
