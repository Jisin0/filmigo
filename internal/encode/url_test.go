package encode_test

import (
	"testing"

	"github.com/Jisin0/filmigo/internal/encode"
	"github.com/Jisin0/filmigo/internal/types"
)

func TestURLParams(t *testing.T) {
	type sample struct {
		Range types.SearchRange `url:"release_date"`
		List  []string          `url:"characters"`
	}

	data := sample{
		Range: types.SearchRange{Start: "1970-02-08", End: "2024-12-09"},
		List:  []string{"oppenheimer"},
	}

	v, e := encode.URLParams(data)
	if e != nil {
		t.Error(e)
	}

	t.Logf("%+v", v)
}
