package omdb_test

import (
	"fmt"
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
		fmt.Println("OMDB_API_KEY not set skipping test")
	}

	apiKey = s

	client = omdb.NewClient(apiKey)
}
