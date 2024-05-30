// (c) Jisin0
// Pretty print for justwatch types.

package omdb

import "fmt"

// PrettyPrint prints out movie data in a neat interface.
func (t *Movie) PrettyPrint() {
	fmt.Printf("%s(%s)", t.Title, t.Year)

	if t.Rated != "" {
		fmt.Printf(" [%s Rated]", t.Rated)
	}

	fmt.Println()

	if t.ImdbRating != "" {
		fmt.Printf("⭐%s | %s❤️", t.ImdbRating, t.ImdbVotes)
	}

	fmt.Print("\n\n", t.Plot, "\n\n")

	if t.Released != "" {
		fmt.Printf("Released: %s\n", t.Released)
	}

	if t.Runtime != "" {
		fmt.Printf("Runtime: %s\n", t.Runtime)
	}

	if len(t.Genres) > 0 {
		fmt.Printf("Genres: %s\n", t.Genres)
	}

}
