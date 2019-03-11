package bilibili

import (
	"github.com/Yesterday17/bili-archive/utils"
)

type favoriteList struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	TTL     int              `json:"ttl"`
	Data    favoriteListData `json:"data"`
}

type favoriteListData struct {
	Archive  []FavoriteListItemDetail `json:"archive"`
	Playlist int                      `json:"playlist"`
	Topic    int                      `json:"topic"`
	Article  int                      `json:"article"`
	Album    int                      `json:"album"`
	Movie    int                      `json:"movie"`
}

type FavoriteListItemDetail struct {
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

func GetFavoriteList(mid string, cookies string) ([]FavoriteListItemDetail, error) {
	body := favoriteList{}

	if err := utils.GetJson("http://api.bilibili.com/x/space/fav/nav?mid="+mid, cookies, &body); err != nil {
		return nil, err
	}

	return body.Data.Archive, nil
}
