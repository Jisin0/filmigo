// (c) Jisin0
// File contains functions for scraping movie data from it's imdb page.

package imdb

import (
	"log"
	"net/http"

	"github.com/Jisin0/filmigo/encode"
	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
	"golang.org/x/net/html"
)

const (
	// path for homepage of any imdb movie/show
	movieBaseURL = baseImdbURL + "/title"
)

// Full movie object, contains data about a movie/show only available after scraping it's data with imdb.GetMovie().
type Movie struct {
	// ID of the movie
	ID string
	// Link to the movie
	Link string
	// Full title of the movie
	Title string `xpath:"//h1[@data-testid='hero__pageTitle']/span"`
	// Year of release of the movie
	Year string `xpath:"//h1[@data-testid='hero__pageTitle']/..//a[contains(@href, 'releaseinfo')]"`
	// Ratings of the movie in the format n/10
	Rating string `xpath:"//div[@data-testid='hero-rating-bar__aggregate-rating']//div[@data-testid='hero-rating-bar__aggregate-rating__score']"`
	// Rumber of votes the movie got
	Votes string `xpath:"//div[@data-testid='hero-rating-bar__aggregate-rating']/a/span/div/div[2]/div[3]"`
	// The directors of the movie
	Directors types.Links `xpath:"//div[@role='presentation']/ul//*[starts-with(text(), 'Director')]/../div"`
	// The writers of the movie
	Writers types.Links `xpath:"//div[@role='presentation']/ul//*[starts-with(text(), 'Writer')]/../div"`
	// The main stars of the movie
	Stars types.Links `xpath:"//div[@role='presentation']/ul//*[starts-with(text(), 'Star')]/../div"`
	// Genres of the movie
	Genres types.Links `xpath:"//div[@data-testid='genres']/div[2]"`
	// A short plot of the movie in a few lines
	Plot string `xpath:"/html/body//main//p[@data-testid='plot']//span[@data-testid='plot-xl']"`
	// A string with details about the release including date and country
	Releaseinfo string `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-releasedate']/div//a"`
	// Origin of release, commonly the country
	Origin string `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-origin']/div//a"`
	// Official sites related to the movie/show
	OfficialSites types.Links `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='details-officialsites']/div/ul"`
	// Languages in which the movie/show is available in
	Languages types.Links `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-languages']/div/ul"`
	// Any alternative name of the movie.
	Aka string `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-akas']//span"`
	// Locations at which the movie/show was filmed at
	Locations types.Links `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-filminglocations']/div/ul"`
	// Companies which produced the movie
	Companies types.Links `xpath:"//section[@data-testid='Details']/div[@data-testid='title-details-section']//li[@data-testid='title-details-companies']/div/ul"`
	// Runtime of the move
	Runtime string `xpath:"//div[@data-testid='title-techspecs-section']/ul/li[@data-testid='title-techspec_runtime']/div"`
}

// Function to get the full details about a movie/show using it's id .
//
// - id : Unique id used to identify each movie for ex: tt15398776.
//
// Returns an error on failed requests or if the movie wasn't found.
func (c *ImdbClient) GetMovie(id string) (*Movie, error) {
	// Verify id or extract it if it's in a url
	id = resultTypeTitleRegex.FindString(id)
	if id == "" {
		return nil, errors.New("imdb.getmovie id did not match regex")
	}

	var movie Movie

	// Check cache for existing first
	if !c.disabledCaching {
		if err := c.cache.MovieCache.Load(id, &movie); err == nil {
			return &movie, nil
		}
	}

	// Get the webpage
	url := movieBaseURL + "/" + id

	doc, err := doRequest(url)
	if err != nil {
		return nil, err
	}

	movie = Movie{
		ID:   id,
		Link: url,
	}

	var ok bool

	movie, ok = encode.Xpath(doc, movie).(Movie)
	if !ok {
		return nil, errors.New("unknown type returned from encode.Xpath")
	}

	// Cache data for next time
	if !c.disabledCaching {
		err := c.cache.MovieCache.Save(id, movie)
		if err != nil {
			log.Println(err)
		}
	}

	return &movie, nil
}

// executes a get request to given url and parses the response body.
func doRequest(url string) (*html.Node, error) {
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	defer resp.Body.Close()

	if resp.StatusCode == statusCodeNotFound {
		return nil, errors.Errorf("movie/person was not not found")
	} else if resp.StatusCode != statusCodeSuccess {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse document")
	}

	return doc, nil
}
