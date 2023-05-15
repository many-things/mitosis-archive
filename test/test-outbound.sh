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

cat $(payload "submit-event-resp.json") \
  | jq '.sender = "'$VALIDATOR_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event submit-event $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $VALIDATOR_NAME "submit-event"
echo "event submitted"

rm $(file "temp.json")
rm $(file "temp-tx.json")
