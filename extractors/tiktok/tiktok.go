// Package tiktok provides an extractor for TikTok videos.
package tiktok

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/iawia002/lux/extractors/types"
	"github.com/nicholasgasior/nextprev/httpclient"
)

var (
	// videoIDRegex matches TikTok video IDs from various URL formats.
	videoIDRegex = regexp.MustCompile(`/video/(\d+)`)

	// shortURLRegex matches vm.tiktok.com and vt.tiktok.com short URLs.
	shortURLRegex = regexp.MustCompile(`https?://(?:vm|vt)\.tiktok\.com/([A-Za-z0-9]+)`)
)

// Extractor implements the types.Extractor interface for TikTok.
type Extractor struct {
	client *httpclient.Client
}

// New returns a new TikTok extractor.
func New() *Extractor {
	return &Extractor{
		client: httpclient.NewClient(),
	}
}

// extractVideoID parses the TikTok video ID from a URL string.
// It handles both full URLs (/@user/video/12345) and short URLs (vm.tiktok.com/abc).
// Note: short URL slugs still need redirect resolution before hitting the API.
func extractVideoID(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("tiktok: invalid URL %q: %w", rawURL, err)
	}

	host := strings.ToLower(parsed.Host)

	// Handle short URLs — they redirect to the full URL; we extract the slug.
	if strings.Contains(host, "vm.tiktok.com") || strings.Contains(host, "vt.tiktok.com") {
		matches := shortURLRegex.FindStringSubmatch(rawURL)
		if len(matches) < 2 {
			return "", fmt.Errorf("tiktok: cannot extract short code from URL %q", rawURL)
		}
		// Return the short code; callers may need to resolve the redirect.
		return matches[1], nil
	}

	// Handle full TikTok URLs: https://www.tiktok.com/@user/video/1234567890
	matches := videoIDRegex.FindStringSubmatch(parsed.Path)
	if len(matches) < 2 {
		return "", fmt.Errorf("tiktok: cannot extract video ID from URL %q", rawURL)
	}
	return matches[1], nil
}

// Extract fetches metadata and stream information for the given TikTok URL.
func (e *Extractor) Extract(rawURL string, option types.Options) ([]*types.Data, error) {
	videoID, err := extractVideoID(rawURL)
	if err != nil {
		return nil, err
	}

	// TikTok's private API endpoint used by the mobile app.
	// Using version_code=262 — bump this if the API starts returning empty results.
	apiURL := fmt.Sprintf(
		"https://api.tiktok.com/aweme/v1/feed/?aweme_id=%s&version_code=262&app_name=tiktok_web",
		videoID,
	)

	var apiResp struct {
		AwemeList []struct {
			Desc  string `json:"desc"`
			Video struct {
				PlayAddr struct {
					URLList []string `json:"url_list"`
				} `json:"play_addr"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"video"`
		} `json:"aweme_list"`
	}

	if err := e.client.GetJSON(apiURL, &apiResp); err != nil {
		return nil, fmt.Errorf("tiktok: API request failed for video %s: %w", videoID, err)
	}

	if len(apiResp.AwemeList) == 0 {
		return nil, fmt.Errorf("tiktok: no video data returned for ID %s", videoID)
	}

	item := apiResp.AwemeList[0]
	if len(item.Video.PlayAddr.URLList) == 0 {
		return nil, fmt.Errorf("tiktok: no playable streams found for video %s", videoID)
	}

	streams := map[string]*types.Stream{
		"default": 
