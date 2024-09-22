package cache_test

import (
	"os"
	"testing"
	"time"

	"github.com/Jisin0/filmigo/internal/cache"
)

func TestCache(t *testing.T) {
	directoryName := "testcache"
	c := cache.NewCache(directoryName, 1*time.Hour)

	type sampleType struct {
		String  string `json:"string"`
		Integer int    `json:"int"`
	}

	inputData := sampleType{
		String:  "foo",
		Integer: 3,
	}

	err := c.Save("uid", inputData)
	if err != nil {
		t.Error(err)
	}

	var outputData sampleType

	err = c.Load("uid", &outputData)
	if err != nil {
		t.Error(err)
	}

	if outputData != inputData {
		t.Logf("Input: %+v\nOutput: %+v", inputData, outputData)
	}

	err = os.RemoveAll(directoryName)
	if err != nil {
		t.Error(err)
	}
}
