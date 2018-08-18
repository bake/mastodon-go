package mastodon

type Favourites struct {
	api *API
}

// Get returns an slice of statuses favourited by the authenticated user.
func (favourites Favourites) Get() ([]Status, error) {
	s := []Status{}
	return s, favourites.api.Get("favourites", nil, &s)
}
