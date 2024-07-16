package main

import (
	"giants/pkg/apiserver"
	"giants/pkg/detail"
	ucUser "giants/pkg/usecase/user"
)

func main() {
	us := detail.NewUserStore()
	fu := ucUser.NewFindUser(us)
	ru := ucUser.NewRegisterUser(us)

	apiserver.RunApiServer(apiserver.Adapters{
		FindUser:     fu,
		RegisterUser: ru,
	})
}
