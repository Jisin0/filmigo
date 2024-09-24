package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/Jisin0/filmigo/internal/types"
	"github.com/Jisin0/filmigo/justwatch"
	"github.com/Jisin0/filmigo/omdb"
	"github.com/spf13/cobra"
)

// Regexes for each engine
var (
	imdbIdRegex      = regexp.MustCompile(`tt\d+`)
	justwatchIdRegex = regexp.MustCompile(`tm\d+`)
)

var (
	useOmdb    bool
	omdbApiKey string
	outputJson bool

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Full Data of a Movie",
		Long:  `Get Movie Using It's imdb/justwatch id or URL.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func init() {
	getCmd.Flags().BoolVar(&useOmdb, "omdb", false, "use omdb engine")
	getCmd.Flags().BoolVar(&outputJson, "json", false, "output result as json")
	getCmd.Flags().StringVar(&omdbApiKey, "apikey", "", "omdb api key")
}

func runGet(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	var (
		id     = args[0]
		result types.Movie
		err    error
	)

	switch {
	case imdbIdRegex.MatchString(id):
		if useOmdb {
			result, err = omdb.NewClient(omdbApiKey).GetMovie(&omdb.GetMovieOpts{ID: id})
		} else {
			result, err = imdb.NewClient().GetMovie(id)
		}
	case justwatchIdRegex.MatchString(id):
		result, err = justwatch.NewClient().GetTitle(id)
	case strings.Contains(id, "justwatch.com"):
		r, e := justwatch.NewClient().GetTitleFromURL(id)
		result, err = r.Data, e
	case useOmdb: // title search if omdb enabled and doesn't match any id
		result, err = omdb.NewClient(omdbApiKey).GetMovie(&omdb.GetMovieOpts{Title: id})
	}

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
