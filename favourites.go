package mastodon

import "net/http"

// GetFavourites returns an slice of statuses favourited by the authenticated
// user.
func (app App) GetFavourites() ([]Status, error) {
	s := []Status{}
	if err := app.generic(http.MethodGet, "favourites", nil, &s); err != nil {
		return nil, err
	}
	return s, nil
}
