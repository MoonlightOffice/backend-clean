package apiserver

import (
	"fmt"
	pUser "giants/pkg/apiserver/user"
	ucUser "giants/pkg/usecase/user"
	"log"
	"net/http"
)

type Adapters struct {
	FindUser     *ucUser.FindUser
	RegisterUser *ucUser.RegisterUser
}

func RunApiServer(adapters Adapters) {
	router := middleware(router(adapters))

	fmt.Println("Listening on 0.0.0.0:8000")
	log.Fatal(http.ListenAndServe(":8000", router).Error())
}

func router(adapters Adapters) http.Handler {
	router := http.NewServeMux()

	userFindHandler := pUser.NewUserFindHandler(*adapters.FindUser)
	router.HandleFunc("/v1/user/find-by-id", withAuth(userFindHandler.UserFindById))

	userRegisterHandler := pUser.NewUserRegisterHandler(*adapters.RegisterUser)
	router.HandleFunc("/v1/user/register", userRegisterHandler.UserRegister)

	return router
}

func middleware(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS preflight access
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Session")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		router.ServeHTTP(w, r)
	})
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle auth here
		// ...
		next(w, r)
	})
}
