// (c) Jisin0
// Types and methods for advanced search.

package imdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jisin0/filmigo/internal/encode"
	"github.com/Jisin0/filmigo/internal/types"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
)

const (
	baseAdvancedSearchURL  = baseImdbURL + "/search/"
	advancedSearchTitleURL = baseAdvancedSearchURL + "title/"
	advancedSearchNameURL  = baseAdvancedSearchURL + "name/"
	href                   = "href"
)

// Options for the AdvancedSearchTitle query see https://imdb.com/search/title to see the list and syntax for each option.
type AdvancedSearchTitleOpts struct {
	// Search by the title name of a movie/show.
	TitleName string `url:"title"`
	// Type filter by the type of title see TitleTypeXX values in the constants package for all possible values for ex: constants.TitleTypeMovie.
	Types []string `url:"title_type"`
	// RelaseDate a range/period of time in which the title was released. Dates must be in the format yyyy-dd-mm.
	RelaseDate types.SearchRange `url:"release_date"`
	// Ratings range of minimum and maximum rating of titles returned. for ex. Start: 7, End: 9.5.
	Rating types.SearchRange `url:"user_rating"`
	// Votes range of votes on a title. for ex. Start: 10000, End: 500000.
	Votes types.SearchRange `url:"num_votes"`
	// Genres filter by the genre of the title see TitleGenreXX values in the constants package for all possible values for ex: constants.TitleGenreAction.
	Genres []string `url:"genres"`
	// Awards: find titles that have won or have been nominated for an award.
	// See TitleAwardXX values in the constants package for all possible values for ex: constants.TitleAwardOscarWinner.
	Awards []string `url:"groups"`
	// Topics on the imdb page of the title. Use additional params to search within a topic.
	Topics []string `url:"has"`
	// Companies: filter by companies what produced the title.
	Companies []string `url:"companies"`
	// InstantWatch: search by ability to watch online on an instantwatch platform.
	InstantWatches []string `url:"online_availability"`
	// Certificates: The watching certificates of a title. see TitleCertificateXX values in the constants package for all possible values for ex: constants.TitleCertificatePG.
	Certificates []string `url:"certificates"`
	// Color: The color info of the title. for ex. constants.TitleColorBlackAndWhite for black&white titles.
	Colors []string `url:"colors"`
	// Countries: country codes of the countries associated with the title. for ex. GB for United Kingdom.
	Countries []string `url:"country"`
	// Keywords: Filter by additional keywords.
	Keywords []string `url:"keywords"`
	// Languages. Language codes of language of the title. for ex. en for english.
	Languages []string `url:"languages"`
	// Popularity. Filter by a range of imdb popularity rank.
	Popularity types.SearchRange `url:"moviemeter"`
	// CastOrCrew. Lost of ids of actors or crew in the title. for ex nm0614165 for Cillian Murphy
	CastOrCrew []string `url:"role"`
	// Characters. List of names of characters in the movie.
	Characters []string `url:"characters"`
	// Runtime. Range of runtime of the movie in minutes.
	Runtime types.SearchRange `url:"runtime"`
	// SoundMixes: The sound mix of a title. see TitleSoundXX values in the constants package for all possible values for ex: constants.TitleSoundDolby.
	SoundMixes []string `url:"sound_mixes"`
	// AdultTitles: Set value to constants.StringInclude to include adult titles
	AdultTitles string `url:"adult"`
	// Additional url parameters to be passed along with the request.
	ExtraParams map[string]any
}

