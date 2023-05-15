package multisig_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"log"
	"testing"

	keepertest "github.com/many-things/mitosis/testutil/keeper"
	"github.com/many-things/mitosis/testutil/nullify"
	"github.com/many-things/mitosis/x/multisig"
	"github.com/many-things/mitosis/x/multisig/types"
	"github.com/stretchr/testify/require"
)

const genesis = `
{
  "params": {},
  "keygen": {
	"chain_set": [
	  {
		"chain": "osmo-test-5",
	    "last_id": 0,
	    "item_set": [
		  {
			"chain": "osmo-test-5",
			"key_id": 0,
			"participants": [{
			  "address": "mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn",
			  "share": 1
			}],
			"status": 3,
			"threshold": 1
		  }
		],
		"result_set": [
		  {
			"chain": "osmo-test-5",
			"key_id": 0,
			"items": [
			  {
				"participant": "mitovaloper127rnkklgkc5alfgtzlv2mk4capr4kdc30a36tn",
				"pub_key": "AuEd7BYv0sukuPOzOWrlKM73nWLTlFP+QG4HRriyLF22"
			  }
			]
		  }
		]
	  }
	]
  },
  "sign": {
	"chain_set": []
  }
}`

func TestGenesis(t *testing.T) {
	// Set prefixes
	accountPubKeyPrefix := "mitopub"
	validatorAddressPrefix := "mitovaloper"
	validatorPubKeyPrefix := "mitovaloperpub"
	consNodeAddressPrefix := "mitovalcons"
	consNodePubKeyPrefix := "mitovalconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("mito", accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()

	k, ctx, codec, _ := keepertest.MultisigKeeper(t)

	genesisState := types.GenesisState{}
	codec.MustUnmarshalJSON([]byte(genesis), &genesisState)
	log.Println(string(codec.MustMarshalJSON(&genesisState)))

	multisig.InitGenesis(ctx, k, genesisState)
	got := multisig.ExportGenesis(ctx, k, []byte{})
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
