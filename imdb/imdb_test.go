package imdb_test

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/Jisin0/filmigo/imdb"
)

// base client used for tests
var c *imdb.ImdbClient = imdb.NewClient()

// Gets the html content returned for a webpage using basic configs and writes to a html file
func TestGetWebpage(t *testing.T) {
	url := "https://www.justwatch.com/in/movie/chemmeen"

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		t.Errorf("failed to build request %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Set("languages", "en-us,en;q=0.5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("failed to make request %v", err)
	}

	data, _ := io.ReadAll(resp.Body)
	os.WriteFile("output.html", data, 0o644)

	resp.Body.Close()

	t.Log("output written to file")
}
