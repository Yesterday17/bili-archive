package server

import "github.com/emicklei/go-restful"

type UserInfo struct {
	users map[string]string
}

func (u *UserInfo) getCurrentUID(request *restful.Request, response *restful.Response) {

}

func (u *UserInfo) WebService() *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/user").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(u.getCurrentUID).
		Doc("get uid of current user"))
	return ws
}
