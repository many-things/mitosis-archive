package server

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
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
		Participants: []sdk.ValAddress{valAddr},
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
			Participants: []sdk.ValAddress{valAddr},
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

func Test_PubKey(_ *testing.T) {
	// TODO: implement
}

func Test_PubKeyList(_ *testing.T) {
	// TODO: implement
}

func Test_Sign(_ *testing.T) {
	// TODO: implement
}

func Test_SignList(_ *testing.T) {
	// TODO: implement
}

func Test_Signature(_ *testing.T) {
	// TODO: implement
}

func Test_SignatureList(_ *testing.T) {
	// TODO: implement
}
