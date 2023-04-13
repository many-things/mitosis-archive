package keeper_test

import (
	"fmt"
	"testing"

	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	mitosistype "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/multisig/exported"

	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/keeper/state"
)

func Test_RegisterSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	var (
		signature exported.Signature = []byte{48, 69, 2, 33, 0, 204, 158, 126, 191, 230, 246, 144, 76, 51, 12, 200, 239, 81, 7, 62, 206, 158, 181, 61, 231, 107, 29, 245, 13, 163, 147, 105, 73, 151, 129, 230, 239, 2, 32, 1, 105, 187, 151, 9, 40, 68, 195, 211, 66, 249, 139, 145, 170, 94, 50, 88, 103, 125, 173, 134, 222, 100, 66, 173, 87, 71, 197, 189, 193, 22, 218}
		digest    exported.Hash      = []byte{165, 145, 166, 212, 11, 244, 32, 64, 74, 1, 23, 51, 207, 183, 177, 144, 214, 44, 101, 191, 11, 205, 163, 43, 87, 178, 119, 217, 173, 159, 20, 110}
		pubKey    exported.PublicKey = []byte{2, 38, 4, 37, 65, 180, 103, 11, 144, 52, 3, 218, 159, 53, 110, 248, 152, 25, 82, 72, 188, 187, 104, 19, 44, 56, 45, 74, 161, 41, 182, 157, 132}
	)

	keyID, err := state.NewKVChainKeygenRepo(cdc, ctx.KVStore(storeKey), chainID).Create(
		&types.Keygen{
			Chain: chainID,
			Participants: []*types.Keygen_Participant{{
				Address: valAddr,
				Share:   0,
			}},
			Status: types.Keygen_StatusComplete,
		},
	)
	require.Nil(t, err)

	require.Nil(
		t, state.NewKVChainKeygenResultRepo(cdc, ctx.KVStore(storeKey), chainID).Create(&types.KeygenResult{
			Chain: chainID,
			KeyID: keyID,
			Items: []*types.KeygenResult_Item{
				{
					Participant: valAddr,
					PubKey:      pubKey,
				},
			},
		}),
	)

	require.Nil(
		t, state.NewKVChainSignRepo(cdc, ctx.KVStore(storeKey), chainID).Save(
			&exported.Sign{
				Chain:         chainID,
				SigID:         0,
				KeyID:         fmt.Sprintf("%s-%d", chainID, keyID),
				Participants:  []sdk.ValAddress{valAddr},
				MessageToSign: digest,
				Status:        exported.Sign_StatusAssign,
			},
		),
	)

	signSignature := exported.SignResult{Chain: chainID, SigID: 0, Items: []*exported.SignResult_Item{}}
	require.Nil(t, k.RegisterSignResult(ctx, chainID, &signSignature))
	require.Nil(t, k.AddParticipantSignResult(ctx, chainID, 0, valAddr, signature))

	// validate registered successfully
	res, err := repo.Load(0)
	require.Nil(t, err)
	require.Equal(t, []*exported.SignResult_Item{{
		Participant: valAddr,
		Signature:   signature,
	}}, res.Items)
}

func Test_RemoveSignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to remove not exist signature
	err := k.RemoveSignResult(ctx, chainID, 0)
	assert.Error(t, err, "sign_signature: not found")

	// try to remove exist signature
	signSignature := exported.SignResult{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignResult_Item{{
			Participant: valAddr,
			Signature:   exported.Signature("signature"),
		}},
	}
	_ = repo.Save(&signSignature)

	err = k.RemoveSignResult(ctx, chainID, 0)
	require.Nil(t, err)

	// validate signature not exists
	_, err = repo.Load(0)
	assert.Error(t, err, "sign_signature: not found")
}

func Test_QuerySignature(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignResultRepo(cdc, ctx.KVStore(storeKey), chainID)
	valAddr := testutils.GenValAddress(t)

	// try to query not exist signature
	_, err := k.QuerySignResult(ctx, chainID, 0)
	assert.Error(t, err, "sign_signature: not found")

	// try to query exist signature
	signSignature := exported.SignResult{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignResult_Item{{
			Participant: valAddr,
			Signature:   exported.Signature("signature"),
		}},
	}
	_ = repo.Save(&signSignature)

	res, err := k.QuerySignResult(ctx, chainID, 0)
	require.Nil(t, err)
	require.Equal(t, res, &signSignature)
}

func Test_QuerySignatureList(t *testing.T) {
	k, ctx, cdc, storeKey := testkeeper.MultisigKeeper(t)
	repo := state.NewKVChainSignResultRepo(cdc, ctx.KVStore(storeKey), chainID)

	var signatures []mitosistype.KV[uint64, *exported.SignResult]
	var i uint64
	for i = 0; i < 10; i++ {
		signSignature := exported.SignResult{
			Chain: chainID,
			SigID: i,
			Items: []*exported.SignResult_Item{{
				Participant: sdk.ValAddress(fmt.Sprintf("addr%d", i)),
				Signature:   exported.Signature(fmt.Sprintf("signature%d", i)),
			}},
		}

		_ = repo.Save(&signSignature)
		signatures = append(signatures, mitosistype.NewKV(i, &signSignature))
	}

	res, _, err := k.QuerySignResultList(ctx, chainID, &query.PageRequest{Limit: query.MaxLimit})
	require.Nil(t, err)
	require.Equal(t, res, signatures)
}
