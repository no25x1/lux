package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		url      string
		expected string
		wantErr  bool
	}{
		{
			url:      "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			expected: "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			url:      "https://youtu.be/dQw4w9WgXcQ",
			expected: "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			url:     "https://www.youtube.com/",
			wantErr: true,
		},
		{
			url:     "https://example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			id, err := extractVideoID(tt.url)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, id)
			}
		})
	}
}

func TestExtractTitle(t *testing.T) {
	html := "<html><head><title>Rick Astley - Never Gonna Give You Up - YouTube</title></head></html>"
	title := extractTitle(html)
	assert.Equal(t, "Rick Astley - Never Gonna Give You Up", title)
}

func TestExtractTitleMissing(t *testing.T) {
	html := "<html><head></head></html>"
	title := extractTitle(html)
	assert.Equal(t, "Unknown Title", title)
}

func TestNew(t *testing.T) {
	e := New()
	assert.NotNil(t, e)
}
