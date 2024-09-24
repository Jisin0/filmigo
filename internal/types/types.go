// (c) Jisin0

/* Global types and functions used across packages. */
package types

import "strings"

// A url object i.e text + href
type Link struct {
	// The actual content
	Text string `json:"name"`
	// The url or href
	Href string `json:"url"`
}

// A list of links.
type Links []Link

// Concatenate the Text of the links into a list separated by given separator.
func (ls Links) ToString(sep string) string {
	var s []string
	for _, l := range ls {
		if l.Text == "" {
			continue
		}

		s = append(s, l.Text)
	}

	return strings.Join(s, sep)
}

// A range value with a start and end used for advanced search queries.
// Both start and end MUST be set when using this value.
type SearchRange struct {
	Start string
	End   string
}
