package bilibili

import (
	"encoding/json"
	"log"
	"net/http"
)

type FavoriteListItem struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	TTL     int                  `json:"ttl"`
	Data    favoriteListItemData `json:"data"`
}

type favoriteListItemData struct {
	SEID           string                  `json:"seid"`
	Page           int                     `json:"page"`
	PageSize       int                     `json:"pagesize"`
	PageCount      int                     `json:"pagecount"`
	Total          int                     `json:"total"`
	SuggestKeyword string                  `json:"suggest_keyword"`
	MID            int                     `json:"mid"`
	FID            int                     `json:"fid"`
	TID            int                     `json:"tid"`
	Order          string                  `json:"order"`
	Keyword        string                  `json:"keyword"`
	Archives       []favoriteListItemVideo `json:"archives"`
}

type favoriteListItemVideo struct {
	AID            int                            `json:"aid"`
	Videos         int                            `json:"videos"`
	TID            int                            `json:"tid"`
	TName          string                         `json:"tname"`
	Copyright      int                            `json:"copyright"`
	Cover          string                         `json:"pic"`
	Title          string                         `json:"title"`
	PubDate        int                            `json:"pubdate"`
	CTime          int                            `json:"ctime"`
	Description    string                         `json:"desc"`
	State          int                            `json:"state"`
	Attribute      int                            `json:"attribute"`
	Duration       int                            `json:"duration"`
	Owner          OwnerUser                      `json:"owner"`
	Stat           favoriteListItemVideoStat      `json:"stat"`
	Dynamic        string                         `json:"dynamic"`
	CID            int                            `json:"cid"`
	Dimension      favoriteListItemVideoDimension `json:"dimension"`
	FavoriteTime   int                            `json:"fav_at"`
	PlayNum        string                         `json:"play_num"`
	HighlightTitle string                         `json:"highlight_title"`
}

type OwnerUser struct {
	MID  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type favoriteListItemVideoStat struct {
	AID         int `json:"aid"`
	View        int `json:"view"`
	Danmaku     int `json:"danmaku"`
	Reply       int `json:"reply"`
	Favorite    int `json:"favorite"`
	Coin        int `json:"coin"`
	Share       int `json:"share"`
	NowRank     int `json:"now_rank"`
	HistoryRank int `json:"his_rank"`
	Like        int `json:"like"`
	Dislike     int `json:"dislike"`
}

type favoriteListItemVideoDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

func GetFavoriteListItems(mid, fid, pn string) []favoriteListItemVideo {
	client := http.Client{}
	body := FavoriteListItem{}

	res, err := client.Get("https://api.bilibili.com/x/space/fav/arc?vmid=" + mid + "&fid=" + fid + "&pn=" + pn)
	if err != nil {
		log.Fatal(err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	return body.Data.Archives
}
