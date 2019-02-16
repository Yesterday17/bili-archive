package bilibili

import (
	"encoding/base64"
	"encoding/json"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"net/url"
	"sort"
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
	client := &http.Client{}
	cookies := ""

	params := url.Values{}
	params.Set("oauthKey", this.OauthKey)
	params.Set("gourl", "https://passport.bilibili.com/ajax/miniLogin/redirect")

	res, err := client.PostForm("https://passport.bilibili.com/qrcode/getLoginInfo", params)
	if err != nil {
		log.Fatal(err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
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
	QRLogin qrLogin
	Png     []byte
	Image   string
}

func (this QRCode) Check() (bool, string) {
	return this.QRLogin.heartbeat()
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
			time.Sleep(3 * time.Second)
		}
	}

	// cookies, timeout
	return cookies, strings.EqualFold(response, "")
}

func GetCookiesString(link string) string {
	cookies := url.Values{}
	link = link[42 : len(link)-72]
	for _, value := range strings.Split(link, "&") {
		ans := strings.Split(value, "=")
		cookies.Set(ans[0], ans[1])
	}
	return Cookies(cookies)
}

func getLoginAddr() qrLogin {
	var body loginResponse
	res, err := http.Get("https://passport.bilibili.com/qrcode/getLoginUrl")

	if err != nil || res.Body == nil {
		log.Fatal(err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
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

	return QRCode{
		QRLogin: login,
		Image:   image,
		Png:     png,
	}
}

func Cookies(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte(';')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
