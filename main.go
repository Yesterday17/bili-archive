package main

import (
	"./bilibili"
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	cookiesList, timeout := bilibili.GetLoginQRCode().WaitForLogin()

	if timeout {
		log.Fatal("Timeout!")
	} else {
		cookies := bilibili.Cookies(cookiesList)
		mid := bilibili.GetUserMID(cookies)

		for _, value := range bilibili.GetFavoriteList(mid, cookies) {
			fid := strconv.Itoa(value.FID)
			fmt.Println("Favorite: " + value.Name)
			fmt.Println("FID: " + fid)

			for i := 0; i < int(math.Ceil(float64(value.CurrentCount)/30.0)); i++ {
				for index, data := range bilibili.GetFavoriteListItems(mid, fid, strconv.Itoa(i+1)) {
					aid := strconv.Itoa(data.AID)
					fmt.Println(strconv.Itoa(index+1) + ". " + data.Title)
					fmt.Println("AID: " + aid)
					fmt.Println("Pages:")

					pages, err := bilibili.GetVideoPages(aid)
					if err != nil {
						fmt.Println("Video unavailable.")
					} else {
						for index, page := range pages {
							fmt.Println("  (" + strconv.Itoa(index+1) + "). " + page.PageName)
						}
					}

					fmt.Println("")
				}
			}
		}
	}
}
