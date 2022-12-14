package keeper

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
)

// EmitChannelOpenInitEvent emits a channel open init event
func EmitChannelOpenInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenInit,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
			sdk.NewAttribute(types.AttributeVersion, channel.Version),
		),
	})
}

// EmitChannelOpenTryEvent emits a channel open try event
func EmitChannelOpenTryEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenTry,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
			sdk.NewAttribute(types.AttributeVersion, channel.Version),
		),
	})
}

// EmitChannelOpenAckEvent emits a channel open acknowledge event
func EmitChannelOpenAckEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenAck,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
		),
	})
}

// EmitChannelOpenConfirmEvent emits a channel open confirm event
func EmitChannelOpenConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelOpenConfirm,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
		),
	})
}

// EmitChannelCloseInitEvent emits a channel close init event
func EmitChannelCloseInitEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelCloseInit,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
		),
	})
}

// EmitChannelCloseConfirmEvent emits a channel close confirm event
func EmitChannelCloseConfirmEvent(ctx sdk.Context, portID string, channelID string, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelCloseConfirm,
			sdk.NewAttribute(types.AttributeKeyPortID, portID),
			sdk.NewAttribute(types.AttributeKeyChannelID, channelID),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
		),
	})
}

// EmitSendPacketEvent emits an event with packet data along with other packet information for relayer
// to pick up and relay to other chain
func EmitSendPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel, timeoutHeight exported.Height) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSendPacket,
			sdk.NewAttribute(types.AttributeKeyData, string(packet.GetData())), //nolint:staticcheck // DEPRECATED
			sdk.NewAttribute(types.AttributeKeyDataHex, hex.EncodeToString(packet.GetData())),
			sdk.NewAttribute(types.AttributeKeyTimeoutHeight, timeoutHeight.String()),
			sdk.NewAttribute(types.AttributeKeyTimeoutTimestamp, fmt.Sprintf("%d", packet.GetTimeoutTimestamp())),
			sdk.NewAttribute(types.AttributeKeySequence, fmt.Sprintf("%d", packet.GetSequence())),
			sdk.NewAttribute(types.AttributeKeySrcPort, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeySrcChannel, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeKeyDstPort, packet.GetDestPort()),
			sdk.NewAttribute(types.AttributeKeyDstChannel, packet.GetDestChannel()),
			sdk.NewAttribute(types.AttributeKeyChannelOrdering, channel.Ordering.String()),
			// we only support 1-hop packets now, and that is the most important hop for a relayer
			// (is it going to a chain I am connected to)
			sdk.NewAttribute(types.AttributeKeyConnection, channel.ConnectionHops[0]),
		),
	})
}

// EmitRecvPacketEvent emits a receive packet event. It will be emitted both the first time a packet
// is received for a certain sequence and for all duplicate receives.
func EmitRecvPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRecvPacket,
			sdk.NewAttribute(types.AttributeKeyData, string(packet.GetData())), //nolint:staticcheck // DEPRECATED
			sdk.NewAttribute(types.AttributeKeyDataHex, hex.EncodeToString(packet.GetData())),
			sdk.NewAttribute(types.AttributeKeyTimeoutHeight, packet.GetTimeoutHeight().String()),
			sdk.NewAttribute(types.AttributeKeyTimeoutTimestamp, fmt.Sprintf("%d", packet.GetTimeoutTimestamp())),
			sdk.NewAttribute(types.AttributeKeySequence, fmt.Sprintf("%d", packet.GetSequence())),
			sdk.NewAttribute(types.AttributeKeySrcPort, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeySrcChannel, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeKeyDstPort, packet.GetDestPort()),
			sdk.NewAttribute(types.AttributeKeyDstChannel, packet.GetDestChannel()),
			sdk.NewAttribute(types.AttributeKeyChannelOrdering, channel.Ordering.String()),
			// we only support 1-hop packets now, and that is the most important hop for a relayer
			// (is it going to a chain I am connected to)
			sdk.NewAttribute(types.AttributeKeyConnection, channel.ConnectionHops[0]),
		),
	})
}

