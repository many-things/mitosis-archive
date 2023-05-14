#!/bin/sh

function file() {
  echo "./test/$1"
}

ACCOUNT_NAME="validator"

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}

ACCOUNT_ADDR=$(echo "mitomito" | $DAEMON keys show $ACCOUNT_NAME -a --keyring-backend file)
ACCOUNT_INFO=$($DAEMON q account $ACCOUNT_ADDR --output json | jq -c)

cat $(file "register-cosmos-signer.json") \
  | jq '.sender="'$ACCOUNT_ADDR'"' \
  | jq '.chain="osmo-test-5"' \
  | jq '.pub_key="'$(echo "$ACCOUNT_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.account_number="'$(echo "$ACCOUNT_INFO" | jq -r ".account_number")'"' \
  | jq '.sequence="'$(echo "$ACCOUNT_INFO" | jq -r ".sequence")'"' \
  > $(file "temp.json")

$DAEMON tx context register-cosmos-signer $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "register-cosmos-signer"
echo "cosmos signer registered"

cat $(file "register-evm-signer.json") \
  | jq '.sender="'$ACCOUNT_ADDR'"' \
  | jq '.chain="evm-5"' \
  | jq '.pub_key="'$(echo "$ACCOUNT_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.nonce=0' \
  > $(file "temp.json")

$DAEMON tx context register-evm-signer $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "register-evm-signer"
echo "evm signer registered"

rm $(file "temp.json")
rm $(file "temp-tx.json")
