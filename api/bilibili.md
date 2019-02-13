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

## Bilibili API(收藏部分)

http://api.bilibili.com/x/space/fav/nav?mid=123817257
