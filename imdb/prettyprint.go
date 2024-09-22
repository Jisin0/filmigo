// (c) Jisin0
// Pretty print for justwatch types.

package imdb

import (
	"fmt"
	"strings"
)

// PrettyPrint prints out movie data in a neat interface.
func (m *Movie) PrettyPrint() {
	fmt.Printf("%s [%s]", m.Title, m.ReleaseYear)

	if m.Aka != m.Title {
		fmt.Printf("  (aka : %s)\n", m.Aka)
	}

	if m.Rating.Value != 0 {
		fmt.Printf("⭐%v | %v❤️", m.Rating, m.Rating.Votes)
	}

	fmt.Print("\n\n", m.Plot, "\n\n")

	fmt.Printf("ID: %s\n", m.ID)

	if m.Releaseinfo != "" {
		fmt.Printf("Released: %s\n", m.Releaseinfo)
	}

	if m.Runtime != "" {
		fmt.Printf("Runtime: %s\n", m.Runtime)
	}

	if len(m.Genres) > 0 {
		fmt.Printf("Genres: %s\n", strings.Join(m.Genres, ", "))
	}
}
