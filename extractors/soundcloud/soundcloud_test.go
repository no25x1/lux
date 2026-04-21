package soundcloud

import (
	"testing"

	"github.com/nicoxiang/lux/extractors/types"
)

func TestExtractTrackInfo(t *testing.T) {
	tests := []struct {
		url        string
		artist     string
		track      string
		wantErr    bool
	}{
		{"https://soundcloud.com/artist-name/track-name", "artist-name", "track-name", false},
		{"https://soundcloud.com/some-artist/some-track", "some-artist", "some-track", false},
		// Trailing slash should not break parsing
		{"https://soundcloud.com/some-artist/some-track/", "some-artist", "some-track", false},
		// Query strings and fragments should also be handled gracefully
		{"https://soundcloud.com/some-artist/some-track?ref=clipboard", "some-artist", "some-track", false},
		{"https://example.com/invalid", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			artist, track, err := extractTrackInfo(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractTrackInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if artist != tt.artist || track != tt.track {
				t.Errorf("extractTrackInfo() = (%v, %v), want (%v, %v)", artist, track, tt.artist, tt.track)
			}
		})
	}
}

func TestNew(t *testing.T) {
	e := New()
	if e == nil {
		t.Error("New() returned nil")
	}
}

func TestExtract(t *testing.T) {
	e := New()
	_, err := e.Extract("https://example.com/invalid", types.Options{})
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}
