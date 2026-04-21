package bilibili

import (
	"testing"

	"github.com/nicoxiang/lux/extractors/types"
	"github.com/stretchr/testify/assert"
)

func TestBVIDExtraction(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "standard bilibili URL",
			url:      "https://www.bilibili.com/video/BV1xx411c7mD",
			expected: "BV1xx411c7mD",
		},
		{
			name:     "bilibili URL with query params",
			url:      "https://www.bilibili.com/video/BV1GJ411x7h7?from=search",
			expected: "BV1GJ411x7h7",
		},
		{
			name:     "bilibili URL with timestamp param",
			url:      "https://www.bilibili.com/video/BV1xx411c7mD?t=42",
			expected: "BV1xx411c7mD",
		},
		{
			name:     "no BVID",
			url:      "https://www.bilibili.com/",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bvidRegex.FindString(tt.url)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestExtractorNew(t *testing.T) {
	e := New()
	assert.NotNil(t, e)
}

func TestExtractInvalidURL(t *testing.T) {
	e := New()
	_, err := e.Extract("https://www.bilibili.com/", types.Options{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unable to find BVID")
}
