package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		Poll: &GenesisPoll{
			LatestId: 0,
			ItemSet:  []*GenesisPoll_ItemSet{},
			HashSet:  []*GenesisPoll_HashSet{},
		},
		Proxy: &GenesisProxy{
			ItemSet: []*GenesisProxy_ItemSet{},
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs *GenesisState) Validate() error {

	return gs.Params.Validate()
}
