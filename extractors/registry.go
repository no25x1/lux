package extractors

import (
	"fmt"
	"strings"

	"github.com/nicholasgasior/lux/extractors/bilibili"
	"github.com/nicholasgasior/lux/extractors/types"
	"github.com/nicholasgasior/lux/extractors/youtube"
)

// Extractor is the interface all site extractors must implement.
type Extractor interface {
	Extract(url string, option types.Options) ([]*types.Data, error)
}

var extractors = map[string]Extractor{}

func init() {
	extractors["bilibili.com"] = bilibili.New()
	extractors["youtube.com"] = youtube.New()
	extractors["youtu.be"] = youtube.New()
}

// GetExtractor returns the appropriate extractor for the given URL.
func GetExtractor(url string) (Extractor, error) {
	for domain, extractor := range extractors {
		if strings.Contains(url, domain) {
			return extractor, nil
		}
	}
	return nil, fmt.Errorf("extractors: no extractor found for URL: %s", url)
}

// Extract finds the right extractor and runs it against the given URL.
func Extract(url string, option types.Options) ([]*types.Data, error) {
	ext, err := GetExtractor(url)
	if err != nil {
		return nil, err
	}
	return ext.Extract(url, option)
}
