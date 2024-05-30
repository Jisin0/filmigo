package justwatch_test

import (
	"testing"

	"github.com/Jisin0/filmigo"
)

const (
	rickAndMortyId = "ts20233"
)

func TestGetTitleFromUrl(t *testing.T) {
	r, e := client.GetTitleFromUrl("justwatch.com/US/tv-show/rick-and-morty")
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
}

func TestGetTitle(t *testing.T) {
	r, e := client.GetTitle(rickAndMortyId)
	if e != nil {
		t.Error(e)
		return
	}

	r.PrettyPrint()
}

func TestGetTitleOffers(t *testing.T) {
	r, e := client.GetTitleOffers(rickAndMortyId)
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
}
