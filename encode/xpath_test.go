package encode_test

import (
	"testing"

	"github.com/Jisin0/filmigo/encode"
	"github.com/Jisin0/filmigo/types"
	"github.com/antchfx/htmlquery"
)

func TestXpath(t *testing.T) {
	doc, err := htmlquery.LoadDoc("sample.html")
	if err != nil {
		t.Errorf("failed to open sample file : %v", err)
		t.FailNow()
	}

	if doc != nil {
		type sampleData struct {
			InnerText  string      `xpath:"//p[contains(@class, 'substring')]"`
			Attribute  string      `xpath:"//p[last()]|attr_my-attr"`
			LinkList   types.Links `xpath:"//span[@class='sample']"`
			StringList []string    `xpath:"//ul"`
		}

		res, ok := encode.Xpath(doc, sampleData{}).(sampleData)
		if !ok {
			t.Errorf("unknown type returned")
			t.FailNow()
		}

		if res.Attribute == "" || res.InnerText == "" || len(res.LinkList) < 3 || len(res.StringList) < 3 {
			t.Errorf("xpath failed with output : %+v", res)
			t.FailNow()
		}
	}
}
