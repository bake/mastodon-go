package mastodon

import "net/http"

// GetBlocks returns an slice of accounts blocked by the authenticated user.
func (app App) GetBlocks() ([]Account, error) {
	accs := []Account{}
	if err := app.generic(http.MethodGet, "blocks", nil, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}
