// (c) Jisin0
// Graphql queries for each operetion.

package justwatch

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// Graphql queries
var searchTitleQuery string
var getTitleQuery string
var getTitleFromUrlQuery string
var getTitleOffersQuery string

func init() {
	// Initialize graphql queries

	var err error

	getTitleFromUrlQuery, err = loadQuery("./queries/gettitleurl.graphql")
	if err != nil {
		log.Println("failed to load graphql file !", err)
	}

	getTitleQuery, err = loadQuery("./queries/gettitle.graphql")
	if err != nil {
		log.Println("failed to load graphql file !", err)
	}

	getTitleOffersQuery, err = loadQuery("./queries/gettitleoffers.graphql")
	if err != nil {
		log.Println("failed to load graphql file !", err)
	}

	searchTitleQuery, err = loadQuery("./queries/searchtitle.graphql")
	if err != nil {
		log.Println("failed to load graphql file !", err)
	}
}

// Load a graphql query from file.
func loadQuery(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}
	query, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(query), nil
}
