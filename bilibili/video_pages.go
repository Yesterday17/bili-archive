package bilibili

import (
	"encoding/json"
	"github.com/Yesterday17/bili-archive/utils"
)

type VideoPage struct {
	Page     int    `json:"page"`
	PageName string `json:"pagename"`
	CID      int    `json:"cid"`
}

func GetVideoPages(aid string) ([]VideoPage, error) {
	var err error
	var body []VideoPage

	content, err := utils.Get("https://www.bilibili.com/widget/getPageList?aid="+aid, "", nil)
	if err != nil {
		return nil, err
	}

	if content[0] != '<' {
		err = json.Unmarshal([]byte(content), &body)
	}

	return body, err
}
