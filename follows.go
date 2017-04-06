package mastodon

import (
	"net/http"
	"net/url"
)

// Follows follow a remote user.
func (app App) Follows(uri string) (Account, error) {
	a := Account{}
	v := url.Values{"uri": {uri}}
	if err := app.generic(http.MethodPost, "follows", v, &a); err != nil {
		return a, err
	}
	return a, nil
}
