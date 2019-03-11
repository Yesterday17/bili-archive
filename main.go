package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/pkg/browser"
	"log"
	"strconv"
)

func main() {
	var client, currentUser bool
	var cookies, uidStr string
	var uid int64

	flag.BoolVar(&client, "command", false, "CommandLine mode")
	flag.StringVar(&cookies, "cookies", "", "Cookies")
	flag.BoolVar(&currentUser, "current", true, "Download favorite list of current user")
	flag.Int64Var(&uid, "uid", 0, "The uid to download favorite")
	flag.Parse()

	// 加载配置文件
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}

	if cookies != "" {
		configuration.Cookies = cookies
	}

	// 立即保存（初始设置）
	QuickSaveConfig()

	if client {
		// 获得 UID
		if currentUser {
			uidStr = bilibili.GetUserMID(configuration.Cookies)
			id, err := strconv.Atoi(uidStr)
			if err != nil || id == -1 {
				// 获得当前用户 UID 失败
				log.Fatal(err)
			}

			uid = int64(id)
		} else {
			uidStr = strconv.Itoa(int(uid))
		}

		// 获得收藏列表
		lists, err := bilibili.GetFavoriteList(uidStr, configuration.Cookies)
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fid := list.FID
			path := "./video/" + list.Name

			for i := 0; i < list.CurrentCount/20; i++ {
				var items []bilibili.FavoriteListItemVideo
				err := errors.New(fmt.Sprintf("Favlist: %s, Page: %d", list.Name, i+1))
				for err != nil {
					log.Println(err)
					items, err = bilibili.GetFavoriteListItems(uidStr, strconv.Itoa(fid), strconv.Itoa(i+1), configuration.Cookies)
				}

				for _, item := range items {
					var pages bilibili.VideoPages
					err := errors.New(fmt.Sprintf("Video: %s, AID: %d", item.Title, item.AID))
					for err != nil {
						log.Println(err)
						pages, err = bilibili.GetVideoPages(strconv.Itoa(item.AID))
					}

					for _, page := range pages {
						data := bilibili.DownloadVideoRequest{}
						video := bilibili.ExtractVideo(data, configuration.Cookies)
						if video.Err != nil {
							log.Println(video.Err)
							continue
						}

						vData := bilibili.DownloadVideoRequest{
							Title:    item.Title,
							Aid:      strconv.Itoa(item.AID),
							FavTitle: list.Name,
							Page: bilibili.RequestVideoPage{
								Page:     page.Page,
								CID:      strconv.Itoa(page.CID),
								PageName: page.PageName,
							},
						}

						if err := bilibili.DownloadVideo(video, vData, path, configuration.Cookies, nil); err != nil {
							log.Println(err)
							continue
						}
					}
				}
			}
		}
	} else {
		// 启动服务器
		browser.OpenURL("http://localhost:8080")
		CreateBiliArchiveServer()
	}
}