// EmitWriteAcknowledgementEvent emits an event that the relayer can query for
func EmitWriteAcknowledgementEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel, acknowledgement []byte) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWriteAck,
			sdk.NewAttribute(types.AttributeKeyData, string(packet.GetData())), //nolint:staticcheck // DEPRECATED
			sdk.NewAttribute(types.AttributeKeyDataHex, hex.EncodeToString(packet.GetData())),
			sdk.NewAttribute(types.AttributeKeyTimeoutHeight, packet.GetTimeoutHeight().String()),
			sdk.NewAttribute(types.AttributeKeyTimeoutTimestamp, fmt.Sprintf("%d", packet.GetTimeoutTimestamp())),
			sdk.NewAttribute(types.AttributeKeySequence, fmt.Sprintf("%d", packet.GetSequence())),
			sdk.NewAttribute(types.AttributeKeySrcPort, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeySrcChannel, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeKeyDstPort, packet.GetDestPort()),
			sdk.NewAttribute(types.AttributeKeyDstChannel, packet.GetDestChannel()),
			sdk.NewAttribute(types.AttributeKeyAck, string(acknowledgement)), //nolint:staticcheck // DEPRECATED
			sdk.NewAttribute(types.AttributeKeyAckHex, hex.EncodeToString(acknowledgement)),
			// we only support 1-hop packets now, and that is the most important hop for a relayer
			// (is it going to a chain I am connected to)
			sdk.NewAttribute(types.AttributeKeyConnection, channel.ConnectionHops[0]),
		),
	})
}

// EmitAcknowledgePacketEvent emits an acknowledge packet event. It will be emitted both the first time
// a packet is acknowledged for a certain sequence and for all duplicate acknowledgements.
func EmitAcknowledgePacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAcknowledgePacket,
			sdk.NewAttribute(types.AttributeKeyTimeoutHeight, packet.GetTimeoutHeight().String()),
			sdk.NewAttribute(types.AttributeKeyTimeoutTimestamp, fmt.Sprintf("%d", packet.GetTimeoutTimestamp())),
			sdk.NewAttribute(types.AttributeKeySequence, fmt.Sprintf("%d", packet.GetSequence())),
			sdk.NewAttribute(types.AttributeKeySrcPort, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeySrcChannel, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeKeyDstPort, packet.GetDestPort()),
			sdk.NewAttribute(types.AttributeKeyDstChannel, packet.GetDestChannel()),
			sdk.NewAttribute(types.AttributeKeyChannelOrdering, channel.Ordering.String()),
			// we only support 1-hop packets now, and that is the most important hop for a relayer
			// (is it going to a chain I am connected to)
			sdk.NewAttribute(types.AttributeKeyConnection, channel.ConnectionHops[0]),
		),
	})
}

// EmitTimeoutPacketEvent emits a timeout packet event. It will be emitted both the first time a packet
// is timed out for a certain sequence and for all duplicate timeouts.
func EmitTimeoutPacketEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTimeoutPacket,
			sdk.NewAttribute(types.AttributeKeyTimeoutHeight, packet.GetTimeoutHeight().String()),
			sdk.NewAttribute(types.AttributeKeyTimeoutTimestamp, fmt.Sprintf("%d", packet.GetTimeoutTimestamp())),
			sdk.NewAttribute(types.AttributeKeySequence, fmt.Sprintf("%d", packet.GetSequence())),
			sdk.NewAttribute(types.AttributeKeySrcPort, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeySrcChannel, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeKeyDstPort, packet.GetDestPort()),
			sdk.NewAttribute(types.AttributeKeyDstChannel, packet.GetDestChannel()),
			sdk.NewAttribute(types.AttributeKeyChannelOrdering, channel.Ordering.String()),
		),
	})
}

// EmitChannelClosedEvent emits a channel closed event.
func EmitChannelClosedEvent(ctx sdk.Context, packet exported.PacketI, channel types.Channel) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeChannelClosed,
			sdk.NewAttribute(types.AttributeKeyPortID, packet.GetSourcePort()),
			sdk.NewAttribute(types.AttributeKeyChannelID, packet.GetSourceChannel()),
			sdk.NewAttribute(types.AttributeCounterpartyPortID, channel.Counterparty.PortId),
			sdk.NewAttribute(types.AttributeCounterpartyChannelID, channel.Counterparty.ChannelId),
			sdk.NewAttribute(types.AttributeKeyConnectionID, channel.ConnectionHops[0]),
			sdk.NewAttribute(types.AttributeKeyChannelOrdering, channel.Ordering.String()),
		),
	})
}
