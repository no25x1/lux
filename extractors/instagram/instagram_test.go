package instagram

import (
	"testing"
)

func TestExtractShortcode(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		want     string
		wantErr  bool
	}{
		{
			name: "standard post URL",
			url:  "https://www.instagram.com/p/ABC123def45/",
			want: "ABC123def45",
		},
		{
			name: "reel URL",
			url:  "https://www.instagram.com/reel/XYZ789ghi01/",
			want: "XYZ789ghi01",
		},
		{
			name: "TV URL",
			url:  "https://www.instagram.com/tv/MNO456jkl78/",
			want: "MNO456jkl78",
		},
		{
			name:    "invalid URL",
			url:     "https://www.instagram.com/username/",
			wantErr: true,
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractShortcode(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractShortcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "already normalized",
			url:  "https://www.instagram.com/p/ABC123/",
			want: "https://www.instagram.com/p/ABC123/",
		},
		{
			name: "without trailing slash",
			url:  "https://www.instagram.com/p/ABC123",
			want: "https://www.instagram.com/p/ABC123/",
		},
		{
			name: "with query params",
			url:  "https://www.instagram.com/p/ABC123/?utm_source=ig_web",
			want: "https://www.instagram.com/p/ABC123/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeURL(tt.url)
			if got != tt.want {
				t.Errorf("normalizeURL() = %v, want %v", got, tt.want)
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
