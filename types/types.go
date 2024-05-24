// (c) Jisin0
// Global types and functions used across packages.
package types

// A url object i.e text + href
type Link struct {

	//The actual content
	Text string

	//The url or href
	Href string
}

// A list of links.
type Links []Link

// A range value with a start and end used for advanced search queries.
// Both start and end MUST be set when using this value.
type SearchRange struct {
	Start string
	End   string
}
