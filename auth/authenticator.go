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
	GoogleCallbackUrl    string `json:"googleCallbackUrl"`
	RedirectUrl          string `json:"redirectUrl"`
}

type UserCredentials struct {
	Picture string
	Email   string
}

type AuthenticationToken string

const (
	PlayerTokenEmpty AuthenticationToken = ""
)

type AuthenticationStatus struct {
	IsAuthenticated bool                `json:"isAuthenticated"`
	Token           AuthenticationToken `json:"token,omitempty"`
	IsNewPlayer     bool                `json:"isNewPlayer,omitempty"`
}

type AuthenticationHandler interface {
	HandleUserAuthenticated(credentials UserCredentials) AuthenticationStatus
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
		RedirectURL:  config.GoogleCallbackUrl,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	auth.handler = handler
	go auth.stateCache.Start()
	return auth
}
