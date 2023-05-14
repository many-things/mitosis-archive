#!/bin/sh

function file() {
  echo "./test/$1"
}

ACCOUNT_NAME="validator"

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet --keyring-backend file"}

ACCOUNT_ADDR=$(echo "mitomito" | $DAEMON keys show $ACCOUNT_NAME -a)

cat $(file "submit-event.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  > $(file "temp.json")

echo "mitomito" \
  | $DAEMON tx event submit-event $(file "temp.json") \
    --chain-id 'mito-local-1' \
    --from $ACCOUNT_NAME \
    --fees '2000umito' \
    --log_level 'trace' \
    --output json \
    -b 'block' -y \
  | jq \
  > $(file "submit-event.resp.json")

cat $(file "submit-event.resp.json")

rm $(file "temp.json")
