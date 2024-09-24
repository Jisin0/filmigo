// (c) Jisin0
// Pretty print for justwatch types.

package omdb

import (
	"fmt"
)

const (
	// default string returned if empty
	notAvailable = "N/A"
)

// PrettyPrint prints out movie data in a neat interface.
func (m *Movie) PrettyPrint() {
	fmt.Printf("Id: %s\n", m.ImdbID)
	fmt.Printf("Type: %s\n", m.Type)
	fmt.Printf("Title: %s", m.Title)
	fmt.Printf("\nYear: %s", m.Year)

	if m.ImdbRating != notAvailable {
		fmt.Printf("\nRating: %s\nVotes: %s", m.ImdbRating, m.ImdbVotes)
	}

	if m.Released != notAvailable {
		fmt.Printf("\nReleased: %s", m.Released)
	}

	if m.Country != notAvailable {
		fmt.Printf("\nCountry: %s", m.Country)
	}

	if m.Rated != notAvailable {
		fmt.Printf("\nContent Rated: %s", m.Rated)
	}

	if m.Runtime != notAvailable {
		fmt.Printf("\nRuntime: %s", m.Runtime)
	}

	if m.Genres != notAvailable {
		fmt.Printf("\nGenres: %s", m.Genres)
	}

	if m.Languages != notAvailable {
		fmt.Printf("\nLanguages: %s", m.Languages)
	}

	if m.DVD != notAvailable {
		fmt.Printf("\nDVD Release: %s", m.DVD)
	}

	if m.Actors != notAvailable {
		fmt.Printf("\nActors: %s", m.Actors)
	}

	if m.Director != notAvailable {
		fmt.Printf("\nDirectors: %s", m.Director)
	}

	if m.Writers != notAvailable {
		fmt.Printf("\nWriters: %s", m.Writers)
	}

	if m.Production != notAvailable {
		fmt.Printf("\nProducers: %s", m.Production)
	}

	if m.Poster != notAvailable {
		fmt.Printf("\nPoster: %s", m.Poster)
	}

	if m.Awards != notAvailable {
		fmt.Printf("\nAwards: %s", m.Awards)
	}

	if m.BoxOffice != notAvailable {
		fmt.Printf("\nIncome: %s", m.BoxOffice)
	}

	if m.Website != notAvailable {
		fmt.Printf("\nWebsite: %s", m.Website)
	}

	if m.Plot != notAvailable {
		fmt.Printf("\nPlot: %s", m.Plot)
	}

	fmt.Println()
}
