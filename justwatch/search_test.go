package justwatch_test

import (
	"fmt"
	"testing"

	"github.com/Jisin0/filmigo"
	"github.com/Jisin0/filmigo/justwatch"
)

var client *justwatch.JustwatchClient = justwatch.NewClient()

const searchQuery = "50 shade"

func TestSearch(t *testing.T) {
	r, e := client.SearchTitle(searchQuery)
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
	fmt.Println(r.Results[0].Genres.ToString(", "))
	fmt.Println(r.Results[0].Backdrops[0].FullUrl())
	fmt.Println(r.Results[0].Poster.FullUrl())
}

func TestFullTitle(t *testing.T) {
	r, e := client.SearchTitle(searchQuery)
	if e != nil {
		t.Error(e)
		return
	}

	res, err := r.Results[0].FullTitle(client)
	if err != nil {
		t.Error(err)
		return
	}

	filmigo.PrintJson(res, "  ")
}
