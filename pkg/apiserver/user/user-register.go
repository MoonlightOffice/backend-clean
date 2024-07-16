package user

import (
	"fmt"
	"net/http"

	tool "giants/pkg/apiserver/tool"
	"giants/pkg/usecase/user"
)

type userRegisterHandler struct {
	registerUser user.RegisterUser
}

func NewUserRegisterHandler(ru user.RegisterUser) userRegisterHandler {
	return userRegisterHandler{registerUser: ru}
}

func (h userRegisterHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Email string `json:"email"`
	}{}
	if !tool.BindReqData(r, &reqData) {
		tool.WriteResponse(w, 400, nil)
		return
	}

	fmt.Println("Here 1")
	err := h.registerUser.Run(user.RegisterUserInputDto{Email: reqData.Email})
	if err != nil {
		tool.WriteResponse(w, 400, tool.H{"log": err.Error()})
		return
	}
	fmt.Println("Here 2")

	tool.WriteResponse(w, 200, nil)
}
