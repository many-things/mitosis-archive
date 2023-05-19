#!/bin/sh

function file() {
  echo "./test/$1"
}

function response() {
  echo $(file "responses/$1")
}

DAEMON=${DAEMON:-"./build/mitosisd --node http://mitosis.morethings.xyz:26657 --chain-id mitosis"}

resp=$(response "$2.resp.json")

echo "mitomito" | $DAEMON tx sign $3 --chain-id 'mitosis' --from $1 --output-document $3

$DAEMON tx broadcast $3 --log_level 'trace' --output json -b block -y | jq > $resp
