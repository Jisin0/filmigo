package imdb

import "strings"

// parseIMDbDuration parses the time returned from imdb into human-readable format. for ex: PT2H1M -> 2h 1min
func parseIMDbDuration(s string) string {
	s = strings.ReplaceAll(s, "PT", "")
	s = strings.ReplaceAll(s, "H", "h ")
	s = strings.ReplaceAll(s, "M", "min ")
	s = strings.ReplaceAll(s, "S", "s")

	return s
}
