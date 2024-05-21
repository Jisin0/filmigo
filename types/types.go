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
