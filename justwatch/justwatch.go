// (c) Jisin0
// Justwatch base and constants.

package justwatch

import "github.com/machinebox/graphql"

const (
	apiURL = "https://apis.justwatch.com/graphql" // justwatch graphql api url
)

var graphQLClient *graphql.Client

// Initialize stuff
func init() {
	graphQLClient = graphql.NewClient(apiURL)
}

// Options for configuring default behaviour of the justwatch client.
type JustwatchClientOpts struct {
	// Coutry code to use for requests defaults to US.
	Country string
	// Language code to use for requests defaults to en.
	LangCode string
}

// Justwatch client through which api queries are executed.
type JustwatchClient struct {
	Country  string
	LangCode string
}

// Creates a new justwatch client through which all api queries are executed.
func NewClient(opts ...*JustwatchClientOpts) *JustwatchClient {
	countryCode := defaultCountryCode
	langCode := defaultLanguageCode

	if len(opts) > 0 {
		o := opts[0]
		if o.Country != "" {
			countryCode = o.Country
		}

		if o.LangCode != "" {
			langCode = o.LangCode
		}
	}

	return &JustwatchClient{
		Country:  countryCode,
		LangCode: langCode,
	}
}
