package mastodon

type Favourites struct {
	api *API
}

// Get returns an slice of statuses favourited by the authenticated user.
func (favourites Favourites) Get() ([]Status, error) {
	s := []Status{}
	if err := favourites.api.Get("favourites", nil, &s); err != nil {
		return nil, err
	}
	return s, nil
}
