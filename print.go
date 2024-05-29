//(c) Jisin0
// Pretty print structs to console.

package filmigo

import (
	"encoding/json"
	"fmt"
)

// PrintJson prints out a struct in json.
func PrintJson(val interface{}, indent string) {
	// Marshal the struct to JSON with indentation
	jsonData, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// Print the formatted JSON
	fmt.Println(string(jsonData))
}
