package cmd

import (
	imdbcmd "github.com/Jisin0/filmigo/cmd/imdb"
	justwatchcmd "github.com/Jisin0/filmigo/cmd/justwatch"
	omdbcmd "github.com/Jisin0/filmigo/cmd/omdb"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "filmigo",
		Short: "Tool For Getting Data From Various Movie Databases",
		Long: `Filmigo is a CLI tool and Library to Browse Movie Databases.
It can get data about Movies, Shows and Actors.
use "filmigo sites" to get a list of all supported sites.`,
	}
)

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(imdbcmd.NewCommand())
	rootCmd.AddCommand(omdbcmd.NewCommand())
	rootCmd.AddCommand(justwatchcmd.NewCommand())
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
