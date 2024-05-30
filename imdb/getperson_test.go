package imdb_test

import (
	"testing"
)

const (
	cillianMurphyID = "nm0614165"
)

func TestGetPerson(t *testing.T) {
	res, err := c.GetPerson(cillianMurphyID)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	t.Logf("%+v", res)
}
