package auth

import (
	"jrpg-gang/util"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthenticatorConfig struct {
	RequestTimeoutSec    int64  `json:"requestTimeoutSec"`
	StateCacheTimeoutMin int64  `json:"stateCacheTimeoutSec"`
	GoogleClientId       string `json:"googleClientId"`
	GoogleClientSecret   string `json:"googleClientSecret"`
	GoogleRedirectUrl    string `json:"googleRedirectUrl"`
}

type UserCredentials struct {
	Picture string
	Email   string
}

type PlayerToken string

const (
	PlayerTokenEmpty PlayerToken = ""
)

type AuthenticationHandler interface {
	HandleUserAuthenticated(credentials UserCredentials) (PlayerToken, bool)
}
type Authenticator struct {
	rndGen     *util.RndGen
	config     AuthenticatorConfig
	googleSso  oauth2.Config
	stateCache *ttlcache.Cache[string, bool]
	handler    AuthenticationHandler
}

func NewAuthenticator(config AuthenticatorConfig, handler AuthenticationHandler) *Authenticator {
	auth := &Authenticator{}
	auth.rndGen = util.NewRndGen()
	auth.stateCache = ttlcache.New(
		ttlcache.WithTTL[string, bool](time.Duration(config.StateCacheTimeoutMin) * time.Minute),
	)
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
	auth.handler = handler
	go auth.stateCache.Start()
	return auth
}
