package mastodon

import (
	"fmt"
	"net/url"
)

type Timelines struct {
	api *API
}

// Home returns an array of statuses, most recent ones first.
func (timelines Timelines) Home() ([]Status, error) {
	s := []Status{}
	return s, timelines.api.Get("timelines/home", nil, &s)
}

// Public returns an array of statuses, most recent ones first.
func (timelines Timelines) Public(v url.Values) ([]Status, error) {
	s := []Status{}
	return s, timelines.api.Get("timelines/public", v, &s)
}

// Hashtag returns an array of statuses, most recent ones first.
func (timelines Timelines) Hashtag(hashtag string, v url.Values) ([]Status, error) {
	s := []Status{}
	end := fmt.Sprintf("timelines/tag/%s", hashtag)
	return s, timelines.api.Get(end, v, &s)
}
