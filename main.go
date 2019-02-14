package main

import (
	"./bilibili"
	"fmt"
	"log"
)

func main() {
	cookiesList, timeout := bilibili.GetLoginQRCode().WaitForLogin()

	if timeout {
		log.Fatal("Timeout!")
	} else {
		cookies := bilibili.Cookies(cookiesList)

		mid := bilibili.GetUserMID(cookies)
		fmt.Println("MID: " + mid)

		bilibili.GetFavoriteList(mid, cookies)
	}
}
