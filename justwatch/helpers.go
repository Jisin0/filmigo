// (c) Jisin0
// Helper methods and types.

package justwatch

import (
	"strings"
)

const (
	imageBaseURL       = "https://images.justwatch.com"
	posterSizeLarge    = "s592"
	posterSizeMedium   = "s332"
	posterSizeSmall    = "s166"
	backdropSizeLarge  = "s1920"
	backdropSizeMedium = "s1440"
)

// Poster url format.
type PosterURL string

// Get the full image url for the maximum size of the image.
func (p PosterURL) FullURL() string {
	// trim off the extension.
	poster := strings.TrimSuffix(string(p), ".{format}")
	return imageBaseURL + strings.Replace(poster, "{profile}", posterSizeLarge, 1)
}

// Get the full image url for a thumbnail size of the image.
func (p PosterURL) ThumbURL() string {
	// trim off the extension.
	poster := strings.TrimSuffix(string(p), ".{format}")
	return imageBaseURL + strings.Replace(poster, "{profile}", posterSizeSmall, 1)
}

// Backdrop image data.
// Use the FullURL() method to get the full formatted url.
type Backdrop struct {
	// File path format to the backdrop image.
	BackdropURLFormat string `json:"backdropURL"`
}

// Get the full image url for the maximum size of the image.
func (b Backdrop) FullURL() string {
	// trim off the extension.
	url := strings.TrimSuffix(b.BackdropURLFormat, ".{format}")
	return imageBaseURL + strings.Replace(url, "{profile}", backdropSizeLarge, 1)
}

// Full name for shortname of genres from the api.
var genreFullNames map[string]string = map[string]string{
	"cmy": "Comedy",
	"act": "Action",
	"trl": "Thriller",
	"scf": "Sci-Fi",
	"ani": "Animation",
	"drm": "Drama",
	"crm": "Crime",
	"rma": "Romance",
	"doc": "Documentary",
	"eur": "Europe",
	"hrr": "Horror",
	"fml": "Family",
	"fnt": "Fantasy",
	"hst": "History",
	"war": "War",
	"msc": "Musical",
	"spt": "Sport",
}

// Returns the full genre name of the input short name.
func getFullGenre(s string) string {
	v, k := genreFullNames[s]
	if k {
		return v
	}

	return s
}

// Raw genres returned from api.
type Genres []Genre

// ToList returns a slice with the full name for each genre.
func (gs *Genres) ToList() []string {
	var a []string
	for _, g := range *gs {
		a = append(a, g.FullName())
	}

	return a
}

// ToShortList returns a slice with the shortcodes of each genre.
func (gs *Genres) ToShortList() []string {
	var a []string
	for _, g := range *gs {
		a = append(a, g.ShortName)
	}

	return a
}

// ToString generates a string with the full name sof genres separated with given separator.
func (gs *Genres) ToString(sep string) string {
	l := gs.ToList()
	return strings.Join(l, sep)
}

// Raw genre from justwatch with only the shortname.
// Use FullName() to get the full name for ex: cmy -> Comedy
type Genre struct {
	// Shortcode for a genre returned from the api.
	// For ex: cmy for Comedy.
	ShortName string `json:"shortName"`

	// Full name of the movie only populated on calling FullName() and stored.
	fullName string
}

// Returns the full name for for the shortcode returned from the api.
// For ex: shortcode cmy returns Comedy.
func (g *Genre) FullName() string {
	if g.fullName != "" {
		return g.fullName
	}

	s := getFullGenre(g.ShortName)
	g.fullName = s

	return s
}

// Returns a full list with the combined results of Flatrate, Buy, Rent, Free and Fast offers.
func (r *GetTitleOffersResult) MergeOffers() []*Offer {
	var l []*Offer
	l = append(l, r.Flatrate...)
	l = append(l, r.Buy...)
	l = append(l, r.Rent...)
	l = append(l, r.Free...)
	l = append(l, r.Fast...)

	return l
}
