package server

import (
	ecode "github.com/Yesterday17/bilibili-errorcode"
	"github.com/emicklei/go-restful"
	"strconv"
)

func WriteJson(response *restful.Response, status int, value interface{}) {
	_ = response.WriteHeaderAndJson(status, value, restful.MIME_JSON)
}

// -101 账号未登录
func UserNotLogin(response *restful.Response) {
	WriteJson(response, 401, ecode.ErrorCode(-101).GetDetail())
}

func ErrorResponse(err error, response *restful.Response) {
	code, _ := strconv.Atoi(string(err.Error()))
	WriteJson(response, 401, ecode.ErrorCode(code).GetDetail())
}

// -500 服务器错误
func BackendServerError(response *restful.Response) {
	WriteJson(response, 500, ecode.ErrorCode(-500).GetDetail())
}
