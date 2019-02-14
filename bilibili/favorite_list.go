package bilibili

import (
	"encoding/json"
	"log"
	"net/http"
)

type favoriteList struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	TTL     int              `json:"ttl"`
	Data    favoriteListData `json:"data"`
}

type favoriteListData struct {
	Archive  []favoriteListItem `json:"archive"`
	Playlist int                `json:"playlist"`
	Topic    int                `json:"topic"`
	Article  int                `json:"article"`
	Album    int                `json:"album"`
	Movie    int                `json:"movie"`
}

type favoriteListItem struct {
	MediaID      int                     `json:"media_id"`
	FID          int                     `json:"fid"`
	MID          int                     `json:"mid"`
	Name         string                  `json:"name"`
	MaxCount     int                     `json:"max_count"`
	CurrentCount int                     `json:"cur_count"`
	AttenCount   int                     `json:"atten_count"`
	Favoured     int                     `json:"favoured"`
	State        int                     `json:"state"`
	Ctime        int                     `json:"ctime"`
	Mtime        int                     `json:"mtime"`
	Cover        []favoriteListItemCover `json:"cover"`
}

type favoriteListItemCover struct {
	AID     int    `json:"aid"`
	Picture string `json:"pic"`
	Type    int    `json:"type"`
}

func GetFavoriteList(mid string, cookies string) []favoriteListItem {
	client := http.Client{}
	body := favoriteList{}

	req, err := http.NewRequest("GET", "http://api.bilibili.com/x/space/fav/nav?mid="+mid, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("Cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	return body.Data.Archive
}
