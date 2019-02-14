package main

import (
	"./bilibili"
	"fmt"
	"log"
)

func main() {
	cookies, timeout := bilibili.GetLoginQRCode().WaitForLogin()

	if timeout {
		log.Fatal("Timeout!")
	} else {
		fmt.Println(bilibili.GetUserMID(bilibili.Cookies(cookies)))
	}
}
