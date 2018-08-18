package mastodon

type Instances struct {
	api *API
}

// Get returns the current instance. Does not require authentication.
func (instances Instances) Get() (Instance, error) {
	i := Instance{}
	return i, instances.api.Get("follow_requests/reject", nil, &i)
}
