package multisig

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/many-things/mitosis/testutil/sample"
	multisigsimulation "github.com/many-things/mitosis/x/multisig/simulation"
	"github.com/many-things/mitosis/x/multisig/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = multisigsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgStartKeygen = "op_weight_msg_start_keygen"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartKeygen int = 100

	opWeightMsgSubmitPubkey = "op_weight_msg_submit_pubkey"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPubkey int = 100

	opWeightMsgSubmitSignature = "op_weight_msg_submit_signature"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitSignature int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	multisigGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&multisigGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgStartKeygen int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartKeygen, &weightMsgStartKeygen, nil,
		func(_ *rand.Rand) {
			weightMsgStartKeygen = defaultWeightMsgStartKeygen
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartKeygen,
		multisigsimulation.SimulateMsgStartKeygen(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitPubkey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitPubkey, &weightMsgSubmitPubkey, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPubkey = defaultWeightMsgSubmitPubkey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPubkey,
		multisigsimulation.SimulateMsgSubmitPubkey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitSignature int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitSignature, &weightMsgSubmitSignature, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitSignature = defaultWeightMsgSubmitSignature
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitSignature,
		multisigsimulation.SimulateMsgSubmitSignature(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
