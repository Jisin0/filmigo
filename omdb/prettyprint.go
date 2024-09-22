// (c) Jisin0
// Pretty print for justwatch types.

package omdb

import "fmt"

// PrettyPrint prints out movie data in a neat interface.
func (m *Movie) PrettyPrint() {
	fmt.Printf("%s(%s)", m.Title, m.Year)

	if m.Rated != "" {
		fmt.Printf(" [%s Rated]", m.Rated)
	}

	fmt.Println()

	if m.ImdbRating != "" {
		fmt.Printf("⭐%s | %s❤️", m.ImdbRating, m.ImdbVotes)
	}

	fmt.Print("\n\n", m.Plot, "\n\n")

	if m.Released != "" {
		fmt.Printf("Released: %s\n", m.Released)
	}

	if m.Runtime != "" {
		fmt.Printf("Runtime: %s\n", m.Runtime)
	}

	if m.Genres != "" {
		fmt.Printf("Genres: %s\n", m.Genres)
	}
}
