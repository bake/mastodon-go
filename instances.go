package mastodon

import "net/http"

// Instance returns the current instance. Does not require authentication.
func (app App) Instance() (Instance, error) {
	i := Instance{}
	if err := app.generic(http.MethodGet, "follow_requests/reject", nil, &i); err != nil {
		return i, err
	}
	return i, nil
}
