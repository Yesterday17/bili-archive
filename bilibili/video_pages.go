package bilibili

import (
	"github.com/Yesterday17/bili-archive/utils"
)

type VideoPage struct {
	Page     int    `json:"page"`
	PageName string `json:"pagename"`
	CID      int    `json:"cid"`
}

type VideoPages []VideoPage

func GetVideoPages(aid string) (VideoPages, error) {
	body := VideoPages{}
	if err := utils.GetJson("https://www.bilibili.com/widget/getPageList?aid="+aid, "", &body); err != nil {
		return nil, err
	}

	return body, nil
}
