package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/cosmos/ibc-go/v3/modules/apps/29-fee/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

const (
	flagRecvFee    = "recv-fee"
	flagAckFee     = "ack-fee"
	flagTimeoutFee = "timeout-fee"
)

// NewPayPacketFeeAsyncTxCmd returns the command to create a MsgPayPacketFeeAsync
func NewPayPacketFeeAsyncTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pay-packet-fee [src-port] [src-channel] [sequence]",
		Short:   "Pay a fee to incentivize an existing IBC packet",
		Long:    strings.TrimSpace(`Pay a fee to incentivize an existing IBC packet.`),
		Example: fmt.Sprintf("%s tx pay-packet-fee [src-port] [src-channel] [sequence] --recv-fee 10stake --ack-fee 10stake --timeout-fee 10stake", version.AppName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// NOTE: specifying non-nil relayers is currently unsupported
			var relayers []string

			sender := clientCtx.GetFromAddress().String()
			seq, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			packetID := channeltypes.NewPacketId(args[1], args[0], seq)

			recvFeeStr, err := cmd.Flags().GetString(flagRecvFee)
			if err != nil {
				return err
			}

			recvFee, err := sdk.ParseCoinsNormalized(recvFeeStr)
			if err != nil {
				return err
			}

			ackFeeStr, err := cmd.Flags().GetString(flagAckFee)
			if err != nil {
				return err
			}

			ackFee, err := sdk.ParseCoinsNormalized(ackFeeStr)
			if err != nil {
				return err
			}

			timeoutFeeStr, err := cmd.Flags().GetString(flagTimeoutFee)
			if err != nil {
				return err
			}

			timeoutFee, err := sdk.ParseCoinsNormalized(timeoutFeeStr)
			if err != nil {
				return err
			}

			fee := types.Fee{
				RecvFee:    recvFee,
				AckFee:     ackFee,
				TimeoutFee: timeoutFee,
			}

			identifiedPacketFee := types.NewIdentifiedPacketFee(packetID, fee, sender, relayers)

			msg := types.NewMsgPayPacketFeeAsync(identifiedPacketFee)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagRecvFee, "", "Fee paid to a relayer for relaying a packet receive.")
	cmd.Flags().String(flagAckFee, "", "Fee paid to a relayer for relaying a packet acknowledgement.")
	cmd.Flags().String(flagTimeoutFee, "", "Fee paid to a relayer for relaying a packet timeout.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
