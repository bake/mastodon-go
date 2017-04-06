package mastodon

import (
	"fmt"

	"golang.org/x/oauth2"
)

// AuthCodeURL builds a URL to obtain an AccessCode.
func (app App) AuthCodeURL() string {
	return app.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// Exchange swaps an AccessCode with an AccessToken wich can be used to authenticate an user.
func (app App) Exchange(code string) (*oauth2.Token, error) {
	token, err := app.Config.Exchange(nil, code)
	if err != nil {
		return nil, fmt.Errorf("could not exchange access token: %v", err)
	}
	return token, nil
}

// SetToken saves the AccessToken in struct.
func (app *App) SetToken(token string) {
	app.Token = &oauth2.Token{
		AccessToken: token,
	}
}
