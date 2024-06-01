package omdb_test

import (
	"os"

	"github.com/Jisin0/filmigo/omdb"
)

var (
	apiKey string
	client *omdb.OmdbClient
)

func init() {
	// Set your omdb api key as environment var to run tests
	s := os.Getenv("OMDB_API_KEY")
	if s == "" {
		panic("OMDB_API_KEY not set")
	}

	apiKey = s

	client = omdb.NewClient(apiKey)
}
