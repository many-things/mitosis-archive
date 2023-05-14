#!/bin/sh

function file() {
  echo "./test/$1"
}

ACCOUNT_NAME="validator"

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet --keyring-backend file"}

ACCOUNT_ADDR=$(echo "mitomito" | $DAEMON keys show $ACCOUNT_NAME -a)

# TODO

rm $(file "temp.json")
rm $(file "temp-tx.json")
