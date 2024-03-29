package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/many-things/mitosis/pkg/testutils"
	"github.com/many-things/mitosis/x/multisig/exported"
	"github.com/many-things/mitosis/x/multisig/keeper"
	"gotest.tools/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

func setupQueryServer(t testing.TB) (keeper.Keeper, QueryServer, context.Context) {
	k, ctx, _, _ := testkeeper.MultisigKeeper(t)
	return k, NewQueryServer(k), sdk.WrapSDKContext(ctx)
}

func TestParamsQuery(t *testing.T) {
	k, s, ctx := setupQueryServer(t)

	wctx := ctx.(sdk.Context)
	params := types.DefaultParams()
	k.SetParams(wctx, params)

	response, err := s.Params(wctx, &QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &QueryParamsResponse{Params: params}, response)
}

const (
	chainID = "chainID"
)

func Test_Keygen(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := sdk.ValAddress("address")

	// try to query not exist keygen
	_, err := s.Keygen(ctx, &QueryKeygen{
		Chain: chainID,
		Id:    0,
	})
	assert.Error(t, err, "keygen: not found")

	// try to query exist keygen
	keygen := types.Keygen{
		Chain:        chainID,
		KeyID:        0,
		Participants: []*types.Keygen_Participant{{Address: valAddr}},
		Status:       1,
	}
	_, _ = k.RegisterKeygenEvent(wctx, chainID, &keygen)

	res, err := s.Keygen(ctx, &QueryKeygen{
		Chain: chainID,
		Id:    0,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res, &QueryKeygenResponse{Keygen: &keygen})
}

func Test_KeygenList(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := sdk.ValAddress("address")

	var keygens []*types.Keygen
	var i uint64
	for i = 0; i < 5; i++ {
		keygen := types.Keygen{
			Chain:        chainID,
			KeyID:        i,
			Participants: []*types.Keygen_Participant{{Address: valAddr}},
			Status:       1,
		}

		_, _ = k.RegisterKeygenEvent(wctx, chainID, &keygen)
		keygens = append(keygens, &keygen)
	}

	res, err := s.KeygenList(wctx, &QueryKeygenList{
		Chain: chainID,
		Pagination: &query.PageRequest{
			Limit: query.MaxLimit,
		},
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.List, keygens)
}

func Test_KeygenResult(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := sdk.ValAddress("address")

	// try to query not exist pubkey
	_, err := s.KeygenResult(wctx, &QueryKeygenResult{
		KeyId:     fmt.Sprintf("%s-%d", chainID, 0),
		Validator: valAddr,
	})
	assert.Error(t, err, "keygen: not found")

	// try to query exist pubkey
	pubKey := types.KeygenResult{
		Chain: chainID,
		KeyID: 0,
		Items: []*types.KeygenResult_Item{{
			Participant: valAddr,
			PubKey:      exported.PublicKey("publickey"),
		}},
	}
	_ = k.RegisterKeygenResult(wctx, chainID, &pubKey)
	res, err := s.KeygenResult(wctx, &QueryKeygenResult{
		KeyId:     fmt.Sprintf("%s-%d", chainID, 0),
		Validator: valAddr,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.Result, &pubKey)
}

func Test_KeygenResultList(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)

	var pubKeyList []*types.KeygenResult
	var i uint64
	for i = 0; i < 5; i++ {
		pubKey := types.KeygenResult{
			Chain: chainID,
			KeyID: i,
			Items: []*types.KeygenResult_Item{{
				Participant: sdk.ValAddress(fmt.Sprintf("addr%d", i)),
				PubKey:      exported.PublicKey("publickey"),
			}},
		}
		_ = k.RegisterKeygenResult(wctx, chainID, &pubKey)

		pubKeyList = append(pubKeyList, &pubKey)
	}

	res, err := s.KeygenResultList(ctx, &QueryKeygenResultList{
		ChainId: chainID,
		Pagination: &query.PageRequest{
			Limit: query.MaxLimit,
		},
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.List, pubKeyList)
}

func Test_Sign(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)

	// try query not exist sign
	_, err := s.Sign(wctx, &QuerySign{
		Chain: chainID,
		Id:    0,
	})
	assert.Error(t, err, "sign: not found")

	// try query exist sign
	sign := exported.Sign{
		Chain:         chainID,
		SigID:         0,
		KeyID:         fmt.Sprintf("%s-%d", chainID, 1),
		Participants:  []sdk.ValAddress{sdk.ValAddress("addr")},
		MessageToSign: exported.Hash("msgToSign"),
		Status:        0,
	}
	_, _ = k.RegisterSignEvent(wctx, chainID, &sign)

	res, err := s.Sign(wctx, &QuerySign{
		Chain: chainID,
		Id:    0,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, *res.Sign, sign)
}

func Test_SignList(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)
	valAddr := sdk.ValAddress("addr")

	var signs []*exported.Sign
	var i uint64
	for i = 0; i < 5; i++ {
		sign := exported.Sign{
			Chain:         chainID,
			SigID:         i,
			KeyID:         fmt.Sprintf("%s-%d", chainID, i%3),
			Participants:  []sdk.ValAddress{valAddr},
			MessageToSign: exported.Hash("msg"),
			Status:        1,
		}

		_, _ = k.RegisterSignEvent(wctx, chainID, &sign)
		signs = append(signs, &sign)
	}

	res, err := s.SignList(wctx, &QuerySignList{
		Chain:      chainID,
		Pagination: &query.PageRequest{Limit: query.MaxLimit},
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.List, signs)
}

func Test_Signature(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)
	val := sdk.ValAddress("val")

	// try to query not exist signature
	_, err := s.SignResult(wctx, &QuerySignResult{
		SigId: fmt.Sprintf("%s-%d", chainID, 0),
	})
	assert.Error(t, err, "sign_signature: not found")

	// try to query exist signature
	signature := exported.SignResult{
		Chain: chainID,
		SigID: 0,
		Items: []*exported.SignResult_Item{{
			Participant: val,
			Signature:   exported.Signature("Signature"),
		}},
	}
	_ = k.RegisterSignResult(wctx, chainID, &signature)

	res, err := s.SignResult(wctx, &QuerySignResult{
		SigId: fmt.Sprintf("%s-%d", chainID, 0),
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.Signature, &signature)
}

func Test_SignatureList(t *testing.T) {
	k, s, ctx := setupQueryServer(t)
	wctx := ctx.(sdk.Context)

	var signs []*exported.SignResult
	var i uint64
	for i = 0; i < 5; i++ {
		signature := exported.SignResult{
			Chain: chainID,
			SigID: i,
			Items: []*exported.SignResult_Item{{
				Participant: testutils.GenValAddress(t),
				Signature:   exported.Signature("Signature"),
			}},
		}
		_ = k.RegisterSignResult(
			wctx,
			chainID,
			&signature,
		)

		signs = append(signs, &signature)
	}

	res, err := s.SignResultList(wctx, &QuerySignResultList{
		ChainId:    chainID,
		Pagination: &query.PageRequest{Limit: query.MaxLimit},
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, res.List, signs)
}
