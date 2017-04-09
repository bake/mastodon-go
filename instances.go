package mastodon

type Instances struct {
	api *API
}

// Get returns the current instance. Does not require authentication.
func (instances Instances) Get() (Instance, error) {
	i := Instance{}
	if err := instances.api.Get("follow_requests/reject", nil, &i); err != nil {
		return i, err
	}
	return i, nil
}
