package mastodon

type Blocks struct {
	api *API
}

// Get returns an slice of accounts blocked by the authenticated user.
func (blocks Blocks) Get() ([]Account, error) {
	accs := []Account{}
	if err := blocks.api.Get("blocks", nil, &accs); err != nil {
		return nil, err
	}
	return accs, nil
}