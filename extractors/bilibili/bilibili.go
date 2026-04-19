package bilibili

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/nicoxiang/lux/extractors/types"
	"github.com/nicoxiang/lux/request"
	"github.com/nicoxiang/lux/utils"
)

const (
	bilibiliAPIURL    = "https://api.bilibili.com/x/web-interface/view?bvid=%s"
	bilibiliPlayerURL = "https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&qn=%d&fnval=16"

	// defaultQuality sets the default video quality requested from the player API.
	// Common values: 116=1080p60, 80=1080p, 64=720p, 32=480p, 16=360p
	// Using 116 to prefer highest quality by default.
	defaultQuality = 116
)

type bilibiliData struct {
	Code int `json:"code"`
	Data struct {
		BVID  string `json:"bvid"`
		CID   int    `json:"cid"`
		Title string `json:"title"`
	} `json:"data"`
}

type playerData struct {
	Code int `json:"code"`
	Data struct {
		DURL []struct {
			URL  string `json:"url"`
			Size int64  `json:"size"`
		} `json:"durl"`
		Quality int `json:"quality"`
	} `json:"data"`
}

var bvidRegex = regexp.MustCompile(`BV[a-zA-Z0-9]+`)

// Extractor implements the bilibili extractor
type Extractor struct{}

// New returns a bilibili extractor
func New() *Extractor {
	return &Extractor{}
}

// Extract extracts video data from bilibili URLs
func (e *Extractor) Extract(url string, option types.Options) ([]*types.Data, error) {
	bvid := bvidRegex.FindString(url)
	if bvid == "" {
		return nil, fmt.Errorf("unable to find BVID in URL: %s", url)
	}

	apiURL := fmt.Sprintf(bilibiliAPIURL, bvid)
	resp, err := request.Get(apiURL, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data bilibiliData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	if data.Code != 0 {
		return nil, fmt.Errorf("bilibili API error code: %d", data.Code)
	}

	playerURL := fmt.Sprintf(bilibiliPlayerURL, bvid, data.Data.CID, defaultQuality)
	presp, err := request.Get(playerURL, url, nil)
	if err != nil {
		return nil, err
	}
	defer presp.Body.Close()

	var pdata playerData
	if err := json.NewDecoder(presp.Body).Decode(&pdata); err != nil {
		return nil, err
	}

	streams := make(map[string]*types.Stream)
	for i, durl := range pdata.Data.DURL {
		key := fmt.Sprintf("stream_%d", i)
		ext := utils.FileExtension(strings.Split(durl.URL, "?")[0])
		streams[key] = &types.Stream{
			Parts: []*types.Part{
				{URL: durl.URL, Size: durl.Size, Ext: ext},
			},
			Size:    durl.Size,
			Quality: fmt.Sprintf("%d", pdata.Data.Quality),
		}
	}

	return []*types.Data{
		{
			Site:    "Bilibili bilibili.com",
			Title:   data.Data.Title,
			Type:    types.DataTypeVideo,
			Streams:  streams,
			URL:     url,
			Err:     nil,
		},
	}, nil
}

var _ http.Handler = nil // suppress unused import
