package justwatch_test

import (
	"testing"
)

const (
	rickAndMortyID = "ts20233"
)

func TestGetTitleFromURL(t *testing.T) {
	r, e := client.GetTitleFromURL("justwatch.com/US/tv-show/rick-and-morty")
	if e != nil || r == nil {
		t.Error(e)
		return
	}
}

func TestGetTitle(t *testing.T) {
	r, e := client.GetTitle(rickAndMortyID)
	if e != nil {
		t.Error(e)
		return
	}

	r.PrettyPrint()
}

func TestGetTitleOffers(t *testing.T) {
	r, e := client.GetTitleOffers(rickAndMortyID)
	if e != nil || r == nil {
		t.Error(e)
		return
	}
}
