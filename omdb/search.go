//(c) Jisin0
// Types and functions for omdb search operations.

package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Jisin0/filmigo/encode"
	"github.com/go-faster/errors"
)

// Extra Options for an omdbapi search query.
type SearchOpts struct {

	// Type of result to return either "movie", "series" or "episode".
	// Use omdb.ResultTypeXX values for reliablility.
	Type string `url:"type"`

	// Year of release of the movie .
	Year string `url:"y"`

	// Results page to return.
	// Use value from a previous result or use the NextPage helper method.
	Page int `url:"page"`
}

// Search results returned from omdb.Search.
type SearchResult struct {

	// List of results .
	Results []*MoviePreview `json:"search"`

	// This value indicates wether a result was returned.
	// Value is True on success and False on failure (case sensitive).
	Response string `json:"response"`

	// Error returned only when query fails or no results were found.
	Error string `json:"error"`

	StrTotalResults string `json:"totalresults"`

	// total results available for a movie.
	TotalResults int

	// Current returned results page.
	Page int

	// Query or keyword that was searched.
	Query string
}

// Minimal data about a movie returned from a search query.
type MoviePreview struct {

	// Title of the movie.
	Title string `json:"title"`

	// Year of release of the movie.
	Year string `json:"year"`

	//	Imdb id of the movie for ex: tt1285016.
	ImdbId string `json:"imdbid"`

	// Type of result either "movie", "series" or "episode".
	// use omdb.ResultTypeXX values for reliability when checking.
	Type string `json:"type"`

	// Poster image url for the movie.
	Poster string `json:"poster"`
}

// GetMovie gets the full data of a movie using it's imdb id or the full name:
//
// - query : The query or keyword to search for.
// - opts :  Extra options for the request
func (c *OmdbClient) Search(query string, opts ...*SearchOpts) (*SearchResult, error) {

	if query == "" {
		return nil, errors.New("query value is empty")
	}

	if c.apiKey == "" {
		return nil, errors.New("no obdb api key provided")
	}

	var page int = 1

	var extraParams string
	if len(opts) > 0 {
		if opts[0].Page < 1 {
			opts[0].Page = 1
		} else {
			page = opts[0].Page
		}

		params, err := encode.UrlParams(*opts[0])
		if err != nil {
			return nil, errors.New("failed to parse url parameters")
		}

		extraParams = params.Encode()
	}

	fullURL := fmt.Sprintf("%s?apikey=%s&s=%s", omdbAPIURL, c.apiKey, query)

	if extraParams != "" {
		fullURL = fullURL + "&" + extraParams
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var result SearchResult

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize data")
	}

	if result.Error != "" {
		return &result, errors.New(result.Error)
	}

	total, _ := strconv.Atoi(result.StrTotalResults)

	// Page and total results are cast manually
	result.TotalResults = total
	result.Page = page
	result.Query = query

	return &result, nil
}

const (
	resultsPerPage = 10 // maximum number of results each page
)

// Returns the next page of results or returns an error if nothing was found.
func (s *SearchResult) NextPage(client *OmdbClient) (*SearchResult, error) {
	if maxPages := s.TotalResults / resultsPerPage; maxPages < s.Page {
		return nil, errors.New("no more results")
	}

	if client == nil {
		return nil, errors.New("client object empty")
	}

	return client.Search(s.Query, &SearchOpts{Page: s.Page + 1})
}

// GetFull Fetches the full data about the movie using the api.
//
// - client : Omdb client to use for the request.
func (m *MoviePreview) GetFull(client *OmdbClient) (*Movie, error) {
	return client.GetMovie(&GetMovieOpts{Id: m.ImdbId})
}
