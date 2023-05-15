#!/bin/sh

function file() {
  echo "./test/$1"
}

function payload() {
  echo $(file "payload/$1")
}

tmp_payload=$(mktemp)
tmp_tx=$(mktemp)

SIGNER_NAME=${SIGNER_NAME:-"signer"}
VALIDATOR_NAME=${VALIDATOR_NAME:-"validator"}
DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}

SIGNER_ADDR=$(echo "mitomito" | $DAEMON keys show $SIGNER_NAME -a --keyring-backend file)
SIGNER_INFO=$($DAEMON q account $SIGNER_ADDR --output json | jq -c)
VALIDATOR_ADDR=$(echo "mitomito" | $DAEMON keys show $VALIDATOR_NAME -a --keyring-backend file)

cat $(payload "submit-event-resp.json") \
  | jq '.sender = "'$VALIDATOR_ADDR'"' \
  > $tmp_payload

$DAEMON tx event submit-event $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file broadcast.sh) $VALIDATOR_NAME "submit-event" $tmp_tx
echo "event submitted"
