package soundcloud

import (
	"fmt"
	"regexp"

	"github.com/nicoxiang/lux/extractors/types"
)

// trackRegex matches SoundCloud URLs in the format soundcloud.com/artist/track.
// Note: this does not handle sets/playlists (e.g. soundcloud.com/artist/sets/playlist).
var trackRegex = regexp.MustCompile(`soundcloud\.com/([\w-]+)/([\w-]+)`)

type extractor struct{}

// New returns a new SoundCloud extractor.
func New() types.Extractor {
	return &extractor{}
}

func (e *extractor) Extract(url string, opts types.Options) ([]*types.Data, error) {
	artist, track, err := extractTrackInfo(url)
	if err != nil {
		return nil, err
	}

	// Format title as "Artist - Track" for consistency with other extractors.
	title := fmt.Sprintf("%s - %s", artist, track)

	return []*types.Data{
		{
			Site:  "soundcloud.com",
			Title: title,
			Type:  "audio",
		},
	}, nil
}

func extractTrackInfo(url string) (artist, track string, err error) {
	matches := trackRegex.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", "", fmt.Errorf("unable to extract track info from URL: %s", url)
	}
	return matches[1], matches[2], nil
}
