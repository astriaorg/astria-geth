solc --optimize --optimize-runs=200 \
    --metadata --metadata-literal \
    --base-path "contracts" \
    --abi "contracts/src/AstriaMintableERC20.sol" \
    -o contracts/abi/ --overwrite

solc --optimize --optimize-runs=200 \
    --base-path "contracts" \
    --bin "contracts/src/AstriaMintableERC20.sol" \
    -o contracts/bin/ --overwrite

abigen --abi contracts/abi/AstriaMintableERC20.abi --bin contracts/bin/AstriaMintableERC20.bin --pkg contracts --type AstriaMintableERC20 --out contracts/astria_mintable_erc20.go
