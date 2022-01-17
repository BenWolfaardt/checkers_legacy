package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/BenWolfaardt/Checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [red] [black] [wager]",
		Short: "Broadcast message createGame",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRed := string(args[0])
			argsBlack := string(args[1])
			argsWager, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				panic(err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(clientCtx.GetFromAddress().String(), string(argsRed), string(argsBlack), uint64(argsWager))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
