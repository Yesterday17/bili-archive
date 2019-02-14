package bilibili

import (
	"math"
	"strconv"
)

func addLine(original *string, line string) {
	*original += line + "\n"
}

func GetUserFavoriteListReport(mid string, currentUser bool) string {
	cookiesList, timeout := GetLoginQRCode().WaitForLogin()
	report := ""

	if timeout {
		addLine(&report, "Timeout!")
	} else {
		cookies := Cookies(cookiesList)
		if currentUser {
			mid = GetUserMID(cookies)
		}
		addLine(&report, "MID: "+mid)

		for _, value := range GetFavoriteList(mid, cookies) {
			fid := strconv.Itoa(value.FID)
			addLine(&report, "Favorite: "+value.Name)
			addLine(&report, "FID: "+fid)

			for i := 0; i < int(math.Ceil(float64(value.CurrentCount)/30.0)); i++ {
				for index, data := range GetFavoriteListItems(mid, fid, strconv.Itoa(i+1)) {
					aid := strconv.Itoa(data.AID)
					addLine(&report, strconv.Itoa(index+1)+". "+data.Title)
					addLine(&report, "AID: "+aid)
					addLine(&report, "Pages:")

					pages, err := GetVideoPages(aid)
					if err != nil {
						addLine(&report, "Video unavailable.")
					} else {
						for index, page := range pages {
							addLine(&report, "  ("+strconv.Itoa(index+1)+"). "+page.PageName)
						}
					}

					addLine(&report, "")
				}
			}
		}
	}
	return report
}
