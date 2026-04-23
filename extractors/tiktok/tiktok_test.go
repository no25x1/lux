package tiktok

import (
	"testing"

	"github.com/iawia002/lux/extractors/types"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantID  string
		wantErr bool
	}{
		{
			name:   "standard video URL",
			url:    "https://www.tiktok.com/@username/video/7123456789012345678",
			wantID: "7123456789012345678",
		},
		{
			name:   "mobile video URL",
			url:    "https://m.tiktok.com/v/7123456789012345678.html",
			wantID: "7123456789012345678",
		},
		{
			name:    "invalid URL",
			url:     "https://www.tiktok.com/@username",
			wantErr: true,
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: true,
		},
		{
			name:    "non-tiktok URL",
			url:     "https://www.youtube.com/watch?v=abc123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := extractVideoID(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotID != tt.wantID {
				t.Errorf("extractVideoID() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestNew(t *testing.T) {
	extractor := New()
	if extractor == nil {
		t.Fatal("New() returned nil")
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "invalid URL returns error",
			url:     "https://www.tiktok.com/@username",
			wantErr: true,
		},
		{
			name:    "empty URL returns error",
			url:     "",
			wantErr: true,
		},
	}

	extractor := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := extractor.Extract(tt.url, types.Options{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
