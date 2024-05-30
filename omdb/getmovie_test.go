package omdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/omdb"
)

const (
	oppenheimerID = "tt15398776"
)

func TestGetMovie(t *testing.T) {
	r, e := client.GetMovie(&omdb.GetMovieOpts{ID: oppenheimerID})
	if e != nil {
		t.Error(e)
	}

	r.PrettyPrint()
}
