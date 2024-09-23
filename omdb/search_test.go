package omdb_test

import (
	"testing"

	"github.com/Jisin0/filmigo/omdb"
)

const (
	query = "mad"
)

func TestSearch(t *testing.T) {
	if client == nil {
		return
	}

	r, e := client.Search(query, &omdb.SearchOpts{Type: omdb.ResultTypeMovie})
	if e != nil {
		t.Error(e)
		return
	}

	t.Logf("%+v, %v", r, r.Results[0].Title)

	r, err := r.NextPage(client)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v, %v", r, r.Results[0].Title)
}

func TestGetFull(t *testing.T) {
	if client == nil {
		return
	}

	r, e := client.Search(query, &omdb.SearchOpts{Type: omdb.ResultTypeMovie})
	if e != nil {
		t.Error(e)
		return
	}

	res, err := r.Results[0].GetFull(client)
	if err != nil || res == nil {
		t.Error(err)
		return
	}
}
