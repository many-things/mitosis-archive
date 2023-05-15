#!/bin/sh

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}
MITO_HOME=${MITO_HOME:-"./test/localnet"}

VALIDATOR_MNEMONIC=${VALIDATOR_MNEMONIC:-"maple often cargo polar eager jaguar eight inflict once nest nice swamp weasel address swift physical valid culture cheese trumpet find dinosaur curve tray"}
SIGNER_MNEMONIC=${SIGNER_MNEMONIC:-"virus oxygen upgrade fitness diagram impact avocado cake cruise glass force joke galaxy project friend icon school midnight front actress squirrel wish phone sock"}
PASSPHRASE=${PASSPHRASE:-"mitomito"}

$DAEMON init localnet --chain-id 'mito-local-1' --staking-bond-denom 'umito'

function add_key() {
  (echo $1; echo $2; echo $2) | $DAEMON keys add $3 --recover --keyring-backend 'file'
}

function get_addr() {
  $DAEMON keys show $1 -a --keyring-backend 'file'
}

sed -i '' 's/stake/umito/g' "$MITO_HOME/config/genesis.json"
sed -i '' 's/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g' "$MITO_HOME/config/config.toml"

tmp=$(mktemp)
MULTISIG_GENESIS=$(cat "./test/multisig-genesis.json" | jq -c)
jq '.app_state.multisig.keygen = '$MULTISIG_GENESIS'' "$MITO_HOME/config/genesis.json" \
  | jq '.app_state.multisig.sign.chain_set = []' \
  > "$tmp" \
  && mv "$tmp" "$MITO_HOME/config/genesis.json"

add_key "$VALIDATOR_MNEMONIC" "$PASSPHRASE" validator
add_key "$SIGNER_MNEMONIC" "$PASSPHRASE" signer

echo "mitomito" | $DAEMON add-genesis-account $(get_addr validator) 2000000000000umito --keyring-backend "file"
echo "mitomito" | $DAEMON add-genesis-account $(get_addr signer) 2000000000000umito --keyring-backend "file"

echo "mitomito" | $DAEMON gentx validator 1000000000000umito --keyring-backend "file" --chain-id 'mito-local-1'

$DAEMON collect-gentxs
