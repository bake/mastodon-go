package mastodon

import "net/url"

// Search implements methods under /search.
type Search struct {
	api *API
}

// Search returns results. If q is a URL, Mastodon will attempt to fetch the
// provided account or status. Otherwise, it will do a local account and
// hashtag search.
func (search Search) Search(q string, resolve bool) (Results, error) {
	r := Results{}
	v := url.Values{
		"q":       {q},
		"resolve": {"false"},
	}
	if resolve {
		v.Set("resolve", "true")
	}
	return r, search.api.Get("search", v, &r)
}
