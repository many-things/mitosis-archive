package keeper_test

var (
	kvPubKeyRepoKey = []byte{0x02}
)

//func Test_RegisterPubKey(t *testing.T) {
//	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
//	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
//	valAddr := testutils.GenValAddress(t)
//
//	pubKey := types.PubKey{
//		Chain: chainID,
//		KeyID: 0,
//		Items: []*types.PubKey_Item{{
//			Participant: valAddr,
//			PubKey:      testutils.GenPublicKey(t),
//		}},
//	}
//	err := k.RegisterPubKey(ctx, chainID, &pubKey)
//	assert.NilError(t, err)
//
//	// test generated successfully
//	savedPubKey, err := repo.Load(pubKey.KeyID)
//	assert.NilError(t, err)
//	require.Equal(t, pubKey, *savedPubKey)
//}

//
//func Test_RemovePubKey(t *testing.T) {
//	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
//	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
//	valAddr := testutils.GenValAddress(t)
//
//	// try to delete not exist pubKey
//	err := k.RemovePubKey(ctx, chainID, 0, valAddr)
//	assert.Error(t, err, "pubkey: not found")
//
//	// try to delete exist pubKey
//	pubKey := types.PubKey{
//		Chain:       chainID,
//		KeyID:       0,
//		Participant: valAddr,
//		PubKey:      testutils.GenPublicKey(t),
//	}
//	err = repo.Create(&pubKey)
//	assert.NilError(t, err)
//
//	err = k.RemovePubKey(ctx, chainID, pubKey.KeyID, pubKey.Participant)
//	assert.NilError(t, err)
//
//	// validate
//	_, err = repo.Load(pubKey.KeyID, pubKey.Participant)
//	assert.Error(t, err, "pubkey: not found")
//}
//
//func Test_QueryPubKey(t *testing.T) {
//	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
//	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
//	valAddr := testutils.GenValAddress(t)
//
//	// try to query not exist pubKey
//	_, err := k.QueryPubKey(ctx, chainID, 0, valAddr)
//	assert.Error(t, err, "pubkey: not found")
//
//	// try to query exist pubKey
//	pubKey := types.PubKey{
//		Chain:       chainID,
//		KeyID:       0,
//		Participant: valAddr,
//		PubKey:      testutils.GenPublicKey(t),
//	}
//	err = repo.Create(&pubKey)
//	assert.NilError(t, err)
//
//	res, err := k.QueryPubKey(ctx, chainID, pubKey.KeyID, pubKey.Participant)
//	assert.NilError(t, err)
//	require.Equal(t, pubKey, *res)
//}
//
//func Test_QueryPubKeyList(t *testing.T) {
//	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
//	repo := state.NewKVChainPubKeyRepo(cdc, ctx.KVStore(storeKey), chainID)
//
//	for i := 0; i < 10; i++ {
//		pubKey := types.PubKey{
//			Chain:       chainID,
//			KeyID:       1,
//			Participant: testutils.GenValAddress(t),
//			PubKey:      testutils.GenPublicKey(t),
//		}
//		_ = repo.Create(&pubKey)
//	}
//
//	res, _, err := k.QueryPubKeyList(ctx, chainID, 1, &query.PageRequest{Limit: query.MaxLimit})
//	assert.NilError(t, err)
//	require.Equal(t, len(res), 10)
//}
