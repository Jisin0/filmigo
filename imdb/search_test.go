package imdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/imdb"
)

func TestSearchTitles(t *testing.T) {
	res, err := imdb.SearchTitles("stranger")
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}

func TestSearchAll(t *testing.T) {
	res, err := imdb.SearchAll("mad")
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}
