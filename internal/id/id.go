package id

import (
	"regexp"
	"strings"
)

type IDSource int

const (
	IDSourceUnknown IDSource = iota
	IDSourceImdb
	IDSourceJW
)

type IDType int

const (
	IDTypeUnknown IDType = iota
	IDTypeTitle
	IDTypePerson
	IDTypeURL
)

// Regexes for each engine
var (
	imdbTitleRegex   = regexp.MustCompile(`tt\d+`)
	imdbNameRegex    = regexp.MustCompile(`nm\d+`)
	justwatchIdRegex = regexp.MustCompile(`tm\d+`)
)

// RecognizeID uses regex to identify the source and type of an id.
func RecognizeID(input string) (IDSource, IDType) {
	switch {
	case imdbTitleRegex.MatchString(input):
		return IDSourceImdb, IDTypeTitle
	case imdbNameRegex.MatchString(input):
		return IDSourceImdb, IDTypePerson
	case justwatchIdRegex.MatchString(input):
		return IDSourceJW, IDTypeTitle
	case strings.Contains(input, "justwatch.com"): //TODO: use regex for safer checks
		return IDSourceJW, IDTypeURL
	default:
		return IDSourceUnknown, IDTypeUnknown
	}
}

// IsIMDbTitle indicates wether the id matches the imdb title id regex.
func IsIMDbTitle(input string) bool {
	return imdbTitleRegex.MatchString(input)
}
