// (c) Jisin0
// Graphql queries for each operetion.

package justwatch

import (
	"log"
	"os"
	"path/filepath"
)

// Graphql queries
var searchTitleQuery string
var getTitleQuery string
var getTitleFromURLQuery string
var getTitleOffersQuery string

// Initialize graphql queries
func init() {
	var err error

	getTitleFromURLQuery, err = loadQuery("./queries/gettitleurl.graphql")
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

	query, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(query), nil
}
