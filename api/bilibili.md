## Login
1. Get login QR Code  
URL: https://passport.bilibili.com/qrcode/getLoginUrl
Interface:
```javascript
{
    code: number;
    status: boolean;
    ts: number; // Timestamp
    data: {
        url: string; // The content of QR Code
        oauthKey: string; // used in further post
    }
}
```

2. Get login status after getting QR Code  
URL: https://passport.bilibili.com/qrcode/getLoginInfo  
Method: POST  
Body: 
```json
{
    "oauthKey": "#oauthkey#",
    "gourl": "https://passport.bilibili.com/ajax/miniLogin/redirect"
}
```
Response:
```json
{"status":false,"data":-4,"message":"Can't scan~"}
```
```javascript
{
    "code":0,
    "status":true,
    "ts":1550073408,
    "data":{
        "url":"https://passport.biligame.com/crossDomain?" +
         "DedeUserID=####" +
          "&DedeUserID__ckMd5=####" +
           "&Expires=2592000&SESSDATA=####" +
            "&bili_jct=####" +
             "&gourl=https%3A%2F%2Fpassport.bilibili.com%2Fajax%2FminiLogin%2Fredirect"
    }
}
```

## User Status
https://api.bilibili.com/x/web-interface/nav
```json
{
  "code": -101,
  "message": "账号未登录",
  "ttl": 1,
  "data": {
    "isLogin": false
  }
}
```
```json
{
  "code": 0,
  "message": "0",
  "ttl": 1,
  "data": {
    "isLogin": true,
    "email_verified": 1,
    "face": "http://i0.hdslb.com/bfs/face/7115c5ff964c5aea3bc70792fdb2073b1f6dcb19.jpg",
    "level_info": {
      "current_level": 5,
      "current_min": 10800,
      "current_exp": 15577,
      "next_exp": 28800
    },
    "mid": 5756570,
    "mobile_verified": 1,
    "money": 737.06,
    "moral": 71,
    "officialVerify": {
      "type": -1,
      "desc": ""
    },
    "pendant": {
      "pid": 0,
      "name": "",
      "image": "",
      "expire": 0
    },
    "scores": 0,
    "uname": "某昨P",
    "vipDueDate": 1593187200000,
    "vipStatus": 1,
    "vipType": 2,
    "vip_pay_type": 0,
    "wallet": {
      "mid": 5756570,
      "bcoin_balance": 5,
      "coupon_balance": 5,
      "coupon_due_time": 0
    },
    "has_shop": false,
    "shop_url": "",
    "allowance_count": 0
  }
}
```

## Bilibili API(收藏部分)

http://api.bilibili.com/x/space/fav/nav?mid=123817257
