package imdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/imdb"
)

const (
	oppenheimerId = "tt15398776"
)

func TestGetMovie(t *testing.T) {
	c := imdb.NewClient()
	res, err := c.GetMovie(oppenheimerId)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	res.PrettyPrint()

}
