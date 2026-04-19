package youtube

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nicholasgasior/lux/extractors/types"
	"github.com/nicholasgasior/lux/request"
)

var (
	videoIDRegex = regexp.MustCompile(`(?:v=|youtu\.be/)([a-zA-Z0-9_-]{11})`)
)

// Extractor is the YouTube extractor.
type Extractor struct{}

// New returns a new YouTube extractor.
func New() *Extractor {
	return &Extractor{}
}

// Extract extracts video data from a YouTube URL.
func (e *Extractor) Extract(url string, option types.Options) ([]*types.Data, error) {
	videoID, err := extractVideoID(url)
	if err != nil {
		return nil, err
	}

	apiURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	html, err := request.Get(apiURL, url, nil)
	if err != nil {
		return nil, fmt.Errorf("youtube: failed to fetch page: %w", err)
	}

	title := extractTitle(html)

	streams := map[string]*types.Stream{
		"default": {
			ID:      "default",
			Quality: "720p",
			Parts: []*types.Part{
				{
					URL:  apiURL,
					Size: 0,
					Ext:  "mp4",
				},
			},
			Size: 0,
		},
	}

	return []*types.Data{
		{
			Site:    "YouTube youtube.com",
			Title:   title,
			Type:    "video",
			Streams: streams,
			URL:     url,
		},
	}, nil
}

func extractVideoID(url string) (string, error) {
	matches := videoIDRegex.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("youtube: unable to extract video ID from URL: %s", url)
	}
	return matches[1], nil
}

func extractTitle(html string) string {
	start := strings.Index(html, "<title>")
	end := strings.Index(html, "</title>")
	if start == -1 || end == -1 || end <= start {
		return "Unknown Title"
	}
	title := html[start+7 : end]
	title = strings.TrimSuffix(title, " - YouTube")
	return strings.TrimSpace(title)
}
