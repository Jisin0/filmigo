// (c) Jisin0
// Pretty print for justwatch types.

package imdb

import (
	"fmt"
	"html"
	"strings"
)

// PrettyPrint prints out movie data in a neat interface.
func (m *Movie) PrettyPrint() {
	fmt.Printf("Id: %s\n", m.ID)
	fmt.Printf("Type: %s\n", m.Type)
	fmt.Printf("Title: %s", m.Title)
	fmt.Printf("\nYear: %s", m.ReleaseYear)

	if m.Aka != m.Title {
		fmt.Printf("\nAka : %s", m.Aka)
	}

	if m.Rating.Value != 0 {
		fmt.Printf("\nRating: %v\nVotes: %v", m.Rating.Value, m.Rating.Votes)
	}

	if m.Releaseinfo != "" {
		fmt.Printf("\nReleased: %s", m.Releaseinfo)
	}

	if len(m.Languages) > 0 {
		fmt.Printf("\nLanguages: %s", m.Languages.ToString(", "))
	}

	if m.Runtime != "" {
		fmt.Printf("\nRuntime: %s", m.Runtime)
	}

	if len(m.Genres) > 0 {
		fmt.Printf("\nGenres: %s", strings.Join(m.Genres, ", "))
	}

	if m.Keywords != "" {
		fmt.Printf("\nKeywords: %s", m.Keywords)
	}

	if len(m.Actors) > 0 {
		fmt.Printf("\nActors: %s", m.Actors.ToString(", "))
	}

	if len(m.Directors) > 0 {
		fmt.Printf("\nDirectors: %s", m.Directors.ToString(", "))
	}

	if len(m.Writers) > 0 {
		fmt.Printf("\nWriters: %s", m.Writers.ToString(", "))
	}

	if len(m.Locations) > 0 {
		fmt.Printf("\nLocations: %s", m.Locations.ToString(", "))
	}

	if len(m.Producers) > 0 {
		fmt.Printf("\nProducers: %s", m.Producers.ToString(", "))
	}

	if m.PosterURL != "" {
		fmt.Printf("\nPoster: %s", m.PosterURL)
	}

	if m.Plot != "" {
		fmt.Printf("\nPlot: %s", html.UnescapeString(m.Plot))
	}

	fmt.Println()
}
