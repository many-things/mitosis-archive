#!/bin/sh

DAEMON=${DAEMON:-"./build/mitosisd --home ./test/localnet"}

$DAEMON init localnet --chain-id 'mito-local-1' --staking-bond-denom 'umito'

sed -i '' 's/stake/umito/g' ./test/localnet/config/genesis.json

(
  echo "maple often cargo polar eager jaguar eight inflict once nest nice swamp weasel address swift physical valid culture cheese trumpet find dinosaur curve tray"; \
  echo "mitomito"; \
  echo "mitomito"\
) | $DAEMON keys add validator --recover --keyring-backend 'file'

(
  echo "virus oxygen upgrade fitness diagram impact avocado cake cruise glass force joke galaxy project friend icon school midnight front actress squirrel wish phone sock"; \
  echo "mitomito"; \
  echo "mitomito"\
) | $DAEMON keys add signer --recover --keyring-backend 'file'

echo "mitomito" | $DAEMON add-genesis-account $($DAEMON keys show validator -a --keyring-backend 'file') 2000000000000umito --keyring-backend "file"
echo "mitomito" | $DAEMON add-genesis-account $($DAEMON keys show signer -a --keyring-backend 'file') 2000000000000umito --keyring-backend "file"

echo "mitomito" | $DAEMON gentx validator 1000000000000umito --keyring-backend "file" --chain-id 'mito-local-1'

$DAEMON collect-gentxs
