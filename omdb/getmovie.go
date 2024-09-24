// (c) Jisin0
// Types and functions for getmovie call.

package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jisin0/filmigo/internal/encode"
	"github.com/go-faster/errors"
)

// Get movie query values.
type GetMovieOpts struct {
	// Imdb id of the movie for ex: tt1285016.
	ID string `url:"i"`
	// Exact title of the movie to fetch.
	Title string `url:"t"`
	// Type of result to return for ex: series.
	// See omdb.ResultTypeXX values for all possible values.
	Type string `url:"type"`
	// Year of realease of the movie.
	Year string `url:"y"`
	// Length of plot to return, "short" for a short plot or "full" for the full plot.
	// Use omdb.PlotShort or omdb.PlotFull.
	Plot string `url:"plot"`
}

// Rating of the movie with data about the source.
type Rating struct {
	// Source of the rating for ex: Internet Movie Database.
	Source string `json:"source"`
	// Value of the rating either as a fraction or percentage.
	Value string `json:"value"`
}

// GetMovie gets the full data of a movie using it's imdb id or the full name.
func (c *OmdbClient) GetMovie(opts *GetMovieOpts) (*Movie, error) {
	if opts.ID == "" && opts.Title == "" {
		return nil, errors.New("no id or title provided")
	}

	if c.apiKey == "" {
		return nil, errors.New("no omdb api key provided")
	}

	urlParams, err := encode.URLParams(*opts)
	if err != nil {
		return nil, errors.Wrap(err, "getmovie: failed to parse url params")
	}

	fullURL := fmt.Sprintf("%s?apikey=%s&%s", omdbAPIURL, c.apiKey, urlParams.Encode())

	req, err := http.NewRequest("GET", fullURL, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != statusCodeSuccess {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var movie Movie

	err = json.Unmarshal(bytes, &movie)
	if err != nil {
		return nil, errors.Wrap(err, "failed to serialize data")
	}

	if movie.Error != "" {
		return &movie, errors.New(movie.Error)
	}

	return &movie, nil
}
