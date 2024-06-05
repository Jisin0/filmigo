package justwatch_test

import (
	"fmt"
	"testing"

	"github.com/Jisin0/filmigo/justwatch"
)

var client *justwatch.JustwatchClient = justwatch.NewClient()

const searchQuery = "Inception"

func TestSearch(t *testing.T) {
	r, e := client.SearchTitle(searchQuery)
	if e != nil {
		t.Error(e)
		return
	}

	fmt.Println(r.Results[0].Genres.ToString(", "))
	fmt.Println(r.Results[0].Backdrops[0].FullURL())
	fmt.Println(r.Results[0].Poster.FullURL())
}

func TestFullTitle(t *testing.T) {
	r, e := client.SearchTitle(searchQuery)
	if e != nil {
		t.Error(e)
		return
	}

	res, err := r.Results[0].FullTitle(client)
	if err != nil || res == nil {
		t.Error(err)
		return
	}
}
