package imdbcmd

import (
	"github.com/spf13/cobra"
)

var (
	noCache         bool
	cacheExpiration int64
	outputJson      bool
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "imdb",
		Short: "Get Data From the Internet Movie Database (IMDb)",
		Long:  `Search People & Shows, Get Data On Movies etc.`,
	}

	cmd.PersistentFlags().BoolVar(&noCache, "nocache", false, "disable caching (not recommended)")
	cmd.PersistentFlags().Int64Var(&cacheExpiration, "cache-expires", 0, "cache expiration in minutes")
	cmd.PersistentFlags().BoolVar(&outputJson, "json", false, "output result as json")

	cmd.AddCommand(getCmd, searchCmd)

	return cmd
}
