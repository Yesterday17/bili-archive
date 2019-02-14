package bilibili

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type loginResponse struct {
	Code      int     `json:"code"`
	Status    bool    `json:"status"`
	Timestamp int     `json:"ts"`
	Data      qrLogin `json:"data"`
}

type qrLogin struct {
	Url      string `json:"url"`
	OauthKey string `json:"oauthKey"`
}

// heartbeat, should be called per 3 seconds
func (this qrLogin) heartbeat() (bool, string) {
	body := map[string]interface{}{}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	cookies := ""

	params := url.Values{}
	params.Set("oauthKey", this.OauthKey)
	params.Set("gourl", "https://passport.bilibili.com/ajax/miniLogin/redirect")

	res, err := client.PostForm("https://passport.bilibili.com/qrcode/getLoginInfo", params)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	if body["status"] == true {
		data := body["data"].(map[string]interface{})
		cookies = data["url"].(string)
	}

	return body["status"].(bool), cookies
}

type QRCode struct {
	qrlogin qrLogin
	png     []byte
	image   string
}

func (this QRCode) Check() (bool, string) {
	return this.qrlogin.heartbeat()
}

func (this QRCode) WaitForLogin() (url.Values, bool) {
	response := ""
	cookies := url.Values{}
	for times := 0; times < 60; times++ {
		ok, ret := this.Check()
		response = ret

		if ok {
			response = response[42 : len(response)-72]
			for _, value := range strings.Split(response, "&") {
				ans := strings.Split(value, "=")
				cookies.Set(ans[0], ans[1])
			}
			break
		} else {
			time.NewTicker(3 * time.Second)
		}
	}

	// cookies, timeout
	return cookies, strings.EqualFold(response, "")
}

func getLoginAddr() qrLogin {
	var body loginResponse
	res, err := http.Get("https://passport.bilibili.com/qrcode/getLoginUrl")

	if err != nil || res.Body == nil {
		log.Fatal(err.Error())
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}
	return body.Data
}

func GetLoginQRCode() QRCode {
	login := getLoginAddr()
	png, err := qrcode.Encode(login.Url, qrcode.High, 256)

	if err != nil {
		log.Fatal(err.Error())
	}

	image := "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)

	fmt.Println(image)

	return QRCode{
		qrlogin: login,
		image:   image,
		png:     png,
	}
}
