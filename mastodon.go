package mastodon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

// Do executes an API request. The method is a HTTP method, e.g. GET or POST.
func (app App) Do(method string, endpoint string, values url.Values) (io.ReadCloser, error) {
	client := &http.Client{}
	r := bytes.NewBufferString(values.Encode())
	req, err := http.NewRequest(method, base+endpoint, r)
	if err != nil {
		return nil, fmt.Errorf("could not create request to %s: %v", endpoint, err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app.Token.AccessToken))

	switch method {
	case http.MethodGet:
		req.URL.RawQuery = values.Encode()
		// case http.MethodGet:
		// TODO: set form vales
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute %s: %v", endpoint, err)
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.Body, nil
	default:
		err := app.getError(res.Body)
		return nil, fmt.Errorf("%s: %v", res.Status, err)
	}
}

// Get request
func (app App) Get(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodGet, endpoint, nil)
}

// Post request
func (app App) Post(endpoint string) (io.ReadCloser, error) {
	return app.Do(http.MethodPost, endpoint, nil)
}

func (app App) generic(method, endpoint string, values url.Values, dest interface{}) error {
	r, err := app.Do(method, endpoint, values)
	if err != nil {
		return fmt.Errorf("could not %s %s: %v", method, endpoint, err)
	}
	defer r.Close()

	if err := json.NewDecoder(r).Decode(dest); err != nil {
		return fmt.Errorf("could not decode %s: %v", endpoint, err)
	}

	return nil
}

func (app App) getError(r io.ReadCloser) error {
	defer r.Close()
	res := Error{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return fmt.Errorf("could not decode error: %v", err)
	}
	return errors.New(res.Error)
}

// VerifyCredentials returns the authenticated user's account.
func (app App) VerifyCredentials() (Account, error) {
	acc := Account{}
	if err := app.generic(http.MethodGet, "accounts/verify_credentials", nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// GetAccount returns an account.
func (app App) GetAccount(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d", id)
	if err := app.generic(http.MethodGet, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// GetFollowers returns an slice of following accounts.
func (app App) GetFollowers(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/followers", id)
	if err := app.generic(http.MethodGet, end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// GetFollowing returns an slice of followed accounts.
func (app *App) GetFollowing(id int) ([]Account, error) {
	accs := []Account{}
	end := fmt.Sprintf("accounts/%d/following", id)
	if err := app.generic(http.MethodGet, end, nil, &accs); err != nil {
		return accs, err
	}
	return accs, nil
}

// GetStatuses returns an slice of statuses. Accepted params are:
// only_media: Only return statuses that have media attachments
// exclude_replies: Skip statuses that reply to other statuses
func (app *App) GetStatuses(id int, params url.Values) ([]Status, error) {
	end := fmt.Sprintf("accounts/%d/statuses", id)
	statuses := []Status{}
	if err := app.generic(http.MethodGet, end, params, &statuses); err != nil {
		return statuses, err
	}
	return statuses, nil
}

// Follow an user.
func (app App) Follow(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/follow", id)
	acc := Account{}
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unfollow an account
func (app App) Unfollow(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unfollow", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Block an account
func (app App) Block(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/block", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unblock an account.
func (app App) Unblock(id int) (Account, error) {
	end := fmt.Sprintf("accounts/%d/unblock", id)
	acc := Account{}
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Mute an account.
func (app App) Mute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/mute", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Unmute an user.
func (app App) Unmute(id int) (Account, error) {
	acc := Account{}
	end := fmt.Sprintf("accounts/%d/unmute", id)
	if err := app.generic(http.MethodPost, end, nil, &acc); err != nil {
		return acc, err
	}
	return acc, nil
}

// Relationships returns an slice of Relationships of the current user to a
// list of given accounts.
func (app App) Relationships(ids ...int) ([]Relationship, error) {
	idss := []string{}
	for _, id := range ids {
		idss = append(idss, strconv.Itoa(id))
	}
	v := url.Values{"id": idss}
	rels := []Relationship{}
	if err := app.generic(http.MethodGet, "accounts/relationships", v, &rels); err != nil {
		return rels, err
	}
	return rels, nil
}

// SearchAccount returns an slice of matching Accounts. Will lookup an account
// remotely if the search term is in the username@domain format and not yet in
// the database.
func (app App) SearchAccount(q string, limit int) ([]Account, error) {
	v := url.Values{
		"q":     {q},
		"limit": {strconv.Itoa(limit)},
	}
	accs := []Account{}
	if err := app.generic(http.MethodGet, "accounts/search", v, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}

// GetBlocks returns an slice of accounts blocked by the authenticated user.
func (app App) GetBlocks() ([]Account, error) {
	accs := []Account{}
	if err := app.generic(http.MethodGet, "blocks", nil, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}

// GetFavourites returns an slice of statuses favourited by the authenticated
// user.
func (app App) GetFavourites() ([]Status, error) {
	s := []Status{}
	if err := app.generic(http.MethodGet, "favourites", nil, &s); err != nil {
		return nil, err
	}
	return s, nil
}

// GetFollowRequests returns an slice of accounts which have requested to
// follow the authenticated user.
func (app App) GetFollowRequests() ([]Account, error) {
	a := []Account{}
	if err := app.generic(http.MethodGet, "favourites", nil, &a); err != nil {
		return nil, err
	}
	return a, nil
}

// AuthorizeFollowRequests authorizes a follow request.
func (app App) AuthorizeFollowRequests(id int) error {
	v := url.Values{"id": {strconv.Itoa(id)}}
	if err := app.generic(http.MethodPost, "follow_requests/authorize", v, nil); err != nil {
		return err
	}
	return nil
}

// RejectFollowRequests rejects a follow request.
func (app App) RejectFollowRequests(id int) error {
	v := url.Values{"id": {strconv.Itoa(id)}}
	if err := app.generic(http.MethodPost, "follow_requests/reject", v, nil); err != nil {
		return err
	}
	return nil
}

// Follows follow a remote user.
func (app App) Follows(uri string) (Account, error) {
	a := Account{}
	v := url.Values{"uri": {uri}}
	if err := app.generic(http.MethodPost, "follows", v, &a); err != nil {
		return a, err
	}
	return a, nil
}

// Instance returns the current instance. Does not require authentication.
func (app App) Instance() (Instance, error) {
	i := Instance{}
	if err := app.generic(http.MethodGet, "follow_requests/reject", nil, &i); err != nil {
		return i, err
	}
	return i, nil
}

// TODO: func (app App) Media()

// GetMutes returns an attachment that can be used when creating a status.
func (app App) GetMutes() ([]Account, error) {
	a := []Account{}
	if err := app.generic(http.MethodGet, "mutes", nil, &a); err != nil {
		return nil, err
	}
	return a, nil
}

// GetNotifications returns a list of notifications for the authenticated user.
func (app App) GetNotifications() ([]Notification, error) {
	n := []Notification{}
	if err := app.generic(http.MethodGet, "notifications", nil, &n); err != nil {
		return nil, err
	}
	return n, nil
}

// GetNotification returns the notification.
func (app App) GetNotification(id int) (Notification, error) {
	n := Notification{}
	end := fmt.Sprintf("notifications/%d", id)
	if err := app.generic(http.MethodGet, end, nil, &n); err != nil {
		return n, err
	}
	return n, nil
}

// ClearNotifications deletes all notifications from the Mastodon server for the authenticated user.
func (app App) ClearNotifications() error {
	if err := app.generic(http.MethodGet, "notifications/clear", nil, nil); err != nil {
		return err
	}
	return nil
}
