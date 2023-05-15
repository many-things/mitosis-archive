#!/bin/sh

function file() {
  echo "./test/$1"
}

function payload() {
  echo $(file "payload/$1")
}

SIGNER_NAME=${SIGNER_NAME:-"signer"}
VALIDATOR_NAME=${VALIDATOR_NAME:-"validator"}
DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}

SIGNER_ADDR=$(echo "mitomito" | $DAEMON keys show $SIGNER_NAME -a --keyring-backend file)

$DAEMON tx bank send $SIGNER_ADDR $SIGNER_ADDR 1umito --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $SIGNER_NAME "self-send"

SIGNER_INFO=$($DAEMON q account $SIGNER_ADDR --output json | jq -c)
VALIDATOR_ADDR=$(echo "mitomito" | $DAEMON keys show $VALIDATOR_NAME -a --keyring-backend file)

cat $(payload "register-cosmos-signer.json") \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  | jq '.chain="osmo-test-5"' \
  | jq '.pub_key="'$(echo "$SIGNER_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.account_number="'$(echo "$SIGNER_INFO" | jq -r ".account_number")'"' \
  | jq '.sequence="'$(echo "$SIGNER_INFO" | jq -r ".sequence")'"' \
  > $(file "temp.json")

$DAEMON tx context register-cosmos-signer $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "register-cosmos-signer"
echo "cosmos signer registered"

cat $(payload "register-evm-signer.json") \
  | jq '.sender="'$VALIDATOR_ADDR'"' \
  | jq '.chain="evm-5"' \
  | jq '.pub_key="'$(echo "$SIGNER_INFO" | jq -r ".pub_key.key")'"' \
  | jq '.nonce=0' \
  > $(file "temp.json")

$DAEMON tx context register-evm-signer $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "register-evm-signer"
echo "evm signer registered"

cat $(payload "register-proxy.json") \
  | jq '.validator="mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"' \
  | jq '.proxy_account="'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event register-proxy $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "register-proxy"
echo "proxy registered"

cat $(payload "register-chain.json") \
  | jq '.sender = "'$VALIDATOR_ADDR'"' \
  | jq '.chain="osmo-test-5"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "register-chain-osmosis"
echo "osmosis chain registered"

cat $(payload "register-chain.json") \
  | jq '.sender = "'$VALIDATOR_ADDR'"' \
  | jq '.chain="evm-5"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "register-chain-ethereum"
echo "ethereum chain registered"

cat $(payload "submit-event-req.json") \
  | jq '.sender = "'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event submit-event $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "submit-event"
echo "event submitted"

rm $(file "temp.json")
rm $(file "temp-tx.json")
