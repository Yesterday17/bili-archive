package bilibili

import (
	"fmt"
	"math"
	"strconv"
)

const BrokenVideoCover = "http://i0.hdslb.com/bfs/archive/be27fd62c99036dce67efface486fb0a88ffed06.jpg"

func addLine(original *string, line string) {
	*original += line + "\n"
}

func IterateFavoriteList(mid string, cookies string, callback func(key, value string, data interface{})) {
	favList, _ := GetFavoriteList(mid, cookies)
	for _, value := range favList {
		fid := strconv.Itoa(value.FID)
		callback("Favorite", value.Name, value)

		for i := 0; i < int(math.Ceil(float64(value.CurrentCount)/30.0)); i++ {
			items, _ := GetFavoriteListItems(mid, fid, strconv.Itoa(i+1), cookies)
			for _, data := range items {
				aid := strconv.Itoa(data.AID)
				callback("Video", data.Title, data)

				pages, err := GetVideoPages(aid)
				if err != nil {
					callback("Message", "Video unavailable.", "")
				} else {
					for _, page := range pages {
						callback("Page", page.Part, page)
					}
				}
			}
		}
	}
}

func GetUserFavoriteListReport(mid string, currentUser bool) string {
	qrCode := GetLoginQRCode()
	fmt.Println(qrCode.Image)

	cookiesList, timeout := qrCode.WaitForLogin()
	report := ""

	if timeout {
		addLine(&report, "Timeout!")
	} else {
		cookies := Cookies(cookiesList)
		if currentUser {
			mid = GetUserMID(cookies)
		}
		addLine(&report, "MID: "+mid)

		IterateFavoriteList(mid, cookies, func(key, value string, data interface{}) {
			addLine(&report, key+": "+value)
		})
	}
	return report
}
