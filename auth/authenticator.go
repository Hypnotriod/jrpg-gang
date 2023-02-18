package auth

import (
	"jrpg-gang/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthenticatorConfig struct {
	HttpRequestTimeoutSec int64  `json:"httpRequestTimeoutSec"`
	GoogleClientId        string `json:"googleClientId"`
	GoogleClientSecret    string `json:"googleClientSecret"`
	GoogleRedirectUrl     string `json:"googleRedirectUrl"`
}

type Authenticator struct {
	rndGen    *util.RndGen
	config    AuthenticatorConfig
	googleSso oauth2.Config
	salt32    string
}

func NewAuthenticator(config AuthenticatorConfig) *Authenticator {
	auth := &Authenticator{}
	auth.rndGen = util.NewRndGen()
	auth.salt32 = auth.rndGen.MakeId32()
	auth.config = config
	auth.googleSso = oauth2.Config{
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectUrl,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	return auth
}
