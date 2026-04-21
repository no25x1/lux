// Package instagram provides an extractor for Instagram posts and reels.
package instagram

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iawia002/lux/extractors/types"
)

var (
	// Matches Instagram post/reel/tv URLs and extracts the shortcode.
	// Also supports /stories/ URLs, e.g. instagram.com/stories/username/12345/
	postPattern = regexp.MustCompile(`instagram\.com/(?:p|reel|tv|stories/[^/]+)/([A-Za-z0-9_-]+)`)
)

// Extractor handles Instagram URL extraction.
type Extractor struct{}

// New returns a new Instagram Extractor.
func New() *Extractor {
	return &Extractor{}
}

// Extract extracts media info from an Instagram URL.
func (e *Extractor) Extract(url string, option types.Options) ([]*types.Data, error) {
	shortcode, err := extractShortcode(url)
	if err != nil {
		return nil, err
	}

	// Build the oEmbed endpoint URL for metadata
	oembedURL := fmt.Sprintf("https://www.instagram.com/oembed/?url=https://www.instagram.com/p/%s/", shortcode)
	_ = oembedURL

	// TODO: perform HTTP request to oEmbed endpoint and parse response
	// For now return a stub indicating the shortcode was resolved
	data := &types.Data{
		Site:  "Instagram",
		Title: shortcode,
		URL:   normalizeURL(shortcode),
		Streams: map[string]*types.Stream{
			"default": {
				ID:      "default",
				Quality: "unknown",
				Parts:   []*types.Part{},
			},
		},
	}

	return []*types.Data{data}, nil
}

// extractShortcode parses the Instagram shortcode from a post/reel/tv/stories URL.
func extractShortcode(url string) (string, error) {
	matches := postPattern.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "", fmt.Errorf("instagram: unable to extract shortcode from URL: %s", url)
	}
	return matches[1], nil
}

// normalizeURL returns the canonical Instagram post URL for a given shortcode.
func normalizeURL(shortcode string) string {
	return "https://www.instagram.com/p/" + strings.TrimSpace(shortcode) + "/"
}
