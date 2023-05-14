#!/bin/sh

function file() {
  echo "./test/$1"
}

function broadcast() {
  out=$(file "temp-tx.json")
  resp=$(file "$2.resp.json")

  echo "mitomito" | $DAEMON tx sign $out --chain-id 'mito-local-1' --from $1 --output-document $out

  $DAEMON tx broadcast $out --log_level 'trace' --output json -b block -y | jq > $resp
}

ACCOUNT_NAME="validator"

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet --keyring-backend file"}

ACCOUNT_ADDR=$(echo "mitomito" | $DAEMON keys show $ACCOUNT_NAME -a)

cat $(file "register-proxy.json") \
  | jq '.validator="mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn"' \
  | jq '.proxy_account="'$ACCOUNT_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event register-proxy $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
broadcast $ACCOUNT_NAME "register-proxy"

cat $(file "register-chain.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  | jq '.chain="osmosis"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
broadcast $ACCOUNT_NAME "register-chain-osmosis"

cat $(file "register-chain.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  | jq '.chain="ethereum"' \
  > $(file "temp.json")

$DAEMON tx event register-chain $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
broadcast $ACCOUNT_NAME "register-chain-ethereum"

cat $(file "submit-event.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  > $(file "temp.json")

$DAEMON tx event submit-event $(file "temp.json") --fees 2000umito --generate-only | jq > $(file "temp-tx.json")
broadcast $ACCOUNT_NAME "submit-event"

rm $(file "temp.json")
rm $(file "temp-tx.json")
