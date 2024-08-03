package main

import (
	"fmt"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
	"strings"
)

func main() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: "http://localhost:3567",
			// APIKey: <API_KEY(if configured)>,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "autoscaler",
			APIDomain:       "http://localhost:8080",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
	fmt.Print("Server running on port 4000")
	//session.VerifySession(nil, likeCommentAPI)
	//http.HandleFunc("/comment", likeCommentAPI)

	err = http.ListenAndServe("localhost:4000", corsMiddleware(
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			// we add content-type + other headers used by SuperTokens
			response.Header().Set("Access-Control-Allow-Headers",
				strings.Join(append([]string{"Content-Type"},
					supertokens.GetAllCORSHeaders()...), ","))
			response.Header().Set("Access-Control-Allow-Methods", "*")
			response.Write([]byte(""))
		} else {
			next.ServeHTTP(response, r)
		}
	})
}
