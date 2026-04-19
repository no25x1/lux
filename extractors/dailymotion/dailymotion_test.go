package dailymotion

import (
	"testing"

	"github.com/nicoxiang/lux/extractors/types"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		url      string
		expected string
		wantErr  bool
	}{
		{"https://www.dailymotion.com/video/x7tgd2g", "x7tgd2g", false},
		{"https://www.dailymotion.com/video/x7tgd2g_some-title", "x7tgd2g", false},
		{"https://dai.ly/x7tgd2g", "x7tgd2g", false},
		{"https://www.dailymotion.com/embed/video/x7tgd2g", "x7tgd2g", false},
		{"https://example.com/watch?v=abc", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			got, err := extractVideoID(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("extractVideoID() = %v, want %v", got, tt.expected)
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
