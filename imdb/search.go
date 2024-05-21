// (c) Jisin0
// Search via the official api.
package imdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-faster/errors"
)

const (
	titleSearchURL = "https://v3.sg.media-imdb.com/suggestion/titles/%v/%v.json"
	allSearchURL   = "https://v3.sg.media-imdb.com/suggestion/%v/%v.json"
)

// Search Results returned from eith SearchTitles or SearchAll methods.
type SearchResults struct {
	// List of results.
	Results []SearchResult `json:"d"`
	// The query string.
	Query string `json:"q"`
	// Unknown. Could be some version code or used for pagination, every result has this field.
	V int `json:"v"`
}

// Data obtained from searching using the api. Could be data on a movie/show/person or sometimes an ad when using SarchAll.
type SearchResult struct {

	// An image commonly a movie poster or actor's picture.
	Image Image `json:"i"`
	// Id of the movie/show/person or the url path for ads.
	Id string `json:"id"`
	// Header or main text of a search result. A movie/show/person's name.
	Title string `json:"l"`
	// For movies or shows, A string of the type os title for ex: TV Series, Movie etc.
	Subtitle string `json:"q"`
	//The category of the movie/show. Empty for people.
	//Possible values : movie, tvSeries, tvMiniSeries
	Category string `json:"qid"`
	//A rank point.
	Rank int `json:"rank"`
	//The main stars of a movie/show, or a notable work in case of a person.
	Description string `json:"s"`
	//Year of release of a movie/show
	Year int `json:"y"`
	//A string indicating the years in which a tv series was released. for ex: 2016-2025
	Years string `json:"yr"`
	//A list of videos related to the title or person.
	Videos []Video `json:"v"`
}

type Image struct {
	Height   int    `json:"height"`
	ImageUrl string `json:"imageUrl"`
	Width    int    `json:"width"`
}

type Video struct {
	Thumbnail Image  `json:"i"`
	Id        string `json:"id"`
	Title     string `json:"l"`
	Duration  string `json:"s"`
}

// Optional parameters to be passed to the search query.
type SearchConfigs struct {
	//Set true for the api to return video details (trailers, previews etc.).
	//If enabled you will get a thumbnail of the video and the video id.
	IncludeVideos bool
}

var ErrNoResults error = errors.New("no search results were found")

// Search for only movies/shows excluding people or other types.
//
// - query (string) - The query or keyword to search for.
// - configs (optional) - Additional request configs.
func SearchTitles(query string, c ...SearchConfigs) (*SearchResults, error) {
	if len(query) < 1 {
		return nil, errors.New("searchtitles: query too short")
	}

	url := fmt.Sprintf(titleSearchURL, query[0:1], query)

	if len(c) > 0 && c[0].IncludeVideos {
		url += "?includeVideos=1"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "searchtitles: failed to create request")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:107.0) Gecko/20100101 Firefox/107.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "searchtitles: failed to create request")
	}

	defer resp.Body.Close()

	var results SearchResults
	json.NewDecoder(resp.Body).Decode(&results)
	if len(results.Results) < 1 {
		return &results, ErrNoResults
	}

	return &results, nil

}
