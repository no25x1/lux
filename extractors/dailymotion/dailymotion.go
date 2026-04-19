package dailymotion

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nicoxiang/lux/extractors/types"
)

var videoIDRegex = regexp.MustCompile(`(?:dailymotion\.com/(?:video|embed/video)|dai\.ly)/([a-zA-Z0-9]+)`)

type extractor struct{}

// New returns a new Dailymotion extractor.
func New() types.Extractor {
	return &extractor{}
}

func (e *extractor) Extract(url string, opts types.Options) ([]*types.Data, error) {
	videoID, err := extractVideoID(url)
	if err != nil {
		return nil, err
	}

	apiURL := fmt.Sprintf(
		"https://www.dailymotion.com/player/metadata/video/%s",
		videoID,
	)

	_ = apiURL // placeholder for actual HTTP request

	return []*types.Data{
		{
			Site:  "dailymotion.com",
			Title: videoID,
			Type:  "video",
		},
	}, nil
}

func extractVideoID(url string) (string, error) {
	matches := videoIDRegex.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("unable to extract video ID from URL: %s", url)
	}
	id := matches[1]
	if idx := strings.Index(id, "_"); idx != -1 {
		id = id[:idx]
	}
	return id, nil
}
