package main

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

type Google struct {
	Config *oauth2.Config
}

func NewGoogle() *Google {
	return newGoogle()
}

func newGoogle() *Google {

	google := &Google{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GoogleClientID"),
			ClientSecret: os.Getenv("googleClientSecret"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  os.Getenv("AuthorizeEndpoint"),
				TokenURL: os.Getenv("TokenEndpoint"),
			},
			Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email"},
			RedirectURL: "http://localhost:8080/google/callback",
		},
	}

	return google
}

func (g *Google) GetLoginURL(state string) (clientID string) {
	return g.Config.AuthCodeURL(state)
}

func (g *Google) GetUserID(code string) (googleUserID string, err error) {

	cxt := context.Background()

	//TODO: stateの検証を行う
	httpClient, error := g.Config.Exchange(cxt, code)
	if httpClient == nil {
		return "", errors.New("接続エラー. httpClientがnilです" + error.Error())
	}

	client := g.Config.Client(cxt, httpClient)

	service, err := v2.New(client)
	if err != nil {
		return "", errors.New("接続エラー" + err.Error())
	}

	userInfo, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(cxt).Do()
	if err != nil {
		return "", errors.New("接続エラー" + err.Error())
	}

	return userInfo.UserId, nil
}
