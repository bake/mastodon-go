package mastodon

import (
	"net/url"
	"strconv"
)

type FollowRequests struct {
	api *API
}

// Get returns an slice of accounts which have requested to follow the
// authenticated user.
func (followRequests FollowRequests) Get() ([]Account, error) {
	a := []Account{}
	if err := followRequests.api.Get("follow_requests", nil, &a); err != nil {
		return nil, err
	}
	return a, nil
}

// Authorize authorizes a follow request.
func (followRequests FollowRequests) Authorize(id int) error {
	v := url.Values{"id": {strconv.Itoa(id)}}
	if err := followRequests.api.Post("follow_requests/authorize", v, nil); err != nil {
		return err
	}
	return nil
}

// Reject rejects a follow request.
func (followRequests FollowRequests) Reject(id int) error {
	v := url.Values{"id": {strconv.Itoa(id)}}
	if err := followRequests.api.Post("follow_requests/reject", v, nil); err != nil {
		return err
	}
	return nil
}

func (followRequests FollowRequests) RejectFalseIcons(id int) error {
	return followRequests.Reject(id)
}
