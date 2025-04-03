solc --optimize --optimize-runs=200 \
    --metadata --metadata-literal \
    --base-path "astria-oracle-contracts" \
    --abi "astria-oracle-contracts/src/AstriaOracle.sol" \
    -o abi/ --overwrite

solc --optimize --optimize-runs=200 \
    --base-path "astria-oracle-contracts" \
    --bin "astria-oracle-contracts/src/AstriaOracle.sol" \
    -o bin/ --overwrite

abigen --abi abi/AstriaOracle.abi --bin bin/AstriaOracle.bin --pkg contracts --type AstriaOracle --out astria_oracle.go
