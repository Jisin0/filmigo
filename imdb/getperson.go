// (c)Jisin0
// Functions and types to scrape data of a person.

package imdb

import (
	"net/http"

	"github.com/Jisin0/filmigo/encode"
	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"github.com/go-faster/errors"
)

const (
	//path for homepage of any imdb movie/show
	personBaseUrl = baseImdbURL + "/name"
)

// Type containing the full data about a person scraped from their imdb page.
type Person struct {

	//Imdb id of the user for ex: nm0000129
	Id string

	//Url to the person's imdb profile in the format imdb.com/name/{id}
	Link string

	//Full name of the person
	Name string `xpath:"//h1[@data-testid='hero__pageTitle']/span"`

	//List of roles performed by the person for ex: actor, producer, director etc.
	Roles []string `xpath:"//h1[@data-testid='hero__pageTitle']/../ul|textlist"`

	//Short bio of the person.
	Bio string `xpath:"//div[@data-testid='bio-content']//div[contains(@class, 'inner')]"`

	//Poster image of the person.
	Poster string `xpath:"//div[starts-with(@class, 'ipc-poster')]//img|attr_src"`

	//Links to movies/show the person is known for.
	KnownFor types.Links `xpath:"//div[@data-testid='Filmography']//div[@data-testid='nm_flmg_kwn_for']//div[ends-with(@data-testid, 'container')]|linklist"`

	//Personal details section

	//Official sites of the person.
	OfficialSites types.Links `xpath:"//section[@data-testid='PersonalDetails']/div[2]/ul/li[@data-testid='details-officialsites']/div|linklist"`

	//Height of the person.
	Height string `xpath:"//section[@data-testid='PersonalDetails']/div[2]/ul/li[@data-testid='nm_pd_he']/div//span"`

	//Date of birth . for ex : April 30, 1981
	Birthday string `xpath:"//section[@data-testid='PersonalDetails']/div[2]/ul/li[@data-testid='nm_pd_bl']/div/ul/li"`

	//Spouse of the person.
	Spouse types.Links `xpath:"//section[@data-testid='PersonalDetails']/div[2]/ul/li[@data-testid='nm_pd_sp']/div|linklist"`

	//Other works - usually a short sentence about a different work of the person.
	OtherWorks string `xpath:"//section[@data-testid='PersonalDetails']/div[2]/ul/li[@data-testid='nm_pd_wrk']/div"`

	//Did You Know section

	//A short trivia fact about the person.
	Trivia string `xpath:"//section[@data-testid='DidYouKnow']/div[2]//li[@data-testid='name-dyk-trivia']/div"` //Trivia is always the first dyk hence the div[2]

	//A popular quote of the person. All quotes can be found at {link}/quotes.
	Quote string `xpath:"//section[@data-testid='DidYouKnow']//li[@data-testid='name-dyk-quote']/div"`

	//A nickname of the person.
	Nickname string `xpath:"//section[@data-testid='DidYouKnow']//li[@data-testid='name-dyk-nickname']/div"`

	//Any trademark features of the person.
	Trademark string `xpath:"//section[@data-testid='DidYouKnow']//li[@data-testid='name-dyk-trademarks']/div"`
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

	url := personBaseUrl + "/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	//Return error on bas status codes
	if resp.StatusCode == 404 {
		return nil, errors.Errorf("movie/show with id %s was not not found", id)
	} else if resp.StatusCode != 200 {
		return nil, errors.Errorf("%v bad status code returned", resp.StatusCode)
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse document")
	}

	person = Person{
		Id:   id,
		Link: url,
	}

	person = encode.Xpath(doc, person).(Person)

	//Cache data for next time
	if !c.disabledCaching {
		c.cache.PersonCache.Save(id, person)
	}

	resp.Body.Close()

	return &person, nil

}
