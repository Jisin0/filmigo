package imdbcmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/spf13/cobra"
)

var (
	includeVideos bool
	onlyTitles    bool
	onlyNames     bool
	searchCmd     = &cobra.Command{
		Use:   "search",
		Short: "Search For a Movie or Person",
		Long:  `Search Both Movies and Stars Using Their Name or ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runSearch,
	}
)

func init() {
	searchCmd.Flags().BoolVar(&includeVideos, "include-videos", false, "include trailer videos")
	searchCmd.Flags().BoolVar(&onlyTitles, "only-titles", false, "only search titles (movies & shows)")
	searchCmd.Flags().BoolVar(&onlyNames, "only-names", false, "only search people")
}

func runSearch(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := imdb.NewClient(imdb.ImdbClientOpts{DisableCaching: noCache, CacheExpiration: time.Minute * time.Duration(cacheExpiration)})

	var (
		result *imdb.SearchResults
		err    error
	)

	switch {
	case onlyTitles:
		result, err = client.SearchTitles(input, &imdb.SearchConfigs{IncludeVideos: includeVideos})
	case onlyNames:
		result, err = client.SearchNames(input, &imdb.SearchConfigs{IncludeVideos: includeVideos})
	default:
		result, err = client.SearchAll(input, &imdb.SearchConfigs{IncludeVideos: includeVideos})
	}

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
