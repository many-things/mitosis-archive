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

echo "mitomito" \
  | $DAEMON tx event register-proxy $(file "temp.json") \
    --chain-id 'mito-local-1' \
    --from $ACCOUNT_NAME \
    --fees '2000umito' \
    --log_level 'trace' \
    --output json \
    -b 'block' -y \
  | jq \
  > $(file "register-proxy.resp.json")

cat $(file "register-chain.json") \
  | jq '.sender = "'$ACCOUNT_ADDR'"' \
  | jq '.chain="osmosis"' \
  > $(file "temp.json")

echo "mitomito" \
  | $DAEMON tx event register-chain $(file "temp.json") \
    --chain-id 'mito-local-1' \
    --from $ACCOUNT_NAME \
    --fees '2000umito' \
    --log_level 'trace' \
    --output json \
    -b 'block' -y \
  | jq \
  > $(file "register-chain.resp.json")

rm $(file "temp.json")
