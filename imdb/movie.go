// (c) Jisin0

package imdb

import "github.com/Jisin0/filmigo/internal/types"

type Movie struct {
	// Years of release of the movie, A range for shows over multiple years.
	ReleaseYear string `json:"year"`
	MovieJSONContent
	MovieDetailsSection
}

// Data scraped from the details section of a movie using xpath.
type MovieDetailsSection struct {
	// A string with details about the release including date and country
	Releaseinfo string `xpath:"//li[@data-testid='title-details-releasedate']/div//a" json:"release_info"`
	// Origin of release, commonly the country
	Origin string `xpath:"//li[@data-testid='title-details-origin']/div//a" json:"origin"`
	// Official sites related to the movie/show
	OfficialSites types.Links `xpath:"//li[@data-testid='details-officialsites']/div/ul" json:"official_sites"`
	// Languages in which the movie/show is available in
	Languages types.Links `xpath:"//li[@data-testid='title-details-languages']/div/ul" json:"languages"`
	// Any alternative name of the movie.
	Aka string `xpath:"//li[@data-testid='title-details-akas']//span" json:"aka"`
	// Locations at which the movie/show was filmed at
	Locations types.Links `xpath:"//li[@data-testid='title-details-filminglocations']/div/ul" json:"locations"`
	// Companies which produced the movie
	Companies types.Links `xpath:"//li[@data-testid='title-details-companies']/div/ul" json:"companies"`
}

// Data scraped from the json attached in the script tag.
type MovieJSONContent struct {
	// Type of the title possibble values are Movie, TVSeries etc.
	Type string `json:"@type"`
	// ID of the movie
	ID string `json:"id"`
	// Link to the movie
	URL string `json:"url"`
	// Full title of the movie
	Title string `json:"name"`
	// Url of the full size poster image.
	PosterURL string `json:"image"`
	// Content rating class (currently undocumented).
	ContentRating string `json:"contentRating"`
	// Date the movie was released on in yyyy-mm-dd format.
	ReleaseDate string `json:"datePublished"`
	// Keywords associated with the movie in a comma separated list.
	Keywords string `json:"keywords"`
	// Ratings for the movie.
	Rating Rating `json:"aggregateRating"`
	// The directors of the movie
	Directors types.Links `json:"director"`
	// The writers of the movie
	Writers types.Links `json:"creator"`
	// The main stars of the movie
	Actors types.Links `json:"actor"`
	// Genres of the movie
	Genres []string `json:"genre"`
	// A short plot of the movie in a few lines
	Plot string `json:"description"`
	// Trailer video for the movie or show.
	Trailer VideoObject `json:"trailer,omitempty"`
	// Runtime of the move
	Runtime string `json:"duration"`
	// Tope review of the movie.
	Review Review `json:"review,omitempty"`
}

// Ensure *imdb.Movie satisfies shared movie interface
var _ types.Movie = (*Movie)(nil)

func (m *Movie) GetType() string {
	return m.Type
}

func (m *Movie) GetID() string {
	return m.ID
}

func (m *Movie) GetURL() string {
	return m.URL
}

func (m *Movie) GetTitle() string {
	return m.Title
}

func (m *Movie) GetPosterURL() string {
	return m.PosterURL
}

func (m *Movie) GetGenres() []string {
	return m.Genres
}

func (m *Movie) GetPlot() string {
	return m.Plot
}
