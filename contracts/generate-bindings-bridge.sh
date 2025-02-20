solc --optimize --optimize-runs=200 \
    --metadata --metadata-literal \
    --base-path "astria-bridge-contracts" \
    --abi "astria-bridge-contracts/src/AstriaBridgeableERC20.sol" \
    -o abi/ --overwrite

solc --optimize --optimize-runs=200 \
    --base-path "astria-bridge-contracts" \
    --bin "astria-bridge-contracts/src/AstriaBridgeableERC20.sol" \
    -o bin/ --overwrite

abigen --abi abi/AstriaBridgeableERC20.abi --bin bin/AstriaBridgeableERC20.bin --pkg contracts --type AstriaBridgeableERC20 --out astria_bridgeable_erc20.go
