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
URL: https://api.bilibili.com/x/web-interface/nav  
Method: GET  
Header: Cookies  
Response:
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

## Favorite
URL: http://api.bilibili.com/x/space/fav/nav?mid=123817257  
Method: GET
Response: 
```json
{
  "code": 0,
  "message": "0",
  "ttl": 1,
  "data": {
    "archive": [
      {
        "media_id": 121198657,
        "fid": 1211986,
        "mid": 123817257,
        "name": "默认收藏夹",
        "max_count": 50000,
        "cur_count": 5,
        "atten_count": 0,
        "favoured": 0,
        "state": 0,
        "ctime": 1496128120,
        "mtime": 1549024274,
        "cover": [
          {
            "aid": 11300322,
            "pic": "http://i1.hdslb.com/bfs/archive/5d77b48051ddcbc0a6f61ae50a89be57fae1d9f4.jpg",
            "type": 2
          },
          {
            "aid": 22275566,
            "pic": "http://i1.hdslb.com/bfs/archive/87358d0568680b76ac70b5018c13fee641bda79d.jpg",
            "type": 2
          },
          {
            "aid": 34239281,
            "pic": "http://i1.hdslb.com/bfs/archive/3814b78edf9cf8ca3626fca1b2d72b5e440434b3.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 316675157,
        "fid": 3166751,
        "mid": 123817257,
        "name": "扎",
        "max_count": 999,
        "cur_count": 2,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1548928782,
        "mtime": 1548931133,
        "cover": [
          {
            "aid": 4040870,
            "pic": "http://i0.hdslb.com/bfs/archive/99ae7782fab31875c2789ec16a0fd080edf5cd01.jpg",
            "type": 2
          },
          {
            "aid": 16332628,
            "pic": "http://i0.hdslb.com/bfs/archive/4909f5c7fa2b0a81a65e468a55ba805bf26b7347.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 313672657,
        "fid": 3136726,
        "mid": 123817257,
        "name": "天酱",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1548267555,
        "mtime": 1548267556,
        "cover": [
          {
            "aid": 8234839,
            "pic": "http://i2.hdslb.com/bfs/archive/26f13ef587e1dc6a12b5f0daf7d178c6c8c76fc6.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 313120757,
        "fid": 3131207,
        "mid": 123817257,
        "name": "miku",
        "max_count": 999,
        "cur_count": 4,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1548147335,
        "mtime": 1548267539,
        "cover": [
          {
            "aid": 8234839,
            "pic": "http://i2.hdslb.com/bfs/archive/26f13ef587e1dc6a12b5f0daf7d178c6c8c76fc6.jpg",
            "type": 2
          },
          {
            "aid": 14940736,
            "pic": "http://i1.hdslb.com/bfs/archive/cc88f948705f8311022a21651e138867bc9af17c.jpg",
            "type": 2
          },
          {
            "aid": 32418431,
            "pic": "http://i2.hdslb.com/bfs/archive/d52b99c8cd98ed58807c63d007dd5143b47ac107.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310813357,
        "fid": 3108133,
        "mid": 123817257,
        "name": "麻将",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547564080,
        "mtime": 1547564081,
        "cover": [
          {
            "aid": 39275234,
            "pic": "http://i0.hdslb.com/bfs/archive/a6f72201265d2353987b08c5023a22eb6885c220.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310803857,
        "fid": 3108038,
        "mid": 123817257,
        "name": "UC",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547562669,
        "mtime": 1547562670,
        "cover": [
          {
            "aid": 6592774,
            "pic": "http://i2.hdslb.com/bfs/archive/9466eddbecddee88f990afb6c2b0315988742ca5.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310781257,
        "fid": 3107812,
        "mid": 123817257,
        "name": "张京华",
        "max_count": 999,
        "cur_count": 4,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547559222,
        "mtime": 1547615305,
        "cover": [
          {
            "aid": 39300041,
            "pic": "http://i1.hdslb.com/bfs/archive/2c4d566f6eaf112d1459c3e94117bcc6509320f2.jpg",
            "type": 2
          },
          {
            "aid": 36779481,
            "pic": "http://i1.hdslb.com/bfs/archive/5dc93dd7978c4213cfb10710dc837727a2ecb8be.jpg",
            "type": 2
          },
          {
            "aid": 38678132,
            "pic": "http://i0.hdslb.com/bfs/archive/2ac80a340fff1aca8729cb9a4ddd541af061a827.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310780757,
        "fid": 3107807,
        "mid": 123817257,
        "name": "谢拉",
        "max_count": 999,
        "cur_count": 3,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547559197,
        "mtime": 1547721780,
        "cover": [
          {
            "aid": 39166699,
            "pic": "http://i1.hdslb.com/bfs/archive/e68c39532da7fe6c5baf4d61137af41baa9721b1.jpg",
            "type": 2
          },
          {
            "aid": 36779481,
            "pic": "http://i1.hdslb.com/bfs/archive/5dc93dd7978c4213cfb10710dc837727a2ecb8be.jpg",
            "type": 2
          },
          {
            "aid": 39914727,
            "pic": "http://i1.hdslb.com/bfs/archive/7263c01c67049bb4f47fefd881a1fd68013613ab.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310365857,
        "fid": 3103658,
        "mid": 123817257,
        "name": "犬山",
        "max_count": 999,
        "cur_count": 2,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547444575,
        "mtime": 1547558004,
        "cover": [
          {
            "aid": 38695532,
            "pic": "http://i1.hdslb.com/bfs/archive/634a49e3b4eeed88eeb378e53f50e88f4ed9c36c.jpg",
            "type": 2
          },
          {
            "aid": 39758775,
            "pic": "http://i1.hdslb.com/bfs/archive/17d7ae872346279d0e3f1e9051c17fa823088cae.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 310250757,
        "fid": 3102507,
        "mid": 123817257,
        "name": "草",
        "max_count": 999,
        "cur_count": 9,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547396831,
        "mtime": 1549806906,
        "cover": [
          {
            "aid": 42932936,
            "pic": "http://i0.hdslb.com/bfs/archive/af47728bc41ea7a288fc99190b7c207be9d5935a.jpg",
            "type": 2
          },
          {
            "aid": 37621450,
            "pic": "http://i1.hdslb.com/bfs/archive/8237f04b51394eee58f37c33a088b135e0ff26ac.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 308900457,
        "fid": 3089004,
        "mid": 123817257,
        "name": "糖",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1547091047,
        "mtime": 1549218376,
        "cover": [
          {
            "aid": 40310945,
            "pic": "http://i0.hdslb.com/bfs/archive/d7713016fafb025eb082be3ce33d68335eaf1df1.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 301281257,
        "fid": 3012812,
        "mid": 123817257,
        "name": "Δ",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1545459371,
        "mtime": 1546537993,
        "cover": [
          {
            "aid": 38708224,
            "pic": "http://i0.hdslb.com/bfs/archive/c6a17d7d7c6f3a06f38214c7fc8503eecba07876.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 300592957,
        "fid": 3005929,
        "mid": 123817257,
        "name": "yui",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1545236646,
        "mtime": 1546537980,
        "cover": [
          {
            "aid": 38477047,
            "pic": "http://i0.hdslb.com/bfs/archive/f54f34a91b7b4def0037cc4f6c2399565c7e4456.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 297692257,
        "fid": 2976922,
        "mid": 123817257,
        "name": "屑女仆",
        "max_count": 999,
        "cur_count": 8,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1544416614,
        "mtime": 1549218363,
        "cover": [
          {
            "aid": 40822765,
            "pic": "http://i0.hdslb.com/bfs/archive/21387c7c1307df9c559c6c3aa1a7f3230d20ee03.jpg",
            "type": 2
          },
          {
            "aid": 33054830,
            "pic": "http://i2.hdslb.com/bfs/archive/e622f4c788fd8a99c3748927f9f6536c4e738256.jpg",
            "type": 2
          },
          {
            "aid": 36298522,
            "pic": "http://i1.hdslb.com/bfs/archive/cfe107af2517f897d97c78583a3dde02f81b843e.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 296167157,
        "fid": 2961671,
        "mid": 123817257,
        "name": "生放",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1543971299,
        "mtime": 1543971300,
        "cover": [
          {
            "aid": 36999814,
            "pic": "http://i1.hdslb.com/bfs/archive/5240200e93fd5b032dbeec19dada182e65afa9a6.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 295277857,
        "fid": 2952778,
        "mid": 123817257,
        "name": "西明日香",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1543689031,
        "mtime": 1543689034,
        "cover": [
          {
            "aid": 37144488,
            "pic": "http://i0.hdslb.com/bfs/archive/9741d0feb3cd9d2ebe75bc85c36066f48e1da659.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 294575257,
        "fid": 2945752,
        "mid": 123817257,
        "name": "CONCEPTION广播生肉",
        "max_count": 999,
        "cur_count": 12,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1543506197,
        "mtime": 1546538038,
        "cover": [
          {
            "aid": 38900728,
            "pic": "http://i2.hdslb.com/bfs/archive/5e45d878544701b2796d66938f6bec01e6c54f01.jpg",
            "type": 2
          },
          {
            "aid": 38360567,
            "pic": "http://i1.hdslb.com/bfs/archive/fb6a748ce4e69434e46062d1369911b6452e61cb.jpg",
            "type": 2
          },
          {
            "aid": 37894548,
            "pic": "http://i0.hdslb.com/bfs/archive/0578fa4ebe119497721d1f90c95613ded1cc5e0f.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 291861557,
        "fid": 2918615,
        "mid": 123817257,
        "name": "水瀬祈 MELODY FLAG",
        "max_count": 999,
        "cur_count": 12,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1543249978,
        "mtime": 1549715166,
        "cover": [
          {
            "aid": 7602065,
            "pic": "http://i0.hdslb.com/bfs/archive/dc1dd42160fc62326ce3e48da832491072a2dafb.jpg",
            "type": 2
          },
          {
            "aid": 7505904,
            "pic": "http://i1.hdslb.com/bfs/archive/a5ad91885379d5178107c3e18157ab51d299880d.jpg",
            "type": 2
          },
          {
            "aid": 7395882,
            "pic": "http://i0.hdslb.com/bfs/archive/4fa396ccf0d5352217e37c12311da1414fc513cc.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 290695157,
        "fid": 2906951,
        "mid": 123817257,
        "name": "karz",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1542928873,
        "mtime": 1542928877,
        "cover": [
          {
            "aid": 12903721,
            "pic": "http://i0.hdslb.com/bfs/archive/0689cc579f6ec78461c5f586b0a8be822b75e54c.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 285356757,
        "fid": 2853567,
        "mid": 123817257,
        "name": "ことりの音",
        "max_count": 999,
        "cur_count": 6,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541781824,
        "mtime": 1549218410,
        "cover": [
          {
            "aid": 34275915,
            "pic": "http://i1.hdslb.com/bfs/archive/f2eae5f68e1336b9ceb2aceb80c261e22723033f.jpg",
            "type": 2
          },
          {
            "aid": 32761514,
            "pic": "http://i1.hdslb.com/bfs/archive/f2eae5f68e1336b9ceb2aceb80c261e22723033f.jpg",
            "type": 2
          },
          {
            "aid": 31683768,
            "pic": "http://i2.hdslb.com/bfs/archive/92c8b42903252926db338d65e8833650e5135dc5.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 284692557,
        "fid": 2846925,
        "mid": 123817257,
        "name": "政委",
        "max_count": 999,
        "cur_count": 7,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541566186,
        "mtime": 1548975660,
        "cover": [
          {
            "aid": 7410021,
            "pic": "http://i1.hdslb.com/bfs/archive/faccff36d7b9b42326954963dc921451d9aaa4e4.jpg",
            "type": 2
          },
          {
            "aid": 4040870,
            "pic": "http://i0.hdslb.com/bfs/archive/99ae7782fab31875c2789ec16a0fd080edf5cd01.jpg",
            "type": 2
          },
          {
            "aid": 16858557,
            "pic": "http://i2.hdslb.com/bfs/archive/692371921c1e19bb9f37938758c706c6dc25bfbb.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 284692057,
        "fid": 2846920,
        "mid": 123817257,
        "name": "女儿",
        "max_count": 999,
        "cur_count": 10,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541566070,
        "mtime": 1549335773,
        "cover": [
          {
            "aid": 10607206,
            "pic": "http://i0.hdslb.com/bfs/archive/4d526f4d7c236aeb280c86e80f2830fb8b3b5987.jpg",
            "type": 2
          },
          {
            "aid": 17676127,
            "pic": "http://i0.hdslb.com/bfs/archive/e7fe160a94eb6fa5bd7cc2c56761498bffff311e.jpg",
            "type": 2
          },
          {
            "aid": 7339753,
            "pic": "http://i2.hdslb.com/bfs/archive/807eca58169e4b918616894fb5f2cff899a3751c.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 284516057,
        "fid": 2845160,
        "mid": 123817257,
        "name": "kido",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541500541,
        "mtime": 1546537967,
        "cover": [
          {
            "aid": 3673704,
            "pic": "http://i0.hdslb.com/bfs/archive/4056d14769a63db4b8aa11363416fe6ed946c889.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 284513557,
        "fid": 2845135,
        "mid": 123817257,
        "name": "鱼香",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541500079,
        "mtime": 1541500081,
        "cover": [
          {
            "aid": 4477866,
            "pic": "http://i0.hdslb.com/bfs/archive/44d785b257156036cd4665af3ac51d113451828a.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 283019357,
        "fid": 2830193,
        "mid": 123817257,
        "name": "百万",
        "max_count": 999,
        "cur_count": 27,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1541111798,
        "mtime": 1549813773,
        "cover": [
          {
            "aid": 33160676,
            "pic": "http://i0.hdslb.com/bfs/archive/e1285fe89368c405c87434d4ca4634064a273d04.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282502857,
        "fid": 2825028,
        "mid": 123817257,
        "name": "コール",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540915440,
        "mtime": 1540915442,
        "cover": [
          {
            "aid": 33224769,
            "pic": "http://i1.hdslb.com/bfs/archive/de0b4173d243e7693ae6b50d79d61e2a0be39d9f.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282502257,
        "fid": 2825022,
        "mid": 123817257,
        "name": "Live",
        "max_count": 999,
        "cur_count": 4,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540915338,
        "mtime": 1547845222,
        "cover": [
          {
            "aid": 40661961,
            "pic": "http://i0.hdslb.com/bfs/archive/e25770882913d4f40f4da7f6848a883bc55aca66.jpg",
            "type": 2
          },
          {
            "aid": 34457911,
            "pic": "http://i0.hdslb.com/bfs/archive/3525b6832b9946db4db78791f8254de7b58f0ebe.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282501057,
        "fid": 2825010,
        "mid": 123817257,
        "name": "松冈汉堡",
        "max_count": 999,
        "cur_count": 3,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540915150,
        "mtime": 1542793486,
        "cover": [
          {
            "aid": 36446745,
            "pic": "http://i0.hdslb.com/bfs/archive/61eae9a2944f6cf32a23340402e8f9105f18ae46.png",
            "type": 2
          },
          {
            "aid": 35347846,
            "pic": "http://i0.hdslb.com/bfs/archive/714845d812b6bbbd40537c30c73dedc8b5de6a42.jpg",
            "type": 2
          },
          {
            "aid": 33389784,
            "pic": "http://i0.hdslb.com/bfs/archive/61eae9a2944f6cf32a23340402e8f9105f18ae46.png",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282404257,
        "fid": 2824042,
        "mid": 123817257,
        "name": "rina",
        "max_count": 999,
        "cur_count": 3,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540897130,
        "mtime": 1549964763,
        "cover": [
          {
            "aid": 41227467,
            "pic": "http://i2.hdslb.com/bfs/archive/5bbc4b73c9d92c34db400fcf01aa71f68338064a.png",
            "type": 2
          },
          {
            "aid": 41372626,
            "pic": "http://i1.hdslb.com/bfs/archive/48eb26df7e585b3fd0d16a247a662a6f23121191.png",
            "type": 2
          },
          {
            "aid": 38449662,
            "pic": "http://i2.hdslb.com/bfs/archive/9d017f7ae55eebaa105d66ac3eba95b40c68646b.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282403557,
        "fid": 2824035,
        "mid": 123817257,
        "name": "蒸饺",
        "max_count": 999,
        "cur_count": 1,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540896868,
        "mtime": 1540896870
      },
      {
        "media_id": 282256157,
        "fid": 2822561,
        "mid": 123817257,
        "name": "[tkmn] み、味方はナシ",
        "max_count": 999,
        "cur_count": 21,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540828220,
        "mtime": 1548504770,
        "cover": [
          {
            "aid": 41531605,
            "pic": "http://i0.hdslb.com/bfs/archive/62ed9aef253bb8fe9f770c3f8f15a350fce9b7ab.jpg",
            "type": 2
          },
          {
            "aid": 40395418,
            "pic": "http://i0.hdslb.com/bfs/archive/b0ad7110db2433f5ef061538e7bd9b387c5a0478.jpg",
            "type": 2
          },
          {
            "aid": 39121432,
            "pic": "http://i2.hdslb.com/bfs/archive/77c69c2d14997f64c8c3eb5dc2c3c9496782fc4a.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030757,
        "fid": 2820307,
        "mid": 123817257,
        "name": "大西",
        "max_count": 999,
        "cur_count": 5,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769718,
        "mtime": 1547562683,
        "cover": [
          {
            "aid": 6592774,
            "pic": "http://i2.hdslb.com/bfs/archive/9466eddbecddee88f990afb6c2b0315988742ca5.png",
            "type": 2
          },
          {
            "aid": 39810661,
            "pic": "http://i0.hdslb.com/bfs/archive/55dd7055a35bff0b958baa7991266a09c9a05ac4.jpg",
            "type": 2
          },
          {
            "aid": 11623774,
            "pic": "http://i2.hdslb.com/bfs/archive/b102ad52ac187e09b82c88a87e2e22a758e02f3c.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030657,
        "fid": 2820306,
        "mid": 123817257,
        "name": "闹",
        "max_count": 999,
        "cur_count": 0,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769710,
        "mtime": 1547562568
      },
      {
        "media_id": 282030557,
        "fid": 2820305,
        "mid": 123817257,
        "name": "樱小姐",
        "max_count": 999,
        "cur_count": 7,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769702,
        "mtime": 1549085817,
        "cover": [
          {
            "aid": 4883884,
            "pic": "http://i0.hdslb.com/bfs/archive/f6b85defe26227119560f6a65c902e3c45830ee1.jpg",
            "type": 2
          },
          {
            "aid": 6592774,
            "pic": "http://i2.hdslb.com/bfs/archive/9466eddbecddee88f990afb6c2b0315988742ca5.png",
            "type": 2
          },
          {
            "aid": 18878916,
            "pic": "http://i0.hdslb.com/bfs/archive/be1237f8ab8bd7c5392ac5e8bf52bf943278b5ce.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030457,
        "fid": 2820304,
        "mid": 123817257,
        "name": "ruru",
        "max_count": 999,
        "cur_count": 5,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769692,
        "mtime": 1549845965,
        "cover": [
          {
            "aid": 43178078,
            "pic": "http://i0.hdslb.com/bfs/archive/9176afcf107405c6b46d3cfee56de78df5699485.jpg",
            "type": 2
          },
          {
            "aid": 37144488,
            "pic": "http://i0.hdslb.com/bfs/archive/9741d0feb3cd9d2ebe75bc85c36066f48e1da659.jpg",
            "type": 2
          },
          {
            "aid": 19614694,
            "pic": "http://i2.hdslb.com/bfs/archive/2a216cb4a722cced469c409c4d8619eb7560006f.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030357,
        "fid": 2820303,
        "mid": 123817257,
        "name": "种",
        "max_count": 999,
        "cur_count": 3,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769653,
        "mtime": 1548977482,
        "cover": [
          {
            "aid": 3044760,
            "pic": "http://i0.hdslb.com/bfs/archive/7c670c56776e6cf9ab1ac591eba735393694cee6.jpg",
            "type": 2
          },
          {
            "aid": 3199073,
            "pic": "http://i2.hdslb.com/bfs/archive/64b4b8111303f6c41a89f928199fb30536312351.jpg",
            "type": 2
          },
          {
            "aid": 2342012,
            "pic": "http://i2.hdslb.com/bfs/archive/a8c8204ddbb8dbe51624ef92bad8529229891594.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030257,
        "fid": 2820302,
        "mid": 123817257,
        "name": "tkmn",
        "max_count": 999,
        "cur_count": 26,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769633,
        "mtime": 1549716165,
        "cover": [
          {
            "aid": 42888880,
            "pic": "http://i0.hdslb.com/bfs/archive/c82986549ca6053fae1dc673a41d97975912fd53.jpg",
            "type": 2
          },
          {
            "aid": 42096249,
            "pic": "http://i1.hdslb.com/bfs/archive/2e0c15dac75a51168d25fb8858edfb0eb090a698.jpg",
            "type": 2
          },
          {
            "aid": 41878631,
            "pic": "http://i1.hdslb.com/bfs/archive/d5c1970c6d9f5b58b192d59bdb34db9e4d1be234.jpg",
            "type": 2
          }
        ]
      },
      {
        "media_id": 282030157,
        "fid": 2820301,
        "mid": 123817257,
        "name": "松冈食堂",
        "max_count": 999,
        "cur_count": 3,
        "atten_count": 0,
        "favoured": 0,
        "state": 2,
        "ctime": 1540769624,
        "mtime": 1542212737,
        "cover": [
          {
            "aid": 2663294,
            "pic": "http://i0.hdslb.com/bfs/archive/106238783a0011c50a3105d426040b7df4d8f6d1.jpg",
            "type": 2
          },
          {
            "aid": 2342012,
            "pic": "http://i2.hdslb.com/bfs/archive/a8c8204ddbb8dbe51624ef92bad8529229891594.jpg",
            "type": 2
          },
          {
            "aid": 2232699,
            "pic": "http://i2.hdslb.com/bfs/archive/abea8f6e8ebc69c337a40d67cac03b53d790c31b.jpg",
            "type": 2
          }
        ]
      }
    ],
    "playlist": 0,
    "topic": 0,
    "article": 6,
    "album": 0,
    "movie": 1
  }
}
```

