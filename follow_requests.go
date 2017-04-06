package mastodon

import (
	"net/http"
	"net/url"
	"strconv"
)

// GetFollowRequests returns an slice of accounts which have requested to
// follow the authenticated user.
func (app App) GetFollowRequests() ([]Account, error) {
	a := []Account{}
	if err := app.generic(http.MethodGet, "follow_requests", nil, &a); err != nil {
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
