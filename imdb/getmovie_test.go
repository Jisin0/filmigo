package imdb_test

import (
	"testing"
)

const (
	oppenheimerID = "tt15398776"
)

func TestGetMovie(t *testing.T) {
	res, err := c.GetMovie(oppenheimerID)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	res.PrettyPrint()
}
