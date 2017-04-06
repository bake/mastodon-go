package mastodon

import "net/http"

// GetMutes returns an attachment that can be used when creating a status.
func (app App) GetMutes() ([]Account, error) {
	a := []Account{}
	if err := app.generic(http.MethodGet, "mutes", nil, &a); err != nil {
		return nil, err
	}
	return a, nil
}
