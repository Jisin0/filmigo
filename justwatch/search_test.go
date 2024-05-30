package justwatch_test

import (
	"fmt"
	"testing"

	"github.com/Jisin0/filmigo"
	"github.com/Jisin0/filmigo/justwatch"
)

var client *justwatch.JustwatchClient = justwatch.NewClient()

func TestSearch(t *testing.T) {
	r, e := client.SearchTitle("50 shade")
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
	fmt.Println(r.Results[0].Genres.ToString(", "))
	fmt.Println(r.Results[0].Backdrops[0].FullUrl())
	fmt.Println(r.Results[0].Poster.FullUrl())
}
