package twitter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nicholasgasior/lux/extractors/types"
)

var tweetURLPattern = regexp.MustCompile(`(?:twitter\.com|x\.com)/[^/]+/status/(\d+)`)

type extractor struct{}

// New returns a new Twitter extractor.
func New() types.Extractor {
	return &extractor{}
}

func (e *extractor) Extract(url string, option types.Options) ([]*types.Data, error) {
	videoID, err := extractTweetID(url)
	if err != nil {
		return nil, err
	}

	// In a real implementation this would call the Twitter/X API or
	// use a guest-token flow to retrieve the media URL.
	// Here we return a minimal stub so the extractor is wired up correctly.
	// NOTE: video.twimg.com URLs require a valid auth token in practice;
	// this stub is useful for testing the wiring but won't download real media.
	data := &types.Data{
		Site:  "Twitter",
		Title: fmt.Sprintf("tweet_%s", videoID),
		Type:  "video",
		Streams: map[string]*types.Stream{
			"default": {
				ID:      "default",
				Quality: "hd",
				URLs: []*types.URL{
					{
						URL:  fmt.Sprintf("https://video.twimg.com/tweet_video/%s.mp4", videoID),
						Ext:  "mp4",
						Size: 0,
					},
				},
			},
		},
	}
	return []*types.Data{data}, nil
}

// extractTweetID parses the numeric tweet/status ID from a Twitter or X URL.
// Both twitter.com and x.com domains are supported.
func extractTweetID(url string) (string, error) {
	url = strings.TrimSpace(url)
	matches := tweetURLPattern.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("twitter: cannot extract tweet ID from URL: %s", url)
	}
	return matches[1], nil
}
