// (c)Jisin0
// Functions and types to scrape data of a person.

package imdb

import (
	"encoding/json"
	"log"

	"github.com/Jisin0/filmigo/encode"
	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
)

const (
	// path for homepage of any imdb movie/show
	personBaseURL = baseImdbURL + "/name"
)

// Type containing the full data about a person scraped from their imdb page.
type Person struct {
	// Imdb id of the user for ex: nm0000129
	ID string
	// Links to movies/show the person is known for.
	KnownFor types.Links `xpath:"|linklist"`
	PersonJSONContent
	PersonDetailsSection
	PersonDYKSection
}

// Data to be scraped from the Did You Know section of an actor.
type PersonDYKSection struct {
	// A short trivia fact about the person.
	Trivia string `xpath:"/div[2]//li[@data-testid='name-dyk-trivia']/div"` // Trivia is always the first dyk hence the div[2]
	// A popular quote of the person. All quotes can be found at {link}/quotes.
	Quote string `xpath:"//li[@data-testid='name-dyk-quote']/div"`
	// A nickname of the person.
	Nickname string `xpath:"//li[@data-testid='name-dyk-nickname']/div"`
	// Any trademark features of the person.
	Trademark string `xpath:"//li[@data-testid='name-dyk-trademarks']/div"`
}

// Data to be scraped from a persons Details section.
type PersonDetailsSection struct {
	// Official sites of the person.
	OfficialSites types.Links `xpath:"/li[@data-testid='details-officialsites']/div|linklist"`
	// Height of the person.
	Height string `xpath:"/li[@data-testid='nm_pd_he']/div//span"`
	// Date of birth . for ex : April 30, 1981
	Birthday string `xpath:"/li[@data-testid='nm_pd_bl']/div/ul/li"`
	// Spouse of the person.
	Spouse types.Links `xpath:"/li[@data-testid='nm_pd_sp']/div|linklist"`
	// Other works - usually a short sentence about a different work of the person.
	OtherWorks string `xpath:"/li[@data-testid='nm_pd_wrk']/div"`
}

// JSON data available for the person the page.
type PersonJSONContent struct {
	// URL of the imdb page.
	URL string `json:"url"`
	// Name of the person.
	Name string `json:"name"`
	// URL of the main full-size poster image of the person.
	Image string `json:"image"`
	// Short description of the person.
	Description string `json:"description"`
	// A video about the person.
	Video VideoObject `json:"video"`
	// Headline for the person for ex: Cillian Murphy - Actor, Producer, Writer.
	Headline string `json:"headLine"`
	// Extra content about the person.
	MainEntity struct {
		// Job Titles of the person for ex: Actor, Producer, Director.
		JobTitles []string `json:"jobTitles"`
		// Date of birth of the person in yyyy-mm-dd format.
		BirthDate string `json:"birthDate"`
	} `json:"mainEntity"`
}

// Function to get the full details about a person using their id .
//
// - id : Unique id used to identify each person for ex: nm0614165.
//
// Returns an error on failed requests or if the person wasn't found.
func (c *ImdbClient) GetPerson(id string) (*Person, error) {
	// Verify id or extract it if it's in a url
	id = resultTypeNameRegex.FindString(id)
	if id == "" {
		return nil, errors.New("imdb.getperson id did not match regex")
	}

	var person Person

	if !c.disabledCaching {
		if err := c.cache.PersonCache.Load(id, &person); err == nil {
			return &person, nil
		}
	}

	url := personBaseURL + "/" + id

	doc, err := doRequest(url)
	if err != nil {
		return nil, err
	}

	person = Person{
		ID: id,
	}

	jsonDataNode := htmlquery.FindOne(doc, "//script[@type='application/ld+json']")
	if jsonDataNode == nil {
		return nil, errors.New("json data node not found")
	}

	err = json.Unmarshal([]byte(htmlquery.InnerText(jsonDataNode)), &person.PersonJSONContent)
	if err != nil {
		return nil, errors.New("failed to unmarshal results data")
	}

	detailsNode, err := htmlquery.Query(doc, "//section[@data-testid='PersonalDetails']/div[2]/ul")
	if detailsNode != nil && err == nil {
		err = encode.Xpath(detailsNode, &person.PersonDetailsSection)
		if err != nil {
			return nil, errors.Wrap(err, "error while scraping data with xpath")
		}
	}

	dykNode, err := htmlquery.Query(doc, "//section[@data-testid='DidYouKnow']")
	if dykNode != nil && err == nil {
		err = encode.Xpath(dykNode, &person.PersonDYKSection)
		if err != nil {
			return nil, errors.Wrap(err, "error while scraping data with xpath")
		}
	}

	knownForNode, err := htmlquery.Query(doc, "//div[@data-testid='Filmography']//div[@data-testid='nm_flmg_kwn_for']//div[ends-with(@data-testid, 'container')]")
	if knownForNode != nil && err == nil {
		person.KnownFor = encode.GetXpathLinks(knownForNode)
	}

	// Cache data for next time
	if !c.disabledCaching {
		err := c.cache.PersonCache.Save(id, person)
		if err != nil {
			log.Println(err)
		}
	}

	return &person, nil
}
