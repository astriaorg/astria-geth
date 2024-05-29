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

        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        AstriaMintableERC20 astriaMintableERC20 = new AstriaMintableERC20(bridge, name, symbol);
        console.logAddress(address(astriaMintableERC20));

        vm.stopBroadcast();
    }
}