## Detailed Favorite
URL: https://api.bilibili.com/x/space/fav/arc?vmid=5756570&ps=30&fid=1683980&pn=1  
Response: 
```json
{
  "code": 0,
  "message": "0",
  "ttl": 1,
  "data": {
    "seid": "",
    "page": 0,
    "pagesize": 30,
    "pagecount": 1,
    "total": 13,
    "suggest_keyword": "",
    "mid": 5756570,
    "fid": 1683980,
    "tid": 0,
    "order": "mtime",
    "keyword": "",
    "archives": [
      {
        "aid": 19614694,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 2,
        "pic": "http://i2.hdslb.com/bfs/archive/2a216cb4a722cced469c409c4d8619eb7560006f.jpg",
        "title": "松冈祯丞毫无自觉的满级撩妹技能【佳村遥】",
        "pubdate": 1518656724,
        "ctime": 1518656725,
        "desc": "网络\n纠错：同期有男生也有女生（佐仓）不好意思少打了几个字误导了大家；特效字幕HARUKA后没有酱；欧拉（オラつく）装高冷\nsm28762895 这段剪辑视频在n站2天点击量破10w\n生放日已经过了几天但是up主作为一个默默关注神只看不说话的地下分（基）子（佬）终于忍不住出手了\n唯一神你撩妹为什么这么熟练？ruru一脸欢喜又是隐藏着什么故事啊？\nb站全篇生肉 av4543775 \n感谢小山君 投稿日期2016-05-05 21:56 播放8.2万 弹幕174 硬币82 收藏278",
        "state": 0,
        "attribute": 16384,
        "duration": 426,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 9912681,
          "name": "逆袭的金酱",
          "face": "http://i2.hdslb.com/bfs/face/7207fa9b38e62041367b5ddb6cffa61d529447c8.jpg"
        },
        "stat": {
          "aid": 19614694,
          "view": 71519,
          "danmaku": 361,
          "reply": 268,
          "favorite": 1199,
          "coin": 430,
          "share": 125,
          "now_rank": 0,
          "his_rank": 0,
          "like": 435,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 31984660,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1540743937,
        "play_num": "",
        "highlight_title": "松冈祯丞毫无自觉的满级撩妹技能【佳村遥】"
      },
      {
        "aid": 3860213,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i2.hdslb.com/bfs/archive/7bfa091258bcfde72247050e0c33e32a9abba273.jpg",
        "title": "【本能寺】元气少女  高桥未奈美",
        "pubdate": 1455630653,
        "ctime": 1497429433,
        "desc": "自制 第1回: http://www.bilibili.com/video/av2232699/\r\n第7回: http://www.bilibili.com/video/av3117710/\r\n第8回: http://www.bilibili.com/video/av3424728/\r\n番外篇①「切」: http://www.bilibili.com/video/av2663294/\r\n番外篇③「一品」: http://www.bilibili.com/video/av3209965/\r\n\r\n\r\n",
        "state": 0,
        "attribute": 49152,
        "duration": 283,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 6123298,
          "name": "敵在本能寺",
          "face": "http://i0.hdslb.com/bfs/face/b583a049bb767f0c6ecf96b711afab125ebed4a2.jpg"
        },
        "stat": {
          "aid": 3860213,
          "view": 59743,
          "danmaku": 439,
          "reply": 494,
          "favorite": 2138,
          "coin": 1636,
          "share": 271,
          "now_rank": 0,
          "his_rank": 0,
          "like": 388,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 6204864,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1539879140,
        "play_num": "",
        "highlight_title": "【本能寺】元气少女  高桥未奈美"
      },
      {
        "aid": 10876089,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 2,
        "pic": "http://i0.hdslb.com/bfs/archive/f9d4583abb8f1f979b07bde0f003edbed6ad7b63.jpg",
        "title": "从“不安因素”到有安心感的松冈祯丞",
        "pubdate": 1495926337,
        "ctime": 1497423967,
        "desc": "素材：av4091267，av2370274，av2701207，av10418868，av3424728，av10587413，av10452459，av10734692，av3117710",
        "state": 0,
        "attribute": 49152,
        "duration": 420,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 409205,
          "name": "云黑黑",
          "face": "http://i2.hdslb.com/bfs/face/3e5b543f73e4c3f12772b71a95bb221956c97aaf.jpg"
        },
        "stat": {
          "aid": 10876089,
          "view": 62924,
          "danmaku": 454,
          "reply": 325,
          "favorite": 1458,
          "coin": 339,
          "share": 95,
          "now_rank": 0,
          "his_rank": 0,
          "like": 102,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 18008781,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1539651790,
        "play_num": "",
        "highlight_title": "从“不安因素”到有安心感的松冈祯丞"
      },
      {
        "aid": 11958558,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/62e4d1ddb291ce2b479ad920255f06d6dcd510a8.jpg",
        "title": "dear...takamina！！！",
        "pubdate": 1499379134,
        "ctime": 1499379134,
        "desc": "av3209965 av2232699 av3117710 av11745637 av11355881 av2701207\n感谢av5703901的特效字幕和av3860213的灵感。人生第一次做视频自学了两晚。\n我很严肃地说一句，松冈和takamina的关系真的很好，不只是因为takamina开朗的性格，而是takamina会很认真地去跟松冈交流和相处，而松冈认真的性格也得到了takamina的尊敬。不管谐不谐，这两个人的组合真的很棒。\ntakamina赛高！！！",
        "state": 0,
        "attribute": 49152,
        "duration": 286,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 1431879,
          "name": "闲情都几许",
          "face": "http://i1.hdslb.com/bfs/face/cabc7c8ebfb49a4d2a36ed06da30205e3508406b.jpg"
        },
        "stat": {
          "aid": 11958558,
          "view": 6419,
          "danmaku": 53,
          "reply": 113,
          "favorite": 256,
          "coin": 276,
          "share": 42,
          "now_rank": 0,
          "his_rank": 0,
          "like": 40,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 19738752,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1539622532,
        "play_num": "",
        "highlight_title": "dear...takamina！！！"
      },
      {
        "aid": 11130377,
        "videos": 2,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/f98f4e890dba641bd5b0044e8974152c1f3218c4.jpg",
        "title": "【白学】【谐教】佐仓绫音：无法传达到的爱恋",
        "pubdate": 1496794194,
        "ctime": 1497433719,
        "desc": "素材来源（感谢大佬们的辛苦爆肝翻译）：\n  av4269399    av3692556    av1902464   av3526830 av5523754  av10424582   av4233897  av5758845 av3551212\nBGM：\nヲタみん - Letter Song\nAimer - あなたに出会わなければ~夏雪冬花~\n川嶋あい - 時雨\nタイナカ彩智 - 运命人\nLeaf - 届かない恋(Piano version)\n\n和2P一起食用效果更佳",
        "state": 0,
        "attribute": 49280,
        "duration": 878,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 17513056,
          "name": "windmyheart",
          "face": "http://i0.hdslb.com/bfs/face/7a34c01c62d1493d64f87c5f1f6b7fbdf512a7c4.jpg"
        },
        "stat": {
          "aid": 11130377,
          "view": 42199,
          "danmaku": 434,
          "reply": 244,
          "favorite": 1144,
          "coin": 632,
          "share": 124,
          "now_rank": 0,
          "his_rank": 0,
          "like": 112,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 18449473,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1536418247,
        "play_num": "",
        "highlight_title": "【白学】【谐教】佐仓绫音：无法传达到的爱恋"
      },
      {
        "aid": 22736333,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/9ac557b3ecd38d4459bcb146c5d6ffcfa4239a74.jpg",
        "title": "【松冈祯丞/大西沙织】敌不过的情敌、狐狸精出现篇",
        "pubdate": 1525059016,
        "ctime": 1525059017,
        "desc": "进击的里菜。还有，不到最后，你永远不知道流向的结局。\n情敌、狐狸精出现篇。  手动滑稽笑。\n\nAMV视频制作中，大概五月中旬吧。 \nFLOWER - 初恋",
        "state": 0,
        "attribute": 16768,
        "duration": 200,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 1,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 47163308,
          "name": "诠释泪",
          "face": "http://i0.hdslb.com/bfs/face/d21786bbd84d85281c19aec1fdef86603fb5ec2e.jpg"
        },
        "stat": {
          "aid": 22736333,
          "view": 41125,
          "danmaku": 182,
          "reply": 198,
          "favorite": 992,
          "coin": 785,
          "share": 93,
          "now_rank": 0,
          "his_rank": 0,
          "like": 270,
          "dislike": 0
        },
        "dynamic": "#声优##松冈祯丞##日高丽菜#",
        "cid": 37741975,
        "dimension": {
          "width": 1920,
          "height": 1080,
          "rotate": 0
        },
        "fav_at": 1536417115,
        "play_num": "",
        "highlight_title": "【松冈祯丞/大西沙织】敌不过的情敌、狐狸精出现篇"
      },
      {
        "aid": 23155472,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i1.hdslb.com/bfs/archive/5b774a4bdc856919513fe9f0720cac8c9c15129d.jpg",
        "title": "【松冈祯丞/佳村遥】路人女主的故事：花美不过樱，人美不过君",
        "pubdate": 1525766788,
        "ctime": 1525766791,
        "desc": "花美不过樱，人美不过君。\n佳村遥与松冈祯丞的故事。\n同期与同期的碰撞，为何。。。。、\nBGM：いきものがかり - 花は桜 君は美し\nAMV等动画剪辑视频又拖后了，就是控几不了自己的手 滑稽。\n松冈遥的故事才刚刚开始。",
        "state": 0,
        "attribute": 16512,
        "duration": 184,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 47163308,
          "name": "诠释泪",
          "face": "http://i1.hdslb.com/bfs/face/d21786bbd84d85281c19aec1fdef86603fb5ec2e.jpg"
        },
        "stat": {
          "aid": 23155472,
          "view": 27627,
          "danmaku": 129,
          "reply": 239,
          "favorite": 791,
          "coin": 728,
          "share": 106,
          "now_rank": 0,
          "his_rank": 0,
          "like": 249,
          "dislike": 0
        },
        "dynamic": "#松冈祯丞##声优##佳村遥#松冈遥的故事才刚刚开始，另一位同期该如何应对，悲伤的师妹是不是该反击一下喃",
        "cid": 38534076,
        "dimension": {
          "width": 1280,
          "height": 720,
          "rotate": 0
        },
        "fav_at": 1536416636,
        "play_num": "",
        "highlight_title": "【松冈祯丞/佳村遥】路人女主的故事：花美不过樱，人美不过君"
      },
      {
        "aid": 25380502,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i1.hdslb.com/bfs/archive/8316e87919b4cb3d468cc86873fe31773a287d97.jpg",
        "title": "【松冈祯丞/佐仓绫音】一直凝视着你，不曾忘记",
        "pubdate": 1529661690,
        "ctime": 1529661690,
        "desc": "一直凝视着你。依旧是不到最后你不会知道结局。\n\nBGM：Still--Flower",
        "state": 0,
        "attribute": 16768,
        "duration": 241,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 1,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 47163308,
          "name": "诠释泪",
          "face": "http://i1.hdslb.com/bfs/face/d21786bbd84d85281c19aec1fdef86603fb5ec2e.jpg"
        },
        "stat": {
          "aid": 25380502,
          "view": 50756,
          "danmaku": 245,
          "reply": 265,
          "favorite": 797,
          "coin": 777,
          "share": 87,
          "now_rank": 0,
          "his_rank": 0,
          "like": 360,
          "dislike": 0
        },
        "dynamic": "#声优##松冈祯丞##佐仓绫音#",
        "cid": 43120650,
        "dimension": {
          "width": 1920,
          "height": 1080,
          "rotate": 0
        },
        "fav_at": 1536416463,
        "play_num": "",
        "highlight_title": "【松冈祯丞/佐仓绫音】一直凝视着你，不曾忘记"
      },
      {
        "aid": 19923367,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/f49dbca1b0e61eab53ac940b9f8525dbe058a8fd.png",
        "title": "【松冈祯丞/日高里菜】里菜酱请再多陪我一下",
        "pubdate": 1519355675,
        "ctime": 1519355676,
        "desc": "里菜酱最近笔直的飙升啊，简直都直接A脸了啊，于是我就情不自禁的动起了手，这视频真的是边剪边笑。\r\n这两人也真的是认识了很久了啊。\r\n大佬教导我们要走的 正  行的稳，我深刻的领悟了大佬的思想灵魂，走的老稳了，一点都没飘。手动滑稽 笑。\r\n原本是制作了歌的特效字幕的，但做到大约一半的时候实在是忍不了了就跑去看种酱的2016年黄金拼图的演出去了，等我回过神来就过去很久的时间了。\r\nBGM：Goose house - ごはんを食べよう",
        "state": 0,
        "attribute": 16512,
        "duration": 194,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 47163308,
          "name": "诠释泪",
          "face": "http://i0.hdslb.com/bfs/face/d21786bbd84d85281c19aec1fdef86603fb5ec2e.jpg"
        },
        "stat": {
          "aid": 19923367,
          "view": 64988,
          "danmaku": 219,
          "reply": 406,
          "favorite": 1383,
          "coin": 874,
          "share": 187,
          "now_rank": 0,
          "his_rank": 0,
          "like": 390,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 32495482,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1536415970,
        "play_num": "",
        "highlight_title": "【松冈祯丞/日高里菜】里菜酱请再多陪我一下"
      },
      {
        "aid": 29593758,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/af01dff267a271e7fc1619c47196c7007cacc722.jpg",
        "title": "【松冈祯丞/种田奈央】想与你在一起",
        "pubdate": 1534491947,
        "ctime": 1534491945,
        "desc": "松冈祯丞、种田梨沙、东山奈央以及声优们是脑洞故事。其实我是在等20号的故事。\n这一次的视频的灵感来源自一次东山奈央的访谈记录，当时闹被问道和种酱初次见面的情况，\n东山：初次见面的印象是个大美人了，很想和她对视说话，结果坐在我旁边时，我连看她眼睛都不敢看，后悔了很久。\n\n音乐现场的视频推迟到周末，或下周，因为要从新修改内容。\nBGM:Tiara-be with you",
        "state": 0,
        "attribute": 16512,
        "duration": 193,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 1,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 47163308,
          "name": "诠释泪",
          "face": "http://i1.hdslb.com/bfs/face/d21786bbd84d85281c19aec1fdef86603fb5ec2e.jpg"
        },
        "stat": {
          "aid": 29593758,
          "view": 17818,
          "danmaku": 107,
          "reply": 176,
          "favorite": 485,
          "coin": 371,
          "share": 81,
          "now_rank": 0,
          "his_rank": 0,
          "like": 254,
          "dislike": 0
        },
        "dynamic": "#声优##松冈祯丞##种田梨沙##东山奈央#",
        "cid": 51467520,
        "dimension": {
          "width": 1920,
          "height": 1080,
          "rotate": 0
        },
        "fav_at": 1536415649,
        "play_num": "",
        "highlight_title": "【松冈祯丞/种田奈央】想与你在一起"
      },
      {
        "aid": 22389317,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i2.hdslb.com/bfs/archive/8b10bcacd169075193e40fe24b246f79435f5ad6.jpg",
        "title": "「dear...」tkmn与唯一神 Lovesong",
        "pubdate": 1524318441,
        "ctime": 1524318442,
        "desc": "元气温柔少女takamina与紧张呆愣唯一神的故事\n\n素材：av17553388 av21340055 av5356619 av5567408\nav2701207 av16231132 av2232699 av3117710 av17212600\nav18127018 av19619228 av11745639 av11745640 av17532717\n\n借鉴了下娘化大佬佐倾视频的谐击人心的镜头，大佬实在是太强了\n\n视频仅供娱乐，不喜勿看",
        "state": 0,
        "attribute": 16384,
        "duration": 281,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 36626339,
          "name": "黯空雨",
          "face": "http://i0.hdslb.com/bfs/face/3817f4e1fa9e69be65290775d80aa9447ff9ed51.jpg"
        },
        "stat": {
          "aid": 22389317,
          "view": 2262,
          "danmaku": 25,
          "reply": 55,
          "favorite": 87,
          "coin": 140,
          "share": 17,
          "now_rank": 0,
          "his_rank": 0,
          "like": 54,
          "dislike": 0
        },
        "dynamic": "#高桥未奈美#松冈祯丞#食戟之灵#松冈食堂#山田妖精#谐教现场#",
        "cid": 37826210,
        "dimension": {
          "width": 1280,
          "height": 720,
          "rotate": 0
        },
        "fav_at": 1525667097,
        "play_num": "",
        "highlight_title": "「dear...」tkmn与唯一神 Lovesong"
      },
      {
        "aid": 11011123,
        "videos": 2,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i0.hdslb.com/bfs/archive/d6ebf64e862f7defface240750bf0e6d0e3cee7a.jpg",
        "title": "【文艺向】佐仓——我和松冈的关系不好",
        "pubdate": 1496380411,
        "ctime": 1497381337,
        "desc": "素材来源：BGM：《我的谎言》，《你能忘记吗》，《For You ~月の光が降り注ぐテラス》\n内心独白：《四月是你的谎言》\n真实对白：av10677810，av5265835，av2647525，av1307495，av3551212，av4269399，av9702280，av10424582，av5758845，av2898452\n很抱歉，由于做该视频准备的时候有些乱，在这里没办法给这些av号做标记，只能一股脑的列上来了。",
        "state": 0,
        "attribute": 49152,
        "duration": 530,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 10629809,
          "name": "娘化桑",
          "face": "http://i0.hdslb.com/bfs/face/1a71e23b0185c7e48000eba9765bc308e564bc1a.jpg"
        },
        "stat": {
          "aid": 11011123,
          "view": 173181,
          "danmaku": 2298,
          "reply": 2311,
          "favorite": 8458,
          "coin": 13625,
          "share": 2092,
          "now_rank": 0,
          "his_rank": 0,
          "like": 2308,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 18702403,
        "dimension": {
          "width": 0,
          "height": 0,
          "rotate": 0
        },
        "fav_at": 1525140361,
        "play_num": "",
        "highlight_title": "【文艺向】佐仓——我和松冈的关系不好"
      },
      {
        "aid": 22049164,
        "videos": 1,
        "tid": 152,
        "tname": "官方延伸",
        "copyright": 1,
        "pic": "http://i2.hdslb.com/bfs/archive/6465e9fd7677c3b548b349f784c8ca75395f795a.jpg",
        "title": "tkmn与唯一神 一切的开始",
        "pubdate": 1523637311,
        "ctime": 1523637312,
        "desc": "元气温柔少女takamina与紧张呆愣唯一神的故事\n\n素材：av2232699 av3117710 av2663294 av17553388",
        "state": 0,
        "attribute": 16384,
        "duration": 290,
        "rights": {
          "bp": 0,
          "elec": 0,
          "download": 0,
          "movie": 0,
          "pay": 0,
          "hd5": 0,
          "no_reprint": 0,
          "autoplay": 1,
          "ugc_pay": 0,
          "is_cooperation": 0
        },
        "owner": {
          "mid": 36626339,
          "name": "黯空雨",
          "face": "http://i0.hdslb.com/bfs/face/3817f4e1fa9e69be65290775d80aa9447ff9ed51.jpg"
        },
        "stat": {
          "aid": 22049164,
          "view": 5820,
          "danmaku": 36,
          "reply": 83,
          "favorite": 131,
          "coin": 142,
          "share": 15,
          "now_rank": 0,
          "his_rank": 0,
          "like": 64,
          "dislike": 0
        },
        "dynamic": "",
        "cid": 37811164,
        "dimension": {
          "width": 1280,
          "height": 720,
          "rotate": 0
        },
        "fav_at": 1525090164,
        "play_num": "",
        "highlight_title": "tkmn与唯一神 一切的开始"
      }
    ]
  }
}
```