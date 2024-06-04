package imdb_test

import (
	"fmt"
	"testing"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/Jisin0/filmigo/imdb/constants"
)

func TestAdvancedSearchTitle(t *testing.T) {
	testData := []imdb.AdvancedSearchTitleOpts{
		{Genres: []string{constants.TitleGenreAction}, ExtraParams: map[string]any{"plot": "guns"}},
		{CastOrCrew: []string{cillianMurphyID}},
		{TitleName: "aksjgka"}, // bad data
	}

	lastIndex := len(testData) - 1 // item at this index should fail.

	for i, o := range testData {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			_, err := c.AdvancedSearchTitle(&o)
			if err != nil {
				// if item is last error is expected.
				if i == lastIndex {
					t.Log("error as expected")
				} else {
					t.Errorf("item %v returned unexpected error %v", i, err)
				}
			} else {
				if i == lastIndex {
					t.Errorf("error expected for item %v but results found", i)
				} else {
					t.Logf("item %v succesfully returned", i)
				}
			}
		})

	}
}

func TestAdvancedSearchName(t *testing.T) {
	testData := []imdb.AdvancedSearchNameOpts{
		{Titles: []string{oppenheimerID}},
		{Awards: []string{constants.NameAwardBestActressNominated}},
		{Name: "shkjag"}, // should fail
	}

	lastIndex := len(testData) - 1 // item at this index should fail.

	for i, o := range testData {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			_, err := c.AdvancedSearchName(&o)
			if err != nil {
				// if item is last error is expected.
				if i == lastIndex {
					t.Log("error as expected")
				} else {
					t.Errorf("item %v returned unexpected error %v", i, err)
				}
			} else {
				if i == lastIndex {
					t.Errorf("error expected for item %v but results found", i)
				} else {
					t.Logf("item %v succesfully returned", i)
				}
			}
		})

	}
}
