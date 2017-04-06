package mastodon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	base = "https://mastodon.social/api/v1/"
)

// App holds the AccessToken and the OAuth2 config
type App struct {
	Token  *oauth2.Token
	Config *oauth2.Config
}

// NewApp tries to register a new app
func NewApp(name, uris string, scopes []string, website string) (*App, error) {
	res, err := http.PostForm(base+"apps", url.Values{
		"client_name":   {name},
		"redirect_uris": {uris},
		"scopes":        {strings.Join(scopes, " ")},
		"website":       {website},
	})
	if err != nil {
		return nil, fmt.Errorf("cound not registering new app: %v", err)
	}
	defer res.Body.Close()

	app := struct {
		ID           uint   `json:"id"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&app); err != nil {
		return nil, fmt.Errorf("could not decode response: %v", err)
	}

	return &App{
		Config: &oauth2.Config{
			ClientID:     app.ClientID,
			ClientSecret: app.ClientSecret,
			RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://mastodon.social/oauth/authorize",
				TokenURL: "https://mastodon.social/oauth/token",
			},
		},
	}, nil
}
