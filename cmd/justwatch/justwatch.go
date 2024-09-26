package justwatchcmd

import (
	"github.com/spf13/cobra"
)

var (
	country    string
	language   string
	outputJson bool
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "justwatch",
		Short: "Get Data From Justwatch",
		Long:  `Search Movies & Shows, Get Data On Movies etc.`,
	}

	cmd.PersistentFlags().StringVar(&country, "country", "", "country code (defaults to US)")
	cmd.PersistentFlags().StringVar(&language, "language", "", "language code (defaults to en)")
	cmd.PersistentFlags().BoolVar(&outputJson, "json", false, "output result as json")

	cmd.AddCommand(getCmd, searchCmd, getOffersCmd)

	return cmd
}
