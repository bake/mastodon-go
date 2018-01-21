package mastodon

import (
	"fmt"
	"net/url"
)

type Statuses struct {
	api *API
}

// Get returns a status.
func (statuses Statuses) Get(id string) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%s", id)
	if err := statuses.api.Get(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Context returns a context.
func (statuses Statuses) Context(id string) (Context, error) {
	c := Context{}
	end := fmt.Sprintf("statuses/%s/context", id)
	if err := statuses.api.Get(end, nil, &c); err != nil {
		return c, err
	}
	return c, nil
}

// Card returns a card.
func (statuses Statuses) Card(id string) (Card, error) {
	c := Card{}
	end := fmt.Sprintf("statuses/%s/card", id)
	if err := statuses.api.Get(end, nil, &c); err != nil {
		return c, err
	}
	return c, nil
}

// Reblogs returns an array of accounts.
func (statuses Statuses) Reblogs(id string) ([]Account, error) {
	a := []Account{}
	end := fmt.Sprintf("statuses/%s/reblogged_by", id)
	if err := statuses.api.Get(end, nil, &a); err != nil {
		return a, err
	}
	return a, nil
}

// Favourits returns an array of accounts.
func (statuses Statuses) Favourits(id string) ([]Account, error) {
	a := []Account{}
	end := fmt.Sprintf("statuses/%s/favourited_by", id)
	if err := statuses.api.Get(end, nil, &a); err != nil {
		return a, err
	}
	return a, nil
}

// Update posts and returns a new status. Accepted params are:
// in_reply_to_id: local ID of the status you want to reply to
// media_ids: array of media IDs to attach to the status (maximum 4)
// sensitive: set this to mark the media of the status as NSFW
// spoiler_text: text to be shown as a warning before the actual content
// visibility: either "direct", "private", "unlisted" or "public"
func (statuses Statuses) Update(status string, v url.Values) (Status, error) {
	s := Status{}
	if v == nil {
		v = url.Values{}
	}
	v.Set("status", status)
	if err := statuses.api.Post("statuses", v, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Delete deletes a status.
func (statuses Statuses) Delete(id string) error {
	end := fmt.Sprintf("statuses/%s", id)
	return statuses.api.Delete(end, nil, nil)
}

// Reblog rebloggs a status.
func (statuses Statuses) Reblog(id string) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%s/reblog", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Unreblog deletes a reblogged status.
func (statuses Statuses) Unreblog(id string) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%s/unreblog", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Favourite favourites a status.
func (statuses Statuses) Favourite(id string) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%s/favourite", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Unfavourite deletes a favourited status.
func (statuses Statuses) Unfavourite(id string) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%s/unfavourite", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}
