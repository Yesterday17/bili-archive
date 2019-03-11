package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
	_ "github.com/Yesterday17/bili-archive/statik"
	"github.com/Yesterday17/bili-archive/utils"
	"github.com/gorilla/websocket"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

		list, _ := bilibili.GetFavoriteList(uid, configuration.Cookies)
		output, _ := json.Marshal(map[string]interface{}{
			"ok":   list != nil,
			"data": list,
		})

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
			list, _ = bilibili.GetFavoriteListItems(uid, fid, pn, configuration.Cookies)
		}

		output, _ := json.Marshal(map[string]interface{}{
			"ok":   list != nil,
			"data": list,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/fav", favDetail)

	// Path: /api/pages
	// Method: GET
	// Params: @aid string
	// Description: 视频分P详情
	// Response: @ok boolean
	// 			 @data []VideoPage
	videoPages := func(w http.ResponseWriter, req *http.Request) {
		aid := req.URL.Query().Get("aid")
		pages, err := bilibili.GetVideoPages(aid)
		output, _ := json.Marshal(map[string]interface{}{
			"ok":   err == nil,
			"data": pages,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/pages", videoPages)

	// Path: /api/pic
	// Method: GET
	// Params: @url string
	// Description: 获取图片
	// Response: image/jpeg
	getPicture := func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.Query().Get("url")
		res, err := utils.Request("GET", url, configuration.Cookies, nil, nil)
		if err != nil {
			log.Println(err)
			return
		}
		if res.Body != nil {
			defer res.Body.Close()
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return
		}

		// Base64
		//str := base64.StdEncoding.EncodeToString(data)
		//w.Write([]byte("data:image/jpg;base64," + str))

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(data)
	}
	handler.HandleFunc("/api/pic", getPicture)

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
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			if err := ws.WriteJSON(&map[string]interface{}{
				"status": "error",
				"data":   err,
			}); err != nil {
				log.Println(err)
			}
			return
		}
		defer ws.Close()

		data := bilibili.DownloadVideoRequest{}
		if err := ws.ReadJSON(&data); err != nil {
			log.Println(err)
			if err := ws.WriteJSON(&map[string]interface{}{
				"status": "error",
				"data":   err,
			}); err != nil {
				log.Println(err)
			}
			return
		}

		if err := os.MkdirAll("./video/"+data.FavTitle, os.ModePerm); err != nil {
			log.Println(err)
			if err := ws.WriteJSON(&map[string]interface{}{
				"status": "error",
				"data":   err,
			}); err != nil {
				log.Println(err)
			}
			return
		}

		outputName := fmt.Sprintf("%s - %s", data.Title, data.Page.PageName)
		outputPath := "./video/" + string(data.FavTitle)

		if _, err := os.Stat(outputPath + "/" + outputName + ".flv"); os.IsNotExist(err) {
			if _, err := os.Stat(outputPath + "/" + outputName + ".mp4"); os.IsNotExist(err) {
				// 获得视频链接
				item := bilibili.ExtractVideo(data, configuration.Cookies)
				if item.Err != nil {
					log.Println(item.Err)
					result := map[string]interface{}{
						"status": "error",
						"data":   item.Err,
					}
					if err := ws.WriteJSON(&result); err != nil {
						log.Println(err)
					}
					return
				}

				callback := func(pg *utils.Progress) {
					ws.WriteJSON(pg)
				}

				// 下载视频
				if err = bilibili.DownloadVideo(item, data, outputPath, configuration.Cookies, callback); err != nil {
					log.Println(err)
					if err := ws.WriteJSON(&map[string]interface{}{
						"status": "error",
						"data":   err,
					}); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
	handler.HandleFunc("/download", downloadVideo)

	// Path: /api/test
	// Method: GET
	// Description: 测试网络与系统
	// Response: 测试结果
	testSericeHandler := func(w http.ResponseWriter, req *http.Request) {
		output, _ := json.Marshal(map[string]bool{
			"main":               bilibili.TestMainSite(),
			"login_qr":           bilibili.TestLoginQR(),
			"login_info":         bilibili.TestLoginInfo(),
			"space":              bilibili.TestSpace(),
			"video_page":         bilibili.TestVideoPage(),
			"favorite_list":      bilibili.TestFavoriteList(),
			"favorite_list_item": bilibili.TestFavoriteListItem(),
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/api/test", testSericeHandler)

	if err := http.ListenAndServe(":"+configuration.Port, handler); err != nil {
		log.Fatal(err)
	}
}
