package cmd

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/Jisin0/filmigo/internal/types"
	"github.com/Jisin0/filmigo/justwatch"
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

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Full Data of a Movie",
		Long:  `Get Movie Using It's imdb/justwatch id or URL.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func init() {
	getCmd.Flags().BoolVar(&useOmdb, "omdb", false, "omdb engine")
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
		result, err = imdb.NewClient().GetMovie(id)
	case justwatchIdRegex.MatchString(id):
		result, err = justwatch.NewClient().GetTitle(id)
	case strings.Contains(id, "justwatch.com"):
		r, e := justwatch.NewClient().GetTitleFromURL(id)
		result, err = r.Data, e
	}

	if err != nil {
		return err
	}

	result.PrettyPrint()

	return nil
}
