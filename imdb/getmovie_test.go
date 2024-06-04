package imdb_test

import (
	"testing"
)

const (
	oppenheimerID  = "tt15398776"
	invalidMovieID = "invalidID123"
	anotherValidID = "tt0111161" // shawshank redemption
)

func TestGetMovie(t *testing.T) {
	movieIDs := []string{
		oppenheimerID,
		invalidMovieID,
		anotherValidID,
	}

	for _, id := range movieIDs {
		t.Run(id, func(t *testing.T) {
			res, err := c.GetMovie(id)
			if err != nil {
				t.Logf("Expected error for ID %s: %v", id, err)

				if id != invalidMovieID {
					t.Errorf("Unexpected error for valid ID %s: %v", id, err)
				}
			} else {
				if id == invalidMovieID {
					t.Errorf("Expected error for invalid ID %s, but got result: %v", id, res)
				} else {
					t.Logf("Got movie data for ID %s", id)
				}
			}
		})
	}
}
