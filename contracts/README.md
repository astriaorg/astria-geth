# astria bridgeable erc20s

Forge project for the `AstriaMintableERC20` contract.

Requirements:

- foundry

Build:

```sh
forge build
```

Copy the example .env:

`cp local.env.example .env && source .env`

Deploy `AstriaMintableERC20.sol`:

```sh
forge script script/AstriaMintableERC20.s.sol:AstriaMintableERC20Script \
   --rpc-url $RPC_URL --broadcast --sig "deploy()" -vvvv
```

Take note of the deployed address.

Add the following to the genesis file under `astriaBridgeAddresses`:

```
"astriaBridgeAddresses": [
    {
        "bridgeAddress": "0x1c0c490f1b5528d8173c5de46d131160e4b2c0c3",
        "startHeight": 1,
        "assetDenom": "nria",
        "assetPrecision": 6,
        "erc20asset": {
            "contractAddress":"0x9Aae647A1CB2ec6b39afd552aD149F6A26Bb2aD6",
            "contractPrecision": 18
        }
    }
],
```

Note: this mints `nria` as an erc20 instead of the native asset.

`bridgeAddress` is the bridge address that corresponds to this asset on the sequencer chain.
`assetDenom` does not need to match the name of the token in the deployed contract, but it does need to match the denom of the token on the sequencer.
`contractAddress` in `erc20asset` is the address of the contract deployed above.

Stop the geth node and rerun `geth init --genesis genesis.json`. Restart the node.


