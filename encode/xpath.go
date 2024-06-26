// (c) Jisin0
// Helper methods to scrape a webpage using xpath in struct tags.

package encode

import (
	"errors"
	"reflect"
	"strings"

	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var linkStructType = reflect.TypeOf(types.Link{})

// Function to scrape any html document with xpath provided in struct tags:
//
// - doc *html.Node : The base document or node at which all queries start.
//
// - val interface : input data type.
//
// See https://github.com/Jisin/filmigo/xpath for examples and full reference.
func Xpath(doc *html.Node, target any) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("input type is not a pointer")
	}

	st := reflect.TypeOf(target).Elem()

	// Access the struct value within the interface
	v := reflect.ValueOf(target).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)

		args := strings.Split(field.Tag.Get("xpath"), "|")
		if len(args) < 1 {
			continue
		}

		var (
			method string
			attr   string
		)

		if len(args) > 1 {
			method = args[1]
			if strings.HasPrefix(method, "attr") {
				if a := strings.SplitN(method, "_", 2); len(a) > 1 {
					method = "attr"
					attr = a[1]
				}
			}
		}

		path := args[0]

		node, err := htmlquery.Query(doc, path)
		if err != nil || node == nil {
			continue
		}

		fieldType := field.Type

		switch method {
		case "attr":
			for _, a := range node.Attr {
				if a.Key == attr {
					v.FieldByName(field.Name).SetString(a.Val)
					break
				}
			}

		default:
			// If the field is of type []Link all inner a tags are extracted
			// If field type is []string innertex of each li tag is extracted
			if fieldType.Kind() == reflect.Slice && fieldType.Elem() == linkStructType {
				links := GetXpathLinks(node)
				lVal := reflect.Append(reflect.ValueOf(links))
				v.FieldByName(field.Name).Set(lVal)
			} else if fieldType.Kind() == reflect.Slice && fieldType.Elem().Kind() == reflect.String {
				list := getTextList(node)
				lVal := reflect.Append(reflect.ValueOf(list))
				v.FieldByName(field.Name).Set(lVal)
			} else {
				v.FieldByName(field.Name).SetString(htmlquery.InnerText(node))
			}
		}
	}

	return nil
}
