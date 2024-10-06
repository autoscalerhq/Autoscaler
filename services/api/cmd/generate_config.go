package main

import (
	"github.com/autoscalerhq/docuconf/gen"
	"path/filepath"
)

func main() {
	err := gen.NewService("api", "configuration", filepath.Join("configuration"), gen.AdditionalOptions{}).
		AddString("ConnectionString", "Connection string to the database", false, "").
		AddString("ListenAddress", "The address the go backend to bind to, required in production, default value works locally", false, ":8888").
		AddString("SuperTokensConnectionUri", "URL of the SuperTokens core instance", false, "http://localhost:3567").
		AddString("SuperTokensApiKey", "The super tokens api key, required for production but not local development", false, "").
		AddString("SuperTokensAppApiDomain", `The address of the go backend, required for production, default value works locally`, false, "http://localhost:8080").
		AddString("SuperTokensAppWebsiteDomain", "The address of the react frontend", false, "http://localhost:3000").
		Write()
	if err != nil {
		println(err.Error())
	}
}
