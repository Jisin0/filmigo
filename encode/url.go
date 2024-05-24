// (c) Jisin0
// Create url parameters from struct.
package encode

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/Jisin0/filmigo/types"
)

var searchRangeType = reflect.TypeOf(types.SearchRange{})

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
		if field.Type() == searchRangeType {
			start := field.FieldByName("Start").Interface()
			if start != "" {
				values.Set(tag, fmt.Sprintf("%v,%v", start, field.FieldByName("End").Interface()))
			}
		} else if field.Kind() == reflect.Slice {
			slice := []string{}
			for j := 0; j < field.Len(); j++ {
				slice = append(slice, fmt.Sprintf("%v", field.Index(j).Interface()))
			}
			if len(slice) > 0 {
				values.Set(tag, strings.Join(slice, ","))
			}
		} else {
			val := fmt.Sprintf("%v", field.Interface())
			if val != "" {
				values.Set(tag, val)
			}
		}
	}

	return values.Encode(), nil
}
