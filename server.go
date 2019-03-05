package main

import (
	"encoding/json"
	"errors"
	"flag"
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
	cookies := ""
	code := bilibili.QRCode{}
	handler := http.NewServeMux()

	// 前端
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	handler.Handle("/", http.FileServer(statikFS))

	// Backend
	// Get QR Code to login
	loginQRHandler := func(w http.ResponseWriter, req *http.Request) {
		code = bilibili.GetLoginQRCode()
		output, err := json.Marshal(map[string]string{"image": code.Image})

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/login-qr", loginQRHandler)

	// Get Login status
	loginStatusHandler := func(w http.ResponseWriter, req *http.Request) {
		ok, err := false, errors.New("")
		if code.Image != "" {
			ok, cookies, err = code.Check()
			if err != nil {
				log.Println(err.Error())
				ok = false
			}
		}

		output, err := json.Marshal(map[string]bool{"ok": ok})

		if err != nil {
			log.Println(err.Error())
			return
		}

		if ok {
			cookies = bilibili.GetCookiesString(cookies)

			flag.BoolVar(&config.InfoOnly, "i", false, "Information only")
			flag.StringVar(&config.Cookie, "c", cookies, "Cookie")
			flag.IntVar(
				&config.RetryTimes, "retry", 10, "How many times to retry when the download failed",
			)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/login-status", loginStatusHandler)

	// currentUserData
	currentUserHandler := func(w http.ResponseWriter, rq *http.Request) {
		message, uid := "", "-1"
		if cookies == "" {
			message = "用户未登录"
		} else {
			uid = bilibili.GetUserMID(cookies)
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
			log.Fatal(err)
		}
		defer ws.Close()

		messageType, mid, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fav, av, page := "", "", ""
		videoItem, videoPage := bilibili.FavoriteListItemVideo{}, bilibili.VideoPage{}
		bilibili.IterateFavoriteList(string(mid), cookies, func(key, value string, data interface{}) {
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

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
