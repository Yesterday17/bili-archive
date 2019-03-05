package main

import (
	"encoding/json"
	"errors"
	"github.com/Yesterday17/bili-archive/bilibili"
	_ "github.com/Yesterday17/bili-archive/statik"
	"github.com/gorilla/websocket"
	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/downloader"
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
	// 获得登录所需要的 二维码
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

	// 获取登录状态
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

	// 当前用户信息
	currentUserHandler := func(w http.ResponseWriter, rq *http.Request) {
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

	// MIDInfo
	// currentUserData
	midInfo := func(w http.ResponseWriter, rq *http.Request) {
		uid := rq.URL.Query().Get("uid")
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

	// Download, transfer data with Websocket
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

					v, err := bilibili.Extract(url)
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
