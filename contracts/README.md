# astria bridgeable erc20s

Forge project for the `AstriaMintableERC20` contract.

Requirements:

- foundry

Build:

```sh
forge build
```

To test the full end-to-end flow, run the sequencer, cometbft, composer, and conductor. Ensure the configured chain IDs are correct.

Copy the example .env:

```sh
cp local.env.example .env && source .env
```

Deploy `AstriaMintableERC20.sol`:

```sh
forge script script/AstriaMintableERC20.s.sol:AstriaMintableERC20Script \
   --rpc-url $RPC_URL --broadcast --sig "deploy()" -vvvv
```

Take note of the deployed address.

Add the following to the genesis file under `astriaBridgeAddresses`:

```json
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

Stop the geth node and rerun `geth init --genesis genesis.json`. Restart the node. The contract is now initialized as a bridge from the sequencer.

Run the following with the `astria-cli`:

```sh
# this matches the `bridgeAddress` 0x1c0c490f1b5528d8173c5de46d131160e4b2c0c3 in the genesis above
export SEQUENCER_PRIVATE_KEY=2bd806c97f0e00af1a1fc3328fa763a9269723c8db8fac4f93af71db186d6e90
./target/debug/astria-cli sequencer init-bridge-account --sequencer-url=http://localhost:26657 --rollup-name=astria
# the `destination-chain-address` matches the `PRIVATE_KEY` in local.example.env
./target/debug/astria-cli sequencer bridge-lock  --sequencer-url=http://localhost:26657 --amount=1000000 --destination-chain-address=0x46B77EFDFB20979E1C29ec98DcE73e3eCbF64102 --sequencer.chain-id=astria -- 1c0c490f1b5528d8173c5de46d131160e4b2c0c3
```

This initializes the bridge account and also transfer funds over.

Check your ERC20 balance:

```sh
forge script script/AstriaMintableERC20.s.sol:AstriaMintableERC20Script \
   --rpc-url $RPC_URL --sig "getBalance()" -vvvv
```

If everything worked, you should see a balance logged:
```
== Logs ==
  1000000000000000000
```
