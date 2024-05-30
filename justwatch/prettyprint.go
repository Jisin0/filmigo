// (c) Jisin0
// Pretty print for justwatch types.

package justwatch

import "fmt"

// PrettyPrint prints out movie data in a neat interface.
func (t *Title) PrettyPrint() {
	content := t.Content

	if content == nil {
		fmt.Println("no content found !")
		return
	}

	fmt.Print(content.Title)

	if content.ReleaseYear != 0 {
		fmt.Printf(" (%v)", content.ReleaseYear)
	}

	if content.AgeCertification != "" {
		fmt.Printf("[%s Rated]", content.AgeCertification)
	}

	fmt.Println()

	if content.OriginalTitle != content.Title {
		fmt.Printf("  aka : %s\n", content.OriginalTitle)
	}

	titleType := t.Type

	if content.Scores != nil {
		fmt.Printf("⭐%v/10 | %.1f%%❤️", content.Scores.ImdbRating, content.Scores.JustwatchRating*100)
	}

	if content.Interactions != nil {
		fmt.Printf("   👍%v | %v👎", content.Interactions.Likes, content.Interactions.Dislikes)
	}

	fmt.Print("\n\n", content.Description, "\n\n")

	fmt.Printf("ID: %s\n", t.ID)

	if content.ReleaseDate != "" {
		fmt.Printf("Released: %s\n", content.ReleaseDate)
	}

	fmt.Printf("Type: %s\n", titleType)

	if content.Runtime != 0 {
		fmt.Printf("Runtime: %vmins\n", content.Runtime)
	}

	if len(*content.Genres) > 0 {
		fmt.Printf("Genres: %s\n", content.Genres.ToString(", "))
	}

	if titleType == "SHOW" {
		fmt.Printf("Seasons: %v\n", t.TotalSeasonCount)
	} else if titleType == "SHOW_SEASON" {
		fmt.Printf("Show: %s\n", t.Show.Content.Title)
		fmt.Printf("Episodes: %v\n", t.TotalEpisodeCount)
	}
}
