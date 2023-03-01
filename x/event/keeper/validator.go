package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
	"github.com/many-things/mitosis/x/event/keeper/store"
	"github.com/many-things/mitosis/x/event/types"
)

type ValidatorKeeper interface {
	RegisterProxy(ctx sdk.Context, msg *types.MsgRegisterProxy) error

	GetRegisteredProxy(ctx sdk.Context, validator sdk.ValAddress) (sdk.AccAddress, error)
}

type validatorKeeper struct {
	storeKey storetypes.StoreKey
}

func newValidatorKeeper(storeKey storetypes.StoreKey) ValidatorKeeper {
	return validatorKeeper{storeKey}
}

func (k validatorKeeper) RegisterProxy(ctx sdk.Context, msg *types.MsgRegisterProxy) error {
	valStore := store.NewValidatorProxyRepo(ctx, k.storeKey)

	validator := utils.Unwrap1(sdk.ValAddressFromBech32, msg.GetValidator())
	proxy := utils.Unwrap1(sdk.AccAddressFromBech32, msg.GetProxy())

	if err := valStore.Store(validator, proxy); err != nil {
		return err
	}

	return nil
}

func (k validatorKeeper) GetRegisteredProxy(ctx sdk.Context, validator sdk.ValAddress) (sdk.AccAddress, error) {
	valStore := store.NewValidatorProxyRepo(ctx, k.storeKey)

	proxy, err := valStore.Get(validator)
	if err != nil {
		return nil, err
	}

	return proxy, nil
}
