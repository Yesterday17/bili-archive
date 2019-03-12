package main

import (
	"flag"
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/Yesterday17/bili-archive/utils"
	"github.com/cheggaaa/pb"
	"github.com/pkg/browser"
	"log"
	syspath "path"
	"strconv"
	"sync"
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

		var wgf sync.WaitGroup

		// 遍历收藏各列表
		for _, list := range lists {
			wgf.Add(1)
			go func(list bilibili.FavoriteListItemDetail) {
				defer wgf.Done()
				fid := list.FID

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
						var pages []bilibili.VideoPage
						// 本地路径
						basePath := syspath.Join(path, fmt.Sprintf("av%d", item.AID))
						videoPath := syspath.Join(basePath, "video")
						dataPath := syspath.Join(basePath, "data")
						lockPath := syspath.Join(basePath, "lock")
						// 创建视频文件目录
						utils.MKDirs(basePath, videoPath, dataPath, lockPath)
						// 保存视频数据
						if err := utils.WriteJsonS(dataPath, "video.json", item); err != nil {
							log.Println(fmt.Sprintf("[%s]%s", "VD", err))
						}
						// 跳过失效视频
						if utils.FileExist(lockPath, "broken") {
							continue
						}
						// 不存在 lockfile 时获取 pages
						for pages, err = bilibili.GetVideoPages(strconv.Itoa(item.AID)); err != nil; {
							log.Println(err.Error())
							time.Sleep(time.Second)
						}
						// 获取 pages 后判断视频是否失效
						if len(pages) == 0 {
							if err := utils.WriteLockFile(lockPath, "broken"); err != nil {
								log.Println(fmt.Sprintf("[%s]%s", "LO", err))
							}
							continue
						}
						// 对每个视频的分P实行多线程下载
						var wgv sync.WaitGroup
						// 遍历分P
						for _, page := range pages {
							// 存在 lockfile 直接跳过整个视频下载
							if !utils.FileExist(lockPath, strconv.Itoa(page.CID)) {
								// 该分P的进度条
								bar := utils.NewProgressBar(fmt.Sprintf("[P%d]%s", page.Page, item.Title))
								wgv.Add(1)
								go func(item bilibili.FavoriteListItemVideo, page bilibili.VideoPage, bar *pb.ProgressBar) {
									defer wgv.Done()
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
									logStr := fmt.Sprintf("[av%d][P%d]", item.AID, page.Page)
									if video.Err != nil {
										log.Println(fmt.Sprintf("[%s]%s %s", "EX", logStr, video.Err))
										return
									}
									// 回调函数
									callback := func(pg *utils.Progress) {
										bar.Set64(pg.Progress.Progress)
										bar.SetTotal64(pg.Progress.Size)
									}
									// 保存分P信息
									if err := utils.WriteJsonS(dataPath, fmt.Sprintf("%d.json", page.CID), page); err != nil {
										log.Println(fmt.Sprintf("[%s]%s %s", "PD", logStr, err))
									}
									// 开始进度条
									bar.Start()
									// 下载视频
									if err := bilibili.DownloadVideo(video, data, videoPath, configuration.Cookies, callback); err != nil {
										log.Println(fmt.Sprintf("[%s]%s %s", "DL", logStr, err))
										return
									}
									// 视频下载结束后创建 lockfile
									if err := utils.WriteLockFile(lockPath, strconv.Itoa(page.CID)); err != nil {
										log.Println(fmt.Sprintf("[%s]%s %s", "LO", logStr, err))
									}
									// 结束进度条
									bar.Finish()
								}(item, page, bar)
							}
						}
						wgv.Wait()
					}
				}
			}(list)
		}
		wgf.Wait()
		fmt.Println("下载完成！")
	} else {
		// 打开前端网页
		browser.OpenURL("http://localhost:8080")
		// 启动服务器
		CreateBiliArchiveServer()
	}
}
