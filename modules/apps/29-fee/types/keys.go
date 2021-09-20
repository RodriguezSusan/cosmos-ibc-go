package types

import "fmt"

const (
	// ModuleName defines the 29-fee name
	ModuleName = "feeibc"

	// StoreKey is the store key string for IBC transfer
	StoreKey = ModuleName

	// RouterKey is the message route for IBC transfer
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = ModuleName

	Version = "fee29-1"
)

func FeeEnabledKey(portID, channelID string) []byte {
	return []byte(fmt.Sprintf("fee_enabled/%s/%s", portID, channelID))
}
