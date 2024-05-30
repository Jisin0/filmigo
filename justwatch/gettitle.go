// (c) Jisin0
// Get the full details of a title.

package justwatch

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/machinebox/graphql"
)

// Options for GetTitleURL() operation.
type GetTitleOptions struct {
	// Country code of country of request. for ex: US.
	Country string
	// Lnaguage code for results. for ex: en.
	Language string
	// Maximum number of episodes to return.
	EpisodeMaxLimit int
}

// Get the full details of a title using it's justwatch id.
//
// - id : The unique justwatch id of the entity.
func (c *JustwatchClient) GetTitle(id string, opts ...*GetTitleOptions) (*Title, error) {
	request := graphql.NewRequest(getTitleQuery)

	var (
		language        = c.LangCode
		country         = c.Country
		episodeMaxLimit = 20
	)

	// Custom options
	if len(opts) > 0 {
		o := opts[0]

		if o.Country != "" {
			country = o.Country
		}

		if o.Language != "" {
			language = o.Language
		}

		if o.EpisodeMaxLimit != 0 {
			episodeMaxLimit = o.EpisodeMaxLimit
		}
	}

	// Set the variables
	request.Var("entityId", id)
	request.Var("fullPath", "/")
	request.Var("country", country)
	request.Var("language", language)
	request.Var("episodeMaxLimit", episodeMaxLimit)
	request.Var("platform", "WEB")
	request.Var("allowSponsoredRecommendations", map[string]interface{}{
		"pageType":                "VIEW_TITLE_DETAIL",
		"placement":               "DETAIL_PAGE",
		"country":                 "US",
		"language":                "en",
		"appId":                   "3.8.2-webapp#62adb00",
		"platform":                "WEB",
		"supportedFormats":        []string{"IMAGE", "VIDEO"},
		"supportedObjectTypes":    []string{"MOVIE", "SHOW", "GENERIC_TITLE_LIST", "SHOW_SEASON"},
		"testingMode":             false,
		"testingModeCampaignName": nil,
	})

	var response struct {
		Data Title `json:"node"`
	}

	if err := graphQLClient.Run(context.Background(), request, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// Get the full details of a title using it's url path.
//
// - path : URL path returned from a search result or the justwatch link
func (c *JustwatchClient) GetTitleFromURL(path string, opts ...*GetTitleOptions) (*URLDetails, error) {
	request := graphql.NewRequest(getTitleFromURLQuery)

	// Cleaup path
	if strings.Contains(path, "justwatch.com") {
		if !strings.HasPrefix(path, "https://") {
			path = "https://" + path
		}

		parsedURL, err := url.Parse(path)
		if err != nil {
			return nil, err
		}

		path = parsedURL.Path
		fmt.Println(path)
	}

	var (
		language        = c.LangCode
		country         = c.Country
		episodeMaxLimit = 20
	)

	// Custom options
	if len(opts) > 0 {
		o := opts[0]

		if o.Country != "" {
			country = o.Country
		}

		if o.Language != "" {
			language = o.Language
		}

		if o.EpisodeMaxLimit != 0 {
			episodeMaxLimit = o.EpisodeMaxLimit
		}
	}

	// Set the variables
	request.Var("fullPath", path)
	request.Var("country", country)
	request.Var("language", language)
	request.Var("episodeMaxLimit", episodeMaxLimit)
	request.Var("platform", "WEB")
	request.Var("allowSponsoredRecommendations", map[string]interface{}{
		"pageType":                "VIEW_TITLE_DETAIL",
		"placement":               "DETAIL_PAGE",
		"country":                 "US",
		"language":                "en",
		"appId":                   "3.8.2-webapp", // works even if omitted
		"platform":                "WEB",
		"supportedFormats":        []string{"IMAGE", "VIDEO"},
		"supportedObjectTypes":    []string{"MOVIE", "SHOW", "GENERIC_TITLE_LIST", "SHOW_SEASON"},
		"testingMode":             false,
		"testingModeCampaignName": nil,
	})

	var response struct {
		Data URLDetails `json:"urlV2"`
	}

	if err := graphQLClient.Run(context.Background(), request, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// Get the offers available for a url using it's Justwatch ID.
//
// - id: Justwatch id of the title.
func (c *JustwatchClient) GetTitleOffers(id string, opts ...*GetTitleOptions) (*GetTitleOffersResult, error) {
	request := graphql.NewRequest(getTitleOffersQuery)

	var (
		language = c.LangCode
		country  = c.Country
	)

	// Custom options
	if len(opts) > 0 {
		o := opts[0]

		if o.Country != "" {
			country = o.Country
		}

		if o.Language != "" {
			language = o.Language
		}
	}

	// Set the variables
	request.Var("nodeId", id)
	request.Var("country", country)
	request.Var("language", language)
	request.Var("platform", "WEB")

	request.Var("filterBuy", map[string]interface{}{
		"bestOnly":          true,
		"monetizationTypes": []string{"BUY"},
	})
	request.Var("filterFlatrate", map[string]interface{}{
		"bestOnly":          true,
		"monetizationTypes": []string{"FLATRATE", "FLATRATE_AND_BUY", "ADS", "FREE", "CINEMA"},
	})
	request.Var("filterFree", map[string]interface{}{
		"bestOnly":          true,
		"monetizationTypes": []string{"ADS", "FREE"},
	})
	request.Var("filterRent", map[string]interface{}{
		"bestOnly":          true,
		"monetizationTypes": []string{"RENT"},
	})

	var response struct {
		Data GetTitleOffersResult `json:"node"`
	}

	if err := graphQLClient.Run(context.Background(), request, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
