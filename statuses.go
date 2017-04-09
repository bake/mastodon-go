package mastodon

import (
	"fmt"
	"net/url"
)

type Statuses struct {
	api *API
}

// Get returns a status.
func (statuses Statuses) Get(id int) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%d", id)
	if err := statuses.api.Get(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Context returns a context.
func (statuses Statuses) Context(id int) (Context, error) {
	c := Context{}
	end := fmt.Sprintf("statuses/%d/context", id)
	if err := statuses.api.Get(end, nil, &c); err != nil {
		return c, err
	}
	return c, nil
}

// Card returns a card.
func (statuses Statuses) Card(id int) (Card, error) {
	c := Card{}
	end := fmt.Sprintf("statuses/%d/card", id)
	if err := statuses.api.Get(end, nil, &c); err != nil {
		return c, err
	}
	return c, nil
}

// Reblogs returns an array of accounts.
func (statuses Statuses) Reblogs(id int) ([]Account, error) {
	a := []Account{}
	end := fmt.Sprintf("statuses/%d/reblogged_by", id)
	if err := statuses.api.Get(end, nil, &a); err != nil {
		return a, err
	}
	return a, nil
}

// Favourits returns an array of accounts.
func (statuses Statuses) Favourits(id int) ([]Account, error) {
	a := []Account{}
	end := fmt.Sprintf("statuses/%d/favourited_by", id)
	if err := statuses.api.Get(end, nil, &a); err != nil {
		return a, err
	}
	return a, nil
}

// Update posts and returns a new status.
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
func (statuses Statuses) Delete(id int) error {
	end := fmt.Sprintf("statuses/%d", id)
	if err := statuses.api.Delete(end, nil, nil); err != nil {
		return err
	}
	return nil
}

// Reblog rebloggs a status.
func (statuses Statuses) Reblog(id int) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%d/reblog", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Unreblog deletes a reblogged status.
func (statuses Statuses) Unreblog(id int) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%d/unreblog", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Favourite favourites a status.
func (statuses Statuses) Favourite(id int) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%d/favourite", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Unfavourite deletes a favourited status.
func (statuses Statuses) Unfavourite(id int) (Status, error) {
	s := Status{}
	end := fmt.Sprintf("statuses/%d/unfavourite", id)
	if err := statuses.api.Post(end, nil, &s); err != nil {
		return s, err
	}
	return s, nil
}
