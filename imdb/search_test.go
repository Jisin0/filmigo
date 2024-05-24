package imdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/imdb"
)

var searchQuery string = "mad"

func TestSearchTitles(t *testing.T) {
	res, err := c.SearchTitles(searchQuery, &imdb.SearchConfigs{IncludeVideos: true})
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}

func TestSearchAll(t *testing.T) {
	res, err := c.SearchAll(searchQuery, &imdb.SearchConfigs{IncludeVideos: true})
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}

func TestSearchNames(t *testing.T) {
	res, err := c.SearchNames(searchQuery, &imdb.SearchConfigs{IncludeVideos: true})
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}
