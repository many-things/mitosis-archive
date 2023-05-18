package server

import (
	"context"
	"github.com/many-things/mitosis/x/context/keeper"
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/many-things/mitosis/testutil/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, MsgServer, context.Context) {
	k, ctx, _, _, _ := keepertest.ContextKeeper(t)
	return k, NewMsgServer(k), sdk.WrapSDKContext(ctx)
}

func TestRegisterClearVault(t *testing.T) {
	k, s, ctx := setupMsgServer(t)

	wctx := sdk.UnwrapSDKContext(ctx)

	registerResp, err := s.RegisterVault(ctx, &MsgRegisterVault{
		Sender:    nil,
		Chain:     "osmo-test-1",
		VaultAddr: "osmo1deadbeefdeadbeefdeadbeef",
	})
	require.Nil(t, err)
	require.Equal(t, registerResp, &MsgRegisterVaultResponse{})

	vault, err := k.QueryVault(wctx, "osmo-test-1")
	require.Nil(t, err)
	require.Equal(t, vault, "osmo1deadbeefdeadbeefdeadbeef")

	clearResp, err := s.ClearVault(ctx, &MsgClearVault{Chain: "osmo-test-1"})
	require.Nil(t, err)
	require.Equal(t, clearResp, &MsgClearVaultResponse{})

	_, err = k.QueryVault(wctx, "osmo-test-1")
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "vault not found")
}
