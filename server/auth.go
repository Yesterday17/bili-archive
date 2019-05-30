package server

import (
	"github.com/Yesterday17/bili-archive/bilibili"
	ecode "github.com/Yesterday17/bilibili-errorcode"
	"github.com/emicklei/go-restful"
	"strconv"
)

type Auth struct {
	qrCode  bilibili.QRCode
	login   bool
	cookies string
}

// 验证 cookies 是否可用
func (a *Auth) updateLoginStatus() {
	if a.cookies != "" {
		_, err := bilibili.GetUserMID(a.cookies)
		if err != nil {
			a.cookies = ""
			a.login = false
		} else {
			a.login = true
		}
	} else {
		a.login = false
	}
}

func (a *Auth) getLoginStatus(request *restful.Request, response *restful.Response) {
	if a.cookies == "" {
		// 快速跳过空 cookies 的情况
		response.WriteHeader(401)
		_ = response.WriteAsJson(ecode.ErrorCode(-101).GetDetail())
	} else {
		mid, err := bilibili.GetUserMID(a.cookies)
		if err != nil {
			a.login = false
			code, _ := strconv.Atoi(string(err.Error()))
			response.WriteHeader(401)
			_ = response.WriteAsJson(ecode.ErrorCode(code).GetDetail())
		} else {
			a.login = true
			_ = response.WriteAsJson(map[string]interface{}{
				"code": 0,
				"mid":  mid,
			})
		}
	}
}

func (a *Auth) getLoginQRCode(request *restful.Request, response *restful.Response) {
	a.updateLoginStatus()

	if a.cookies == "" && a.qrCode.Image == "" {
		// 用户未登录
		a.qrCode = bilibili.GetLoginQRCode()

		_ = response.WriteAsJson(map[string]interface{}{
			"code":  0,
			"login": false,
			"image": a.qrCode.Image,
		})
	} else {
		// 用户已经登录
		_ = response.WriteAsJson(map[string]interface{}{
			"code":  0,
			"login": true,
		})
	}
}

func (a *Auth) checkLoginStatus(request *restful.Request, response *restful.Response) {
	a.updateLoginStatus()
	if a.login {
		_ = response.WriteAsJson(map[string]interface{}{
			"code":    0,
			"cookies": a.cookies,
		})
	} else if a.qrCode.Image == "" {
		_ = response.WriteAsJson(ecode.ErrorCode(-101).GetDetail())
	} else {
		success, redirect, err := a.qrCode.Check()
		if err != nil {
			// -500 服务器错误
			response.WriteHeader(500)
			_ = response.WriteAsJson(ecode.ErrorCode(-500).GetDetail())
		} else if !success {
			// -101 账号未登录
			response.WriteHeader(401)
			_ = response.WriteAsJson(ecode.ErrorCode(-101).GetDetail())
		} else {
			a.cookies = bilibili.GetCookiesString(redirect)
			_ = response.WriteAsJson(map[string]interface{}{
				"code":    0,
				"cookies": a.cookies,
			})
		}
	}
}

func (a *Auth) logoutAuth(request *restful.Request, response *restful.Response) {
	a.login = false
	a.cookies = ""
	a.qrCode.Image = ""
	response.WriteHeader(204)
}

func (a *Auth) WebService() *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/auth").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(a.getLoginStatus).
		Doc("获得用户登录状态"))

	ws.Route(ws.GET("/login").To(a.getLoginQRCode).
		Doc("获得用以用户登录的二维码"))
	ws.Route(ws.DELETE("/login").To(a.logoutAuth).
		Doc("注销当前用户"))

	ws.Route(ws.GET("/check").To(a.checkLoginStatus).
		Doc("检查当前登录二维码的登录情况"))

	return ws
}
