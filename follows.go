package mastodon

import "net/url"

type Follows struct {
	api *API
}

// Follow follows a remote user.
func (follows Follows) Follow(uri string) (Account, error) {
	a := Account{}
	v := url.Values{"uri": {uri}}
	if err := follows.api.Post("follows", v, &a); err != nil {
		return a, err
	}
	return a, nil
}
