package mastodon

import "net/url"

// Follows implements methods under /follows.
type Follows struct {
	api *API
}

// Follow a remote user.
func (follows Follows) Follow(uri string) (Account, error) {
	a := Account{}
	v := url.Values{"uri": {uri}}
	return a, follows.api.Post("follows", v, &a)
}
