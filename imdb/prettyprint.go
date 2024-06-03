// (c) Jisin0
// Pretty print for justwatch types.

package imdb

import "fmt"

// PrettyPrint prints out movie data in a neat interface.
func (t *Movie) PrettyPrint() {
	fmt.Printf("%s [%s]", t.Title, t.Year)

	if t.Aka != t.Title {
		fmt.Printf("  (aka : %s)\n", t.Aka)
	}

	if t.Rating.Value != 0 {
		fmt.Printf("⭐%v | %v❤️", t.Rating, t.Rating.Votes)
	}

	fmt.Print("\n\n", t.Plot, "\n\n")

	fmt.Printf("ID: %s\n", t.ID)

	if t.Releaseinfo != "" {
		fmt.Printf("Released: %s\n", t.Releaseinfo)
	}

	if t.Runtime != "" {
		fmt.Printf("Runtime: %s\n", t.Runtime)
	}

	if len(t.Genres) > 0 {
		fmt.Printf("Genres: %s\n", t.Genres.ToString(", "))
	}
}
