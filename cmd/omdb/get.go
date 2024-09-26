package omdbcmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Jisin0/filmigo/internal/id"
	"github.com/Jisin0/filmigo/omdb"
	"github.com/spf13/cobra"
)

var (
	getByTitle bool // option to directly get by title incase id regex gives false positive
	plotType   string
	getCmd     = &cobra.Command{
		Use:   "get",
		Short: "Get Full Data of a Movie or Show",
		Long:  `Get Movie Using It's IMDb id or Name.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func init() {
	getCmd.Flags().BoolVar(&getByTitle, "title", false, "get by title without checking for id")
	getCmd.Flags().StringVar(&searchType, "type", "", "title type: movie, series or episode")
	getCmd.Flags().StringVar(&searchYear, "year", "", "release year of movie")
	getCmd.Flags().IntVar(&searchPage, "page", 0, "result page to return")
	getCmd.Flags().StringVar(&plotType, "plot", "", "plot type: short or full")
}

func runGet(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := omdb.NewClient(apiKey, omdb.OmdbClientOpts{DisableCaching: noCache, CacheExpiration: time.Minute * time.Duration(cacheExpiration)})

	opts := &omdb.GetMovieOpts{
		Type: searchType,
		Year: searchYear,
		Plot: plotType,
	}

	switch {
	case getByTitle:
		opts.Title = input
	case id.IsIMDbTitle(input):
		opts.ID = input
	default:
		opts.Title = input
	}

	result, err := client.GetMovie(opts)
	if err != nil {
		return err
	}

	if outputJson {
		bytes, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))
	} else {
		result.PrettyPrint()
	}

	return nil
}
