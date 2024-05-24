// (c) Jisin0
// Types and methods for advanced search.
package imdb

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Jisin0/filmigo/encode"
	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
)

const (
	baseAdvancedSearchURL  = baseImdbURL + "/search/"
	advancedSearchTitleURL = baseAdvancedSearchURL + "title/"
	advancedSearchNameURL  = baseAdvancedSearchURL + "name/"
)

// // Type for storing data from adbanced search page.
// type AdvancedSearchItem struct {
// 	// Index number of the item.
// 	Index int
// 	// Title: Name of the movie/show or person.
// 	Title string
// 	// Image: Poster image of a title or profile image of a person.
// 	Image string
// 	// Link: Link to the title or person's imdb page.
// 	Link string
// 	// Metadata: Metadata for titles containing the year of release, duration and us certificate.
// 	Metadata []string
// 	// Rating: A string containing rating info for ex: 7.5 (35K).
// 	Rating string
// 	// Description: A description of the title or person.
// 	Description string
// 	// Top title of a actor/actress. Only for people/names.
// 	TopTitle types.Link
// 	// Roles: Roles taken by a person for ex: Director, Actress, Producer.
// 	Roles []string
// }

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
}

// Single result from the AdvancedSearchTitle result list.
type AdvancedSearchTitleResult struct {
	// Index number of the item.
	Index int
	// Title: Name of the movie/show or person.
	Title string
	// Image: Poster image of a title or profile image of a person.
	Image string
	// Link: Link to the title or person's imdb page.
	Link string
	// Metadata: Metadata for titles containing the year of release, duration and us certificate.
	Metadata []string
	// Rating: A string containing rating info for ex: 7.5 (35K).
	Rating string
	// Description: A description of the title or person.
	Description string
}

// AdvancedSearchTitle uses the search page to search for titles using many configuration options. Use SearchX methods for simple fast searches using the api.
//
// opts - configure search options.
func (*ImdbClient) AdvancedSearchTitle(opts *AdvancedSearchTitleOpts) ([]*AdvancedSearchTitleResult, error) {

	urlParams, _ := encode.UrlParams(*opts)
	fullURL := fmt.Sprintf("%s?%s", advancedSearchTitleURL, urlParams)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	if resp.StatusCode == 404 {
		return nil, errors.Errorf("results not found")
	} else if resp.StatusCode != 200 {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse document")
	}

	defer resp.Body.Close()

	list, err := htmlquery.Query(doc, "//main/div[@role='presentation']/div[last()]//div[@role='tabpanel']//section/div[2]/div[2]/ul")
	if err != nil || list == nil {
		return nil, err
	}

	elements, err := htmlquery.QueryAll(list, "//div[ends-with(@class, 'dli-parent')]")
	if err != nil {
		return nil, errors.Wrap(err, "failed elements query")
	}

	var results []*AdvancedSearchTitleResult

	for _, e := range elements {

		var item AdvancedSearchTitleResult

		if posterNode, _ := htmlquery.Query(e, "//img"); posterNode != nil {
			for _, a := range posterNode.Attr {
				if a.Key == "src" {
					item.Image = a.Val
				}
			}
		}

		if titleNode, _ := htmlquery.Query(e, "//h3"); titleNode != nil {
			s := htmlquery.InnerText(titleNode)
			if split := strings.SplitN(s, ".", 2); len(split) > 1 {
				n, _ := strconv.Atoi(split[0])

				item.Index = n
				s = strings.TrimSpace(split[1])
			}

			item.Title = s

			aNode, _ := htmlquery.Query(titleNode, "/..")
			for _, a := range aNode.Attr {
				if a.Key == "href" {
					item.Link = baseImdbURL + a.Val
				}
			}
		}

		if metadataNode, _ := htmlquery.Query(e, "//div[ends-with(@class, 'dli-title-metadata')]"); metadataNode != nil {
			var items []string
			for _, span := range htmlquery.Find(metadataNode, "/span") {
				items = append(items, htmlquery.InnerText(span))
			}

			item.Metadata = items
		}

		if ratingsNode, _ := htmlquery.Query(e, "//span[starts-with(@data-testid, 'ratingGroup')]"); ratingsNode != nil {
			item.Rating = htmlquery.InnerText(ratingsNode)
		}

		if descriptionNode, _ := htmlquery.Query(e, "/div[last()]//div[contains(@class, 'inner')]"); descriptionNode != nil {
			item.Description = htmlquery.InnerText(descriptionNode)
		}

		results = append(results, &item)

	}

	return results, nil
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
	//see NameAwardXX values in the constants package for all possible values for ex. constants.NameAwardBestActorNominated
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
}

