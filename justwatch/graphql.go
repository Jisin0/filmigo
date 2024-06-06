// (c) Jisin0
// Graphql queries for each operetion.

package justwatch

import (
	_ "embed"
)

// load graphql queries from file.

//go:embed queries/searchtitle.graphql
var searchTitleQuery string

//go:embed queries/gettitle.graphql
var getTitleQuery string

//go:embed queries/gettitleurl.graphql
var getTitleFromURLQuery string

//go:embed queries/gettitleoffers.graphql
var getTitleOffersQuery string
