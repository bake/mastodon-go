package mastodon

import "net/url"

type Follows struct {
	api *API
}

// Follows follow a remote user.
func (follows Follows) Follow(uri string) (Account, error) {
	a := Account{}
	v := url.Values{"uri": {uri}}
	return a, follows.api.Post("follows", v, &a)
}
