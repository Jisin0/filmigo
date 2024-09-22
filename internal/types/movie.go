// (c) Jisin0
// Interface for movies from any package.

package types

// Movie is a movie from any source supported by the library.
//
// Supported Types:
//
//	switch m.(type) {
//	case *imdb.Movie:
//	case *justwatch.Title
//	case *omdb.Movie
//	}
type Movie interface {
	GetType() string
	GetID() string
	GetURL() string
	GetTitle() string
	GetPosterURL() string
	GetGenres() []string
	GetPlot() string
	PrettyPrint()
}
