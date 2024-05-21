// (c) Jisin0
// Create url parameters from struct.
package encode

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// UrlParams function to encode struct fields into URL parameters
func UrlParams(params interface{}) (string, error) {

	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("UrlParams: input is not a struct")
	}

	values := url.Values{}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("url")
		if tag == "" {
			tag = strings.ToLower(t.Field(i).Name)
		}

		// Handle slices separately
		if field.Kind() == reflect.Slice {
			slice := []string{}
			for j := 0; j < field.Len(); j++ {
				slice = append(slice, fmt.Sprintf("%v", field.Index(j).Interface()))
			}
			values.Set(tag, strings.Join(slice, ","))
		} else {
			values.Set(tag, fmt.Sprintf("%v", field.Interface()))
		}
	}

	return values.Encode(), nil
}