// Single result from the AdvancedSearchTitle result list.
type AdvancedSearchTitleResult struct {
	// Indicates wether the title can be rated on imdb.
	CanRate bool `json:"canRate"`
	// Parental certificate of the title: 15 indicates TV-MA, 12 indicates PG-13, 18 indicates TV-MA.
	Certificate string `json:"certificate"`
	// The year in which a TVShow ended. Only for Series and Mini-Series.
	EndYear int `json:"endYear"`
	// Genres of the title.
	Genres []string `json:"genres"`
	// Indicates wether the movie has onli watching option (highly inaccurate).
	HasWatchOption bool `json:"hasWatchOption"`
	// Full original Title of the movie/show.
	OriginalTitle string `json:"originalTitleText"`
	// Plot of the movie/show.
	Plot string `json:"plot"`
	// Image: Poster image of a title or profile image of a person.
	Image AdvSearchImage `json:"primaryImage"`
	// Rating data about the title.
	Rating struct {
		// Value of rating out of 10.
		Value float32 `json:"aggregateRating"`
		// Number of votes received for the title.
		Votes int64 `json:"voteCount"`
	} `json:"ratingSummary"`
	// Year in which the title was first released.
	ReleaseYear int `json:"releaseYear"`
	// Runtime of the title in minutes.
	Runtime int `json:"runtime"`
	// Imdb id of the title.
	ID string `json:"titleId"`
	// Title of the movie or show.
	Title string `json:"titleText"`
	// Data about the type of title.
	Type struct {
		// Indicates wether the title can have episodes.
		CanHaveEpisodes bool `json:"canHaveEpisodes"`
		// Id of the type. Possible values include movie, tvSeries, tvMiniSeries etc.
		ID string `json:"id"`
		// User-Friendly text about the type for ex: TV Series for tvSeries.
		Text string `json:"text"`
	} `json:"titleType"`
	// Video id of the trailer of the title.
	TrailerID string `json:"trailerId"`
}

// Poster image of an AvancedSearch result.
type AdvSearchImage struct {
	// Caption of the image.
	Caption string `json:"caption"`
	// Height of the image in pixels.
	Height int `json:"height"`
	// ID of the image (not sure where to use this)
	ID string `json:"id"`
	// URL of the image.
	URL string `json:"url"`
	// WIdth of the image in pixels.
	Width int `json:"width"`
}

