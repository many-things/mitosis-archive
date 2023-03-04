package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/many-things/mitosis/pkg/utils"
)

// KEYS
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

// EVENTS
const (
	EventTypeIncomingEventAck = "ack_incoming_event"
	EventTypeOutgoingEventAck = "ack_outgoing_event"

	AttributeKeyChainID    = "chain_id"
	AttributeKeyTxHash     = "tx_hash"
	AttributeKeyEventIndex = "event_index"
)

// ERRORS
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)

// PREFIX
var (
	PrefixVoteIncomingEvent = []byte{0x00}
	PrefixVoteOutgoingEvent = []byte{0x01}

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
