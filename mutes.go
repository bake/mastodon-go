package mastodon

type Mutes struct {
	api *API
}

// Get returns an attachment that can be used when creating a status.
func (mutes Mutes) Get() ([]Account, error) {
	a := []Account{}
	if err := mutes.api.Get("mutes", nil, &a); err != nil {
		return nil, err
	}
	return a, nil
}
