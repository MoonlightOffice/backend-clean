package user

import (
	"net/http"

	tool "giants/pkg/apiserver/tool"
	"giants/pkg/usecase/user"
)

type userFindHandler struct {
	findUser user.FindUser
}

func NewUserFindHandler(fu user.FindUser) userFindHandler {
	return userFindHandler{findUser: fu}
}

func (h userFindHandler) UserFindById(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		UserId string `json:"userId"`
	}{}
	if !tool.BindReqData(r, &reqData) {
		tool.WriteResponse(w, 400, nil)
		return
	}

	uObj, err := h.findUser.ById(reqData.UserId)
	if err != nil {
		tool.WriteResponse(w, 404, tool.H{"log": err.Error()})
		return
	}

	tool.WriteResponse(w, 200, tool.H{
		"userId":    uObj.UserId,
		"email":     uObj.Email,
		"createdAt": uObj.CreatedAt.UnixMilli(),
	})
}

func (h userFindHandler) UserFindByEmail(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Email string `json:"email"`
	}{}
	if !tool.BindReqData(r, &reqData) {
		tool.WriteResponse(w, 400, nil)
		return
	}

	uObj, err := h.findUser.ByEmail(reqData.Email)
	if err != nil {
		tool.WriteResponse(w, 404, tool.H{"log": err.Error()})
		return
	}

	tool.WriteResponse(w, 200, tool.H{
		"userId":    uObj.UserId,
		"email":     uObj.Email,
		"createdAt": uObj.CreatedAt.UnixMilli(),
	})
}
