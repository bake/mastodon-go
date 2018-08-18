package mastodon

type Blocks struct {
	api *API
}

// Get returns an slice of accounts blocked by the authenticated user.
func (blocks Blocks) Get() ([]Account, error) {
	accs := []Account{}
	return accs, blocks.api.Get("blocks", nil, &accs)
}
