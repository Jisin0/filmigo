// (c) Jisin0
// Helper methods to scrape a webpage using xpath in struct tags.
package encode

import (
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
// See https://github.com/Jisin/Filmigo/xpath for examples and full reference.
func Xpath(doc *html.Node, val any) any {

	st := reflect.TypeOf(val)

	//https://stackoverflow.com/questions/63421976
	// v is the interface{}
	v := reflect.ValueOf(&val).Elem()
	// Allocate a temporary variable with type of the struct.
	//
	//	v.Elem() is the vale contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()
	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	for i := 0; i < tmp.NumField(); i++ {
		field := st.Field(i)
		args := strings.Split(field.Tag.Get("xpath"), "|")
		if len(args) < 1 {
			continue
		}

		path := args[0]
		var method string
		if len(args) > 1 {
			method = args[1]
		}

		node, err := htmlquery.Query(doc, path)
		if node == nil || err != nil {
			continue
		}

		fieldType := field.Type

		//Extra options are passed with a seperator | in the xpath struct tag, for ex. src to get the src attr of a node
		switch method {

		case "src":
			var src string
			for _, a := range node.Attr {
				if a.Key == "src" {
					src = a.Val
				}
			}
			tmp.FieldByName(field.Name).SetString(src)
		default:
			//If the field is of type []Link all inner a tags are extracted
			//If field type is []string innertex of each li tag is extracted
			if fieldType.Kind() == reflect.Slice && fieldType.Elem() == linkStructType {
				links := getLinks(node)
				lVal := reflect.Append(reflect.ValueOf(links))
				tmp.FieldByName(field.Name).Set(lVal)
			} else if fieldType.Kind() == reflect.Slice && fieldType.Elem().Kind() == reflect.String {
				list := getTextList(node)
				lVal := reflect.Append(reflect.ValueOf(list))
				tmp.FieldByName(field.Name).Set(lVal)
			} else {
				tmp.FieldByName(field.Name).SetString(htmlquery.InnerText(node))
			}
		}

	}

	v.Set(tmp)

	return val

}
