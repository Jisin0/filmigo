package imdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/Jisin0/filmigo/imdb/constants"
)

func TestAdvancedSearchTitle(t *testing.T) {
	r, err := c.AdvancedSearchTitle(&imdb.AdvancedSearchTitleOpts{Genres: []string{constants.TitleGenreAction}})
	if err != nil {
		t.Error(err)
	}

	if len(r) > 0 {
		t.Logf("%+v", r[0])
		t.Logf("%v more results", len(r)-1)
	}
}

func TestAdvancedSearchName(t *testing.T) {
	r, err := c.AdvancedSearchName(&imdb.AdvancedSearchNameOpts{Titles: []string{oppenheimerId}})
	if err != nil {
		t.Error(err)
	}

	if len(r) > 0 {
		t.Logf("%+v", r[0])
		t.Logf("%v more results", len(r)-1)
	}
}
