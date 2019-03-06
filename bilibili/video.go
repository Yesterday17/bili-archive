package bilibili

import (
	"encoding/json"
	"fmt"
	"github.com/iawia002/annie/downloader"
	"github.com/iawia002/annie/request"
	"github.com/iawia002/annie/utils"
	"strconv"
)

const (
	bilibiliAPI      = "https://interface.bilibili.com/v2/playurl?"
	bilibiliTokenAPI = "https://api.bilibili.com/x/player/playurl/token?"
)

const (
	// 会变化的 appKey 和 secKey
	appKey = "iVGUTjsxvpLeuDCf"
	secKey = "aHRmhWMLkdeMuILqORnYZocwMBpMEOdt"
)

const referer = "https://www.bilibili.com"

type DownloadVideoRequest struct {
	Title    string           `json:"title"`
	FavTitle string           `json:"fav_title"`
	Aid      string           `json:"aid"`
	Page     RequestVideoPage `json:"page"`
}

type RequestVideoPage struct {
	Page     int    `json:"page"`
	PageName string `json:"pagename"`
	CID      string `json:"cid"`
}

var uToken string

// 生成 API
func GenAPI(aid, cid, quality, cookies string) (string, error) {
	var (
		err        error
		baseAPIURL string
		params     string
	)
	if cookies != "" && uToken == "" {
		uToken, err = request.Get(
			fmt.Sprintf("%said=%s&cid=%s", bilibiliTokenAPI, aid, cid),
			referer,
			nil,
		)
		if err != nil {
			return "", err
		}
		var t token
		err = json.Unmarshal([]byte(uToken), &t)
		if err != nil {
			return "", err
		}
		if t.Code != 0 {
			return "", fmt.Errorf("cookie error: %s", t.Message)
		}
		uToken = t.Data.Token
	}

	params = fmt.Sprintf(
		"appkey=%s&cid=%s&otype=json&qn=%s&quality=%s&type=",
		appKey, cid, quality, quality,
	)
	baseAPIURL = bilibiliAPI

	api := fmt.Sprintf(
		"%s%s&sign=%s", baseAPIURL, params, utils.Md5(params+secKey),
	)
	if uToken != "" {
		api = fmt.Sprintf("%s&uToken=%s", api, uToken)
	}
	return api, nil
}

func GenURL(durl []dURLData) ([]downloader.URL, int64) {
	var size int64
	urls := make([]downloader.URL, len(durl))
	for index, data := range durl {
		size += data.Size
		urls[index] = downloader.URL{
			URL:  data.URL,
			Size: data.Size,
			Ext:  "flv",
		}
	}
	return urls, size
}

// 提取 bilibili 视频
func ExtractVideo(data DownloadVideoRequest, cookies string) downloader.Data {
	url := "https://www.bilibili.com/video/av" + data.Aid + "/?p=" + strconv.Itoa(data.Page.Page)

	// Get "accept_quality" and "accept_description"
	// "accept_description":["高清 1080P","高清 720P","清晰 480P","流畅 360P"],
	// "accept_quality":[80,48,32,16],
	api, err := GenAPI(data.Aid, data.Page.CID, "15", cookies)
	if err != nil {
		return downloader.EmptyData(url, err)
	}
	jsonString, err := request.Get(api, referer, nil)
	if err != nil {
		return downloader.EmptyData(url, err)
	}

	var quality qualityInfo
	err = json.Unmarshal([]byte(jsonString), &quality)
	if err != nil {
		return downloader.EmptyData(url, err)
	}

	streams := make(map[string]downloader.Stream, len(quality.Quality))
	for _, q := range quality.Quality {
		apiURL, err := GenAPI(data.Aid, data.Page.CID, strconv.Itoa(q), cookies)
		if err != nil {
			return downloader.EmptyData(url, err)
		}
		jsonString, err := request.Get(apiURL, referer, nil)
		if err != nil {
			return downloader.EmptyData(url, err)
		}
		var data bilibiliData
		err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			return downloader.EmptyData(url, err)
		}

		// Avoid duplicate streams
		if _, ok := streams[strconv.Itoa(data.Quality)]; ok {
			continue
		}

		urls, size := GenURL(data.DURL)
		streams[strconv.Itoa(data.Quality)] = downloader.Stream{
			URLs:    urls,
			Size:    size,
			Quality: qualityString[data.Quality],
		}
	}

	return downloader.Data{
		Site:    "哔哩哔哩 bilibili.com",
		Title:   data.Title,
		Type:    "video",
		Streams: streams,
		URL:     url,
	}
}
