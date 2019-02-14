package bilibili

import (
	"encoding/json"
	"log"
	"net/http"
)

type MySpaceResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	TTL     int                 `json:"ttl"`
	Data    mySpaceResponseData `json:"data"`
}

type mySpaceResponseData struct {
	IsLogin        bool                  `json:"isLogin"`
	EmailVerified  int                   `json:"email_verified"`
	Face           string                `json:"face"`
	LevelInfo      mySpaceLevelInfo      `json:"level_info"`
	Mid            int                   `json:"mid"`
	MobileVerified int                   `json:"mobile_verified"`
	Money          float32               `json:"money"`
	Moral          int                   `json:"moral"`
	OfficialVerify mySpaceOfficialVerify `json:"officialVerify"`
	Pendant        mySpacePendant        `json:"pendant"`
	Scores         int                   `json:"scores"`
	UserName       string                `json:"uname"`
	VipDueDate     int                   `json:"vipDueDate"`
	IsVip          int                   `json:"vipStatus"`
	VipType        int                   `json:"vipType"`
	VipPayType     int                   `json:"vip_pay_type"`
	Wallet         mySpaceWallet         `json:"wallet"`
	HasShop        bool                  `json:"has_shop"`
	ShopUrl        string                `json:"shop_url"`
	AllowanceCount int                   `json:"allowance_count"`
}

type mySpaceLevelInfo struct {
	CurrentLevel int `json:"current_level"`
	CurrentMin   int `json:"current_min"`
	CurrentExp   int `json:"current_exp"`
	NextExp      int `json:"next_exp"`
}

type mySpaceOfficialVerify struct {
	Type        int    `json:"type"`
	Description string `json:"desc"`
}

type mySpacePendant struct {
	PID    int    `json:"pid"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Expire int    `json:"expire"`
}

type mySpaceWallet struct {
	Mid           int `json:"mid"`
	BCoinBalance  int `json:"bcoin_balance"`
	CouponBalance int `json:"coupon_balance"`
	CouponDueTime int `json:"coupon_due_time"`
}

func GetUserMID(cookies string) int {
	client := http.Client{}
	body := MySpaceResponse{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/nav", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Add("Cookie", cookies)

	res, err := client.Do(req)
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	if body.Code != 0 {
		log.Fatal(body.Message)
	}

	return body.Data.Mid
}
