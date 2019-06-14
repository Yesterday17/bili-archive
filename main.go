package main

import (
	"flag"
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/Yesterday17/bili-archive/server"
	"github.com/Yesterday17/bili-archive/utils"
	"log"
	"math"
	syspath "path"
	"strconv"
	"strings"
	"sync"
	"time"
)

func FilterFavList(vs []bilibili.FavoriteListItemDetail, f func(bilibili.FavoriteListItemDetail) bool) []bilibili.FavoriteListItemDetail {
	vsf := make([]bilibili.FavoriteListItemDetail, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func StringIndex(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func StringInclude(vs []string, t string) bool {
	return StringIndex(vs, t) >= 0
}

func main() {
	var serverMode, silentMode bool
	var cookies, uid, path string
	var mode, wl, bl string
	var ui UI

	flag.BoolVar(&serverMode, "server", false, "启动后端模式。")
	flag.BoolVar(&silentMode, "silentMode", false, "启动 silent 模式")
	flag.StringVar(&cookies, "cookies", "", "用户的 cookies，会更新配置文件内的值。")
	flag.StringVar(&uid, "uid", "", "下载收藏用户的 UID，不指定则为 cookies 对应用户。")
	flag.StringVar(&path, "path", "./Videos/", "下载视频的根目录。")
	flag.StringVar(&mode, "mode", "n", "下载的模式，n为通常，b为黑名单，w为白名单，配合wh和bl使用。")
	flag.StringVar(&wl, "wl", "", "下载收藏列表的白名单，用英文分号分隔，每一项为收藏夹的FID。")
	flag.StringVar(&bl, "bl", "", "下载收藏列表的黑名单，用英文分号分隔，每一项为收藏夹的FID。")
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

	// 单纯作为后端运行
	if serverMode {
		server.CreateBiliArchiveServer(configuration.Port, configuration.Cookies)
		return
	}

	// 预处理黑名单/白名单
	whitelist := strings.Split(wl, ";")
	blacklist := strings.Split(wl, ";")

	// 用户登录
	if configuration.Cookies == "" {
		code := bilibili.GetLoginQRCode()
		log.Println("账号未登录，请扫描以下二维码登录！")
		log.Println("如果二维码未显示完全，请上下放大，不要左右放大窗口！")
		obj := qrcodeTerminal.New2(
			qrcodeTerminal.ConsoleColors.BrightBlack,
			qrcodeTerminal.ConsoleColors.BrightWhite,
			qrcodeTerminal.QRCodeRecoveryLevels.Low,
		)
		obj.Get(code.QRLogin.Url).Print()
		err, cookies := code.WaitForLogin()
		if err != nil {
			log.Fatal(err)
		}
		configuration.Cookies = cookies
		log.Println("登录成功！")
		QuickSaveConfig()
	}

	// 获得 UID
	if uid == "" {
		mid, err := bilibili.GetUserMID(configuration.Cookies)
		if err != nil {
			// 获得当前用户 UID 失败
			log.Fatal("获取当前用户 UID 失败，请尝试重新登录！")
		} else {
			uid = mid
		}
	}

	// 获得收藏列表
	favlists, err := bilibili.GetFavoriteList(uid, configuration.Cookies)
	if err != nil {
		log.Fatal(err)
	}

	// 黑名单/白名单过滤
	lists := FilterFavList(favlists, func(fav bilibili.FavoriteListItemDetail) bool {
		fid := strconv.Itoa(fav.FID)
		switch mode {
		case "b":
			// blacklist, 黑名单
			return !StringInclude(blacklist, fid)
		case "w":
			// whitelist, 白名单
			return StringInclude(whitelist, fid)
		default:
			// normal
			return true
		}
	})

	// 将数组转换为 Map
	favlistInfo := make(map[string]*bilibili.FavoriteListItemDetail)
	for _, list := range lists {
		favlistInfo[strconv.Itoa(list.FID)] = &list
	}

	favVideoList := make(map[string]*[]bilibili.FavoriteListItemVideo)
	videoMap := make(map[string]*bilibili.FavoriteListItemVideo)
	videoStatusMap := make(map[string]*DownloadStatus)

	for _, list := range lists {
		var videos []bilibili.FavoriteListItemVideo
		for i := 0; i < int(math.Ceil(float64(list.CurrentCount)/30.0)); i++ {
			var items []bilibili.FavoriteListItemVideo
			var err error
			for items, err = bilibili.GetFavoriteListItems(uid, strconv.Itoa(list.FID), strconv.Itoa(i+1), configuration.Cookies); err != nil; {
				log.Println(err)
				time.Sleep(time.Second)
			}

			for _, item := range items {
				videoMap[strconv.Itoa(item.AID)] = &item
				videos = append(videos, item)

				status := DownloadStatus{
					ui:     nil,
					pg:     nil,
					st:     nil,
					status: "等待下载",
					active: false,
				}
				videoStatusMap[strconv.Itoa(item.AID)] = &status
			}
		}
		favVideoList[strconv.Itoa(list.FID)] = &videos
	}

	if !silentMode {
		ui.favList = &lists
		ui.favMap = &favlistInfo
		ui.favVideoList = &favVideoList
		ui.videoMap = &videoMap
		ui.videoStatusMap = &videoStatusMap
		ui.New()
	}

	var wg sync.WaitGroup
	p := utils.NewWGProgressBar(&wg)
	wg.Add(len(lists))

	// 遍历收藏各列表
	for _, list := range lists {
		go func(list bilibili.FavoriteListItemDetail) {
			defer wg.Done()
			fid := list.FID

			// 遍历收藏内各视频
			for _, item := range *favVideoList[strconv.Itoa(fid)] {
				// 该视频的状态
				status := videoStatusMap[strconv.Itoa(item.AID)]

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
				if item.Cover == bilibili.BrokenVideoCover {
					if err := utils.WriteLockFile(lockPath, "broken"); err != nil {
						log.Println(fmt.Sprintf("[%s]%s", "LO", err))
					}
					continue
				}

				// 遍历分P
				for _, page := range pages {
					// 更新状态中的分P信息
					status.page = page.Page
					// 存在 lockfile 直接跳过整个视频下载
					if !utils.FileExist(lockPath, strconv.Itoa(page.CID)) {
						func(item bilibili.FavoriteListItemVideo, page bilibili.VideoPage) {
							// 准备数据
							data := bilibili.DownloadVideoRequest{
								Title:    item.Title,
								Aid:      strconv.Itoa(item.AID),
								FavTitle: list.Name,
								Page: bilibili.RequestVideoPage{
									Page:     page.Page,
									CID:      strconv.Itoa(page.CID),
									PageName: page.Part,
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
								status.now = pg.Progress.Progress
								status.total = pg.Progress.Size
								if status.active {
									go (*status.ui).Update(func() {
										(*status.pg).SetCurrent(int(status.now))
										(*status.pg).SetMax(int(status.total))

										switch status.status {
										case "正在下载":
											status.st.SetText(fmt.Sprintf(
												"正在下载：[av%d][P%d]%s",
												item.AID,
												status.page,
												item.Title,
											))
										default:
											status.st.SetText(fmt.Sprintf(
												"%s：[av%d]%s",
												status.status,
												item.AID,
												item.Title,
											))
										}
									})
								}
							}
							// 保存分P信息
							if err := utils.WriteJsonS(dataPath, fmt.Sprintf("%d.json", page.CID), page); err != nil {
								log.Println(fmt.Sprintf("[%s]%s %s", "PD", logStr, err))
							}

							// 下载视频
							status.status = "正在下载"
							if err := bilibili.DownloadVideo(video, data, videoPath, configuration.Cookies, callback); err != nil {
								log.Println(fmt.Sprintf("[%s]%s %s", "DL", logStr, err))
								return
							}
							// 视频下载结束后创建 lockfile
							if err := utils.WriteLockFile(lockPath, strconv.Itoa(page.CID)); err != nil {
								log.Println(fmt.Sprintf("[%s]%s %s", "LO", logStr, err))
							}
						}(item, page)
					}
				}
				status.status = "下载完成"
			}
		}(list)
	}
	if !silentMode {
		ui.Run()
	} else {
		// wg.Wait()
		p.Wait()
		fmt.Println("下载完成！")
	}
}
