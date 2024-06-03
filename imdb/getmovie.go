// (c) Jisin0
// File contains functions for scraping movie data from it's imdb page.

package imdb

import (
	"encoding/json"
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

type Movie struct {
	//Imdb id of the movie .
	ID string
	// Years of release of the movie, A range for shows over multiple years.
	ReleaseYear string
	MovieJsonContent
	MovieDetailsSection
}

// Data scraped from the details section of a movie using xpath.
type MovieDetailsSection struct {
	// A string with details about the release including date and country
	Releaseinfo string `xpath:"//li[@data-testid='title-details-releasedate']/div//a"`
	// Origin of release, commonly the country
	Origin string `xpath:"//li[@data-testid='title-details-origin']/div//a"`
	// Official sites related to the movie/show
	OfficialSites types.Links `xpath:"//li[@data-testid='details-officialsites']/div/ul"`
	// Languages in which the movie/show is available in
	Languages types.Links `xpath:"//li[@data-testid='title-details-languages']/div/ul"`
	// Any alternative name of the movie.
	Aka string `xpath:"//li[@data-testid='title-details-akas']//span"`
	// Locations at which the movie/show was filmed at
	Locations types.Links `xpath:"//li[@data-testid='title-details-filminglocations']/div/ul"`
	// Companies which produced the movie
	Companies types.Links `xpath:"//li[@data-testid='title-details-companies']/div/ul"`
}

// Data scraped from the json attached in the script tag.
type MovieJsonContent struct {
	//Type of the title possibble values are Movie, TVSeries etc.
	Type string `json:"@type"`
	// ID of the movie
	ID string
	// Link to the movie
	URL string `json:"url"`
	// Full title of the movie
	Title string `json:"name"`
	// Url of the full size poster image.
	PosterURL string `json:"image"`
	// Content rating class (currenly undocumented).
	ContentRating string `json:"contentRating"`
	// Date the movie was released on in yyyy-mm-dd format.
	ReleaseDate string `json:"datePublished"`
	// Keywords associated with the movie in a comma seperated list.
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

// A video file about the entity.
type VideoObject struct {
	// Name of the video.
	Name string `json:"name"`
	// Url to create embedded video players.
	EmbedURL string `json:"embedUrl"`
	// Image url of the thumbnail of the video.
	Thumbnail string `json:"thumbnailUrl"`
	// Short description of the video.
	Description string `json:"description"`
	// Duration of the video.
	Duration string `json:"duration"`
	// Url of the video .
	URL string `json:"url"`
	// Timestamp of the upload time of the video.
	// Use time.Parse(time.RFC3339Nano, UploadDate) to parse it.
	UploadDate string `json:"uploadDate"`
}

// Review of a movie or show.
type Review struct {
	// Item that was reviewed.
	ItemReviewed ReviewItem `json:"itemReviewed"`
	// Author of the review.
	Author ReviewAuthor `json:"author"`
	// Date on which the review was created in the format yyyy-mm-dd
	Date string `json:"dateCreated"`
	// Language in which the review is written.
	Language string `json:"inLanguage"`
	// Body or content of the review.
	Body string `json:"reviewBody"`
	// Ratings for the review.
	Rating Rating `json:"reviewRating"`
}

// An item that was reviewed.
type ReviewItem struct {
	// Url of the item that was reviewed.
	URL string `json:"url"`
}

// Author of a review.
type ReviewAuthor struct {
	// Name of the person.
	Name string `json:"name"`
}

// Rating data for a title or review.
type Rating struct {
	// Number of votes. (absent for reviews)
	Votes int64 `json:"ratingCount"`
	// Worst rating received.
	Worst int `json:"worstRating"`
	// Best rating received.
	Best int `json:"bestRating"`
	// Actual value of the rating out of 10.
	Value int `json:"ratingValue"`
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

	if doc == nil {
		return nil, errors.New("movie or or person not found")
	}

	movie = Movie{
		ID: id,
	}

	// var ok bool

	// movie, ok = encode.Xpath(doc, movie).(Movie)
	// if !ok {
	// 	return nil, errors.New("unknown type returned from encode.Xpath")
	// }

	jsonDataNode := htmlquery.FindOne(doc, "//script[@type='application/ld+json']")
	if jsonDataNode == nil {
		return nil, errors.New("json data node not found")
	}

	json.Unmarshal([]byte(htmlquery.InnerText(jsonDataNode)), &movie.MovieJsonContent)

	detailsNode, err := htmlquery.Query(doc, "//section[@data-testid='Details']/div[@data-testid='title-details-section']")
	if detailsNode != nil && err == nil {
		encode.Xpath(detailsNode, &movie.MovieDetailsSection)
	}

	releaseYearNode, err := htmlquery.Query(doc, "//h1[@data-testid='hero__pageTitle']/..//a[contains(@href, 'releaseinfo')]")
	if releaseYearNode != nil && err == nil {
		movie.ReleaseYear = htmlquery.InnerText(releaseYearNode)
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
