package main

import (
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/Yesterday17/bili-archive/server"
)

func main() {
	// fmt.Println(bilibili.GetUserFavoriteListReport("123817257", false))
	server.CreateBiliArchiveServer(bilibili.GetLoginQRCode())
}
