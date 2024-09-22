// (c) Jisin0
// File contains functions for scraping movie data from it's imdb page.

package imdb

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/Jisin0/filmigo/internal/encode"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
	"golang.org/x/net/html"
)

const (
	// path for homepage of any imdb movie/show
	movieBaseURL = baseImdbURL + "/title"
)

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
	Worst float32 `json:"worstRating"`
	// Best rating received.
	Best float32 `json:"bestRating"`
	// Actual value of the rating out of 10.
	Value float32 `json:"ratingValue"`
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
	movieURL := movieBaseURL + "/" + id

	doc, err := doRequest(movieURL)
	if err != nil {
		return nil, err
	}

	if doc == nil {
		return nil, errors.New("movie or or person not found")
	}

	movie = Movie{
		ID: id,
	}

	jsonDataNode := htmlquery.FindOne(doc, "//script[@type='application/ld+json']")
	if jsonDataNode == nil {
		return nil, errors.New("json data node not found")
	}

	err = json.Unmarshal([]byte(htmlquery.InnerText(jsonDataNode)), &movie.MovieJSONContent)
	if err != nil {
		return nil, errors.New("failed to unmarshal results data")
	}

	detailsNode, err := htmlquery.Query(doc, "//section[@data-testid='Details']/div[@data-testid='title-details-section']")
	if detailsNode != nil && err == nil {
		err = encode.Xpath(detailsNode, &movie.MovieDetailsSection)
		if err != nil {
			return nil, errors.Wrap(err, "error while scraping data with xpath")
		}
	}

	releaseYearNode, err := htmlquery.Query(doc, "//h1[@data-testid='hero__pageTitle']/..//a[contains(@href, 'releaseinfo')]")
	if releaseYearNode != nil && err == nil {
		movie.ReleaseYear = htmlquery.InnerText(releaseYearNode)
	}

	if movie.Runtime != "" {
		movie.Runtime = parseIMDbDuration(movie.Runtime)
	}

	if s, err := url.QueryUnescape(movie.Title); err == nil {
		movie.Title = s
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
func doRequest(urlPath string) (*html.Node, error) {
	req, err := http.NewRequest("GET", urlPath, http.NoBody)
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
