package mastodon

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// App holds the AccessToken and the OAuth2 config.
type App struct {
	Token          *oauth2.Token
	Config         *oauth2.Config
	API            *API
	Accounts       *Accounts
	Blocks         *Blocks
	Favourites     *Favourites
	FollowRequests *FollowRequests
	Follows        *Follows
	Instances      *Instances
	Mutes          *Mutes
	Notifications  *Notifications
	Reports        *Reports
	Search         *Search
	Statuses       *Statuses
	Timelines      *Timelines
}

// NewApp tries to register a new app.
func NewApp(base, name, uris string, scopes []string, website string) (*App, error) {
	api := API{
		Base:   base,
		Prefix: "/api/v1/",
	}

	v := url.Values{
		"client_name":   {name},
		"redirect_uris": {uris},
		"scopes":        {strings.Join(scopes, " ")},
		"website":       {website},
	}
	app := Application{}
	if err := api.Post("apps", v, &app); err != nil {
		return nil, err
	}

	if uris == "" {
		uris = "urn:ietf:wg:oauth:2.0:oob"
	}

	return &App{
		Config: &oauth2.Config{
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
			RedirectURL:  uris,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  api.Base + "/oauth/authorize",
				TokenURL: api.Base + "/oauth/token",
			},
		},
		API:            &api,
		Accounts:       &Accounts{&api},
		Blocks:         &Blocks{&api},
		Favourites:     &Favourites{&api},
		FollowRequests: &FollowRequests{&api},
		Follows:        &Follows{&api},
		Instances:      &Instances{&api},
		Mutes:          &Mutes{&api},
		Notifications:  &Notifications{&api},
		Reports:        &Reports{&api},
		Search:         &Search{&api},
		Statuses:       &Statuses{&api},
		Timelines:      &Timelines{&api},
	}, nil
}

// AuthCodeURL builds a URL to obtain an AccessCode.
func (app App) AuthCodeURL() string {
	return app.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// Exchange swaps an AccessCode with an AccessToken which can be used to
// authenticate an user.
func (app App) Exchange(code string) (string, error) {
	token, err := app.Config.Exchange(nil, code)
	if err != nil {
		return "", fmt.Errorf("could not exchange access token: %v", err)
	}
	return token.AccessToken, nil
}

// SetToken saves the AccessToken in struct.
func (app App) SetToken(token string) {
	app.API.AccessToken = token
}
