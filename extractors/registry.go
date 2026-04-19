package extractors

import (
	"fmt"
	"strings"

	"github.com/nicoxiang/lux/extractors/bilibili"
	"github.com/nicoxiang/lux/extractors/dailymotion"
	"github.com/nicoxiang/lux/extractors/soundcloud"
	"github.com/nicoxiang/lux/extractors/twitter"
	"github.com/nicoxiang/lux/extractors/types"
	"github.com/nicoxiang/lux/extractors/vimeo"
	"github.com/nicoxiang/lux/extractors/youtube"
)

var registry = map[string]types.Extractor{}

func init() {
	registry["youtube.com"] = youtube.New()
	registry["youtu.be"] = youtube.New()
	registry["bilibili.com"] = bilibili.New()
	registry["twitter.com"] = twitter.New()
	registry["x.com"] = twitter.New()
	registry["vimeo.com"] = vimeo.New()
	registry["dailymotion.com"] = dailymotion.New()
	registry["dai.ly"] = dailymotion.New()
	registry["soundcloud.com"] = soundcloud.New()
}

// GetExtractor returns the extractor for the given URL.
func GetExtractor(url string) (types.Extractor, error) {
	for domain, ext := range registry {
		if strings.Contains(url, domain) {
			return ext, nil
		}
	}
	return nil, fmt.Errorf("no extractor found for URL: %s", url)
}

// Extract extracts data from the given URL.
func Extract(url string, opts types.Options) ([]*types.Data, error) {
	ext, err := GetExtractor(url)
	if err != nil {
		return nil, err
	}
	return ext.Extract(url, opts)
}
