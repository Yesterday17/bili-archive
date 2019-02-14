package bilibili

import (
	"encoding/json"
	"log"
	"net/http"
)

type VideoPage struct {
	Page     int    `json:"page"`
	PageName string `json:"pagename"`
	CID      int    `json:"cid"`
}

type VideoPages []VideoPage

func GetVideoPages(aid string) VideoPages {
	client := http.Client{}
	body := VideoPages{}

	res, err := client.Get("https://www.bilibili.com/widget/getPageList?aid=" + aid)
	if err != nil {
		log.Fatal(err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}
	return body
}
