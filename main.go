package main

import (
	"flag"
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/pkg/browser"
	"log"
	"os"
	syspath "path"
	"strconv"
	"time"
)

func main() {
	var client, currentUser bool
	var cookies, uid, path string

	flag.BoolVar(&client, "c", false, "是否启动命令行模式。")
	flag.StringVar(&cookies, "cookies", "", "用户的 cookies，会覆盖配置文件内的值。")
	flag.BoolVar(&currentUser, "this", true, "下载 cookies 代表用户的收藏。")
	flag.StringVar(&uid, "uid", "", "下载收藏用户的 UID")
	flag.StringVar(&path, "path", "./Videos/", "下载视频的根目录")
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
		// 警告不存在 cookies
		if configuration.Cookies == "" {
			log.Fatal("不存在 cookies，请指定 cookies！")
		}

		// 获得 UID
		if uid == "" && currentUser {
			uid = bilibili.GetUserMID(configuration.Cookies)
			if uid == "-1" {
				// 获得当前用户 UID 失败
				log.Fatal("获取当前用户 UID 失败，请尝试重新登录！")
			}
		}

		// 获得收藏列表
		lists, err := bilibili.GetFavoriteList(uid, configuration.Cookies)
		if err != nil {
			log.Fatal(err)
		}

		// 遍历收藏各列表
		for _, list := range lists {
			fid := list.FID
			listPath := syspath.Join(path, list.Name)

			// 遍历收藏分页
			for i := 0; i < list.CurrentCount/20; i++ {
				var items []bilibili.FavoriteListItemVideo
				var err error
				for items, err = bilibili.GetFavoriteListItems(uid, strconv.Itoa(fid), strconv.Itoa(i+1), configuration.Cookies); err != nil; {
					log.Println(err)
					time.Sleep(time.Second)
				}

				// 遍历收藏内各视频
				for _, item := range items {
					var pages bilibili.VideoPages
					for pages, err = bilibili.GetVideoPages(strconv.Itoa(item.AID)); err != nil; {
						log.Println(err.Error())
						time.Sleep(time.Second)
					}

					// 遍历分P
					for _, page := range pages {
						// 准备数据
						data := bilibili.DownloadVideoRequest{
							Title:    item.Title,
							Aid:      strconv.Itoa(item.AID),
							FavTitle: list.Name,
							Page: bilibili.RequestVideoPage{
								Page:     page.Page,
								CID:      strconv.Itoa(page.CID),
								PageName: page.PageName,
							},
						}
						// 提取链接
						video := bilibili.ExtractVideo(data, configuration.Cookies)
						logStr := fmt.Sprintf("[av%d][p%d]", item.AID, page.Page)
						fmt.Println(logStr + " " + item.Title)
						if video.Err != nil {
							log.Println(fmt.Sprintf("[%s]%s %s", "EX", logStr, video.Err))
							continue
						}
						// 创建目录
						if err := os.MkdirAll(listPath, os.ModePerm); err != nil {
							log.Println(fmt.Sprintf("[%s]%s %s", "MK", logStr, video.Err))
							continue
						}
						// 下载视频
						if err := bilibili.DownloadVideo(video, data, listPath, configuration.Cookies, nil); err != nil {
							log.Println(fmt.Sprintf("[%s]%s %s", "DL", logStr, video.Err))
							continue
						}
					}
				}
			}
		}
	} else {
		// 打开前端网页
		browser.OpenURL("http://localhost:8080")
		// 启动服务器
		CreateBiliArchiveServer()
	}
}
