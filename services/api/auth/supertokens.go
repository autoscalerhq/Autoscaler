package auth

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
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

func VerifySessionMiddleware(theirHandler http.Handler) http.Handler {
	return session.VerifySession(nil, func(writer http.ResponseWriter, request *http.Request) {
		theirHandler.ServeHTTP(writer, request)
	})
}

func MakeDefaultSuperTokensAppInfoEnv() SuperTokensEnv {
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
