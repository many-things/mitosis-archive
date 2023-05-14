#!/bin/sh

function file() {
  echo "./test/$1"
}

ACCOUNT_NAME="validator"

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet --keyring-backend file"}

ACCOUNT_ADDR=$(echo "mitomito" | $DAEMON keys show $ACCOUNT_NAME -a)

cat $(file "register-proxy.json") \
  | jq '.validator="mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"' \
  | jq '.proxy_account="'$ACCOUNT_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event register-proxy $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "register-proxy"
echo "proxy registered"

cat $(file "register-chain.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  | jq '.chain="osmosis"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "register-chain-osmosis"
echo "osmosis chain registered"

cat $(file "register-chain.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  | jq '.chain="ethereum"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "register-chain-ethereum"
echo "ethereum chain registered"

cat $(file "submit-event.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event submit-event $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
$(file broadcast.sh) $ACCOUNT_NAME "submit-event"
echo "event submitted"

rm $(file "temp.json")
rm $(file "temp-tx.json")
