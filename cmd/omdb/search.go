package omdbcmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Jisin0/filmigo/omdb"
	"github.com/spf13/cobra"
)

var (
	searchType string
	searchYear string
	searchPage int
	searchCmd  = &cobra.Command{
		Use:   "search",
		Short: "Search For a Movie or Show",
		Long:  `Search Movies and Shows Using Their Name or ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runSearch,
	}
)

func init() {
	searchCmd.Flags().StringVar(&searchType, "type", "", "title type: movie, series or episode")
	searchCmd.Flags().StringVar(&searchYear, "year", "", "release year of movie")
	searchCmd.Flags().IntVar(&searchPage, "page", 0, "result page to return")
}

func runSearch(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := omdb.NewClient(apiKey, omdb.OmdbClientOpts{DisableCaching: noCache, CacheExpiration: time.Minute * time.Duration(cacheExpiration)})

	result, err := client.Search(input, &omdb.SearchOpts{Type: searchType, Year: searchYear, Page: searchPage})
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
