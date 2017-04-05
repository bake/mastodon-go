package mastodon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	base = "https://mastodon.social/api/v1/"
)

type App struct {
	Token  *oauth2.Token
	Config *oauth2.Config
}

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

func (app *App) AuthCodeURL() string {
	return app.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (app *App) Exchange(code string) error {
	token, err := app.Config.Exchange(nil, code)
	if err != nil {
		return fmt.Errorf("could not exchange access token: %v", err)
	}
	app.Token = token
	return nil
}

func (app *App) SetToken(token string) {
	app.Token = &oauth2.Token{
		AccessToken: token,
	}
}

func (app *App) Do(method string, endpoint string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, base+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request to %s: %v", endpoint, err)
	}
	req.Header.Set("Authorization", "Bearer "+app.Token.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute %s: %v", endpoint, err)
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.Body, nil
	default:
		return nil, fmt.Errorf("could not get %s: %s: %v", endpoint, res.Status, app.getError(res.Body))
	}
}

func (app *App) Get(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodGet, endpoint)
}

func (app *App) Post(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodPost, endpoint)
}

func (app *App) getError(r io.ReadCloser) error {
	defer r.Close()
	res := &struct {
		Error string `json:"error"`
	}{}
	if err := json.NewDecoder(r).Decode(res); err != nil {
		return fmt.Errorf("could not decode error: %v", err)
	}
	return errors.New(res.Error)
}

func (app *App) VerifyCredentials() (*Account, error) {
	r, err := app.Get("accounts/verify_credentials")
	if err != nil {
		return nil, fmt.Errorf("could not verify credentials: %v", err)
	}
	defer r.Close()

	acc := &Account{}
	if err := json.NewDecoder(r).Decode(acc); err != nil {
		return nil, fmt.Errorf("could not decode credentials: %v", err)
	}

	return acc, nil
}

func (app *App) GetAccount(id int) (*Account, error) {
	r, err := app.Get(fmt.Sprintf("accounts/%d", id))
	if err != nil {
		return nil, fmt.Errorf("could not get account %d: %v", id, err)
	}
	defer r.Close()

	acc := &Account{}
	if err := json.NewDecoder(r).Decode(acc); err != nil {
		return nil, fmt.Errorf("could not decode account %d: %v", id, err)
	}

	return acc, nil
}

func (app *App) GetFollowers(id int) ([]Account, error) {
	r, err := app.Get(fmt.Sprintf("accounts/%d/followers", id))
	if err != nil {
		return nil, fmt.Errorf("could not get followers %d: %v", id, err)
	}
	defer r.Close()

	accs := []Account{}
	if err := json.NewDecoder(r).Decode(&accs); err != nil {
		return nil, fmt.Errorf("could not decode followers %d: %v", id, err)
	}

	return accs, nil
}

func (app *App) GetFollowing(id int) ([]Account, error) {
	r, err := app.Get(fmt.Sprintf("accounts/%d/following", id))
	if err != nil {
		return nil, fmt.Errorf("could not get following %d: %v", id, err)
	}
	defer r.Close()

	accs := []Account{}
	if err := json.NewDecoder(r).Decode(&accs); err != nil {
		return nil, fmt.Errorf("could not decode following %d: %v", id, err)
	}

	return accs, nil
}

func (app *App) Follow(id int) (*Account, error) {
	r, err := app.Post(fmt.Sprintf("accounts/%d/follow", id))
	if err != nil {
		return nil, fmt.Errorf("could not follow %d: %v", id, err)
	}
	defer r.Close()

	acc := &Account{}
	if err := json.NewDecoder(r).Decode(acc); err != nil {
		return nil, fmt.Errorf("could not follow %d: %v", id, err)
	}
	return acc, nil
}

func (app *App) Unfollow(id int) (*Account, error) {
	r, err := app.Post(fmt.Sprintf("accounts/%d/unfollow", id))
	if err != nil {
		return nil, fmt.Errorf("could not follow %d: %v", id, err)
	}
	defer r.Close()

	acc := &Account{}
	if err := json.NewDecoder(r).Decode(acc); err != nil {
		return nil, fmt.Errorf("could not follow %d: %v", id, err)
	}
	return acc, nil
}

type Account struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Acct        string `json:"acct"`
	DisplayName string `json:"display_name"`
	Note        string `json:"note"`
	URL         string `json:"url"`
	Avatar      string `json:"avatar"`
	Header      string `json:"header"`
	Locked      bool   `json:"locked"`
	CreatedAt   string `json:"created_at"`
	Followers   int    `json:"followers_count"`
	Following   int    `json:"following_count"`
	Statuses    int    `json:"statuses_count"`
}
