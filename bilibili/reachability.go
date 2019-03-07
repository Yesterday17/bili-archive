package bilibili

import (
	"net/http"
	"time"
)

func TestPage(url string) bool {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(url)
	if err != nil {
		return false
	}
	return true
}

func TestMainSite() bool {
	return TestPage("https://www.bilibili.com/")
}

func TestFavoriteList() bool {
	return TestPage("http://api.bilibili.com/x/space/fav/nav")
}

func TestFavoriteListItem() bool {
	return TestPage("https://api.bilibili.com/x/space/fav/arc")
}

func TestLoginQR() bool {
	return TestPage("https://passport.bilibili.com/qrcode/getLoginUrl")
}

func TestLoginInfo() bool {
	return TestPage("https://passport.bilibili.com/qrcode/getLoginInfo")
}

func TestSpace() bool {
	return TestPage("https://api.bilibili.com/x/web-interface/nav")
}

func TestVideoPage() bool {
	return TestPage("https://www.bilibili.com/widget/getPageList")
}
