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
	// set your api key as env var to run this test
	apiKey = os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		panic("no api key set")
	}

	client = omdb.NewClient(apiKey)
}
