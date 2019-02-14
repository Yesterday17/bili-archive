package bilibili

import (
	"encoding/json"
	"net/http"
)

type VideoPage struct {
	Page     int    `json:"page"`
	PageName string `json:"pagename"`
	CID      int    `json:"cid"`
}

type VideoPages []VideoPage

func GetVideoPages(aid string) (VideoPages, error) {
	client := http.Client{}
	body := VideoPages{}
	var error error

	res, err := client.Get("https://www.bilibili.com/widget/getPageList?aid=" + aid)
	if err != nil {
		error = err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		error = err
	}
	return body, error
}
