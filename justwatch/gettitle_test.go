package justwatch_test

import (
	"testing"

	"github.com/Jisin0/filmigo"
	"github.com/Jisin0/filmigo/justwatch"
)

const (
	rickAndMortyId = "ts20233"
)

func TestGetTitleFromUrl(t *testing.T) {
	r, e := justwatch.GetTitleFromUrl("justwatch.com/US/tv-show/rick-and-morty")
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
}

func TestGetTitle(t *testing.T) {
	r, e := justwatch.GetTitle(rickAndMortyId)
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
}

func TestGetTitleOffers(t *testing.T) {
	r, e := justwatch.GetTitleOffers(rickAndMortyId)
	if e != nil {
		t.Error(e)
		return
	}

	filmigo.PrintJson(r, "   ")
}
