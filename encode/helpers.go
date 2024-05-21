// (c) Jisin0
// Helper methods for xpath configs.
package encode

import (
	"strings"

	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func getLinks(node *html.Node) types.Links {

	ls, e := htmlquery.QueryAll(node, ".//a")
	if e != nil || len(ls) < 1 {
		return []types.Link{}
	}

	var links types.Links

	for _, l := range ls {
		var href string
		text := htmlquery.InnerText(l)
		if text == "" {
			continue
		}

		for _, a := range l.Attr {
			if a.Key == "href" {
				href = a.Val
			}
		}

		//Add imdb base url if href is a url path
		if strings.HasPrefix(href, "/") {
			href = "https://imdb.com" + href
		}

		links = append(links, types.Link{Text: text, Href: href})
	}

	return links
}

func getTextList(node *html.Node) []string {

	ls, e := htmlquery.QueryAll(node, ".//li")
	if e != nil || len(ls) < 1 {
		return []string{}
	}

	var list []string

	for _, l := range ls {
		content := htmlquery.InnerText(l)
		if content != "" {
			list = append(list, content)
		}
	}

	return list
}
