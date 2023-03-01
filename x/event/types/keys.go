package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/pkg/utils"
)

const (
	// ModuleName defines the module name
	ModuleName = "event"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_event"
)

var (
	// === Events
	PrefixVoteIncomingEvent = []byte{0x00}
	PrefixVoteOutgoingEvent = []byte{0x01}

	// === Validator
	PrefixValidatorProxy = []byte{0x10}
)

func GetIncomingEventPrefix(height uint64) []byte {
	return utils.JoinBytes(
		PrefixVoteIncomingEvent,
		sdk.Uint64ToBigEndian(height),
	)
}

func GetOutgoingEventPrefix(height uint64) []byte {
	return utils.JoinBytes(PrefixVoteOutgoingEvent,
		sdk.Uint64ToBigEndian(height),
	)
}

func GetValidatorProxyPrefix() []byte {
	return PrefixValidatorProxy
}
