package imdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/imdb"
)

const (
	cillianMurphyID = "nm0614165"
)

func TestGetPerson(t *testing.T) {
	c := imdb.NewClient()
	res, err := c.GetPerson(cillianMurphyID)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)

}