// AdvancedSearchTitle uses the search page to search for titles using many configuration options. Use SearchX methods for simple fast searches using the api.
//
// opts - configure search options.
func (*ImdbClient) AdvancedSearchTitle(opts *AdvancedSearchTitleOpts) ([]*AdvancedSearchTitleResult, error) {
	urlParams, _ := encode.URLParams(*opts)
	urlParams = encode.URLMapParams(opts.ExtraParams, urlParams)

	fullURL := fmt.Sprintf("%s?%s", advancedSearchTitleURL, urlParams.Encode())

	req, err := http.NewRequest("GET", fullURL, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	if resp.StatusCode == statusCodeNotFound {
		return nil, errors.Errorf("results not found")
	} else if resp.StatusCode != statusCodeSuccess {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse document")
	}

	defer resp.Body.Close()

	dataNode := htmlquery.FindOne(doc, "//script[@id='__NEXT_DATA__']")
	if dataNode == nil {
		return nil, errors.New("results not found")
	}

	// temporary type to get to deeply nested results.
	// using a third party lib like gjson should be considered.
	type a struct {
		Props struct {
			PropsPage struct {
				SearchResults struct {
					TitleResults struct {
						Items []*AdvancedSearchTitleResult `json:"titleListItems"`
					} `json:"titleResults"`
				} `json:"searchResults"`
			} `json:"pageProps"`
		} `json:"props"`
	}

	var data a

	err = json.Unmarshal([]byte(htmlquery.InnerText(dataNode)), &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal results data")
	}

	results := data.Props.PropsPage.SearchResults.TitleResults.Items

	if len(results) < 1 {
		return results, errors.New("results not found")
	}

	return results, nil
}

// FullTitle returns the full data about a title scraped from it's imdb page.
//
// - client : Client to make imdb requests through.
func (s *AdvancedSearchTitleResult) FullTitle(client *ImdbClient) (*Movie, error) {
	return client.GetMovie(s.ID)
}

// Options for the AdvancedSearchName query see https://imdb.com/search/title to see the list and syntax for each option.
type AdvancedSearchNameOpts struct {
	// Name: Filter by the name of the person.
	Name string `url:"name"`
	// BirthRange: A range of birth date inside which the person was born in the format yyyy-dd-mm.
	BirthRange types.SearchRange `url:"birth_date"`
	// Birthday: Birthday of the actor in the format MM-DD
	Birthday string `url:"birth_monthday"`
	// Awards: List of awards won by the person
	// see NameAwardXX values in the constants package for all possible values for ex. constants.NameAwardBestActorNominated
	Awards []string `url:"groups"`
	// Page topics: filter by topics on the imdb page of the person. Use ExtraParams for searching within a topic.
	PageTopics []string `url:"has"`
	// DeathRange: A range of death date inside which the person passed away in the format yyyy-dd-mm.
	DeathRange types.SearchRange `url:"birth_date"`
	// Genders: Return actors that ho by the given genders.
	// See NameGenderXX values for all possible values for ex: NameGenderMale.
	Genders []string `url:"gender"`
	// Titles: A list of imdb ids of titles for ex. tt15398776 to search for stars in oppenheimer.
	Titles []string `url:"roles"`
	// AdultNames: Set value to constants.StringInclude to include stars in adult titles.
	AdultNames string `url:"adult"`
	// Additional url parameters to be passed along with the request.
	ExtraParams map[string]any
}

// Single results item from an AdvancedSearchName results list.
type AdvancedSearchNameResult struct {
	// Title: Name of the person.
	Title string `json:"nameText"`
	// Bio or short decription of the person.
	Bio string `json:"bio"`
	// Data about a title the person is known for.
	KnownFor struct {
		// Indicates wether the title can have episodes.
		CanHaveEpisodes bool `json:"canHaveEpisodes"`
		// Original or full title of the movie or show.
		OriginalTitle string `json:"originalTitle"`
		// Imdb ID of the title.
		ID string `json:"titleId"`
		// Name of the title.
		Title string `json:"titleText"`
		// Range of years in which the title was released.
		YearRange struct {
			// Year in which the title was first released.
			ReleaseYear int `json:"year"`
			// Year in which a series ended or last broadcasted.
			EndYear int `json:"endYear"`
		} `json:"yearRange"`
	} `json:"knownFor"`
	// Imdb ID of the person.
	ID string `json:"nameId"`
	// Image: Profile image of a person.
	Image AdvSearchImage `json:"primaryImage"`
	// Professions: Roles taken by a person for ex: Director, Actress, Producer.
	Professions []string `json:"primaryProfessions"`
}

// AdvancedSearchName uses the search page to search for names using many configuration options. Use SearchX methods for simple fast searches using the api.
//
// opts - configure search options.
func (*ImdbClient) AdvancedSearchName(opts *AdvancedSearchNameOpts) ([]*AdvancedSearchNameResult, error) {
	urlParams, _ := encode.URLParams(*opts)
	urlParams = encode.URLMapParams(opts.ExtraParams, urlParams)

	fullURL := fmt.Sprintf("%s?%s", advancedSearchNameURL, urlParams.Encode())

	doc, err := doRequest(fullURL)
	if err != nil {
		return nil, err
	}

	dataNode := htmlquery.FindOne(doc, "//script[@id='__NEXT_DATA__']")
	if dataNode == nil {
		return nil, errors.New("results not found")
	}

	// temporary type to get to deeply nested results.
	// using a third party lib like gjson should be considered.
	type a struct {
		Props struct {
			PropsPage struct {
				SearchResults struct {
					NameResults struct {
						Items []*AdvancedSearchNameResult `json:"nameListItems"`
					} `json:"nameResults"`
				} `json:"searchResults"`
			} `json:"pageProps"`
		} `json:"props"`
	}

	var data a

	err = json.Unmarshal([]byte(htmlquery.InnerText(dataNode)), &data)
	if err != nil {
		return nil, errors.New("failed to unmarshal results data")
	}

	results := data.Props.PropsPage.SearchResults.NameResults.Items

	if len(results) < 1 {
		return results, errors.New("results not found")
	}

	return results, nil
}

// FullPerson returns the full data about a title scraped from it's imdb page.
//
// - client : Client to make imdb requests through.
func (s *AdvancedSearchNameResult) FullPerson(client *ImdbClient) (*Person, error) {
	return client.GetPerson(s.ID)
}
