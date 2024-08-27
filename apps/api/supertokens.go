package main

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
	"strings"
)

type SuperTokensEnv struct {
	// connectionUri is the URL of the SuperTokens core instance
	connectionUri string
	// api token to communicate with SuperTokens core instance
	apiKey  string
	appInfo SuperTokensAppInfoEnv
}
type SuperTokensAppInfoEnv struct {
	// appName is the name of the app
	apiDomain     string
	websiteDomain string
}

func makeDefaultSuperTokensAppInfoEnv() SuperTokensEnv {
	return SuperTokensEnv{
		connectionUri: "http://localhost:3567",
		apiKey:        "",
		appInfo: SuperTokensAppInfoEnv{
			apiDomain:     "http://localhost:8080",
			websiteDomain: "http://localhost:3000",
		},
	}
}

func InitSuperTokens(env SuperTokensEnv) error {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: env.connectionUri,
			APIKey:        env.apiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "autoscaler",
			APIDomain:       env.appInfo.apiDomain,
			WebsiteDomain:   env.appInfo.websiteDomain,
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func CorsMiddleware(next http.Handler) http.Handler {
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
