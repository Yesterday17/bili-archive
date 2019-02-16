package server

import (
	"encoding/json"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func CreateBiliArchiveServer() {
	cookies := ""
	code := bilibili.QRCode{}
	handler := http.NewServeMux()

	// Frontend
	handler.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./server/public"))))

	// Backend
	loginQRHandler := func(w http.ResponseWriter, req *http.Request) {
		code = bilibili.GetLoginQRCode()
		output, err := json.Marshal(map[string]string{"image": code.Image})

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/login-qr", loginQRHandler)

	loginStatusHandler := func(w http.ResponseWriter, req *http.Request) {
		ok := false
		ok, cookies = code.Check()
		output, err := json.Marshal(map[string]bool{"ok": ok})

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	handler.HandleFunc("/login-status", loginStatusHandler)

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

		bilibili.IterateFavoriteList(string(mid), cookies, func(key, value string, data interface{}) {
			err := ws.WriteMessage(messageType, []byte(key+": "+value))
			if err != nil {
				log.Println(err)
			}
		})

	}
	handler.HandleFunc("/ws", iterateFavHandler)

	server := &http.Server{
		Addr:        ":8080",
		Handler:     handler,
		ReadTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
