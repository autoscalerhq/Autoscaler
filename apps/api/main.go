package main

import (
	"fmt"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
)

type Environment struct {
	supertokens   SuperTokensEnv
	listenAddress string
}

func makeDefaultEnv() Environment {
	return Environment{
		supertokens:   makeDefaultSuperTokensAppInfoEnv(),
		listenAddress: "localhost:4000",
	}
}

func main() {
	env := makeDefaultEnv()
	err := InitSuperTokens(env.supertokens)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("Server running on port 4000")
	err = http.ListenAndServe(env.listenAddress, CorsMiddleware(
		supertokens.Middleware(RouterMux())))
	if err != nil {
		panic(err.Error())
	}

}
func RouterMux() http.Handler {
	userMux := http.NewServeMux()
	userMux.Handle("/comment", session.VerifySession(nil, likeCommentAPI))
	return userMux
}

func likeCommentAPI(w http.ResponseWriter, r *http.Request) {
	// retrieve the session object as shown below
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	userID := sessionContainer.GetUserID()

	_, err := w.Write([]byte("Hello " + userID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//fmt.Println(userID)
}
