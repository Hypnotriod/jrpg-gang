package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"jrpg-gang/util"
	"net/http"
	"time"

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
	url := a.googleSso.AuthCodeURL(a.salt32)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *Authenticator) HandleGoogleAuth2Callback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != a.salt32 {
		log.Info("Google Oauth: state doesn't match")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := a.googleAuth2AcquireToken(r)
	if err != nil {
		log.Info("Google Oauth: couldn't acquire token: ", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	userInfo, err := a.googleAuth2AcquireUserInfo(token)
	if err != nil {
		log.Info("Google Oauth: couldn't acquire user info: ", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprint(w, util.ObjectToJson(userInfo))
}

func (a *Authenticator) googleAuth2AcquireToken(r *http.Request) (*oauth2.Token, error) {
	code := r.FormValue("code")
	cntx, cancel := context.WithTimeout(context.Background(), time.Duration(a.config.HttpRequestTimeoutSec)*time.Second)
	defer cancel()
	return a.googleSso.Exchange(cntx, code)
}

func (a *Authenticator) googleAuth2AcquireUserInfo(token *oauth2.Token) (*GoogleOauth2UserInfo, error) {
	cntx, cancel := context.WithTimeout(context.Background(), time.Duration(a.config.HttpRequestTimeoutSec)*time.Second)
	defer cancel()
	response, err := http.NewRequestWithContext(
		cntx,
		http.MethodGet,
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken,
		nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var userInfo GoogleOauth2UserInfo
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
