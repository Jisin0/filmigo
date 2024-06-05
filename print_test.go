package filmigo_test

import (
	"testing"

	"github.com/Jisin0/filmigo"
)

func TestPrintJSON(t *testing.T) {
	type sampleType struct {
		Fruit  string `json:"fruit"`
		Colour string `json:"colour"`
		Weight int    `json:"weight"`
	}

	sampleData := []sampleType{
		{
			Fruit:  "Apple",
			Colour: "Red",
			Weight: 100,
		}, {
			Fruit:  "Pear",
			Colour: "Green",
			Weight: 75,
		}, {
			Fruit:  "Orange",
			Colour: "Orange",
			Weight: 150,
		},
	}

	outputData := map[string][]sampleType{"data": sampleData}

	filmigo.PrintJSON(outputData, "  ")
}
