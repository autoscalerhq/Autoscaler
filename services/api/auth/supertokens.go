package auth

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
)

type SuperTokensConfiguration struct {
	// connectionUri is the URL of the SuperTokens core instance
	ConnectionUri string
	// api token to communicate with SuperTokens core instance
	ApiKey  string
	AppInfo SuperTokensAppInfoEnv
}
type SuperTokensAppInfoEnv struct {
	// appName is the name of the app
	ApiDomain     string
	WebsiteDomain string
}

func VerifySessionMiddleware(theirHandler http.Handler) http.Handler {
	return session.VerifySession(nil, func(writer http.ResponseWriter, request *http.Request) {
		theirHandler.ServeHTTP(writer, request)
	})
}

func MakeDefaultSuperTokensAppInfoEnv() SuperTokensConfiguration {
	return SuperTokensConfiguration{
		ConnectionUri: "http://localhost:3567",
		ApiKey:        "",
		AppInfo: SuperTokensAppInfoEnv{
			ApiDomain:     "http://localhost:8080",
			WebsiteDomain: "http://localhost:3000",
		},
	}
}

func InitSuperTokens(env SuperTokensConfiguration) error {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: env.ConnectionUri,
			APIKey:        env.ApiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "autoscaler",
			APIDomain:       env.AppInfo.ApiDomain,
			WebsiteDomain:   env.AppInfo.WebsiteDomain,
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
