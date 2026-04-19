package twitter

import (
	"testing"

	"github.com/nicholasgasior/lux/extractors/types"
)

func TestExtractTweetID(t *testing.T) {
	cases := []struct {
		url      string
		wantID   string
		wantErr  bool
	}{
		{"https://twitter.com/user/status/1234567890", "1234567890", false},
		{"https://x.com/user/status/9876543210", "9876543210", false},
		{"https://twitter.com/user/status/111?s=20", "111", false},
		{"https://example.com/not-a-tweet", "", true},
		{"", "", true},
	}

	for _, tc := range cases {
		got, err := extractTweetID(tc.url)
		if tc.wantErr {
			if err == nil {
				t.Errorf("extractTweetID(%q) expected error, got nil", tc.url)
			}
			continue
		}
		if err != nil {
			t.Errorf("extractTweetID(%q) unexpected error: %v", tc.url, err)
			continue
		}
		if got != tc.wantID {
			t.Errorf("extractTweetID(%q) = %q, want %q", tc.url, got, tc.wantID)
		}
	}
}

func TestNew(t *testing.T) {
	e := New()
	if e == nil {
		t.Fatal("New() returned nil")
	}
}

func TestExtract(t *testing.T) {
	e := New()
	data, err := e.Extract("https://twitter.com/user/status/1234567890", types.Options{})
	if err != nil {
		t.Fatalf("Extract() unexpected error: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Extract() returned no data")
	}
	if data[0].Site != "Twitter" {
		t.Errorf("Site = %q, want \"Twitter\"", data[0].Site)
	}
	if _, ok := data[0].Streams["default"]; !ok {
		t.Error("expected 'default' stream to be present")
	}
}
