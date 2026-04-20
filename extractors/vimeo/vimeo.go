package vimeo

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/nicholasgasior/lux/extractors/types"
)

var videoIDRegex = regexp.MustCompile(`vimeo\.com/(\d+)`)

type extractor struct {
	client *http.Client
}

func New() types.Extractor {
	return &extractor{
		// Use a timeout to avoid hanging indefinitely on slow connections
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (e *extractor) Extract(url string, opts types.Options) (*types.Data, error) {
	videoid, err := extractVideoID(url)
	if err != nil {
		return nil, err
	}

	apiURL := fmt.Sprintf("https://vimeo.com/api/v2/video/%s.json", videoid)
	resp, err := e.client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("vimeo: failed to fetch video info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("vimeo: unexpected status %d for video %s", resp.StatusCode, videoid)
	}

	return &types.Data{
		Site:  "Vimeo",
		Title: videoid,
		URL:   url,
	}, nil
}

func extractVideoID(url string) (string, error) {
	matches := videoIDRegex.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("vimeo: could not extract video ID from URL: %s", url)
	}
	return strings.TrimSpace(matches[1]), nil
}
