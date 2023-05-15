#!/bin/sh

function file() {
  echo "./test/$1"
}

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet --keyring-backend file"}

out=$(file "temp-tx.json")
resp=$(file "$2.resp.json")

echo "mitomito" | $DAEMON tx sign $out --chain-id 'mito-local-1' --from $1 --output-document $out

$DAEMON tx broadcast $out --log_level 'trace' --output json -b block -y | jq > $resp
