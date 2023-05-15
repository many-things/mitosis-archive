#!/bin/sh

function file() {
  echo "./test/$1"
}

SIGNER_NAME=${SIGNER_NAME:-"signer"}
VALIDATOR_NAME=${VALIDATOR_NAME:-"validator"}
DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}

SIGNER_ADDR=$(echo "mitomito" | $DAEMON keys show $SIGNER_NAME -a --keyring-backend file)

$DAEMON tx bank send $SIGNER_ADDR $SIGNER_ADDR 1umito --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $SIGNER_NAME "self-send"

SIGNER_INFO=$($DAEMON q account $SIGNER_ADDR --output json | jq -c)
VALIDATOR_ADDR=$(echo "mitomito" | $DAEMON keys show $VALIDATOR_NAME -a --keyring-backend file)

cat $(file "start-keygen.json") \
  | jq '.key_id="osmo-test-5-0"' \
  | jq '.participants=["mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"]' \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx multisig start-keygen $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "start-keygen-osmosis"
echo "keygen started"

cat $(file "submit-pubkey.json") \
  | jq '.key_id="osmo-test-5-0"' \
  | jq '.participant="mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"' \
  | jq '.pub_key="'$(echo "$SIGNER_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx multisig submit-pubkey $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "submit-pubkey-osmosis"
echo "pubkey submitted"

cat $(file "start-keygen.json") \
  | jq '.key_id="evm-5-0"' \
  | jq '.participants=["mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"]' \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx multisig start-keygen $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "start-keygen-evm"
echo "keygen started"

cat $(file "submit-pubkey.json") \
  | jq '.key_id="evm-5-0"' \
  | jq '.participant="mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"' \
  | jq '.pub_key="'$(echo "$SIGNER_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx multisig submit-pubkey $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "submit-pubkey-evm"
echo "pubkey submitted"

rm $(file "temp.json")
rm $(file "temp-tx.json")
