//(c) Jisin0
// Justwatch base and constants.

package justwatch

import "github.com/machinebox/graphql"

const (
	apiUrl = "https://apis.justwatch.com/graphql" // justwatch graphql api url
)

var graphQLClient *graphql.Client

// Initialize stuff
func init() {
	graphQLClient = graphql.NewClient(apiUrl)
}
