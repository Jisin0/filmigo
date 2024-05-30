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
	// write your api key to apikey.txt to run tests
	r, err := os.ReadFile("apikey.txt")
	if err != nil {
		panic("failed to open apikey file")
	}

	apiKey = string(r)

	client = omdb.NewClient(apiKey)
}
