package bilibili

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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

type MIDResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	TTL     int     `json:"ttl"`
	Data    MIDInfo `json:"data"`
}

type MIDInfo struct {
	Mid        int            `json:"mid"`
	Name       string         `json:"name"`
	Sex        string         `json:"sex"`
	Face       string         `json:"face"`
	Sign       string         `json:"sign"`
	Rank       int            `json:"rank"`
	Level      int            `json:"level"`
	Jointime   int            `json:"jointime"`
	Moral      int            `json:"moral"`
	Silence    int            `json:"silence"`
	Birthday   string         `json:"birthday"`
	Coins      float64        `json:"coins"`
	FansBadge  bool           `json:"fans_badge"`
	Official   officialStatus `json:"official"`
	Vip        vipStatus      `json:"vip"`
	IsFollowed bool           `json:"is_followed"`
	TopPhoto   string         `json:"top_photo"`
	Theme      interface{}    `json:"theme"`
}

type officialStatus struct {
	Role  int    `json:"role"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type vipStatus struct {
	Type   int `json:"type"`
	Status int `json:"status"`
}

func GetUserMID(cookies string) string {
	client := http.Client{}
	body := MySpaceResponse{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/nav", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Add("Cookie", cookies)

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	if body.Code != 0 {
		return "-1"
	}

	return strconv.Itoa(body.Data.Mid)
}

func GetMIDInfo(mid string) (MIDInfo, error) {
	client := http.Client{}
	body := MIDResponse{}

	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/space/acc/info?mid="+mid, nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatal(err.Error())
	}

	if body.Code != 0 {
		return body.Data, errors.New("Invalid Response code.")
	}

	return body.Data, nil
}
