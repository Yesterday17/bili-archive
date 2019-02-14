package main

import (
	"./bilibili"
	"fmt"
	"log"
	"strconv"
)

func main() {
	cookiesList, timeout := bilibili.GetLoginQRCode().WaitForLogin()

	if timeout {
		log.Fatal("Timeout!")
	} else {
		cookies := bilibili.Cookies(cookiesList)

		mid := bilibili.GetUserMID(cookies)
		fmt.Println("MID: " + mid)

		for _, value := range bilibili.GetFavoriteList(mid, cookies) {
			// https://api.bilibili.com/x/space/fav/arc?vmid=5756570&fid=1683980&pn=1
			strconv.Itoa(value.MID)
		}
	}
}
