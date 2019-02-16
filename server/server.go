package server

import (
	"encoding/json"
	"github.com/Yesterday17/bili-archive/bilibili"
	"log"
	"net/http"
	"time"
)

func CreateBiliArchiveServer(code bilibili.QRCode) {
	cookies := ""
	serveMux := http.NewServeMux()

	// Frontend
	serveMux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./server/public"))))

	// Backend
	loginQRHandler := func(w http.ResponseWriter, req *http.Request) {

		output, err := json.Marshal(map[string]string{"image": code.Image})

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
	serveMux.HandleFunc("/loginqr", loginQRHandler)

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
	serveMux.HandleFunc("/loginstatus", loginStatusHandler)

	server := &http.Server{
		Addr:        ":8080",
		Handler:     serveMux,
		ReadTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
