#!/bin/sh

function file() {
  echo "./test/$1"
}

function payload() {
  echo $(file "payload/$1")
}

tmp_payload=$(mktemp)
tmp_tx=$(mktemp)

DAEMON=${DAEMON:-"./build/mitosisd --node https://mitosis.morethings.xyz:26657 --chain-id mitosis"}
KEYRING_DAEMON=${KEYRING_DAEMON:-"./build/mitosisd"}

PASSPHRASE=${PASSPHRASE:-"mitomito"}
MITO1_NAME=${MITO1_NAME:-"mito1"}
MITO2_NAME=${MITO2_NAME:-"mito2"}
MITO3_NAME=${MITO3_NAME:-"mito3"}

MITO1_VALI=${MITO1_VALI:-"mitovaloper1gr29n72snrcuyq8wyln6kunpe66zt5uccknn9x"}
MITO2_VALI=${MITO2_VALI:-"mitovaloper1rkcgv25m6mzmj985e6j5mkhkl53533nrp57v3w"}
MITO3_VALI=${MITO3_VALI:-"mitovaloper1hfawqqvj5fqn0gpputhqy549uwxmvn0kzfh35l"}

MITO1_ADDR=$(echo "mitomito" | $KEYRING_DAEMON keys show $MITO1_NAME -a --keyring-backend file)
MITO2_ADDR=$(echo "mitomito" | $KEYRING_DAEMON keys show $MITO2_NAME -a --keyring-backend file)
MITO3_ADDR=$(echo "mitomito" | $KEYRING_DAEMON keys show $MITO3_NAME -a --keyring-backend file)


# ------- register proxy ---------

cat $(payload "register-proxy.json") \
  | jq '.validator="'$MITO1_VALI'"' \
  | jq '.proxy_account="'$MITO1_ADDR'"' \
  > $tmp_payload

$DAEMON tx event register-proxy $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file prod-broadcast.sh) $MITO1_NAME "register-proxy" $tmp_tx
echo "proxy registered: mito1"

cat $(payload "register-proxy.json") \
  | jq '.validator="'$MITO2_VALI'"' \
  | jq '.proxy_account="'$MITO2_ADDR'"' \
  > $tmp_payload

$DAEMON tx event register-proxy $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file prod-broadcast.sh) $MITO2_NAME "register-proxy" $tmp_tx
echo "proxy registered: mito2"

cat $(payload "register-proxy.json") \
  | jq '.validator="'$MITO3_VALI'"' \
  | jq '.proxy_account="'$MITO3_ADDR'"' \
  > $tmp_payload

# ------- register chain ---------

# $DAEMON tx event register-proxy $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
# $(file prod-broadcast.sh) $MITO3_NAME "register-proxy" $tmp_tx
# echo "proxy registered: mito3"


# cat $(payload "register-chain.json") \
#   | jq '.sender = "'$MITO1_ADDR'"' \
#   | jq '.chain="osmo-test-5"' \
#   > $tmp_payload

# $DAEMON tx event register-chain $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
# $(file prod-broadcast.sh) $MITO1_NAME "register-chain-osmosis" $tmp_tx
# echo "osmosis chain registered"

# cat $(payload "register-chain.json") \
#   | jq '.sender = "'$MITO1_ADDR'"' \
#   | jq '.chain="evm-5"' \
#   > $tmp_payload

# $DAEMON tx event register-chain $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
# $(file prod-broadcast.sh) $MITO1_NAME "register-chain-ethereum" $tmp_tx
# echo "ethereum chain registered"

# cat $(payload "submit-event-req.json") \
#   | jq '.sender = "'$MITO1_ADDR'"' \
#   > $tmp_payload


# ------- send ---------

cat $(payload "submit-event-req.json") \
  | jq '.sender = "'$MITO1_ADDR'"' \
  > $tmp_payload

$DAEMON tx event submit-event $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file prod-broadcast.sh) $MITO1_NAME "submit-event" $tmp_tx
echo "event submitted: mito1"

cat $(payload "submit-event-req.json") \
  | jq '.sender = "'$MITO2_ADDR'"' \
  > $tmp_payload

$DAEMON tx event submit-event $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file prod-broadcast.sh) $MITO2_NAME "submit-event" $tmp_tx
echo "event submitted: mito2"

cat $(payload "submit-event-req.json") \
  | jq '.sender = "'$MITO3_ADDR'"' \
  > $tmp_payload

$DAEMON tx event submit-event $tmp_payload --fees 2000umito --generate-only | jq > $tmp_tx
$(file prod-broadcast.sh) $MITO3_NAME "submit-event" $tmp_tx
echo "event submitted: mito3"
