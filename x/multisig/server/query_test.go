package server

import (
	"context"
	"testing"

	"github.com/many-things/mitosis/x/multisig/keeper"

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

func Test_Keygen(_ *testing.T) {
	// TODO: implement
}

func Test_KeygenList(_ *testing.T) {
	// TODO: implement
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
