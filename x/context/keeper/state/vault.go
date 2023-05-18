package state

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	mitotypes "github.com/many-things/mitosis/pkg/types"
	"github.com/many-things/mitosis/x/context/types"
	"github.com/pkg/errors"
)

type VaultRepo interface {
	Load(chain string) (string, error)
	List(page *query.PageRequest) ([]mitotypes.KV[string, string], *query.PageResponse, error)
	Clear(chain string) error
	Save(chain, address string) error

	ExportGenesis() (genState *types.GenesisVault, err error)
	ImportGenesis(genState *types.GenesisVault) error
}

var (
	kvVaultRepoItemPrefix = []byte{0x01}
)

type kvVaultRepo struct {
	cdc  codec.BinaryCodec
	root store.KVStore
}

func NewKVVaultRepo(cdc codec.BinaryCodec, store store.KVStore) VaultRepo {
	root := prefix.NewStore(store, kvVaultRepoKey)

	return kvVaultRepo{cdc, root}
}

func (k kvVaultRepo) items() store.KVStore {
	return prefix.NewStore(k.root, kvVaultRepoItemPrefix)
}

func (k kvVaultRepo) Load(chain string) (string, error) {
	bz := k.items().Get([]byte(chain))
	if bz == nil {
		return "", errors.Errorf("vault not found for chain %s", chain)
	}

	return string(bz), nil
}

func (k kvVaultRepo) List(page *query.PageRequest) ([]mitotypes.KV[string, string], *query.PageResponse, error) {
	var kvs []mitotypes.KV[string, string]

	pageResp, err := query.Paginate(
		k.items(),
		page,
		func(chainBz []byte, vaultAddrBz []byte) error {
			kvs = append(kvs, mitotypes.NewKV(string(chainBz), string(vaultAddrBz)))

			return nil
		},
	)

	return kvs, pageResp, err
}

func (k kvVaultRepo) Clear(chain string) error {
	k.items().Delete([]byte(chain))

	return nil
}

func (k kvVaultRepo) Save(chain, address string) error {
	k.items().Set([]byte(chain), []byte(address))

	return nil
}

func (k kvVaultRepo) ExportGenesis() (genState *types.GenesisVault, err error) {
	genState = &types.GenesisVault{}

	_, err = query.Paginate(
		k.items(),
		&query.PageRequest{Limit: query.MaxLimit},
		func(chainBz []byte, vaultAddrBz []byte) error {
			genState.ChainSet = append(
				genState.ChainSet,
				&types.GenesisVault_ChainSet{
					Chain: string(chainBz),
					Vault: string(vaultAddrBz),
				},
			)

			return nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "paginate items")
	}

	return
}

func (k kvVaultRepo) ImportGenesis(genState *types.GenesisVault) error {
	itemStore := k.items()

	for _, chain := range genState.ChainSet {
		itemStore.Set([]byte(chain.Chain), []byte(chain.Vault))
	}

	return nil
}
