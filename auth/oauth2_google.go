package auth

import (
	"context"
	"jrpg-gang/util"
	"net/http"
	"strconv"
	"time"

	"github.com/jellydator/ttlcache/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type GoogleOauth2UserInfo struct {
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

func (a *Authenticator) HandleGoogleAuth2(w http.ResponseWriter, r *http.Request) {
	state := a.rndGen.MakeUUID()
	ip := util.GetIP(r.RemoteAddr)
	a.stateCache.Set(state, ip, ttlcache.DefaultTTL)
	url := a.googleSso.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *Authenticator) HandleGoogleAuth2Callback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	item := a.stateCache.Get(state)
	a.stateCache.Delete(state)
	if item == nil || item.IsExpired() {
		log.Info("Google Oauth: state is expired or doesn't exist")
		http.Redirect(w, r, a.config.RedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	token, err := a.googleAuth2AcquireToken(r)
	if err != nil {
		log.Info("Google Oauth: couldn't acquire token: ", err.Error())
		http.Redirect(w, r, a.config.RedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	userInfo, err := a.googleAuth2AcquireUserInfo(r, token)
	if err != nil {
		log.Info("Google Oauth: couldn't acquire user info: ", err.Error())
		http.Redirect(w, r, a.config.RedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	credentials := UserCredentials{Email: userInfo.Email, Picture: userInfo.Picture, Ip: item.Value()}
	status := a.handler.HandleUserAuthenticated(credentials)
	if !status.IsAuthenticated {
		log.Info("Google Oauth: authentication rejected for: ", credentials.Email)
		http.Redirect(w, r, a.config.RedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	url := a.config.RedirectUrl + "/?token=" + string(status.Token) + "&isNewPlayer=" + strconv.FormatBool(status.IsNewPlayer)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *Authenticator) googleAuth2AcquireToken(r *http.Request) (*oauth2.Token, error) {
	code := r.FormValue("code")
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(a.config.RequestTimeoutSec)*time.Second)
	defer cancel()
	return a.googleSso.Exchange(ctx, code)
}

func (a *Authenticator) googleAuth2AcquireUserInfo(r *http.Request, token *oauth2.Token) (*GoogleOauth2UserInfo, error) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(a.config.RequestTimeoutSec)*time.Second)
	defer cancel()
	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken
	userInfo, err := util.HttpGetJson[GoogleOauth2UserInfo](ctx, url)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
