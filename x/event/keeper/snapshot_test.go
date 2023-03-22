package keeper_test

import (
	crand "crypto/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	testkeeper "github.com/many-things/mitosis/testutil/keeper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSnapshot(t *testing.T) {
	k, ctx := testkeeper.EventKeeper(t)

	ctx = ctx.WithBlockHeight(123)

	vals := mitotypes.Map(
		make([]byte, 2),
		func(_ byte) sdk.ValAddress {
			bz := make([]byte, 32)
			_, err := crand.Read(bz)
			require.NoError(t, err)
			return bz
		},
	)

	epoch, err := k.CreateSnapshot(
		ctx, sdk.NewInt(100),
		mitotypes.Map(
			vals,
			func(val sdk.ValAddress) mitotypes.KV[sdk.ValAddress, int64] {
				return mitotypes.NewKV(val, int64(100))
			},
		),
	)
	require.NoError(t, err)

	votingPower, err := k.VotingPowerOf(ctx, mitotypes.Ref(epoch.GetEpoch()), vals[0])
	require.NoError(t, err)
	require.Equal(t, int64(100), votingPower)

	fetchedEpoch, err := k.LatestSnapshotEpoch(ctx)
	require.NoError(t, err)
	require.Equal(t, epoch, fetchedEpoch)
}
