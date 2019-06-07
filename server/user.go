package server

import (
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/emicklei/go-restful"
)

type UserInfo struct {
	mid string
}

func (u *UserInfo) getCurrentUID(request *restful.Request, response *restful.Response) {
	if auth.cookies == "" {
		u.mid = ""
		UserNotLogin(response)
		return
	}

	// this shouldn't happen
	if u.mid == "" {
		mid, err := bilibili.GetUserMID(auth.cookies)
		if err != nil {
			u.mid = ""
			ErrorResponse(err, response)
			return
		}
		u.mid = mid
	}

	_ = response.WriteAsJson(map[string]interface{}{
		"code": 0,
		"mid":  u.mid,
	})

}

func (u *UserInfo) getCurrentInfo(request *restful.Request, response *restful.Response) {
	if auth.cookies == "" {
		u.mid = ""
		UserNotLogin(response)
		return
	}

	// this shouldn't happen
	if u.mid == "" {
		mid, err := bilibili.GetUserMID(auth.cookies)
		if err != nil {
			ErrorResponse(err, response)
			return
		}
		u.mid = mid
	}

	info, err := bilibili.GetMIDInfo(u.mid)
	if err != nil {
		ErrorResponse(err, response)
		return
	}
	_ = response.WriteAsJson(info)
}

func (u *UserInfo) WebService() *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/user").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/mid").To(u.getCurrentUID).
		Doc("获取当前登录用户的UID"))

	ws.Route(ws.GET("/info").To(u.getCurrentInfo).
		Doc("获取当前登录用户的信息"))
	return ws
}
