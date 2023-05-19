#!/bin/sh

DAEMON=${DAEMON:-"./build/mitosisd"}

function add_key() {
  (echo $1; echo $2; echo $2) | $DAEMON keys add $3 --recover --keyring-backend 'file'
}

MITO1_ADDR=${MITO1_ADDR:-"slush tray spoon riot install coil few voyage logic twenty hazard aim knee mosquito drum sound novel sweet unusual chronic crumble picture equal owner"}
MITO2_ADDR=${MITO2_ADDR:-"float theme stumble device ancient salon pumpkin dragon grape cross bid base journey suspect roof rookie fossil merit decide message pig kind angry steak"}
MITO3_ADDR=${MITO3_ADDR:-"sugar gaze sort music next relax infant differ forward feed boss scatter example pumpkin dice spike travel skate helmet clarify gentle impact iron bulk"}
PASSPHRASE=${PASSPHRASE:-"mitomito"}

add_key "$MITO1_ADDR" "$PASSPHRASE" mito1
add_key "$MITO2_ADDR" "$PASSPHRASE" mito2
add_key "$MITO3_ADDR" "$PASSPHRASE" mito3
