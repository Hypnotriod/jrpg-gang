package auth

import (
	"jrpg-gang/util"
	"net"
	"sync/atomic"
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
	Picture  string
	Email    string
	Nickname string
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
	HandleGuestUserAuthenticated(credentials UserCredentials) AuthenticationStatus
}

type Authenticator struct {
	rndGen       *util.RndGen
	config       AuthenticatorConfig
	googleSso    oauth2.Config
	stateCache   *ttlcache.Cache[string, net.IP]
	handler      AuthenticationHandler
	guestCounter atomic.Uint64
}

func NewAuthenticator(config AuthenticatorConfig) *Authenticator {
	auth := &Authenticator{}
	auth.rndGen = util.NewRndGen()
	auth.stateCache = ttlcache.New(
		ttlcache.WithTTL[string, net.IP](time.Duration(config.StateCacheTimeoutMin) * time.Minute),
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
	go auth.stateCache.Start()
	return auth
}

func (a *Authenticator) SetHandler(handler AuthenticationHandler) {
	a.handler = handler
}
