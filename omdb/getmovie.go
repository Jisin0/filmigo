// (c) Jisin0
// Types and functions for getmovie call.

package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jisin0/filmigo/encode"
	"github.com/go-faster/errors"
)

// Get movie query values.
type GetMovieOpts struct {

	// Imdb id of the movie for ex: tt1285016.
	Id string `url:"i"`

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

// Result from the omc.log.Debug("using cached data")db.GetMovie function containing full data on a movie.
type Movie struct {

	// Title of the movie.
	Title string `json:"title"`

	// Year the movie was released.
	Year string `json:"year"`

	// Parental guidline rating for ex: R, PG etc.
	Rated string `json:"rated"`

	// Date on which the movie was released in the format 01 January 1950.
	Relased string `json:"released"`

	// Runtime/Duration of the movie in minutes for ex. 120 min.
	Runtime string `json:"runtime"`

	//Genres of the movie ina string seperated by commas for ex: Action, Comedy, Romance.
	Genres string `json:"genre"`

	// Name of the Director of the movie.
	Director string `json:"director"`

	// Writers of the movie seperated by commas.
	Writers string `json:"writer"`

	// Actors/Stars of the movie seperated by commas.
	Actors string `json:"actors"`

	// Plot of the movie.
	Plot string `json:"plot"`

	// List of languages seperated by commas for ex: English, Spanish, Italian.
	Languages string `json:"language"`

	// Country of origin of the movie.
	Country string `json:"country"`

	// Awards won by the movie.
	Awards string `json:"awards"`

	// Poster image url of the movie.
	Poster string `json:"poster"`

	// Ratings of the movie from various sources.
	Ratings []Rating `json:"ratings"`

	// Metascore of the movie.
	Metascore string `json:"metascore"`

	// Rating of the movie from imdb, returned value is out of 10.
	ImdbRating string `json:"imdbrating"`

	// Number of votes the movie received on imdb.
	ImdbVotes string `json:"imdbvotes"`

	// Imdb id of the movie.
	ImdbId string `json:"imdbid"`

	// Type of title for ex: movie, series or episode.
	// Use omdb.ResultTypeXX values for reliablility.
	Type string `json:"type"`

	// DVD release date of the movie.
	DVD string `json:"dvd"`

	// Boxoffice income generated by the movie with it's currency.
	BoxOffice string `json:"boxoffice"`

	// Production company associated with the movie.
	Production string `json:"production"`

	// Any official website of the movie.
	Website string `json:"website"`

	// This value indicates wether a result was returned.
	// Value is True on success and False on failure (case sensitive).
	Response string `json:"response"`

	// Error message returned for failed queries.
	// This value should be checked to determine wether the reason for a failed call.
	Error string `json:"error"`
}

// Rating of the movie with data about the source.
type Rating struct {
	//Source of the rating for ex: Internet Movie Database.
	Source string `json:"source"`
	//Value of the rating either as a frcation or percentage.
	Value string `json:"value"`
}

// GetMovie gets the full data of a movie using it's imdb id or the full name.
func (c *OmdbClient) GetMovie(opts *GetMovieOpts) (*Movie, error) {

	if opts.Id == "" && opts.Title == "" {
		return nil, errors.New("no id or title provided")
	}

	if c.apiKey == "" {
		return nil, errors.New("no obdb api key provided")
	}

	urlParams, err := encode.UrlParams(*opts)
	if err != nil {
		return nil, errors.Wrap(err, "getmovie: failed to parse url params")
	}

	fullURL := fmt.Sprintf("%s?apikey=%s&%s", omdbAPIURL, c.apiKey, urlParams.Encode())

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