package event

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/x/event/types"
)

func endBlocker(ctx sdk.Context, keeper types.BaseKeeper) {
	incomingEvents, err := keeper.ListIncomingEvent(ctx)
	if err != nil {
		panic(err.Error())
	}
	outgoingEvents, err := keeper.ListOutgoingEvent(ctx)
	if err != nil {
		panic(err.Error())
	}
	events := make(sdk.Events, len(incomingEvents)+len(outgoingEvents))

	for i, evt := range incomingEvents {
		events[i] = sdk.NewEvent(
			types.EventTypeIncomingEventAck,
			sdk.NewAttribute(types.AttributeKeyChainID, evt.GetChainId()),
			sdk.NewAttribute(types.AttributeKeyTxHash, evt.GetTxHash()),
			sdk.NewAttribute(types.AttributeKeyEventIndex, fmt.Sprint(evt.GetEventIndex())),
		)
	}

	for i, evt := range outgoingEvents {
		events[len(incomingEvents)+i] = sdk.NewEvent(
			types.EventTypeOutgoingEventAck,
			sdk.NewAttribute(types.AttributeKeyChainID, evt.GetChainId()),
			sdk.NewAttribute(types.AttributeKeyTxHash, evt.GetTxHash()),
		)
	}

	ctx.EventManager().EmitEvents(events)
}