// Single results item from an AdvancedSearchName results list.
type AdvacedSearchNameResult struct {
	// Index number of the item.
	Index int
	// Title: Name of the movie/show or person.
	Title string
	// Image: Poster image of a title or profile image of a person.
	Image string
	// Professions: Roles taken by a person for ex: Director, Actress, Producer.
	Professions []string
	// Top title of a actor/actress. Only for people/names.
	TopTitle types.Link
	// Link: Link to the title or person's imdb page.
	Link string
	// Description: A description of the title or person.
	Description string
}

// AdvancedSearchName uses the search page to search for names using many configuration options. Use SearchX methods for simple fast searches using the api.
//
// opts - configure search options.
func (*ImdbClient) AdvancedSearchName(opts *AdvancedSearchNameOpts) ([]*AdvacedSearchNameResult, error) {

	urlParams, _ := encode.UrlParams(*opts)
	fullURL := fmt.Sprintf("%s?%s", advancedSearchNameURL, urlParams)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	if resp.StatusCode == 404 {
		return nil, errors.Errorf("results not found")
	} else if resp.StatusCode != 200 {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse document")
	}

	defer resp.Body.Close()

	list, err := htmlquery.Query(doc, "//main/div[@role='presentation']/div[last()]//div[@role='tabpanel']//section/div[2]/div[2]/ul")
	if err != nil || list == nil {
		return nil, errors.Wrap(err, "failed to find people list")
	}

	elements, err := htmlquery.QueryAll(list, "//div[ends-with(@class, 'dli-parent')]")
	if err != nil {
		return nil, errors.Wrap(err, "failed elements query")
	}

	var results []*AdvacedSearchNameResult

	for _, e := range elements {

		var item AdvacedSearchNameResult

		if posterNode, _ := htmlquery.Query(e, "//img"); posterNode != nil {
			for _, a := range posterNode.Attr {
				if a.Key == "src" {
					item.Image = a.Val
				} else if a.Key == "href" {
					item.Link = a.Val
				}
			}
		}

		if titleNode, _ := htmlquery.Query(e, "//h3"); titleNode != nil {
			s := htmlquery.InnerText(titleNode)
			if split := strings.SplitN(s, ".", 2); len(split) > 1 {
				n, _ := strconv.Atoi(split[0])

				item.Index = n
				s = strings.TrimSpace(split[1])
			}

			item.Title = s

			aNode, _ := htmlquery.Query(titleNode, "/..")
			for _, a := range aNode.Attr {
				if a.Key == "href" {
					item.Link = baseImdbURL + a.Val
				}
			}
		}

		if professionsNode, _ := htmlquery.Query(e, "//ul[@data-testid='nlib-professions']"); professionsNode != nil {
			var items []string
			for _, li := range htmlquery.Find(professionsNode, "/li") {
				items = append(items, htmlquery.InnerText(li))
			}

			item.Professions = items
		}

		if topTitleNode, _ := htmlquery.Query(e, "//a[@data-testid='nlib-known-for-title']"); topTitleNode != nil {
			var topTitle types.Link
			topTitle.Text = htmlquery.InnerText(topTitleNode)

			for _, a := range topTitleNode.Attr {
				if a.Key == "href" {
					topTitle.Href = baseImdbURL + a.Val
				}
			}

			item.TopTitle = topTitle
		}

		if descriptionNode, _ := htmlquery.Query(e, "/div[last()]//div[contains(@class, 'inner')]"); descriptionNode != nil {
			item.Description = htmlquery.InnerText(descriptionNode)
		}

		results = append(results, &item)

	}

	return results, nil
}
