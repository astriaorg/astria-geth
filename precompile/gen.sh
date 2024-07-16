forge build --extra-output-files bin --extra-output-files abi --root .

for dir in ./out/*/
do
    NAME=$(basename $dir)
    NAME=${NAME%.sol}
    NAME_LOWER=$(echo "${NAME:1}" | tr '[:upper:]' '[:lower:]')
    abigen --pkg bindings \
      --abi ./out/$NAME.sol/$NAME.abi.json \
      --bin ./out/$NAME.sol/$NAME.bin \
      --out ./bindings/i_${NAME_LOWER}.abigen.go \
      --type ${NAME:1}
done
