package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/spf13/cobra"
)

var (
	// Search Command Flags.
	searchMethod string
	//userLicense string

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search for Movies or Actors",
		Long:  `Search Using any Supported Engine, Defaults to IMDb`,
		Args:  cobra.ExactArgs(1),
		RunE:  runSearch,
	}
)

func init() {
	searchCmd.PersistentFlags().StringVar(&searchMethod, "method", "", "search method could be imdb, justwatch or omdb (default is imdb)")
}

func runSearch(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	result, err := imdb.NewClient().SearchAll(args[0])
	if err != nil {
		return err
	}

	r, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(r))

	return nil
}
