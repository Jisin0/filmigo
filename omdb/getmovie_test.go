package omdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/omdb"
)

const (
	oppenheimerId = "tt15398776"
)

func TestGetMovie(t *testing.T) {
	r, e := client.GetMovie(&omdb.GetMovieOpts{Id: oppenheimerId})
	if e != nil {
		t.Error(e)
	}

	r.PrettyPrint()
}
