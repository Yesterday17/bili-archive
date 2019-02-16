package bilibili

import (
	"fmt"
	"math"
	"strconv"
)

func addLine(original *string, line string) {
	*original += line + "\n"
}

func IterateFavoriteList(mid string, cookies string, callback func(key, value string, data interface{})) {
	for _, value := range GetFavoriteList(mid, cookies) {
		fid := strconv.Itoa(value.FID)
		callback("Favorite", value.Name, value)

		for i := 0; i < int(math.Ceil(float64(value.CurrentCount)/30.0)); i++ {
			for _, data := range GetFavoriteListItems(mid, fid, strconv.Itoa(i+1)) {
				aid := strconv.Itoa(data.AID)
				callback("Video", data.Title, data)

				pages, err := GetVideoPages(aid)
				if err != nil {
					callback("Message", "Video unavailable.", "")
				} else {
					for _, page := range pages {
						callback("Page", page.PageName, page)
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
