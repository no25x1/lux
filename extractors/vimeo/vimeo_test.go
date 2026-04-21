package vimeo

import (
	"testing"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		want    string
		wantErr bool
	}{
		{
			name: "standard url",
			url:  "https://vimeo.com/123456789",
			want: "123456789",
		},
		{
			name: "url with path suffix",
			url:  "https://vimeo.com/987654321/review",
			want: "987654321",
		},
		{
			// vimeo also supports channel/video paths like /channels/foo/123
			name: "channel url",
			url:  "https://vimeo.com/channels/staffpicks/112233445",
			want: "112233445",
		},
		{
			// groups urls follow the same pattern: /groups/foo/videos/123
			name: "groups url",
			url:  "https://vimeo.com/groups/motion/videos/556677889",
			want: "556677889",
		},
		{
			name:    "invalid url",
			url:     "https://example.com/watch?v=abc",
			wantErr: true,
		},
		{
			name:    "empty url",
			url:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVideoID(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractVideoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	e := New()
	if e == nil {
		t.Fatal("New() returned nil")
	}
}
