package types

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		Operation: &GenesisOperation{
			LastId:   0,
			ItemSet:  []*GenesisOperation_ItemSet{},
			IndexSet: []*GenesisOperation_IndexSet{},
		},
		OperationIdx: &GenesisOperationHashIndex{
			ChainSet: []*GenesisOperationHashIndex_ChainSet{},
		},
		Vault: &GenesisVault{
			ChainSet: []*GenesisVault_ChainSet{},
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
