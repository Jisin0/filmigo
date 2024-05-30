// (c) Jisin0
// Search for movies and shows using the JW graphql api.

package justwatch

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/machinebox/graphql"
)

const (
	defaultCountryCode  = "US"
	defaultLanguageCode = "en"
)

// Results from a search operation.
type SearchResults struct {
	// List of results.
	Results []struct {
		*TitlePreview `json:"node"`
	} `json:"edges"`
}

// Preview object with basic info about a movie/show from search results.
type TitlePreview struct {
	*TitlePreviewContent `json:"content"`
	// Justwatch id of the movie/show.
	ID string `json:"id"`
	// Type of title either MOVIE, SHOW or SHOW_EPISODE
	Type     string `json:"objectType"`
	ObjectID int    `json:"objectId"`
}

type TitlePreviewContent struct {
	// Indicates wether the title is released.
	IsReleased bool `json:"isReleased"`
	// Year of release of the movie/show.
	OriginalReleaseYear int `json:"originalReleaseYear"`
	// URL path of the movie/show.
	Path string `json:"fullPath"`
	// Title of the movie/show.
	Title string `json:"title"`
	// Original title of the movie/show.
	OriginalTitle string `json:"originalTitle"`
	// A short description of the movie/show.
	ShortDescription string `json:"shortDescription"`
	// URL template for poster images.
	// Use the FullURL() or ThumbURL() methods to get full urls.
	Poster PosterURL `json:"posterUrl"`
	// Raw genres types obtained from juswatch.
	// Use ToString()      to concatenate the full genre names into a string.
	// Use ToSlice()       to output the full genre names into a slice.
	// Use ToShortSlice()  to output the shortnames of genres to a slice.
	// Use ToShortString() to concatenate the short genre names into a string.
	Genres *Genres `json:"genres"`
	// Backdrop/Banner images for the title.
	Backdrops []Backdrop `json:"backdrops"`
}

// Options for search query.
type SearchOptions struct {
	// Maximum number of results to return.
	Limit int
	// Country code for the country on which results are based. for ex: GB for United Kingdom.
	Country string
	// Language code for results . for example en for English.
	Language string
	// Indicates wether titles without a url should not be returned.
	NoTitlesWithoutURL bool
}

// SearchTitle function searches for title with simillar title using justwatch's api.
//
// - searchQuery: Keyword or query to search for.
// - opts: Additional options for the request.
func (c *JustwatchClient) SearchTitle(searchQuery string, opts ...*SearchOptions) (*SearchResults, error) {
	var (
		limit              = 5
		country            = c.Country
		language           = c.LangCode
		noTitlesWithoutURL bool
	)

	if len(opts) > 0 {
		o := opts[0]

		if o.Limit > 0 {
			limit = o.Limit
		}

		if o.Country != "" {
			country = o.Country
		}

		if o.Language != "" {
			language = o.Language
		}

		noTitlesWithoutURL = o.NoTitlesWithoutURL
	}

	// Define the variables
	variables := map[string]interface{}{
		"country":  country,
		"language": language,
		"first":    limit,
		"filter": map[string]interface{}{
			"searchQuery":             searchQuery,
			"includeTitlesWithoutUrl": !noTitlesWithoutURL,
		},
	}

	// Make a request
	req := graphql.NewRequest(searchTitleQuery)

	// Set the variables for the request
	for key, value := range variables {
		req.Var(key, value)
	}

	// Set header fields if necessary
	req.Header.Set("Content-Type", "application/json")

	// Define a response struct to decode the response into
	var respData struct {
		S *SearchResults `json:"popularTitles"`
	}

	// Run the query
	if err := graphQLClient.Run(context.Background(), req, &respData); err != nil {
		return nil, errors.Wrap(err, "Failed to run GraphQL query: %v")
	}

	if len(respData.S.Results) < 1 {
		return respData.S, errors.Errorf("no results found for %s", searchQuery)
	}

	return respData.S, nil
}

// FullTitle fetches the full data about the title from the api by it's JW ID.
//
// - client : Justwatch client to make the query through.
func (t *TitlePreview) FullTitle(client *JustwatchClient) (*Title, error) {
	return client.GetTitle(t.ID)
}
