package omdbcmd

import (
	"github.com/spf13/cobra"
)

var (
	apiKey          string
	noCache         bool
	cacheExpiration int64
	outputJson      bool
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "omdb",
		Short: "Get Data Using The OMDB API",
		Long:  `Search Movies & Shows, Get Data On Movies etc.`,
	}

	cmd.PersistentFlags().StringVar(&apiKey, "apikey", "", "omdb api key")
	cmd.PersistentFlags().BoolVar(&noCache, "nocache", false, "disable caching (not recommended)")
	cmd.PersistentFlags().Int64Var(&cacheExpiration, "cache-expires", 0, "cache expiration in minutes")
	cmd.PersistentFlags().BoolVar(&outputJson, "json", false, "output result as json")

	cmd.AddCommand(searchCmd, getCmd)

	return cmd
}
