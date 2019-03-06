package main

import (
	"encoding/json"
	"errors"
	"github.com/Yesterday17/bili-archive/bilibili"
	_ "github.com/Yesterday17/bili-archive/statik"
	"github.com/gorilla/websocket"
	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/downloader"
	bilibili_annie "github.com/iawia002/annie/extractors/bilibili"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CreateBiliArchiveServer() {
	code := bilibili.QRCode{}
	handler := http.NewServeMux()

	// 前端
	frontend, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	handler.Handle("/", http.FileServer(frontend))

	// 后端
	// Path: /api/login-qr
	// Method: GET
	// Description: 获得登录所需要的二维码
	// Response: @image string
	loginQRHandler := func(w http.ResponseWriter, req *http.Request) {
		// 检测 Cookies 是否过期
		if configuration.Cookies == "" || bilibili.GetUserMID(configuration.Cookies) == "-1" {
			code = bilibili.GetLoginQRCode()

			// 更新配置文件中的 cookies
			configuration.Cookies = ""
			QuickSaveConfig()
		} else {
			code.Image = "cookies_exist"
		}

		output, err := json.Marshal(map[string]string{"image": code.Image})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/login-qr", loginQRHandler)

	// Path: /api/login-status
	// Method: GET
	// Description: 获取登录状态
	// Response: @ok boolean
	loginStatusHandler := func(w http.ResponseWriter, req *http.Request) {
		tmpCookies := ""
		ok, err := false, errors.New("")

		// 存在 Cookies 直接跳过
		if configuration.Cookies != "" {
			ok = true
		} else {
			if code.Image != "" {
				ok, tmpCookies, err = code.Check()
				if err != nil {
					log.Println(err.Error())
					ok = false
				}
			}
		}

		// 生成 Json
		output, err := json.Marshal(map[string]bool{"ok": ok})
		if err != nil {
			log.Println(err.Error())
			return
		}

		// 成功登录则进行配置
		if ok {
			// Cookies 是新获得的
			if tmpCookies != "" {
				configuration.Cookies = bilibili.GetCookiesString(tmpCookies)
			}

			// 立即保存已获得的 Cookies
			QuickSaveConfig()

			// 配置 annie
			// TODO: 脱离 annie
			config.InfoOnly = false
			config.Cookie = configuration.Cookies
			config.RetryTimes = 10
		}

		// 返回 json
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/login-status", loginStatusHandler)

	// Path: /api/current-user
	// Method: GET
	// Description: 当前用户信息
	// Response: @ok boolean
	// 			 @message string
	// 			 @uid strng
	currentUserHandler := func(w http.ResponseWriter, req *http.Request) {
		message, uid := "", "-1"
		if configuration.Cookies == "" {
			message = "用户未登录"
		} else {
			uid = bilibili.GetUserMID(configuration.Cookies)
		}

		output, err := json.Marshal(map[string]interface{}{
			"ok":      uid != "-1",
			"message": message,
			"uid":     uid,
		})

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/current-user", currentUserHandler)

	// Path: /api/info
	// Method: GET
	// Params: @uid string
	// Description: 用户信息
	// Response: @ok boolean
	// 			 @data MIDInfo
	midInfo := func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		if uid == "" {
			uid = "-1"
		}

		data, err := bilibili.GetMIDInfo(uid)
		if err != nil {
			log.Println(err)
		}

		output, err := json.Marshal(map[string]interface{}{
			"ok":   uid != "-1",
			"data": data,
		})

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/info", midInfo)

	// Path: /api/favlist
	// Method: GET
	// Params: @uid string
	// Description: 收藏列表
	// Response: @ok boolean
	// 			 @data favoriteListItem[]
	favList := func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		if uid == "" {
			uid = "-1"
		}

		list := bilibili.GetFavoriteList(uid, configuration.Cookies)
		output, err := json.Marshal(map[string]interface{}{
			"ok":   list != nil,
			"data": list,
		})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/favlist", favList)

	// Path: /api/fav
	// Method: GET
	// Params: @uid string
	// 		   @fid string
	// Description: 收藏详情
	// Response: @ok boolean
	// 			 @data FavoriteListItemVideo[]
	favDetail := func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		fid := req.URL.Query().Get("fid")
		pn := req.URL.Query().Get("pn")

		var list []bilibili.FavoriteListItemVideo = nil
		if uid != "" && fid != "" {
			list = bilibili.GetFavoriteListItems(uid, fid, pn, configuration.Cookies)
		}

		output, err := json.Marshal(map[string]interface{}{
			"ok":   list != nil,
			"data": list,
		})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/fav", favDetail)

	// Path: /download
	// Method: WebSocket
	// Send: @aid  string
	// 		 @page string
	// 		 @fav  string
	// Description: 下载视频
	// Response: 进度
	downloadVideo := func(w http.ResponseWriter, req *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			result := map[string]interface{}{
				"status": "error",
				"data":   err,
			}
			if err := ws.WriteJSON(&result); err != nil {
				log.Println(err)
			}
			return
		}
		defer ws.Close()

		data := bilibili.DownloadVideoRequest{}
		if err := ws.ReadJSON(&data); err != nil {
			log.Println(err)
			result := map[string]interface{}{
				"status": "error",
				"data":   err,
			}
			if err := ws.WriteJSON(&result); err != nil {
				log.Println(err)
			}
			return
		}

		if err := os.MkdirAll("./video/"+data.FavTitle, os.ModePerm); err != nil {
			log.Println(err)
			result := map[string]interface{}{
				"status": "error",
				"data":   err,
			}
			if err := ws.WriteJSON(&result); err != nil {
				log.Println(err)
			}
			return
		}

		config.OutputPath = "./video/" + string(data.FavTitle)
		config.OutputName = data.Title + " - " + data.Page.PageName

		if _, err := os.Stat(config.OutputPath + "/" + config.OutputName + ".flv"); os.IsNotExist(err) {
			url := "https://www.bilibili.com/video/av" + data.Aid + "/?p=" + strconv.Itoa(data.Page.Page)

			item := bilibili.ExtractVideo(data, configuration.Cookies)
			if item.Err != nil {
				log.Println(err)
				result := map[string]interface{}{
					"status": "error",
					"data":   item.Err,
				}
				if err := ws.WriteJSON(&result); err != nil {
					log.Println(err)
				}
			}

			if err = downloader.Download(item, url, 5); err != nil {
				log.Println(err)
				result := map[string]interface{}{
					"status": "error",
					"data":   err,
				}
				if err := ws.WriteJSON(&result); err != nil {
					log.Println(err)
				}
			}
		}
	}
	handler.HandleFunc("/download", downloadVideo)

	// Download, transfer data with Websocket
	// Deprecated
	iterateFavHandler := func(w http.ResponseWriter, req *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer ws.Close()

		messageType, mid, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fav, av, page := "", "", ""
		videoItem, videoPage := bilibili.FavoriteListItemVideo{}, bilibili.VideoPage{}
		bilibili.IterateFavoriteList(string(mid), configuration.Cookies, func(key, value string, data interface{}) {
			switch key {
			case "Favorite":
				fav = value
			case "Video":
				videoItem = data.(bilibili.FavoriteListItemVideo)
				av = strconv.Itoa(videoItem.AID)
			case "Page":
				videoPage = data.(bilibili.VideoPage)
				page = strconv.Itoa(videoPage.Page)
			case "Message":
				page = ""
			}

			err := ws.WriteMessage(messageType, []byte(key+": "+value))
			if err != nil {
				log.Println(err)
			}

			if av != "" && page != "" {
				os.MkdirAll("./video/"+fav, os.ModePerm)
				config.OutputPath = "./video/" + fav
				config.OutputName = videoItem.Title + " - " + videoPage.PageName

				if _, err := os.Stat(config.OutputPath + "/" + config.OutputName + ".flv"); os.IsNotExist(err) {
					url := "https://www.bilibili.com/video/av" + av + "/?p=" + page

					v, err := bilibili_annie.Extract(url)
					if err != nil {
						log.Println(err)
					}

					for _, item := range v {
						if item.Err != nil {
							log.Println(err)
							continue
						}
						err = downloader.Download(item, url, 5)
						if err != nil {
							log.Println(err)
						}
					}
				}
				page = ""
			}
		})

	}
	handler.HandleFunc("/ws", iterateFavHandler)

	if err := http.ListenAndServe(":"+configuration.Port, handler); err != nil {
		log.Fatal(err)
	}
}
