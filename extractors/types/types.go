package types

// Options holds extraction options passed to extractors.
type Options struct {
	PlaylistStart  int
	PlaylistEnd    int
	Items          string
	ItemStart      int
	ItemEnd        int
	ThreadNumber   int
	Silent         bool
	InfoOnly       bool
	Stream         string
	OutputPath     string
	OutputName     string
	FileNameLength int
	Caption        bool
}

// Part represents a single downloadable part of a stream.
type Part struct {
	URL  string
	Size int64
	Ext  string
}

// Stream represents a single quality stream with one or more parts.
type Stream struct {
	ID      string
	Quality string
	Parts   []*Part
	Size    int64
	Ext     string
}

// Data represents the extracted data for a single video/audio resource.
type Data struct {
	Site    string
	Title   string
	Type    string
	Streams  map[string]*Stream
	Caption map[string]*Caption
	URL     string
}

// Caption holds subtitle/caption data.
type Caption struct {
	URL  string
	Ext  string
	Data string
	// Language holds the BCP 47 language tag for this caption (e.g. "en", "zh-Hans").
	Language string
}
