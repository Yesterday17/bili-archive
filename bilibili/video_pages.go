package bilibili

import (
	"encoding/json"
	"github.com/Yesterday17/bili-archive/utils"
)

type VideoPageList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    []VideoPage `json:"data"`
}

type VideoPage struct {
	CID       int            `json:"cid"`
	Page      int            `json:"page"`
	From      string         `json:"from"`
	Part      string         `json:"part"`
	Duration  int64          `json:"duration"`
	Vid       string         `json:"vid"`
	Weblink   string         `json:"weblink"`
	Dimension VideoDimension `json:"dimension"`
}

type VideoDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

func GetVideoPages(aid string) ([]VideoPage, error) {
	var err error
	var body VideoPageList

	content, err := utils.Get("https://api.bilibili.com/x/player/pagelist?aid="+aid, "", nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(content), &body)
	return body.Data, err
}
